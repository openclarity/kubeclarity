// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sambetts-cisco/kubeclarity/api/v2/server/models"
)

// GetDashboardMostVulnerableOKCode is the HTTP code returned for type GetDashboardMostVulnerableOK
const GetDashboardMostVulnerableOKCode int = 200

/*GetDashboardMostVulnerableOK Success

swagger:response getDashboardMostVulnerableOK
*/
type GetDashboardMostVulnerableOK struct {

	/*
	  In: Body
	*/
	Payload *models.MostVulnerable `json:"body,omitempty"`
}

// NewGetDashboardMostVulnerableOK creates GetDashboardMostVulnerableOK with default headers values
func NewGetDashboardMostVulnerableOK() *GetDashboardMostVulnerableOK {

	return &GetDashboardMostVulnerableOK{}
}

// WithPayload adds the payload to the get dashboard most vulnerable o k response
func (o *GetDashboardMostVulnerableOK) WithPayload(payload *models.MostVulnerable) *GetDashboardMostVulnerableOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dashboard most vulnerable o k response
func (o *GetDashboardMostVulnerableOK) SetPayload(payload *models.MostVulnerable) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDashboardMostVulnerableOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDashboardMostVulnerableDefault unknown error

swagger:response getDashboardMostVulnerableDefault
*/
type GetDashboardMostVulnerableDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetDashboardMostVulnerableDefault creates GetDashboardMostVulnerableDefault with default headers values
func NewGetDashboardMostVulnerableDefault(code int) *GetDashboardMostVulnerableDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDashboardMostVulnerableDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dashboard most vulnerable default response
func (o *GetDashboardMostVulnerableDefault) WithStatusCode(code int) *GetDashboardMostVulnerableDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dashboard most vulnerable default response
func (o *GetDashboardMostVulnerableDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get dashboard most vulnerable default response
func (o *GetDashboardMostVulnerableDefault) WithPayload(payload *models.APIResponse) *GetDashboardMostVulnerableDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dashboard most vulnerable default response
func (o *GetDashboardMostVulnerableDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDashboardMostVulnerableDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
