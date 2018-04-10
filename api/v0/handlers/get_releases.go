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

	var payload []*models.ReleaseRelease
	if len(restapi.TillersList) != 0 {
		var err error
		tillerClient := helm.NewTillerClient(30 * time.Second)
		payload, err = tillerClient.ListReleases(restapi.TillersList[0])
		if err != nil {
			fmt.Printf("Failed to read Releases: %v\n", err)
		}
		fmt.Printf("PAYLOAD: %+v\n", payload)
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
