package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/releases"
)

func NewDeleteRelease() *deleteRelease {
	return &deleteRelease{}
}

type deleteRelease struct{}

func (d *deleteRelease) Handle(params releases.DeleteReleaseParams) middleware.Responder {
	fmt.Printf("deleteRelease: %s\n", params.ReleaseID)
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return releases.NewDeleteReleaseOK()
}

func (d *deleteRelease) MockHandle(params releases.DeleteReleaseParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("delete-release-%s.json", params.ReleaseID)), nil)
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
