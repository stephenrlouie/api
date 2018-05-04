package handlers

import (
	"fmt"
	"path"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/optikon/api/api/v0/convert"
	"github.com/optikon/api/api/v0/mock"
	"github.com/optikon/api/api/v0/models"

	"github.com/optikon/api/api/v0/server/config"
	"github.com/optikon/api/api/v0/server/restapi/operations/clusters"
)

func NewGetClusters() *getClusters {
	return &getClusters{}
}

type getClusters struct{}

func (d *getClusters) Handle(params clusters.GetClustersParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	// Call cluster registry
	allClusters, err := config.ClusterClient.ClusterregistryV1alpha1().Clusters().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return clusters.NewGetClustersInternalServerError()
	}

	conv := convert.RegToOptikonClusters(allClusters)

	// Get # of pods for this cluster ("health")
	for _, cl := range conv {
		if _, ok := config.EdgeClients[cl.Metadata.Name]; !ok {
			fmt.Printf("Cannot get pods for cluster, No client-go API client for edge cluster=%s\n", cl.Metadata.Name)
		} else {
			pods, err := config.EdgeClients[cl.Metadata.Name].CoreV1().Pods("").List(metav1.ListOptions{})
			if err != nil {
				fmt.Println(err)
				return clusters.NewGetClustersInternalServerError()
			}
			cl.Metadata.Annotations["NumPods"] = strconv.Itoa(len(pods.Items))
		}
	}

	return clusters.NewGetClustersOK().WithPayload(conv)
}

func (d *getClusters) MockHandle(params clusters.GetClustersParams) middleware.Responder {
	payload := models.GetClustersOKBody{}
	statusCode, err := mock.GetMock(path.Join(config.MockBasePath, "get-clusters.json"), &payload)
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
