package handlers

import (
	"fmt"
	"path"

	"github.com/go-openapi/runtime/middleware"
	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi"
	"wwwin-github.cisco.com/edge/optikon/api/v0/server/restapi/operations/charts"
)

func NewUpdateChart() *updateChart {
	return &updateChart{}
}

type updateChart struct{}

func (d *updateChart) Handle(params charts.UpdateChartParams) middleware.Responder {
	fmt.Printf("updateChart: %s\n", params.ChartID)
	if restapi.MockBasePath != "" {
		return d.MockHandle(params)
	}
	return charts.NewUpdateChartOK()
}

func (d *updateChart) MockHandle(params charts.UpdateChartParams) middleware.Responder {
	statusCode, err := mock.GetMock(
		path.Join(restapi.MockBasePath, fmt.Sprintf("update-chart-%s.json", params.ChartID)), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return charts.NewUpdateChartInternalServerError()
	}

	switch statusCode {
	case 200:
		return charts.NewUpdateChartOK()
	case 400:
		return charts.NewUpdateChartBadRequest()
	case 401:
		return charts.NewUpdateChartUnauthorized()
	case 404:
		return charts.NewUpdateChartNotFound()
	case 500:
		return charts.NewUpdateChartInternalServerError()
	}
	return charts.NewUpdateChartInternalServerError()
}
