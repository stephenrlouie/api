package helm

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/version"
	"wwwin-github.cisco.com/edge/optikon/api/v0/convert"
	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
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

func (tc *TillerClient) ListReleases(address string) (obj []*models.ReleaseRelease, err error) {
	fmt.Printf("List Release Request\n")
	tc.execute(address, func(rsc tiller.ReleaseServiceClient, ctx context.Context, cancel context.CancelFunc) {
		req := tiller.ListReleasesRequest{}
		defer cancel()
		lrc, err := rsc.ListReleases(ctx, &req)
		if err != nil {
			fmt.Printf("Error: Failed to list releases: %v\n", err)
			return
		}

		res, err := lrc.Recv()
		if err != nil {
			fmt.Printf("Error: Failed to receive releases: %v\n", err)
		}

		for _, v := range res.Releases {
			oneRelease, err := convert.ReleaseToJSON(v)
			if err != nil {
				fmt.Printf("Error: Parsing release: %v\n", err)
			}
			obj = append(obj, oneRelease)
		}
	})
	return
}
