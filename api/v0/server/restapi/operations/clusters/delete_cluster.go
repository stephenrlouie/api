// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteClusterHandlerFunc turns a function with the right signature into a delete cluster handler
type DeleteClusterHandlerFunc func(DeleteClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteClusterHandlerFunc) Handle(params DeleteClusterParams) middleware.Responder {
	return fn(params)
}

// DeleteClusterHandler interface for that can handle valid delete cluster params
type DeleteClusterHandler interface {
	Handle(DeleteClusterParams) middleware.Responder
}

// NewDeleteCluster creates a new http.Handler for the delete cluster operation
func NewDeleteCluster(ctx *middleware.Context, handler DeleteClusterHandler) *DeleteCluster {
	return &DeleteCluster{Context: ctx, Handler: handler}
}

/*DeleteCluster swagger:route DELETE /clusters/{clusterId} clusters deleteCluster

Deletes a cluster

*/
type DeleteCluster struct {
	Context *middleware.Context
	Handler DeleteClusterHandler
}

func (o *DeleteCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
