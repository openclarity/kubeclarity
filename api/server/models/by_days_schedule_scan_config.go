// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ByDaysScheduleScanConfig by days schedule scan config
//
// swagger:model ByDaysScheduleScanConfig
type ByDaysScheduleScanConfig struct {

	// days interval
	// Minimum: 1
	DaysInterval int64 `json:"daysInterval,omitempty"`

	// time of day
	TimeOfDay *TimeOfDay `json:"timeOfDay,omitempty"`
}

// ScheduleScanConfigType gets the schedule scan config type of this subtype
func (m *ByDaysScheduleScanConfig) ScheduleScanConfigType() string {
	return "ByDaysScheduleScanConfig"
}

// SetScheduleScanConfigType sets the schedule scan config type of this subtype
func (m *ByDaysScheduleScanConfig) SetScheduleScanConfigType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ByDaysScheduleScanConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// days interval
		// Minimum: 1
		DaysInterval int64 `json:"daysInterval,omitempty"`

		// time of day
		TimeOfDay *TimeOfDay `json:"timeOfDay,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ScheduleScanConfigType string `json:"ScheduleScanConfigType,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ByDaysScheduleScanConfig

	if base.ScheduleScanConfigType != result.ScheduleScanConfigType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid ScheduleScanConfigType value: %q", base.ScheduleScanConfigType)
	}

	result.DaysInterval = data.DaysInterval
	result.TimeOfDay = data.TimeOfDay

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ByDaysScheduleScanConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// days interval
		// Minimum: 1
		DaysInterval int64 `json:"daysInterval,omitempty"`

		// time of day
		TimeOfDay *TimeOfDay `json:"timeOfDay,omitempty"`
	}{

		DaysInterval: m.DaysInterval,

		TimeOfDay: m.TimeOfDay,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ScheduleScanConfigType string `json:"ScheduleScanConfigType,omitempty"`
	}{

		ScheduleScanConfigType: m.ScheduleScanConfigType(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this by days schedule scan config
func (m *ByDaysScheduleScanConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaysInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOfDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ByDaysScheduleScanConfig) validateDaysInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.DaysInterval) { // not required
		return nil
	}

	if err := validate.MinimumInt("daysInterval", "body", m.DaysInterval, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ByDaysScheduleScanConfig) validateTimeOfDay(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeOfDay) { // not required
		return nil
	}

	if m.TimeOfDay != nil {
		if err := m.TimeOfDay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOfDay")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this by days schedule scan config based on the context it is used
func (m *ByDaysScheduleScanConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimeOfDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ByDaysScheduleScanConfig) contextValidateTimeOfDay(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeOfDay != nil {
		if err := m.TimeOfDay.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOfDay")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ByDaysScheduleScanConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ByDaysScheduleScanConfig) UnmarshalBinary(b []byte) error {
	var res ByDaysScheduleScanConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
