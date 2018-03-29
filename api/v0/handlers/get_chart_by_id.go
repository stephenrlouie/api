package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/charts"
)

func NewGetChartByID() *getChartById {
	return &getChartById{}
}

type getChartById struct{}

func (d *getChartById) Handle(params charts.GetChartByIDParams) middleware.Responder {
	fmt.Printf("GetChartById: %s\n", params.ChartID)
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return charts.NewGetChartByIDOK()
}

func (d *getChartById) MockHandle(params charts.GetChartByIDParams) middleware.Responder {
	payload := models.ChartChart{}
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("get-chart-%s.json", params.ChartID)), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return charts.NewGetChartByIDInternalServerError()
	}

	switch statusCode {
	case 201:
		return charts.NewGetChartByIDOK().WithPayload(&payload)
	case 400:
		return charts.NewGetChartByIDBadRequest()
	case 401:
		return charts.NewGetChartByIDUnauthorized()
	case 404:
		return charts.NewGetChartByIDNotFound()
	case 500:
		return charts.NewGetChartByIDInternalServerError()
	}
	return charts.NewGetChartByIDInternalServerError()
}
