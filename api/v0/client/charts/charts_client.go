// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new charts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for charts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddCharts adds a new cluster to optikon
*/
func (a *Client) AddCharts(params *AddChartsParams) (*AddChartsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddChartsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCharts",
		Method:             "POST",
		PathPattern:        "/charts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddChartsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddChartsCreated), nil

}

/*
DeleteChart deletes chart by ID
*/
func (a *Client) DeleteChart(params *DeleteChartParams) (*DeleteChartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteChartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteChart",
		Method:             "DELETE",
		PathPattern:        "/charts/{chartId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteChartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteChartOK), nil

}

/*
GetChartByID finds chart by ID

Returns a single chart
*/
func (a *Client) GetChartByID(params *GetChartByIDParams) (*GetChartByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChartById",
		Method:             "GET",
		PathPattern:        "/charts/{chartId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartByIDOK), nil

}

/*
GetCharts returns all charts

Returns a list of charts
*/
func (a *Client) GetCharts(params *GetChartsParams) (*GetChartsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCharts",
		Method:             "GET",
		PathPattern:        "/charts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartsOK), nil

}

/*
UpdateChart updates an existing chart
*/
func (a *Client) UpdateChart(params *UpdateChartParams) (*UpdateChartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateChartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateChart",
		Method:             "PUT",
		PathPattern:        "/charts/{chartId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateChartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateChartOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
