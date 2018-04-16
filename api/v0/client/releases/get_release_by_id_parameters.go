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

// NewGetReleaseByIDParams creates a new GetReleaseByIDParams object
// with the default values initialized.
func NewGetReleaseByIDParams() *GetReleaseByIDParams {
	var ()
	return &GetReleaseByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleaseByIDParamsWithTimeout creates a new GetReleaseByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleaseByIDParamsWithTimeout(timeout time.Duration) *GetReleaseByIDParams {
	var ()
	return &GetReleaseByIDParams{

		timeout: timeout,
	}
}

// NewGetReleaseByIDParamsWithContext creates a new GetReleaseByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleaseByIDParamsWithContext(ctx context.Context) *GetReleaseByIDParams {
	var ()
	return &GetReleaseByIDParams{

		Context: ctx,
	}
}

// NewGetReleaseByIDParamsWithHTTPClient creates a new GetReleaseByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleaseByIDParamsWithHTTPClient(client *http.Client) *GetReleaseByIDParams {
	var ()
	return &GetReleaseByIDParams{
		HTTPClient: client,
	}
}

/*GetReleaseByIDParams contains all the parameters to send to the API endpoint
for the get release by Id operation typically these are written to a http.Request
*/
type GetReleaseByIDParams struct {

	/*Labels
	  The node labels to identify applicable clusters

	*/
	Labels *string
	/*ReleaseID
	  ID of release to return

	*/
	ReleaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get release by Id params
func (o *GetReleaseByIDParams) WithTimeout(timeout time.Duration) *GetReleaseByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get release by Id params
func (o *GetReleaseByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get release by Id params
func (o *GetReleaseByIDParams) WithContext(ctx context.Context) *GetReleaseByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get release by Id params
func (o *GetReleaseByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get release by Id params
func (o *GetReleaseByIDParams) WithHTTPClient(client *http.Client) *GetReleaseByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get release by Id params
func (o *GetReleaseByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabels adds the labels to the get release by Id params
func (o *GetReleaseByIDParams) WithLabels(labels *string) *GetReleaseByIDParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the get release by Id params
func (o *GetReleaseByIDParams) SetLabels(labels *string) {
	o.Labels = labels
}

// WithReleaseID adds the releaseID to the get release by Id params
func (o *GetReleaseByIDParams) WithReleaseID(releaseID string) *GetReleaseByIDParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get release by Id params
func (o *GetReleaseByIDParams) SetReleaseID(releaseID string) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleaseByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Labels != nil {

		// query param labels
		var qrLabels string
		if o.Labels != nil {
			qrLabels = *o.Labels
		}
		qLabels := qrLabels
		if qLabels != "" {
			if err := r.SetQueryParam("labels", qLabels); err != nil {
				return err
			}
		}

	}

	// path param releaseId
	if err := r.SetPathParam("releaseId", o.ReleaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
