// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostScanScanUUIDContentAnalysisHandlerFunc turns a function with the right signature into a post scan scan UUID content analysis handler
type PostScanScanUUIDContentAnalysisHandlerFunc func(PostScanScanUUIDContentAnalysisParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostScanScanUUIDContentAnalysisHandlerFunc) Handle(params PostScanScanUUIDContentAnalysisParams) middleware.Responder {
	return fn(params)
}

// PostScanScanUUIDContentAnalysisHandler interface for that can handle valid post scan scan UUID content analysis params
type PostScanScanUUIDContentAnalysisHandler interface {
	Handle(PostScanScanUUIDContentAnalysisParams) middleware.Responder
}

// NewPostScanScanUUIDContentAnalysis creates a new http.Handler for the post scan scan UUID content analysis operation
func NewPostScanScanUUIDContentAnalysis(ctx *middleware.Context, handler PostScanScanUUIDContentAnalysisHandler) *PostScanScanUUIDContentAnalysis {
	return &PostScanScanUUIDContentAnalysis{Context: ctx, Handler: handler}
}

/* PostScanScanUUIDContentAnalysis swagger:route POST /scan/{scan-uuid}/contentAnalysis postScanScanUuidContentAnalysis

Report a content analysis for a specific resource

*/
type PostScanScanUUIDContentAnalysis struct {
	Context *middleware.Context
	Handler PostScanScanUUIDContentAnalysisHandler
}

func (o *PostScanScanUUIDContentAnalysis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostScanScanUUIDContentAnalysisParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
