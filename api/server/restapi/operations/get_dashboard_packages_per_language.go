// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDashboardPackagesPerLanguageHandlerFunc turns a function with the right signature into a get dashboard packages per language handler
type GetDashboardPackagesPerLanguageHandlerFunc func(GetDashboardPackagesPerLanguageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDashboardPackagesPerLanguageHandlerFunc) Handle(params GetDashboardPackagesPerLanguageParams) middleware.Responder {
	return fn(params)
}

// GetDashboardPackagesPerLanguageHandler interface for that can handle valid get dashboard packages per language params
type GetDashboardPackagesPerLanguageHandler interface {
	Handle(GetDashboardPackagesPerLanguageParams) middleware.Responder
}

// NewGetDashboardPackagesPerLanguage creates a new http.Handler for the get dashboard packages per language operation
func NewGetDashboardPackagesPerLanguage(ctx *middleware.Context, handler GetDashboardPackagesPerLanguageHandler) *GetDashboardPackagesPerLanguage {
	return &GetDashboardPackagesPerLanguage{Context: ctx, Handler: handler}
}

/* GetDashboardPackagesPerLanguage swagger:route GET /dashboard/packagesPerLanguage getDashboardPackagesPerLanguage

Get packages count per language

*/
type GetDashboardPackagesPerLanguage struct {
	Context *middleware.Context
	Handler GetDashboardPackagesPerLanguageHandler
}

func (o *GetDashboardPackagesPerLanguage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDashboardPackagesPerLanguageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
