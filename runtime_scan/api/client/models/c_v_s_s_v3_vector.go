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

// CVSSV3Vector c v s s v3 vector
//
// swagger:model CVSSV3Vector
type CVSSV3Vector struct {

	// attack complexity
	AttackComplexity AttackComplexity `json:"attackComplexity,omitempty"`

	// attack vector
	AttackVector AttackVector `json:"attackVector,omitempty"`

	// availability
	Availability Availability `json:"availability,omitempty"`

	// confidentiality
	Confidentiality Confidentiality `json:"confidentiality,omitempty"`

	// integrity
	Integrity Integrity `json:"integrity,omitempty"`

	// privileges required
	PrivilegesRequired PrivilegesRequired `json:"privilegesRequired,omitempty"`

	// scope
	Scope Scope `json:"scope,omitempty"`

	// user interaction
	UserInteraction UserInteraction `json:"userInteraction,omitempty"`

	// vector
	Vector string `json:"vector,omitempty"`
}

// Validate validates this c v s s v3 vector
func (m *CVSSV3Vector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttackComplexity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttackVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfidentiality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivilegesRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInteraction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CVSSV3Vector) validateAttackComplexity(formats strfmt.Registry) error {
	if swag.IsZero(m.AttackComplexity) { // not required
		return nil
	}

	if err := m.AttackComplexity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attackComplexity")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validateAttackVector(formats strfmt.Registry) error {
	if swag.IsZero(m.AttackVector) { // not required
		return nil
	}

	if err := m.AttackVector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attackVector")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validateAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	if err := m.Availability.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("availability")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validateConfidentiality(formats strfmt.Registry) error {
	if swag.IsZero(m.Confidentiality) { // not required
		return nil
	}

	if err := m.Confidentiality.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("confidentiality")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validateIntegrity(formats strfmt.Registry) error {
	if swag.IsZero(m.Integrity) { // not required
		return nil
	}

	if err := m.Integrity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("integrity")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validatePrivilegesRequired(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivilegesRequired) { // not required
		return nil
	}

	if err := m.PrivilegesRequired.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("privilegesRequired")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) validateUserInteraction(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInteraction) { // not required
		return nil
	}

	if err := m.UserInteraction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userInteraction")
		}
		return err
	}

	return nil
}

// ContextValidate validate this c v s s v3 vector based on the context it is used
func (m *CVSSV3Vector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttackComplexity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttackVector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfidentiality(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntegrity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivilegesRequired(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserInteraction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CVSSV3Vector) contextValidateAttackComplexity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AttackComplexity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attackComplexity")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidateAttackVector(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AttackVector.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attackVector")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidateAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Availability.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("availability")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidateConfidentiality(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Confidentiality.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("confidentiality")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidateIntegrity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Integrity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("integrity")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidatePrivilegesRequired(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrivilegesRequired.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("privilegesRequired")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Scope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *CVSSV3Vector) contextValidateUserInteraction(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UserInteraction.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userInteraction")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CVSSV3Vector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CVSSV3Vector) UnmarshalBinary(b []byte) error {
	var res CVSSV3Vector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
