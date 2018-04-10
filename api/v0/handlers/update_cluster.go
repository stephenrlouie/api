package handlers

import (
	"fmt"
	"log"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/convert"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"

	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/clusters"
)

func NewUpdateCluster() *updateCluster {
	return &updateCluster{}
}

type updateCluster struct{}

func (d *updateCluster) Handle(params clusters.UpdateClusterParams) middleware.Responder {
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}

	conv := convert.OptikonToRegCluster(*params.Body)

	updatedCluster, err := restapi.ClusterClient.ClusterregistryV1alpha1().Clusters().Update(conv)
	if err != nil {
		fmt.Println(err)
		return clusters.NewAddClusterInternalServerError()
	}
	log.Printf("Updated cluster: %s", updatedCluster.GetName())
	return clusters.NewUpdateClusterOK()
}

func (d *updateCluster) MockHandle(params clusters.UpdateClusterParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("update-cluster-%s.json", params.ClusterID)), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return clusters.NewUpdateClusterInternalServerError()
	}

	switch statusCode {
	case 200:
		return clusters.NewUpdateClusterOK()
	case 400:
		return clusters.NewUpdateClusterBadRequest()
	case 401:
		return clusters.NewUpdateClusterUnauthorized()
	case 404:
		return clusters.NewUpdateClusterNotFound()
	case 500:
		return clusters.NewUpdateClusterInternalServerError()
	}
	return clusters.NewUpdateClusterInternalServerError()
}
