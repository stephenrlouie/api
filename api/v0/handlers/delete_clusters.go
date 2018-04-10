package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/clusters"
)

func NewDeleteCluster() *deleteCluster {
	return &deleteCluster{}
}

type deleteCluster struct{}

func (d *deleteCluster) Handle(params clusters.DeleteClusterParams) middleware.Responder {
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}

	err := restapi.ClusterClient.ClusterregistryV1alpha1().Clusters().Delete(params.ClusterID, &v1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return clusters.NewDeleteClusterInternalServerError()
	}
	fmt.Printf("Deleted cluster: %s\n", params.ClusterID)
	return clusters.NewDeleteClusterOK()
}

func (d *deleteCluster) MockHandle(params clusters.DeleteClusterParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("delete-cluster-%s.json", params.ClusterID)), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return clusters.NewDeleteClusterInternalServerError()
	}

	switch statusCode {
	case 200:
		return clusters.NewDeleteClusterOK()
	case 400:
		return clusters.NewDeleteClusterBadRequest()
	case 401:
		return clusters.NewDeleteClusterUnauthorized()
	case 404:
		return clusters.NewDeleteClusterNotFound()
	case 500:
		return clusters.NewDeleteClusterInternalServerError()
	}
	return clusters.NewDeleteClusterInternalServerError()
}
