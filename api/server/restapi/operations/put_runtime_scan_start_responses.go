// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sambetts-cisco/kubeclarity/api/v2/server/models"
)

// PutRuntimeScanStartCreatedCode is the HTTP code returned for type PutRuntimeScanStartCreated
const PutRuntimeScanStartCreatedCode int = 201

/*PutRuntimeScanStartCreated Success

swagger:response putRuntimeScanStartCreated
*/
type PutRuntimeScanStartCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPutRuntimeScanStartCreated creates PutRuntimeScanStartCreated with default headers values
func NewPutRuntimeScanStartCreated() *PutRuntimeScanStartCreated {

	return &PutRuntimeScanStartCreated{}
}

// WithPayload adds the payload to the put runtime scan start created response
func (o *PutRuntimeScanStartCreated) WithPayload(payload interface{}) *PutRuntimeScanStartCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put runtime scan start created response
func (o *PutRuntimeScanStartCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRuntimeScanStartCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutRuntimeScanStartBadRequestCode is the HTTP code returned for type PutRuntimeScanStartBadRequest
const PutRuntimeScanStartBadRequestCode int = 400

/*PutRuntimeScanStartBadRequest Scan failed to start

swagger:response putRuntimeScanStartBadRequest
*/
type PutRuntimeScanStartBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutRuntimeScanStartBadRequest creates PutRuntimeScanStartBadRequest with default headers values
func NewPutRuntimeScanStartBadRequest() *PutRuntimeScanStartBadRequest {

	return &PutRuntimeScanStartBadRequest{}
}

// WithPayload adds the payload to the put runtime scan start bad request response
func (o *PutRuntimeScanStartBadRequest) WithPayload(payload string) *PutRuntimeScanStartBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put runtime scan start bad request response
func (o *PutRuntimeScanStartBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRuntimeScanStartBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PutRuntimeScanStartDefault unknown error

swagger:response putRuntimeScanStartDefault
*/
type PutRuntimeScanStartDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPutRuntimeScanStartDefault creates PutRuntimeScanStartDefault with default headers values
func NewPutRuntimeScanStartDefault(code int) *PutRuntimeScanStartDefault {
	if code <= 0 {
		code = 500
	}

	return &PutRuntimeScanStartDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put runtime scan start default response
func (o *PutRuntimeScanStartDefault) WithStatusCode(code int) *PutRuntimeScanStartDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put runtime scan start default response
func (o *PutRuntimeScanStartDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put runtime scan start default response
func (o *PutRuntimeScanStartDefault) WithPayload(payload *models.APIResponse) *PutRuntimeScanStartDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put runtime scan start default response
func (o *PutRuntimeScanStartDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRuntimeScanStartDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
