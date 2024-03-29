// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetRuntimeQuickscanConfigParams creates a new GetRuntimeQuickscanConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuntimeQuickscanConfigParams() *GetRuntimeQuickscanConfigParams {
	return &GetRuntimeQuickscanConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeQuickscanConfigParamsWithTimeout creates a new GetRuntimeQuickscanConfigParams object
// with the ability to set a timeout on a request.
func NewGetRuntimeQuickscanConfigParamsWithTimeout(timeout time.Duration) *GetRuntimeQuickscanConfigParams {
	return &GetRuntimeQuickscanConfigParams{
		timeout: timeout,
	}
}

// NewGetRuntimeQuickscanConfigParamsWithContext creates a new GetRuntimeQuickscanConfigParams object
// with the ability to set a context for a request.
func NewGetRuntimeQuickscanConfigParamsWithContext(ctx context.Context) *GetRuntimeQuickscanConfigParams {
	return &GetRuntimeQuickscanConfigParams{
		Context: ctx,
	}
}

// NewGetRuntimeQuickscanConfigParamsWithHTTPClient creates a new GetRuntimeQuickscanConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuntimeQuickscanConfigParamsWithHTTPClient(client *http.Client) *GetRuntimeQuickscanConfigParams {
	return &GetRuntimeQuickscanConfigParams{
		HTTPClient: client,
	}
}

/* GetRuntimeQuickscanConfigParams contains all the parameters to send to the API endpoint
   for the get runtime quickscan config operation.

   Typically these are written to a http.Request.
*/
type GetRuntimeQuickscanConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runtime quickscan config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeQuickscanConfigParams) WithDefaults() *GetRuntimeQuickscanConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runtime quickscan config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeQuickscanConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runtime quickscan config params
func (o *GetRuntimeQuickscanConfigParams) WithTimeout(timeout time.Duration) *GetRuntimeQuickscanConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime quickscan config params
func (o *GetRuntimeQuickscanConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime quickscan config params
func (o *GetRuntimeQuickscanConfigParams) WithContext(ctx context.Context) *GetRuntimeQuickscanConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime quickscan config params
func (o *GetRuntimeQuickscanConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime quickscan config params
func (o *GetRuntimeQuickscanConfigParams) WithHTTPClient(client *http.Client) *GetRuntimeQuickscanConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime quickscan config params
func (o *GetRuntimeQuickscanConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeQuickscanConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
