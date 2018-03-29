package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/charts"
)

func NewAddChart() *addChart {
	return &addChart{}
}

type addChart struct{}

func (d *addChart) Handle(params charts.AddChartsParams) middleware.Responder {
	fmt.Printf("AddCharts\n")
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return charts.NewAddChartsCreated()
}

func (d *addChart) MockHandle(params charts.AddChartsParams) middleware.Responder {
	statusCode, err := mock.GetMock(path.Join(restapi.MockBasePath, "add-charts.json"), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return charts.NewAddChartsInternalServerError()
	}

	switch statusCode {
	case 201:
		return charts.NewAddChartsCreated()
	case 400:
		return charts.NewAddChartsBadRequest()
	case 401:
		return charts.NewAddChartsUnauthorized()
	case 500:
		return charts.NewAddChartsInternalServerError()
	}
	return charts.NewAddChartsInternalServerError()
}
