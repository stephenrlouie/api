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

func NewGetClusterByID() *getClusterById {
	return &getClusterById{}
}

type getClusterById struct{}

func (d *getClusterById) Handle(params clusters.GetClusterByIDParams) middleware.Responder {
	fmt.Printf("getClusterById: %s\n", params.ClusterID)
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return clusters.NewGetClusterByIDOK()
}

func (d *getClusterById) MockHandle(params clusters.GetClusterByIDParams) middleware.Responder {
	payload := models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster{}
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("get-cluster-%s.json", params.ClusterID)), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return clusters.NewGetClusterByIDInternalServerError()
	}

	switch statusCode {
	case 201:
		return clusters.NewGetClusterByIDOK().WithPayload(&payload)
	case 400:
		return clusters.NewGetClusterByIDBadRequest()
	case 401:
		return clusters.NewGetClusterByIDUnauthorized()
	case 404:
		return clusters.NewGetClusterByIDNotFound()
	case 500:
		return clusters.NewGetClusterByIDInternalServerError()
	}
	return clusters.NewGetClusterByIDInternalServerError()
}
