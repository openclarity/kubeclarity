// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackagesCountPerLanguage packages count per language
//
// swagger:model PackagesCountPerLanguage
type PackagesCountPerLanguage struct {

	// count
	Count uint32 `json:"count,omitempty"`

	// language
	Language string `json:"language,omitempty"`
}

// Validate validates this packages count per language
func (m *PackagesCountPerLanguage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this packages count per language based on context it is used
func (m *PackagesCountPerLanguage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackagesCountPerLanguage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackagesCountPerLanguage) UnmarshalBinary(b []byte) error {
	var res PackagesCountPerLanguage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
