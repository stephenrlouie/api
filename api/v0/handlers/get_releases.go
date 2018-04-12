package handlers

import (
	"fmt"
	"path"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/helm"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/releases"
)

func NewGetReleases() *getReleases {
	return &getReleases{}
}

type getReleases struct{}

func (d *getReleases) Handle(params releases.GetReleasesParams) middleware.Responder {
	fmt.Printf("GetReleases\n")
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}

	// TODO: Async these requests to speed up.
	var payload []*models.ReleaseRelease
	for _, tiller := range restapi.TillersList {
		tillerClient := helm.NewTillerClient(5 * time.Second)
		singleResult, err := tillerClient.ListReleases(tiller)
		if err != nil {
			fmt.Printf("Failed to read Releases: %v\n", err)
		}
		payload = append(payload, singleResult...)
	}

	return releases.NewGetReleasesOK().WithPayload(payload)
}

func (d *getReleases) MockHandle(params releases.GetReleasesParams) middleware.Responder {
	payload := models.GetReleasesOKBody{}
	statusCode, err := mock.GetMock(path.Join(restapi.MockBasePath, "get-releases.json"), &payload)
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
