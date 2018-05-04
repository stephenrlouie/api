package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/optikon/api/api/v0/convert"
	"github.com/optikon/api/api/v0/mock"

	"github.com/optikon/api/api/v0/server/config"
	"github.com/optikon/api/api/v0/server/restapi/operations/clusters"
)

func NewAddCluster() *addCluster {
	return &addCluster{}
}

type addCluster struct{}

func (d *addCluster) Handle(params clusters.AddClusterParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	conv := convert.OptikonToRegCluster(*params.Body)

	// --------------- add cluster to central cluster-registry --------------- --------------- ---------------
	createdCluster, err := config.ClusterClient.ClusterregistryV1alpha1().Clusters().Create(conv)
	if err != nil {
		fmt.Println(err)
		return clusters.NewAddClusterInternalServerError()
	}

	// --------------- add client info to map  --------------- --------------- ---------------

	if conf, ok := createdCluster.GetAnnotations()["Conf"]; ok {
		// TODO - clean this up - the k8s api needs a kubeconfig FILE but the kubeconfig here is posted as a string, in the JSON
		// so write out a temporary file to use (hack for now)
		if apiServerUrl, ok := createdCluster.GetAnnotations()["APIServer"]; ok {
			err := ioutil.WriteFile("/tmp/curconf", []byte(conf), 0644)
			cfg, err := clientcmd.BuildConfigFromFlags(apiServerUrl, "/tmp/curconf")
			if err != nil {
				log.Fatalf("Error building kubeconfig: %s\n", err.Error())
			}
			temp, err := kubernetes.NewForConfig(cfg)
			if err != nil {
				log.Fatalf("Error building kubernetes clientset: %s\n", err.Error())
			}
			config.EdgeClients[createdCluster.GetName()] = temp
			fmt.Println("added new edge client")
		}
	}

	log.Printf("Created cluster: %s", createdCluster.GetName())
	return clusters.NewAddClusterCreated()
}

func (d *addCluster) MockHandle(params clusters.AddClusterParams) middleware.Responder {
	statusCode, err := mock.GetMock(path.Join(config.MockBasePath, "add-clusters.json"), nil)
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
