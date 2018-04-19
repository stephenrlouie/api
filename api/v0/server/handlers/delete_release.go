package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/helm"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/config"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/releases"
)

func NewDeleteRelease() *deleteRelease {
	return &deleteRelease{}
}

type deleteRelease struct{}

func (d *deleteRelease) Handle(params releases.DeleteReleaseParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	tillers, err := GetTillers(params.Labels)
	if err != nil {
		fmt.Printf("Error: Failed to find tillers: %v", err)
		return releases.NewDeleteReleaseInternalServerError()
	}

	err = helm.DeleteAllReleases(tillers, params.ReleaseID)
	if err != nil {
		return releases.NewDeleteReleaseInternalServerError()
	}

	return releases.NewDeleteReleaseOK()
}

func (d *deleteRelease) MockHandle(params releases.DeleteReleaseParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(config.MockBasePath, fmt.Sprintf("delete-release-%s.json", params.ReleaseID)), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return releases.NewDeleteReleaseInternalServerError()
	}

	switch statusCode {
	case 200:
		return releases.NewDeleteReleaseOK()
	case 400:
		return releases.NewDeleteReleaseBadRequest()
	case 401:
		return releases.NewDeleteReleaseUnauthorized()
	case 404:
		return releases.NewDeleteReleaseNotFound()
	case 500:
		return releases.NewDeleteReleaseInternalServerError()
	}
	return releases.NewDeleteReleaseInternalServerError()
}
