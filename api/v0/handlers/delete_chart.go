package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/charts"
)

func NewDeleteChart() *deleteChart {
	return &deleteChart{}
}

type deleteChart struct{}

func (d *deleteChart) Handle(params charts.DeleteChartParams) middleware.Responder {
	fmt.Printf("deleteChart: %s\n", params.ChartID)
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return charts.NewDeleteChartOK()
}

func (d *deleteChart) MockHandle(params charts.DeleteChartParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("delete-chart-%s.json", params.ChartID)), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return charts.NewDeleteChartInternalServerError()
	}

	switch statusCode {
	case 200:
		return charts.NewDeleteChartOK()
	case 400:
		return charts.NewDeleteChartBadRequest()
	case 401:
		return charts.NewDeleteChartUnauthorized()
	case 404:
		return charts.NewDeleteChartNotFound()
	case 500:
		return charts.NewDeleteChartInternalServerError()
	}
	return charts.NewDeleteChartInternalServerError()
}
