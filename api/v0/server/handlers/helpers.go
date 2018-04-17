package handlers

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/convert"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/config"
)

const TILLERKEY = "Tiller"

func GetClusters(labels *string) ([]*models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster, error) {
	list, err := config.ClusterClient.ClusterregistryV1alpha1().Clusters().List(v1.ListOptions{LabelSelector: safeLabels(labels)})

	conv := convert.RegToOptikonClusters(list)

	if err != nil {
		fmt.Printf("Error: Searching for clusters with labels %v", err)
		return conv, err

	}
	return conv, nil
}

func GetTillers(labels *string) ([]string, error) {
	list, err := GetClusters(labels)
	ret := make([]string, len(list))

	if err != nil {
		fmt.Printf("Error: Searching for clusters with labels %v", err)
		return ret, err
	}

	for i, c := range list {
		ret[i] = c.Metadata.Annotations[TILLERKEY]
	}
	return ret, nil
}

func GetTillersToClusterName(labels *string) (map[string]string, error) {
	output := map[string]string{}

	list, err := GetClusters(labels)
	if err != nil {
		fmt.Printf("Error: Searching for clusters with labels %v", err)
		return output, err
	}

	for _, c := range list {
		output[c.Metadata.Annotations[TILLERKEY]] = c.Metadata.Name
	}
	return output, nil
}

func safeLabels(labels *string) string {
	if labels == nil {
		return ""
	}
	return *labels
}
