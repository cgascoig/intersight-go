// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicNvgreSettings NVGRE Settings
//
// Network Virtualization using Generic Routing Encapsulation Settings.
//
// swagger:model vnicNvgreSettings
type VnicNvgreSettings struct {

	// Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface.
	//
	Enabled bool `json:"Enabled,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicNvgreSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Enabled bool `json:"Enabled,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Enabled = dataAO0.Enabled

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicNvgreSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Enabled bool `json:"Enabled,omitempty"`
	}

	dataAO0.Enabled = m.Enabled

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic nvgre settings
func (m *VnicNvgreSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicNvgreSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicNvgreSettings) UnmarshalBinary(b []byte) error {
	var res VnicNvgreSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
