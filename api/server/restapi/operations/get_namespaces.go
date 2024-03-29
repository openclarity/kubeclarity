// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNamespacesHandlerFunc turns a function with the right signature into a get namespaces handler
type GetNamespacesHandlerFunc func(GetNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNamespacesHandlerFunc) Handle(params GetNamespacesParams) middleware.Responder {
	return fn(params)
}

// GetNamespacesHandler interface for that can handle valid get namespaces params
type GetNamespacesHandler interface {
	Handle(GetNamespacesParams) middleware.Responder
}

// NewGetNamespaces creates a new http.Handler for the get namespaces operation
func NewGetNamespaces(ctx *middleware.Context, handler GetNamespacesHandler) *GetNamespaces {
	return &GetNamespaces{Context: ctx, Handler: handler}
}

/* GetNamespaces swagger:route GET /namespaces getNamespaces

Get list of namespaces in kubernetes cluster

*/
type GetNamespaces struct {
	Context *middleware.Context
	Handler GetNamespacesHandler
}

func (o *GetNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetNamespacesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
