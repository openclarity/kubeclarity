// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ScanStatus scan status
//
// swagger:model ScanStatus
type ScanStatus string

func NewScanStatus(value ScanStatus) *ScanStatus {
	v := value
	return &v
}

const (

	// ScanStatusSUCCESS captures enum value "SUCCESS"
	ScanStatusSUCCESS ScanStatus = "SUCCESS"

	// ScanStatusFAILED captures enum value "FAILED"
	ScanStatusFAILED ScanStatus = "FAILED"
)

// for schema
var scanStatusEnum []interface{}

func init() {
	var res []ScanStatus
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scanStatusEnum = append(scanStatusEnum, v)
	}
}

func (m ScanStatus) validateScanStatusEnum(path, location string, value ScanStatus) error {
	if err := validate.EnumCase(path, location, value, scanStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this scan status
func (m ScanStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateScanStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this scan status based on context it is used
func (m ScanStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
