package handlers

import (
	"fmt"
	"path"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/convert"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi"

	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/clusters"
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

	// Call cluster registry
	allClusters, err := restapi.ClusterClient.ClusterregistryV1alpha1().Clusters().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return clusters.NewGetClustersInternalServerError()
	}

	conv := convert.RegToOptikonClusters(allClusters)

	for _, cl := range conv {
		// Use regular kube api client for this cluster, to get # pods
		fmt.Printf("\n\n\n EDGE CLIENTS IS: %+v\n\n\n", restapi.EdgeClients)
		pods, err := restapi.EdgeClients[cl.Metadata.Name].CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			fmt.Println(err)
			return clusters.NewGetClustersInternalServerError()
		}
		cl.Metadata.Annotations["NumPods"] = strconv.Itoa(len(pods.Items))
	}

	return clusters.NewGetClustersOK().WithPayload(conv)
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
