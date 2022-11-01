// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sambetts-cisco/kubeclarity/shared/v2/pkg/scanner/dependency_track/api/client/models"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/* UpdateProjectOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateProjectOK struct {
	Payload *models.Project
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectOK  %+v", 200, o.Payload)
}
func (o *UpdateProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectUnauthorized creates a UpdateProjectUnauthorized with default headers values
func NewUpdateProjectUnauthorized() *UpdateProjectUnauthorized {
	return &UpdateProjectUnauthorized{}
}

/* UpdateProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateProjectUnauthorized struct {
}

func (o *UpdateProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectUnauthorized ", 401)
}

func (o *UpdateProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/* UpdateProjectNotFound describes a response with status code 404, with default header values.

The UUID of the project could not be found
*/
type UpdateProjectNotFound struct {
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectNotFound ", 404)
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectConflict creates a UpdateProjectConflict with default headers values
func NewUpdateProjectConflict() *UpdateProjectConflict {
	return &UpdateProjectConflict{}
}

/* UpdateProjectConflict describes a response with status code 409, with default header values.

A project with the specified name already exists
*/
type UpdateProjectConflict struct {
}

func (o *UpdateProjectConflict) Error() string {
	return fmt.Sprintf("[POST /v1/project][%d] updateProjectConflict ", 409)
}

func (o *UpdateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
