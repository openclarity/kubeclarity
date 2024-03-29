// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/kubeclarity/api/client/models"
)

// DeleteApplicationResourcesIDReader is a Reader for the DeleteApplicationResourcesID structure.
type DeleteApplicationResourcesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationResourcesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteApplicationResourcesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteApplicationResourcesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteApplicationResourcesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteApplicationResourcesIDNoContent creates a DeleteApplicationResourcesIDNoContent with default headers values
func NewDeleteApplicationResourcesIDNoContent() *DeleteApplicationResourcesIDNoContent {
	return &DeleteApplicationResourcesIDNoContent{}
}

/* DeleteApplicationResourcesIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteApplicationResourcesIDNoContent struct {
}

func (o *DeleteApplicationResourcesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /applicationResources/{id}][%d] deleteApplicationResourcesIdNoContent ", 204)
}

func (o *DeleteApplicationResourcesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationResourcesIDNotFound creates a DeleteApplicationResourcesIDNotFound with default headers values
func NewDeleteApplicationResourcesIDNotFound() *DeleteApplicationResourcesIDNotFound {
	return &DeleteApplicationResourcesIDNotFound{}
}

/* DeleteApplicationResourcesIDNotFound describes a response with status code 404, with default header values.

Application resource not found.
*/
type DeleteApplicationResourcesIDNotFound struct {
}

func (o *DeleteApplicationResourcesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /applicationResources/{id}][%d] deleteApplicationResourcesIdNotFound ", 404)
}

func (o *DeleteApplicationResourcesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationResourcesIDDefault creates a DeleteApplicationResourcesIDDefault with default headers values
func NewDeleteApplicationResourcesIDDefault(code int) *DeleteApplicationResourcesIDDefault {
	return &DeleteApplicationResourcesIDDefault{
		_statusCode: code,
	}
}

/* DeleteApplicationResourcesIDDefault describes a response with status code -1, with default header values.

unknown error
*/
type DeleteApplicationResourcesIDDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the delete application resources ID default response
func (o *DeleteApplicationResourcesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteApplicationResourcesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /applicationResources/{id}][%d] DeleteApplicationResourcesID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteApplicationResourcesIDDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *DeleteApplicationResourcesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
