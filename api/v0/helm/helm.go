package helm

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"k8s.io/helm/pkg/chartutil"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/version"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/convert"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

type TillerClient struct {
	timeout time.Duration
}

// Test
func NewTillerClient(timeout time.Duration) *TillerClient {
	return &TillerClient{timeout}
}

func (tc *TillerClient) execute(address string, request func(tiller.ReleaseServiceClient, context.Context, context.CancelFunc)) error {
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

func (tc *TillerClient) ListReleases(address string) (obj []*models.ReleaseRelease, retErr error) {
	fmt.Printf("List Release Request\n")
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

func (tc *TillerClient) InstallRelease(address string, chartTar io.Reader, name string, namespace string) (retErr error) {
	fmt.Printf("Install Release Request\n")
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {
		ch, err := chartutil.LoadArchive(chartTar)
		if err != nil {
			retErr = err
			fmt.Printf("Error: Failed to read tar file %v\n", err)
			return
		}

		req := tiller.InstallReleaseRequest{
			Chart:     ch,
			Values:    ch.GetValues(),
			Name:      name,
			Namespace: namespace,
		}
		defer cancel()

		_, err = rsc.InstallRelease(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to install release: %v\n", err)
			retErr = err
			return
		}
	})
	return
}

func (tc *TillerClient) DeleteRelease(address string, releaseID string) (retErr error) {
	fmt.Printf("Delete Release Request\n")
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

func (tc *TillerClient) UpdateRelease(address string, releaseID string, chartTar io.Reader) (retErr error) {
	fmt.Printf("Update Release Request\n")
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {
		ch, err := chartutil.LoadArchive(chartTar)
		if err != nil {
			retErr = err
			fmt.Printf("Error: Failed to read tar file %v\n", err)
			return
		}

		req := tiller.UpdateReleaseRequest{
			Chart:  ch,
			Values: ch.GetValues(),
			Name:   releaseID,
		}
		defer cancel()

		_, err = rsc.UpdateRelease(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to install release: %v\n", err)
			retErr = err
			return
		}
	})
	return
}
