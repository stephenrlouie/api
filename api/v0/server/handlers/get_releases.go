package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"github.com/optikon/api/api/v0/helm"
	"github.com/optikon/api/api/v0/mock"
	"github.com/optikon/api/api/v0/models"
	"github.com/optikon/api/api/v0/server/config"
	"github.com/optikon/api/api/v0/server/restapi/operations/releases"
)

func NewGetReleases() *getReleases {
	return &getReleases{}
}

type getReleases struct{}

func (d *getReleases) Handle(params releases.GetReleasesParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	tillersMap, err := GetTillersToClusterName(params.Labels)
	if err != nil {
		fmt.Printf("Error: Failed to get map of tiller -> cluster name: %v", err)
		return releases.NewGetReleasesInternalServerError()
	}

	payload, err := helm.ListAllReleases(tillersMap)
	if err != nil {
		return releases.NewGetReleasesInternalServerError()
	}

	return releases.NewGetReleasesOK().WithPayload(payload)
}

func (d *getReleases) MockHandle(params releases.GetReleasesParams) middleware.Responder {
	payload := models.GetReleasesOKBody{}
	statusCode, err := mock.GetMock(path.Join(config.MockBasePath, "get-releases.json"), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return releases.NewGetReleasesInternalServerError()
	}

	switch statusCode {
	case 200:
		return releases.NewGetReleasesOK().WithPayload(payload)
	case 400:
		return releases.NewGetReleasesBadRequest()
	case 401:
		return releases.NewGetReleasesUnauthorized()
	case 500:
		return releases.NewGetReleasesInternalServerError()
	}
	return releases.NewGetReleasesInternalServerError()
}
