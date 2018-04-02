// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterByIDParams creates a new GetClusterByIDParams object
// with the default values initialized.
func NewGetClusterByIDParams() GetClusterByIDParams {
	var ()
	return GetClusterByIDParams{}
}

// GetClusterByIDParams contains all the bound params for the get cluster by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getClusterById
type GetClusterByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of cluster to return
	  Required: true
	  In: path
	*/
	ClusterID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetClusterByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rClusterID, rhkClusterID, _ := route.Params.GetOK("clusterId")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetClusterByIDParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ClusterID = raw

	return nil
}