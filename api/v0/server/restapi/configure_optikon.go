// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	flag "github.com/spf13/pflag"
	graceful "github.com/tylerb/graceful"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/cluster-registry/pkg/client/clientset_generated/clientset"

	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/clusters"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/releases"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../../server --name  --spec ../swagger.yaml --skip-models --exclude-main

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

func configureFlags(api *operations.OptikonAPI) {
	//api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{}
}

func configureAPI(api *operations.OptikonAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ReleasesAddReleasesHandler = releases.AddReleasesHandlerFunc(func(params releases.AddReleasesParams) middleware.Responder {
		return middleware.NotImplemented("operation releases.AddReleases has not yet been implemented")
	})
	api.ClustersAddClusterHandler = clusters.AddClusterHandlerFunc(func(params clusters.AddClusterParams) middleware.Responder {
		return middleware.NotImplemented("operation clusters.AddCluster has not yet been implemented")
	})
	api.ReleasesDeleteReleaseHandler = releases.DeleteReleaseHandlerFunc(func(params releases.DeleteReleaseParams) middleware.Responder {
		return middleware.NotImplemented("operation releases.DeleteRelease has not yet been implemented")
	})
	api.ClustersDeleteClusterHandler = clusters.DeleteClusterHandlerFunc(func(params clusters.DeleteClusterParams) middleware.Responder {
		return middleware.NotImplemented("operation clusters.DeleteCluster has not yet been implemented")
	})
	api.ReleasesGetReleaseByIDHandler = releases.GetReleaseByIDHandlerFunc(func(params releases.GetReleaseByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation releases.GetReleaseByID has not yet been implemented")
	})
	api.ReleasesGetReleasesHandler = releases.GetReleasesHandlerFunc(func(params releases.GetReleasesParams) middleware.Responder {
		return middleware.NotImplemented("operation releases.GetReleases has not yet been implemented")
	})
	api.ClustersGetClusterByIDHandler = clusters.GetClusterByIDHandlerFunc(func(params clusters.GetClusterByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation clusters.GetClusterByID has not yet been implemented")
	})
	api.ClustersGetClustersHandler = clusters.GetClustersHandlerFunc(func(params clusters.GetClustersParams) middleware.Responder {
		return middleware.NotImplemented("operation clusters.GetClusters has not yet been implemented")
	})
	api.ReleasesUpdateReleaseHandler = releases.UpdateReleaseHandlerFunc(func(params releases.UpdateReleaseParams) middleware.Responder {
		return middleware.NotImplemented("operation releases.UpdateRelease has not yet been implemented")
	})
	api.ClustersUpdateClusterHandler = clusters.UpdateClusterHandlerFunc(func(params clusters.UpdateClusterParams) middleware.Responder {
		return middleware.NotImplemented("operation clusters.UpdateCluster has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
	if MockBasePath != "" {
		return
	}
	// Read in central kubeconfig
	cfg, err := clientcmd.BuildConfigFromFlags(CentralKubeAPIUrl, CentralKubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s\n", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s\n", err.Error())
	}

	// Verify that we can reach the central cluster
	pods, err := kubeClient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Connected to central cluster- There are %d total pods\n", len(pods.Items))

	//  set up cluster registry client connection
	ClusterClient, err = clientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error reaching cluster-registry API: %s\n", err.Error())
	}
	log.Println("Connected to central cluster registry")

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.AllowAll()
	return corsHandler.Handler(handler)
}
