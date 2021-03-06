// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdapterAdapterConfig Adapter Configuration
//
// Global adapter level settings.
//
// swagger:model adapterAdapterConfig
type AdapterAdapterConfig struct {

	// Global Ethernet settings for this adapter.
	//
	EthSettings *AdapterEthSettings `json:"EthSettings,omitempty"`

	// Global Fibre Channel settings for this adapter.
	//
	FcSettings *AdapterFcSettings `json:"FcSettings,omitempty"`

	// PCIe slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
	//
	SlotID string `json:"SlotId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AdapterAdapterConfig) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		EthSettings *AdapterEthSettings `json:"EthSettings,omitempty"`

		FcSettings *AdapterFcSettings `json:"FcSettings,omitempty"`

		SlotID string `json:"SlotId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.EthSettings = dataAO0.EthSettings

	m.FcSettings = dataAO0.FcSettings

	m.SlotID = dataAO0.SlotID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AdapterAdapterConfig) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		EthSettings *AdapterEthSettings `json:"EthSettings,omitempty"`

		FcSettings *AdapterFcSettings `json:"FcSettings,omitempty"`

		SlotID string `json:"SlotId,omitempty"`
	}

	dataAO0.EthSettings = m.EthSettings

	dataAO0.FcSettings = m.FcSettings

	dataAO0.SlotID = m.SlotID

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this adapter adapter config
func (m *AdapterAdapterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFcSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdapterAdapterConfig) validateEthSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.EthSettings) { // not required
		return nil
	}

	if m.EthSettings != nil {
		if err := m.EthSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EthSettings")
			}
			return err
		}
	}

	return nil
}

func (m *AdapterAdapterConfig) validateFcSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.FcSettings) { // not required
		return nil
	}

	if m.FcSettings != nil {
		if err := m.FcSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FcSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdapterAdapterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdapterAdapterConfig) UnmarshalBinary(b []byte) error {
	var res AdapterAdapterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
