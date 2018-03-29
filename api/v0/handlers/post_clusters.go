package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/clusters"
)

func NewAddCluster() *addCluster {
	return &addCluster{}
}

type addCluster struct{}

func (d *addCluster) Handle(params clusters.AddClusterParams) middleware.Responder {
	fmt.Printf("addClusters\n")
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return clusters.NewAddClusterCreated()
}

func (d *addCluster) MockHandle(params clusters.AddClusterParams) middleware.Responder {
	statusCode, err := mock.GetMock(path.Join(restapi.MockBasePath, "add-clusters.json"), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return clusters.NewAddClusterInternalServerError()
	}

	switch statusCode {
	case 201:
		return clusters.NewAddClusterCreated()
	case 400:
		return clusters.NewAddClusterBadRequest()
	case 401:
		return clusters.NewAddClusterUnauthorized()
	case 500:
		return clusters.NewAddClusterInternalServerError()
	}
	return clusters.NewAddClusterInternalServerError()
}
