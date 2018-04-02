// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateChartParams creates a new UpdateChartParams object
// with the default values initialized.
func NewUpdateChartParams() UpdateChartParams {
	var ()
	return UpdateChartParams{}
}

// UpdateChartParams contains all the bound params for the update chart operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateChart
type UpdateChartParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of chart to return
	  Required: true
	  In: path
	*/
	ChartID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateChartParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rChartID, rhkChartID, _ := route.Params.GetOK("chartId")
	if err := o.bindChartID(rChartID, rhkChartID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateChartParams) bindChartID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ChartID = raw

	return nil
}