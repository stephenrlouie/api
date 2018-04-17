package handlers

import (
	"fmt"
	"path"
	"time"

	"k8s.io/helm/pkg/chartutil"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/helm"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/config"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/releases"
)

func NewAddRelease() *addRelease {
	return &addRelease{}
}

type addRelease struct{}

func (d *addRelease) Handle(params releases.AddReleasesParams) middleware.Responder {
	fmt.Printf("AddReleases\n")
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	tillers, err := GetTillers(params.Labels)
	if err != nil {
		fmt.Printf("Error: Failed to find tillers: %v", err)
		return releases.NewGetReleasesInternalServerError()
	}

	// NOTE - only loading the tar once
	ch, err := chartutil.LoadArchive(&params.ChartTar)
	if err != nil {
		fmt.Printf("Chart load error: %v\n", err)
		return releases.NewGetReleasesInternalServerError()
	}

	// TODO Async this
	for _, tiller := range tillers {
		tillerClient := helm.NewTillerClient(5 * time.Second)
		err = tillerClient.InstallRelease(tiller, ch, params.Name, params.Namespace)
		if err != nil {
			fmt.Printf("Error: Failed to install Release: %v on tiller: %s\n", err, tiller)
			return releases.NewAddReleasesInternalServerError()
		}
	}
	return releases.NewAddReleasesCreated()
}

func (d *addRelease) MockHandle(params releases.AddReleasesParams) middleware.Responder {
	statusCode, err := mock.GetMock(path.Join(config.MockBasePath, "add-releases.json"), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return releases.NewAddReleasesInternalServerError()
	}

	switch statusCode {
	case 201:
		return releases.NewAddReleasesCreated()
	case 400:
		return releases.NewAddReleasesBadRequest()
	case 401:
		return releases.NewAddReleasesUnauthorized()
	case 500:
		return releases.NewAddReleasesInternalServerError()
	}
	return releases.NewAddReleasesInternalServerError()
}
