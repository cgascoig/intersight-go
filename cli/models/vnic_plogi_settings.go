// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicPlogiSettings Plogi Settings
//
// Fibre Channel Plogi Settings.
//
// swagger:model vnicPlogiSettings
type VnicPlogiSettings struct {

	// The number of times that the system tries to log in to a port after the first failure.
	//
	Retries int64 `json:"Retries,omitempty"`

	// The number of milliseconds that the system waits before it tries to log in again.
	//
	Timeout int64 `json:"Timeout,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicPlogiSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Retries int64 `json:"Retries,omitempty"`

		Timeout int64 `json:"Timeout,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Retries = dataAO0.Retries

	m.Timeout = dataAO0.Timeout

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicPlogiSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Retries int64 `json:"Retries,omitempty"`

		Timeout int64 `json:"Timeout,omitempty"`
	}

	dataAO0.Retries = m.Retries

	dataAO0.Timeout = m.Timeout

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic plogi settings
func (m *VnicPlogiSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicPlogiSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicPlogiSettings) UnmarshalBinary(b []byte) error {
	var res VnicPlogiSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
