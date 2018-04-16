package handlers

import (
	"fmt"
	"path"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/clusterregistry"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/helm"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/releases"
)

func NewUpdateRelease() *updateRelease {
	return &updateRelease{}
}

type updateRelease struct{}

func (d *updateRelease) Handle(params releases.UpdateReleaseParams) middleware.Responder {
	fmt.Printf("updateRelease: %s\n", params.ReleaseID)
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}

	tillers, err := clusterregistry.GetTillers(params.Labels)
	if err != nil {
		fmt.Printf("Error: Failed to find tillers: %v", err)
		return releases.NewGetReleasesInternalServerError()
	}

	// TODO Async this
	for _, tiller := range tillers {
		tillerClient := helm.NewTillerClient(5 * time.Second)
		err := tillerClient.UpdateRelease(tiller, params.ReleaseID, &params.ChartTar)
		if err != nil {
			fmt.Printf("Error: Failed to install Release: %v on tiller: %s\n", err, tiller)
			return releases.NewUpdateReleaseInternalServerError()
		}
	}

	return releases.NewUpdateReleaseOK()
}

func (d *updateRelease) MockHandle(params releases.UpdateReleaseParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("update-release-%s.json", params.ReleaseID)), nil)
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
