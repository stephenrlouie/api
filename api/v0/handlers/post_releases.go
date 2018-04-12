package handlers

import (
	"fmt"
	"path"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/helm"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon-api/api/v0/server/restapi/operations/releases"
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

	// TODO Async this
	for _, tiller := range restapi.TillersList {
		tillerClient := helm.NewTillerClient(5 * time.Second)
		err := tillerClient.InstallRelease(tiller, &params.ChartTar, params.Name, params.Namespace)
		if err != nil {
			fmt.Printf("Failed to install Release: %v\n", err)
			return releases.NewAddReleasesInternalServerError()
		}
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
