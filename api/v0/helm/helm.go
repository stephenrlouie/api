package helm

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/optikon/api/api/v0/client/releases"
	"github.com/optikon/api/api/v0/convert"
	"github.com/optikon/api/api/v0/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/version"
)

type tillerClient struct {
	timeout time.Duration
}

var grpcTimeout = 10 * time.Second

func newTillerClient(timeout time.Duration) *tillerClient {
	return &tillerClient{timeout}
}

func (tc *tillerClient) execute(address string, request func(tiller.ReleaseServiceClient, context.Context, context.CancelFunc)) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: unable to dial tiller\n")
		return err
	}
	rsc := tiller.NewReleaseServiceClient(conn)

	md := metadata.Pairs("x-helm-api-client", version.Version)
	ctx := metadata.NewOutgoingContext(context.TODO(), md)
	ctx2, cancel := context.WithTimeout(ctx, tc.timeout)

	request(rsc, ctx2, cancel)
	return nil
}

func ListAllReleases(tillersMap map[string]string) ([]*models.ReleaseRelease, error) {
	wg := sync.WaitGroup{}
	doneChan := make(chan bool, 1)
	errChan := make(chan error, len(tillersMap))
	sharedModels := make([][]*models.ReleaseRelease, len(tillersMap))
	var payload []*models.ReleaseRelease

	index := 0
	for tillerIP, clusterName := range tillersMap {
		wg.Add(1)
		go func(tillerIP string, clusterName string, index int, errChan chan error) {
			defer wg.Done()
			tillerClient := newTillerClient(grpcTimeout)
			singleResult, err := tillerClient.listReleases(tillerIP)
			if err != nil {
				fmt.Printf("Error: Failed to read Releases: %v on tiller: %s\n", err, tillerIP)
				errChan <- err
				return
			}

			for _, r := range singleResult {
				r.OnCluster = clusterName
			}
			sharedModels[index] = singleResult
		}(tillerIP, clusterName, index, errChan)
		index++
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			fmt.Printf("Error from Tiller: %v\n", err)
			return payload, err
		}
	case <-doneChan:
	}

	// Flatten the 2D array of models
	for _, models := range sharedModels {
		payload = append(payload, models...)
	}
	return payload, nil
}

func (tc *tillerClient) listReleases(address string) (obj []*models.ReleaseRelease, retErr error) {
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {
		req := tiller.ListReleasesRequest{}
		defer cancel()
		resp, err := rsc.ListReleases(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to list releases: %v\n", err)
			retErr = err
			return
		}

		res, err := resp.Recv()
		if err != nil {
			fmt.Printf("Error: Failed to receive releases: %v\n", err)
			retErr = err
			return
		}

		for _, v := range res.Releases {
			oneRelease, err := convert.ReleaseToJSON(v)
			if err != nil {
				fmt.Printf("Error: Parsing release: %v\n", err)
				retErr = err
				return
			}
			obj = append(obj, oneRelease)
		}
	})
	return
}

func InstallAllReleases(addresses []string, chartTar io.Reader, name string, namespace string) error {
	// NOTE - only loading the tar once
	ch, err := chartutil.LoadArchive(chartTar)
	if err != nil {
		fmt.Printf("Chart load error: %v\n", err)
		return releases.NewGetReleasesInternalServerError()
	}

	wg := sync.WaitGroup{}
	doneChan := make(chan bool, 1)
	errChan := make(chan error, len(addresses))

	for _, tiller := range addresses {
		wg.Add(1)
		go func(tillerIP string, ch *chart.Chart, name string, namespace string, errChan chan error) {
			defer wg.Done()
			tillerClient := newTillerClient(grpcTimeout)
			err = tillerClient.installRelease(tillerIP, ch, name, namespace)
			if err != nil {
				fmt.Printf("Error: Failed to install Release: %v on tiller: %s\n", err, tillerIP)
				errChan <- err
				return
			}
		}(tiller, ch, name, namespace, errChan)
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			fmt.Printf("Error from Tiller: %v\n", err)
			return err
		}
	case <-doneChan:
	}

	return nil
}

func (tc *tillerClient) installRelease(address string, ch *chart.Chart, name string, namespace string) (retErr error) {
	fmt.Printf("Install Release Request: %s\n", address)
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {
		req := tiller.InstallReleaseRequest{
			Chart:     ch,
			Values:    ch.GetValues(),
			Name:      name,
			Namespace: namespace,
		}
		defer cancel()

		_, err := rsc.InstallRelease(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to install release: %v\n", err)
			retErr = err
			return
		}
	})
	return
}

func DeleteAllReleases(addresses []string, releaseID string) error {
	wg := sync.WaitGroup{}
	doneChan := make(chan bool, 1)
	errChan := make(chan error, len(addresses))

	for _, tiller := range addresses {
		wg.Add(1)
		go func(tillerIP string, releaseID string, errChan chan error) {
			defer wg.Done()
			tillerClient := newTillerClient(grpcTimeout)
			err := tillerClient.deleteRelease(tillerIP, releaseID)
			if err != nil {
				fmt.Printf("Error: Failed to delete Release: %v on tiller: %s\n", err, tillerIP)
				errChan <- err
				return
			}
		}(tiller, releaseID, errChan)
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			fmt.Printf("Error from Tiller: %v\n", err)
			return err
		}
	case <-doneChan:
	}

	return nil
}

func (tc *tillerClient) deleteRelease(address string, releaseID string) (retErr error) {
	fmt.Printf("Delete Release Request: %s\n", address)
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {
		req := tiller.UninstallReleaseRequest{
			Name:  releaseID,
			Purge: true,
		}
		defer cancel()

		_, err := rsc.UninstallRelease(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to uninstall releases: %v\n", err)
			retErr = err
			return
		}
	})
	return
}

func UpdateAllReleases(addresses []string, releaseID string, chartTar io.Reader) error {
	// NOTE - only loading the tar once
	ch, err := chartutil.LoadArchive(chartTar)
	if err != nil {
		fmt.Printf("Chart load error: %v\n", err)
		return err
	}

	wg := sync.WaitGroup{}
	doneChan := make(chan bool, 1)
	errChan := make(chan error, len(addresses))

	for _, tiller := range addresses {
		wg.Add(1)
		go func(tillerIP string, releaseID string, ch *chart.Chart, errChan chan error) {
			defer wg.Done()
			tillerClient := newTillerClient(grpcTimeout)
			err := tillerClient.updateRelease(tillerIP, releaseID, ch)
			if err != nil {
				fmt.Printf("Error: Failed to install Release: %v on tiller: %s\n", err, tillerIP)
				errChan <- err
				return
			}
		}(tiller, releaseID, ch, errChan)
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			fmt.Printf("Error from Tiller: %v\n", err)
			return err
		}
	case <-doneChan:
	}
	return nil
}

func (tc *tillerClient) updateRelease(address string, releaseID string, ch *chart.Chart) (retErr error) {
	fmt.Printf("Update Release Request: %s\n", address)
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {

		req := tiller.UpdateReleaseRequest{
			Chart:  ch,
			Values: ch.GetValues(),
			Name:   releaseID,
		}
		defer cancel()

		_, err := rsc.UpdateRelease(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to install release: %v\n", err)
			retErr = err
			return
		}
	})
	return
}
