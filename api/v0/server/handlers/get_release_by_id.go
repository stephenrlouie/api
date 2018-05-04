package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"github.com/optikon/api/api/v0/mock"
	"github.com/optikon/api/api/v0/models"
	"github.com/optikon/api/api/v0/server/config"
	"github.com/optikon/api/api/v0/server/restapi/operations/releases"
)

func NewGetReleaseByID() *getReleaseById {
	return &getReleaseById{}
}

type getReleaseById struct{}

// Not Implemented
func (d *getReleaseById) Handle(params releases.GetReleaseByIDParams) middleware.Responder {
	if config.MockBasePath != "" {
		return d.MockHandle(params)
	}

	_, err := GetTillers(params.Labels)
	if err != nil {
		return releases.NewGetReleaseByIDBadRequest()
	}
	// do something with this tiller list.

	return releases.NewGetReleaseByIDOK()
}

func (d *getReleaseById) MockHandle(params releases.GetReleaseByIDParams) middleware.Responder {
	payload := models.ReleaseRelease{}
	statusCode, err := mock.GetMock(
		path.Join(config.MockBasePath, fmt.Sprintf("get-release-%s.json", params.ReleaseID)), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return releases.NewGetReleaseByIDInternalServerError()
	}

	switch statusCode {
	case 201:
		return releases.NewGetReleaseByIDOK().WithPayload(&payload)
	case 400:
		return releases.NewGetReleaseByIDBadRequest()
	case 401:
		return releases.NewGetReleaseByIDUnauthorized()
	case 404:
		return releases.NewGetReleaseByIDNotFound()
	case 500:
		return releases.NewGetReleaseByIDInternalServerError()
	}
	return releases.NewGetReleaseByIDInternalServerError()
}
