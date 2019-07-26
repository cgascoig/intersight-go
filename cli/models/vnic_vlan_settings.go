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

// VnicVlanSettings VLAN Settings
//
// VLAN configuration for the virtual interface.
//
// swagger:model vnicVlanSettings
type VnicVlanSettings struct {

	// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
	//
	DefaultVlan int64 `json:"DefaultVlan,omitempty"`

	// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
	//
	// Enum: [ACCESS TRUNK]
	Mode *string `json:"Mode,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicVlanSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		DefaultVlan int64 `json:"DefaultVlan,omitempty"`

		Mode *string `json:"Mode,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.DefaultVlan = dataAO0.DefaultVlan

	m.Mode = dataAO0.Mode

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicVlanSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		DefaultVlan int64 `json:"DefaultVlan,omitempty"`

		Mode *string `json:"Mode,omitempty"`
	}

	dataAO0.DefaultVlan = m.DefaultVlan

	dataAO0.Mode = m.Mode

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic vlan settings
func (m *VnicVlanSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vnicVlanSettingsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCESS","TRUNK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vnicVlanSettingsTypeModePropEnum = append(vnicVlanSettingsTypeModePropEnum, v)
	}
}

// property enum
func (m *VnicVlanSettings) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vnicVlanSettingsTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VnicVlanSettings) validateMode(formats strfmt.Registry) error {

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
func (m *VnicVlanSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicVlanSettings) UnmarshalBinary(b []byte) error {
	var res VnicVlanSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
