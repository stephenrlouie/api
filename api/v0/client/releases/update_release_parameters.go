// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateReleaseParams creates a new UpdateReleaseParams object
// with the default values initialized.
func NewUpdateReleaseParams() *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReleaseParamsWithTimeout creates a new UpdateReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReleaseParamsWithTimeout(timeout time.Duration) *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{

		timeout: timeout,
	}
}

// NewUpdateReleaseParamsWithContext creates a new UpdateReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReleaseParamsWithContext(ctx context.Context) *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{

		Context: ctx,
	}
}

// NewUpdateReleaseParamsWithHTTPClient creates a new UpdateReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReleaseParamsWithHTTPClient(client *http.Client) *UpdateReleaseParams {
	var ()
	return &UpdateReleaseParams{
		HTTPClient: client,
	}
}

/*UpdateReleaseParams contains all the parameters to send to the API endpoint
for the update release operation typically these are written to a http.Request
*/
type UpdateReleaseParams struct {

	/*ChartTar
	  The file to upload

	*/
	ChartTar os.File
	/*Labels
	  The node labels to identify applicable clusters

	*/
	Labels *string
	/*ReleaseID
	  ID of release to update

	*/
	ReleaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update release params
func (o *UpdateReleaseParams) WithTimeout(timeout time.Duration) *UpdateReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update release params
func (o *UpdateReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update release params
func (o *UpdateReleaseParams) WithContext(ctx context.Context) *UpdateReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update release params
func (o *UpdateReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update release params
func (o *UpdateReleaseParams) WithHTTPClient(client *http.Client) *UpdateReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update release params
func (o *UpdateReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChartTar adds the chartTar to the update release params
func (o *UpdateReleaseParams) WithChartTar(chartTar os.File) *UpdateReleaseParams {
	o.SetChartTar(chartTar)
	return o
}

// SetChartTar adds the chartTar to the update release params
func (o *UpdateReleaseParams) SetChartTar(chartTar os.File) {
	o.ChartTar = chartTar
}

// WithLabels adds the labels to the update release params
func (o *UpdateReleaseParams) WithLabels(labels *string) *UpdateReleaseParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the update release params
func (o *UpdateReleaseParams) SetLabels(labels *string) {
	o.Labels = labels
}

// WithReleaseID adds the releaseID to the update release params
func (o *UpdateReleaseParams) WithReleaseID(releaseID string) *UpdateReleaseParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the update release params
func (o *UpdateReleaseParams) SetReleaseID(releaseID string) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param chartTar
	if err := r.SetFileParam("chartTar", &o.ChartTar); err != nil {
		return err
	}

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
