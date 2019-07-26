// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexValidateServerFirmwareVersionCompatibilityTaskRef hyperflex validate server firmware version compatibility task ref
// swagger:model hyperflexValidateServerFirmwareVersionCompatibilityTaskRef
type HyperflexValidateServerFirmwareVersionCompatibilityTaskRef struct {
	MoMoRef
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexValidateServerFirmwareVersionCompatibilityTaskRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoMoRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoMoRef = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexValidateServerFirmwareVersionCompatibilityTaskRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MoMoRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex validate server firmware version compatibility task ref
func (m *HyperflexValidateServerFirmwareVersionCompatibilityTaskRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoMoRef
	if err := m.MoMoRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexValidateServerFirmwareVersionCompatibilityTaskRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexValidateServerFirmwareVersionCompatibilityTaskRef) UnmarshalBinary(b []byte) error {
	var res HyperflexValidateServerFirmwareVersionCompatibilityTaskRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}