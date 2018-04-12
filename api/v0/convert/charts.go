package convert

import (
	"github.com/golang/protobuf/ptypes/any"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

func ChartToJSON(chart *chart.Chart) *models.ChartChart {
	deps := []*models.ChartChart{}
	for _, dep := range chart.Dependencies {
		deps = append(deps, ChartToJSON(dep))
	}

	ret := models.ChartChart{
		Config:       ConfigToJSON(chart.Values),
		Metadata:     MetadataToJSON(chart.Metadata),
		Template:     TemplateToJSON(chart.Templates),
		Files:        ProtoBufToJSON(chart.Files),
		Dependencies: deps,
	}
	return &ret
}

func ProtoBufToJSON(filesList []*any.Any) []*models.ProtobufAny {
	ret := []*models.ProtobufAny{}
	for _, item := range filesList {
		ret = append(ret, &models.ProtobufAny{
			TypeURL: item.TypeUrl,
			Value:   item.Value,
		})
	}
	return ret

}
func TemplateToJSON(tempList []*chart.Template) []*models.ChartTemplate {
	ret := []*models.ChartTemplate{}
	for _, item := range tempList {
		ret = append(ret, &models.ChartTemplate{
			Name: item.Name,
			Data: item.Data,
		})
	}
	return ret
}

// TODO Annotations
func MetadataToJSON(meta *chart.Metadata) *models.ChartMetadata {
	return &models.ChartMetadata{
		Name:          meta.Name,
		Home:          meta.Home,
		Sources:       meta.Sources,
		Version:       meta.Version,
		Description:   meta.Description,
		Keywords:      meta.Keywords,
		Maintainers:   MaintainersToJSON(meta.Maintainers),
		Engine:        meta.Engine,
		Icon:          meta.Icon,
		APIVersion:    meta.ApiVersion,
		Condition:     meta.Condition,
		Tags:          meta.Tags,
		AppVersion:    meta.AppVersion,
		Deprecated:    meta.Deprecated,
		TillerVersion: meta.TillerVersion,
		KubeVersion:   meta.KubeVersion,
		Annotations:   AnnotationsToJSON(meta.Annotations),
	}
}

func MaintainersToJSON(mainList []*chart.Maintainer) []*models.ChartMaintainer {
	ret := []*models.ChartMaintainer{}
	for _, item := range mainList {
		ret = append(ret, &models.ChartMaintainer{
			Email: item.Email,
			Name:  item.Name,
			URL:   item.Url,
		})
	}
	return ret
}

func AnnotationsToJSON(annotations map[string]string) models.ChartMap {
	ret := []*models.ChartValue{}
	for k, v := range annotations {
		ret = append(ret, &models.ChartValue{
			Key:   k,
			Value: v,
		})
	}
	return ret
}

func ConfigToJSON(cfg *chart.Config) *models.ChartConfig {
	vals := models.ChartMap{}
	for k, v := range cfg.Values {
		vals = append(vals, &models.ChartValue{
			Key:   k,
			Value: v.Value,
		})
	}
	return &models.ChartConfig{Raw: cfg.Raw, Values: vals}
}
