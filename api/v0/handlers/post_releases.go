package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/releases"
)

func NewAddRelease() *addRelease {
	return &addRelease{}
}

type addRelease struct{}

func (d *addRelease) Handle(params releases.AddReleasesParams) middleware.Responder {
	fmt.Printf("AddReleases\n")
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return releases.NewAddReleasesCreated()
}

func (d *addRelease) MockHandle(params releases.AddReleasesParams) middleware.Responder {
	statusCode, err := mock.GetMock(path.Join(restapi.MockBasePath, "add-releases.json"), nil)
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
