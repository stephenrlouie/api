// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddReleasesHandlerFunc turns a function with the right signature into a add releases handler
type AddReleasesHandlerFunc func(AddReleasesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddReleasesHandlerFunc) Handle(params AddReleasesParams) middleware.Responder {
	return fn(params)
}

// AddReleasesHandler interface for that can handle valid add releases params
type AddReleasesHandler interface {
	Handle(AddReleasesParams) middleware.Responder
}

// NewAddReleases creates a new http.Handler for the add releases operation
func NewAddReleases(ctx *middleware.Context, handler AddReleasesHandler) *AddReleases {
	return &AddReleases{Context: ctx, Handler: handler}
}

/*AddReleases swagger:route POST /releases releases addReleases

Add a new release to the clusters

*/
type AddReleases struct {
	Context *middleware.Context
	Handler AddReleasesHandler
}

func (o *AddReleases) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddReleasesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
