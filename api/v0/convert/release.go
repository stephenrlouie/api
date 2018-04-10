package convert

import (
	"fmt"

	"k8s.io/helm/pkg/proto/hapi/release"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// Takes GRPC release structure and converts it to opitkon API structure
func ReleaseToJSON(rel *release.Release) (*models.ReleaseRelease, error) {
	fmt.Printf("ReleaseToJSON\n")
	fmt.Printf("Received: %+v\n", rel)
	hks, err := HooksToJSON(rel.Hooks)
	if err != nil {
		return &models.ReleaseRelease{}, err
	}

	ifo, err := InfoToJSON(rel.Info)
	if err != nil {
		return &models.ReleaseRelease{}, err
	}

	return &models.ReleaseRelease{
		Config:    ConfigToJSON(rel.Config),
		Hooks:     hks,
		Info:      ifo,
		Manifest:  rel.Manifest,
		Name:      rel.Name,
		Namespace: rel.Namespace,
		Version:   rel.Version,
	}, nil
}
