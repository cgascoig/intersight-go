// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VnicFcInterruptSettings Interrupt Settings
//
// Interrupt Settings for the virtual fibre channel interface.
//
// swagger:model vnicFcInterruptSettings
type VnicFcInterruptSettings struct {

	// The preferred driver interrupt mode. This can be one of the following:- MSIx — Message Signaled Interrupts (MSI) with the optional extension. MSI   — MSI only. INTx  — PCI INTx interrupts. MSIx is the recommended option.
	//
	// Enum: [MSIx MSI INTx]
	Mode *string `json:"Mode,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicFcInterruptSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Mode *string `json:"Mode,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Mode = dataAO0.Mode

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicFcInterruptSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Mode *string `json:"Mode,omitempty"`
	}

	dataAO0.Mode = m.Mode

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic fc interrupt settings
func (m *VnicFcInterruptSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vnicFcInterruptSettingsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MSIx","MSI","INTx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vnicFcInterruptSettingsTypeModePropEnum = append(vnicFcInterruptSettingsTypeModePropEnum, v)
	}
}

// property enum
func (m *VnicFcInterruptSettings) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vnicFcInterruptSettingsTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VnicFcInterruptSettings) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("Mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcInterruptSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcInterruptSettings) UnmarshalBinary(b []byte) error {
	var res VnicFcInterruptSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
