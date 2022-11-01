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

	"github.com/sambetts-cisco/kubeclarity/runtime_scan/api/v2/client/models"
)

// NewPostScanScanUUIDResultsParams creates a new PostScanScanUUIDResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScanScanUUIDResultsParams() *PostScanScanUUIDResultsParams {
	return &PostScanScanUUIDResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScanScanUUIDResultsParamsWithTimeout creates a new PostScanScanUUIDResultsParams object
// with the ability to set a timeout on a request.
func NewPostScanScanUUIDResultsParamsWithTimeout(timeout time.Duration) *PostScanScanUUIDResultsParams {
	return &PostScanScanUUIDResultsParams{
		timeout: timeout,
	}
}

// NewPostScanScanUUIDResultsParamsWithContext creates a new PostScanScanUUIDResultsParams object
// with the ability to set a context for a request.
func NewPostScanScanUUIDResultsParamsWithContext(ctx context.Context) *PostScanScanUUIDResultsParams {
	return &PostScanScanUUIDResultsParams{
		Context: ctx,
	}
}

// NewPostScanScanUUIDResultsParamsWithHTTPClient creates a new PostScanScanUUIDResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScanScanUUIDResultsParamsWithHTTPClient(client *http.Client) *PostScanScanUUIDResultsParams {
	return &PostScanScanUUIDResultsParams{
		HTTPClient: client,
	}
}

/* PostScanScanUUIDResultsParams contains all the parameters to send to the API endpoint
   for the post scan scan UUID results operation.

   Typically these are written to a http.Request.
*/
type PostScanScanUUIDResultsParams struct {

	// Body.
	Body *models.ImageVulnerabilityScan

	// ScanUUID.
	//
	// Format: uuid
	ScanUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post scan scan UUID results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScanScanUUIDResultsParams) WithDefaults() *PostScanScanUUIDResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post scan scan UUID results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScanScanUUIDResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) WithTimeout(timeout time.Duration) *PostScanScanUUIDResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) WithContext(ctx context.Context) *PostScanScanUUIDResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) WithHTTPClient(client *http.Client) *PostScanScanUUIDResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) WithBody(body *models.ImageVulnerabilityScan) *PostScanScanUUIDResultsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) SetBody(body *models.ImageVulnerabilityScan) {
	o.Body = body
}

// WithScanUUID adds the scanUUID to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) WithScanUUID(scanUUID strfmt.UUID) *PostScanScanUUIDResultsParams {
	o.SetScanUUID(scanUUID)
	return o
}

// SetScanUUID adds the scanUuid to the post scan scan UUID results params
func (o *PostScanScanUUIDResultsParams) SetScanUUID(scanUUID strfmt.UUID) {
	o.ScanUUID = scanUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PostScanScanUUIDResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scan-uuid
	if err := r.SetPathParam("scan-uuid", o.ScanUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
