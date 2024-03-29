// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CISDockerBenchmarkResultsEX c i s docker benchmark results e x
//
// swagger:model CISDockerBenchmarkResultsEX
type CISDockerBenchmarkResultsEX struct {

	// code
	Code string `json:"code,omitempty"`

	// desc
	Desc string `json:"desc,omitempty"`

	// level
	Level CISDockerBenchmarkLevel `json:"level,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this c i s docker benchmark results e x
func (m *CISDockerBenchmarkResultsEX) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CISDockerBenchmarkResultsEX) validateLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.Level) { // not required
		return nil
	}

	if err := m.Level.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("level")
		}
		return err
	}

	return nil
}

// ContextValidate validate this c i s docker benchmark results e x based on the context it is used
func (m *CISDockerBenchmarkResultsEX) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CISDockerBenchmarkResultsEX) contextValidateLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Level.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("level")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CISDockerBenchmarkResultsEX) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CISDockerBenchmarkResultsEX) UnmarshalBinary(b []byte) error {
	var res CISDockerBenchmarkResultsEX
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
