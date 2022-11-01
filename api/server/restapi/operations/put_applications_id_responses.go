// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sambetts-cisco/kubeclarity/api/v2/server/models"
)

// PutApplicationsIDOKCode is the HTTP code returned for type PutApplicationsIDOK
const PutApplicationsIDOKCode int = 200

/*PutApplicationsIDOK Update Application successful.

swagger:response putApplicationsIdOK
*/
type PutApplicationsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ApplicationInfo `json:"body,omitempty"`
}

// NewPutApplicationsIDOK creates PutApplicationsIDOK with default headers values
func NewPutApplicationsIDOK() *PutApplicationsIDOK {

	return &PutApplicationsIDOK{}
}

// WithPayload adds the payload to the put applications Id o k response
func (o *PutApplicationsIDOK) WithPayload(payload *models.ApplicationInfo) *PutApplicationsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put applications Id o k response
func (o *PutApplicationsIDOK) SetPayload(payload *models.ApplicationInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutApplicationsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutApplicationsIDNotFoundCode is the HTTP code returned for type PutApplicationsIDNotFound
const PutApplicationsIDNotFoundCode int = 404

/*PutApplicationsIDNotFound Application not found.

swagger:response putApplicationsIdNotFound
*/
type PutApplicationsIDNotFound struct {
}

// NewPutApplicationsIDNotFound creates PutApplicationsIDNotFound with default headers values
func NewPutApplicationsIDNotFound() *PutApplicationsIDNotFound {

	return &PutApplicationsIDNotFound{}
}

// WriteResponse to the client
func (o *PutApplicationsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*PutApplicationsIDDefault unknown error

swagger:response putApplicationsIdDefault
*/
type PutApplicationsIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPutApplicationsIDDefault creates PutApplicationsIDDefault with default headers values
func NewPutApplicationsIDDefault(code int) *PutApplicationsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutApplicationsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put applications ID default response
func (o *PutApplicationsIDDefault) WithStatusCode(code int) *PutApplicationsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put applications ID default response
func (o *PutApplicationsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put applications ID default response
func (o *PutApplicationsIDDefault) WithPayload(payload *models.APIResponse) *PutApplicationsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put applications ID default response
func (o *PutApplicationsIDDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutApplicationsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
