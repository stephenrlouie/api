package convert

import (
	"fmt"

	"k8s.io/helm/pkg/proto/hapi/chart"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

func ConfigToJSON(cfg *chart.Config) *models.ReleaseConfig {
	fmt.Printf("ConfigToJSON\n")
	fmt.Printf("Received:  %+v\n", cfg)
	vals := models.ReleaseConfigValues{}
	for k, v := range cfg.Values {
		vals = append(vals, &models.ReleaseConfigValue{
			Key:   k,
			Value: v.Value,
		})
	}
	return &models.ReleaseConfig{Raw: cfg.Raw, Values: vals}
}
