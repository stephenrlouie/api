package convert

// NESTED CONVERT FUNCTIONS -- cluster-reg ---> optikon
import (
	"k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
	"github.com/optikon/api/api/v0/models"
)

func RegToOptikonClusters(input *v1alpha1.ClusterList) []*models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster {
	output := []*models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster{}

	for _, c := range input.Items {
		conv := RegToOptikonCluster(c)
		output = append(output, &conv)
	}
	return output
}

func RegToOptikonCluster(c v1alpha1.Cluster) models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster {
	return models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster{
		APIVersion: c.APIVersion,
		Kind:       c.Kind,
		Metadata:   RegToOptikonClusterMeta(c),
		Spec:       RegToOptikonClusterSpec(c),
	}
}

func RegToOptikonClusterMeta(c v1alpha1.Cluster) *models.IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	meta := c.GetObjectMeta()

	return &models.IoK8sApimachineryPkgApisMetaV1ObjectMeta{
		Name:        meta.GetName(),
		Annotations: meta.GetAnnotations(),
		Labels:      meta.GetLabels(),
		Namespace:   meta.GetNamespace(),
	}
}

// TODO - implement. this lists auth info for cluster + K8s API endpoints
func RegToOptikonClusterSpec(c v1alpha1.Cluster) *models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1ClusterSpec {
	return &models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1ClusterSpec{}
}

// CONVERT FUNCTIONS: optikon API --> Cluster registry
// Note that i'm using the provided Setter functions bc. the v1 object metadata is
func OptikonToRegCluster(c models.IoK8sClusterRegistryPkgApisClusterregistryV1alpha1Cluster) *v1alpha1.Cluster {
	output := v1alpha1.Cluster{}
	output.APIVersion = c.APIVersion
	output.Kind = c.Kind

	output.SetName(c.Metadata.Name)

	output.SetAnnotations(c.Metadata.Annotations)
	output.SetLabels(c.Metadata.Labels)
	output.SetNamespace(c.Metadata.Namespace)
	return &output
}
