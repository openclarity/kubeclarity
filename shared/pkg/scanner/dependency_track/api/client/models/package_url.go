// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageURL package URL
//
// swagger:model PackageURL
type PackageURL struct {

	// coordinates
	Coordinates string `json:"coordinates,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// qualifiers
	Qualifiers map[string]string `json:"qualifiers,omitempty"`

	// scheme
	Scheme string `json:"scheme,omitempty"`

	// subpath
	Subpath string `json:"subpath,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this package URL
func (m *PackageURL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this package URL based on context it is used
func (m *PackageURL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageURL) UnmarshalBinary(b []byte) error {
	var res PackageURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
