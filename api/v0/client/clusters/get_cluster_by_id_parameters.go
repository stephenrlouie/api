// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterByIDParams creates a new GetClusterByIDParams object
// with the default values initialized.
func NewGetClusterByIDParams() *GetClusterByIDParams {
	var ()
	return &GetClusterByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterByIDParamsWithTimeout creates a new GetClusterByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterByIDParamsWithTimeout(timeout time.Duration) *GetClusterByIDParams {
	var ()
	return &GetClusterByIDParams{

		timeout: timeout,
	}
}

// NewGetClusterByIDParamsWithContext creates a new GetClusterByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterByIDParamsWithContext(ctx context.Context) *GetClusterByIDParams {
	var ()
	return &GetClusterByIDParams{

		Context: ctx,
	}
}

// NewGetClusterByIDParamsWithHTTPClient creates a new GetClusterByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterByIDParamsWithHTTPClient(client *http.Client) *GetClusterByIDParams {
	var ()
	return &GetClusterByIDParams{
		HTTPClient: client,
	}
}

/*GetClusterByIDParams contains all the parameters to send to the API endpoint
for the get cluster by Id operation typically these are written to a http.Request
*/
type GetClusterByIDParams struct {

	/*ClusterID
	  ID of cluster to return

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster by Id params
func (o *GetClusterByIDParams) WithTimeout(timeout time.Duration) *GetClusterByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster by Id params
func (o *GetClusterByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster by Id params
func (o *GetClusterByIDParams) WithContext(ctx context.Context) *GetClusterByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster by Id params
func (o *GetClusterByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster by Id params
func (o *GetClusterByIDParams) WithHTTPClient(client *http.Client) *GetClusterByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster by Id params
func (o *GetClusterByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster by Id params
func (o *GetClusterByIDParams) WithClusterID(clusterID string) *GetClusterByIDParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster by Id params
func (o *GetClusterByIDParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
