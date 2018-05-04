package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"github.com/optikon/api/api/v0/helm"
	"github.com/optikon/api/api/v0/mock"
	"github.com/optikon/api/api/v0/server/config"
	"github.com/optikon/api/api/v0/server/restapi/operations/releases"
)

func NewAddRelease() *addRelease {
	return &addRelease{}
}

type addRelease struct{}

func (d *addRelease) Handle(params releases.AddReleasesParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	tillers, err := GetTillers(params.Labels)
	if err != nil {
		fmt.Printf("Error: Failed to find tillers: %v", err)
		return releases.NewGetReleasesInternalServerError()
	}

	err = helm.InstallAllReleases(tillers, &params.ChartTar, params.Name, params.Namespace)
	if err != nil {
		return releases.NewAddReleasesInternalServerError()
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
