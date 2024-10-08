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

// NewGetDashboardPackagesPerLicenseParams creates a new GetDashboardPackagesPerLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardPackagesPerLicenseParams() *GetDashboardPackagesPerLicenseParams {
	return &GetDashboardPackagesPerLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardPackagesPerLicenseParamsWithTimeout creates a new GetDashboardPackagesPerLicenseParams object
// with the ability to set a timeout on a request.
func NewGetDashboardPackagesPerLicenseParamsWithTimeout(timeout time.Duration) *GetDashboardPackagesPerLicenseParams {
	return &GetDashboardPackagesPerLicenseParams{
		timeout: timeout,
	}
}

// NewGetDashboardPackagesPerLicenseParamsWithContext creates a new GetDashboardPackagesPerLicenseParams object
// with the ability to set a context for a request.
func NewGetDashboardPackagesPerLicenseParamsWithContext(ctx context.Context) *GetDashboardPackagesPerLicenseParams {
	return &GetDashboardPackagesPerLicenseParams{
		Context: ctx,
	}
}

// NewGetDashboardPackagesPerLicenseParamsWithHTTPClient creates a new GetDashboardPackagesPerLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardPackagesPerLicenseParamsWithHTTPClient(client *http.Client) *GetDashboardPackagesPerLicenseParams {
	return &GetDashboardPackagesPerLicenseParams{
		HTTPClient: client,
	}
}

/* GetDashboardPackagesPerLicenseParams contains all the parameters to send to the API endpoint
   for the get dashboard packages per license operation.

   Typically these are written to a http.Request.
*/
type GetDashboardPackagesPerLicenseParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard packages per license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardPackagesPerLicenseParams) WithDefaults() *GetDashboardPackagesPerLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard packages per license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardPackagesPerLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard packages per license params
func (o *GetDashboardPackagesPerLicenseParams) WithTimeout(timeout time.Duration) *GetDashboardPackagesPerLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard packages per license params
func (o *GetDashboardPackagesPerLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard packages per license params
func (o *GetDashboardPackagesPerLicenseParams) WithContext(ctx context.Context) *GetDashboardPackagesPerLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard packages per license params
func (o *GetDashboardPackagesPerLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard packages per license params
func (o *GetDashboardPackagesPerLicenseParams) WithHTTPClient(client *http.Client) *GetDashboardPackagesPerLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard packages per license params
func (o *GetDashboardPackagesPerLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardPackagesPerLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
