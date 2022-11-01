// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/sambetts-cisco/kubeclarity/api/v2/client/models"
)

// GetApplicationsReader is a Reader for the GetApplications structure.
type GetApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetApplicationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationsOK creates a GetApplicationsOK with default headers values
func NewGetApplicationsOK() *GetApplicationsOK {
	return &GetApplicationsOK{}
}

/* GetApplicationsOK describes a response with status code 200, with default header values.

Success
*/
type GetApplicationsOK struct {
	Payload *GetApplicationsOKBody
}

func (o *GetApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /applications][%d] getApplicationsOK  %+v", 200, o.Payload)
}
func (o *GetApplicationsOK) GetPayload() *GetApplicationsOKBody {
	return o.Payload
}

func (o *GetApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetApplicationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsDefault creates a GetApplicationsDefault with default headers values
func NewGetApplicationsDefault(code int) *GetApplicationsDefault {
	return &GetApplicationsDefault{
		_statusCode: code,
	}
}

/* GetApplicationsDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetApplicationsDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get applications default response
func (o *GetApplicationsDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationsDefault) Error() string {
	return fmt.Sprintf("[GET /applications][%d] GetApplications default  %+v", o._statusCode, o.Payload)
}
func (o *GetApplicationsDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetApplicationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetApplicationsOKBody get applications o k body
swagger:model GetApplicationsOKBody
*/
type GetApplicationsOKBody struct {

	// List of applications in the given filters and page. List length must be lower or equal to pageSize
	Items []*models.Application `json:"items"`

	// Total applications count under the given filters
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this get applications o k body
func (o *GetApplicationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetApplicationsOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getApplicationsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetApplicationsOKBody) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("getApplicationsOK"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get applications o k body based on the context it is used
func (o *GetApplicationsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetApplicationsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getApplicationsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetApplicationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetApplicationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetApplicationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
