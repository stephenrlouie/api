// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewGetReleasesParams creates a new GetReleasesParams object
// with the default values initialized.
func NewGetReleasesParams() *GetReleasesParams {

	return &GetReleasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleasesParamsWithTimeout creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleasesParamsWithTimeout(timeout time.Duration) *GetReleasesParams {

	return &GetReleasesParams{

		timeout: timeout,
	}
}

// NewGetReleasesParamsWithContext creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleasesParamsWithContext(ctx context.Context) *GetReleasesParams {

	return &GetReleasesParams{

		Context: ctx,
	}
}

// NewGetReleasesParamsWithHTTPClient creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleasesParamsWithHTTPClient(client *http.Client) *GetReleasesParams {

	return &GetReleasesParams{
		HTTPClient: client,
	}
}

/*GetReleasesParams contains all the parameters to send to the API endpoint
for the get releases operation typically these are written to a http.Request
*/
type GetReleasesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) WithTimeout(timeout time.Duration) *GetReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get releases params
func (o *GetReleasesParams) WithContext(ctx context.Context) *GetReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get releases params
func (o *GetReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) WithHTTPClient(client *http.Client) *GetReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}