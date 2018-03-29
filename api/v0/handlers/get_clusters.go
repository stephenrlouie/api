package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/clusters"
)

func NewGetClusters() *getClusters {
	return &getClusters{}
}

type getClusters struct{}

func (d *getClusters) Handle(params clusters.GetClustersParams) middleware.Responder {
	fmt.Printf("getClusters\n")
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	payload := models.GetClustersOKBody{}
	return clusters.NewGetClustersOK().WithPayload(payload)
}

func (d *getClusters) MockHandle(params clusters.GetClustersParams) middleware.Responder {
	payload := models.GetClustersOKBody{}
	statusCode, err := mock.GetMock(path.Join(restapi.MockBasePath, "get-clusters.json"), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return clusters.NewGetClustersInternalServerError()
	}

	switch statusCode {
	case 200:
		return clusters.NewGetClustersOK().WithPayload(payload)
	case 400:
		return clusters.NewGetClustersBadRequest()
	case 401:
		return clusters.NewGetClustersUnauthorized()
	case 500:
		return clusters.NewGetClustersInternalServerError()
	}
	return clusters.NewGetClustersInternalServerError()
}
