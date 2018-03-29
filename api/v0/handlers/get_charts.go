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

func NewGetCharts() *getCharts {
	return &getCharts{}
}

type getCharts struct{}

func (d *getCharts) Handle(params charts.GetChartsParams) middleware.Responder {
	fmt.Printf("GetCharts\n")
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	payload := models.GetChartsOKBody{}
	return charts.NewGetChartsOK().WithPayload(payload)
}

func (d *getCharts) MockHandle(params charts.GetChartsParams) middleware.Responder {
	payload := models.GetChartsOKBody{}
	statusCode, err := mock.GetMock(path.Join(restapi.MockBasePath, "get-charts.json"), &payload)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return charts.NewGetChartsInternalServerError()
	}

	switch statusCode {
	case 200:
		return charts.NewGetChartsOK().WithPayload(payload)
	case 400:
		return charts.NewGetChartsBadRequest()
	case 401:
		return charts.NewGetChartsUnauthorized()
	case 500:
		return charts.NewGetChartsInternalServerError()
	}
	return charts.NewGetChartsInternalServerError()
}
