package config

import (
	flag "github.com/spf13/pflag"

	"k8s.io/client-go/kubernetes"
	"k8s.io/cluster-registry/pkg/client/clientset_generated/clientset"
)

var (
	MockBasePath string

	CentralKubeconfig string
	CentralKubeAPIUrl string
	ClusterClient     *clientset.Clientset             //cluster registry API client	MockBasePath  string
	EdgeClients       map[string]*kubernetes.Clientset //map of edge cluster name --> a client to connect to that cluster

)

func init() {
	flag.StringVar(&MockBasePath, "mock-base-path", "", "Path to the directory containing mock response files.")
	flag.StringVar(&CentralKubeAPIUrl, "central-kube-api", "", "Kubernetes API server URL for the cluster running cluster-registry API")
	flag.StringVar(&CentralKubeconfig, "central-kubeconfig", "", "Path to the kubeconfig running cluster-registry API server")
	EdgeClients = map[string]*kubernetes.Clientset{}
}
