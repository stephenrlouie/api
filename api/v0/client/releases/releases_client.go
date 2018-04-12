// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new releases API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for releases API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddReleases adds a new release to the clusters
*/
func (a *Client) AddReleases(params *AddReleasesParams) (*AddReleasesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addReleases",
		Method:             "POST",
		PathPattern:        "/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddReleasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddReleasesCreated), nil

}

/*
DeleteRelease deletes release by ID
*/
func (a *Client) DeleteRelease(params *DeleteReleaseParams) (*DeleteReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRelease",
		Method:             "DELETE",
		PathPattern:        "/releases/{releaseId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReleaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteReleaseOK), nil

}

/*
GetReleaseByID finds release by ID

Returns a single release
*/
func (a *Client) GetReleaseByID(params *GetReleaseByIDParams) (*GetReleaseByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleaseByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReleaseById",
		Method:             "GET",
		PathPattern:        "/releases/{releaseId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReleaseByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReleaseByIDOK), nil

}

/*
GetReleases returns all releases

Returns a list of releases
*/
func (a *Client) GetReleases(params *GetReleasesParams) (*GetReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReleases",
		Method:             "GET",
		PathPattern:        "/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReleasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReleasesOK), nil

}

/*
UpdateRelease updates an existing release
*/
func (a *Client) UpdateRelease(params *UpdateReleaseParams) (*UpdateReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRelease",
		Method:             "PUT",
		PathPattern:        "/releases/{releaseId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateReleaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateReleaseOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
