// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddClusterHandlerFunc turns a function with the right signature into a add cluster handler
type AddClusterHandlerFunc func(AddClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddClusterHandlerFunc) Handle(params AddClusterParams) middleware.Responder {
	return fn(params)
}

// AddClusterHandler interface for that can handle valid add cluster params
type AddClusterHandler interface {
	Handle(AddClusterParams) middleware.Responder
}

// NewAddCluster creates a new http.Handler for the add cluster operation
func NewAddCluster(ctx *middleware.Context, handler AddClusterHandler) *AddCluster {
	return &AddCluster{Context: ctx, Handler: handler}
}

/*AddCluster swagger:route POST /clusters clusters addCluster

Add a new cluster to optikon

*/
type AddCluster struct {
	Context *middleware.Context
	Handler AddClusterHandler
}

func (o *AddCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}