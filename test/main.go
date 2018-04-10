package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kubernetes/helm/pkg/proto/hapi/services"
	"github.com/kubernetes/helm/pkg/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("192.168.100.102:30134", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the Tiller: %v", err)
	} else {
		fmt.Printf("Connected\n")
	}
	defer conn.Close()
	RSC := services.NewReleaseServiceClient(conn)
	rels := services.ListReleasesRequest{}
	//ctx := context.Background()
	md := metadata.Pairs("x-helm-api-client", version.Version)
	ctx := metadata.NewOutgoingContext(context.TODO(), md)
	ctx2, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	lrc, err := RSC.ListReleases(ctx2, &rels)
	if err != nil {
		fmt.Printf("FAILED TO LIST RELEASES: %v\n", err)
		os.Exit(1)
	}
	stuff, err := lrc.Recv()
	if err != nil {
		fmt.Printf("Failed to receive: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("CTX: %+v\nLRC: %+v\n", ctx, stuff)
}
