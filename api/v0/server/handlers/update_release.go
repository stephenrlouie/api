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

func NewUpdateRelease() *updateRelease {
	return &updateRelease{}
}

type updateRelease struct{}

func (d *updateRelease) Handle(params releases.UpdateReleaseParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	tillers, err := GetTillers(params.Labels)
	if err != nil {
		fmt.Printf("Error: Failed to find tillers: %v", err)
		return releases.NewGetReleasesInternalServerError()
	}

	err = helm.UpdateAllReleases(tillers, params.ReleaseID, &params.ChartTar)

	if err != nil {
		return releases.NewUpdateReleaseInternalServerError()
	}

	return releases.NewUpdateReleaseOK()
}

func (d *updateRelease) MockHandle(params releases.UpdateReleaseParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(config.MockBasePath, fmt.Sprintf("update-release-%s.json", params.ReleaseID)), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return releases.NewUpdateReleaseInternalServerError()
	}

	switch statusCode {
	case 200:
		return releases.NewUpdateReleaseOK()
	case 400:
		return releases.NewUpdateReleaseBadRequest()
	case 401:
		return releases.NewUpdateReleaseUnauthorized()
	case 404:
		return releases.NewUpdateReleaseNotFound()
	case 500:
		return releases.NewUpdateReleaseInternalServerError()
	}
	return releases.NewUpdateReleaseInternalServerError()
}
