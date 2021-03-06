// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdapterExtEthInterface Adapter:Ext Eth Interface
//
// Physical port of a virtual interface card.
//
// swagger:model adapterExtEthInterface
type AdapterExtEthInterface struct {
	EquipmentBase

	// A collection of references to the [adapter.Unit](mo://adapter.Unit) Managed Object.
	//
	// When this managed object is deleted, the referenced [adapter.Unit](mo://adapter.Unit) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	AdapterUnit *AdapterUnitRef `json:"AdapterUnit,omitempty"`

	// admin state
	// Read Only: true
	AdminState string `json:"AdminState,omitempty"`

	// ep dn
	// Read Only: true
	EpDn string `json:"EpDn,omitempty"`

	// ext eth interface Id
	// Read Only: true
	ExtEthInterfaceID string `json:"ExtEthInterfaceId,omitempty"`

	// interface type
	// Read Only: true
	InterfaceType string `json:"InterfaceType,omitempty"`

	// mac address
	// Read Only: true
	MacAddress string `json:"MacAddress,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// peer dn
	// Read Only: true
	PeerDn string `json:"PeerDn,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AdapterExtEthInterface) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		AdapterUnit *AdapterUnitRef `json:"AdapterUnit,omitempty"`

		AdminState string `json:"AdminState,omitempty"`

		EpDn string `json:"EpDn,omitempty"`

		ExtEthInterfaceID string `json:"ExtEthInterfaceId,omitempty"`

		InterfaceType string `json:"InterfaceType,omitempty"`

		MacAddress string `json:"MacAddress,omitempty"`

		OperState string `json:"OperState,omitempty"`

		PeerDn string `json:"PeerDn,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdapterUnit = dataAO1.AdapterUnit

	m.AdminState = dataAO1.AdminState

	m.EpDn = dataAO1.EpDn

	m.ExtEthInterfaceID = dataAO1.ExtEthInterfaceID

	m.InterfaceType = dataAO1.InterfaceType

	m.MacAddress = dataAO1.MacAddress

	m.OperState = dataAO1.OperState

	m.PeerDn = dataAO1.PeerDn

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AdapterExtEthInterface) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AdapterUnit *AdapterUnitRef `json:"AdapterUnit,omitempty"`

		AdminState string `json:"AdminState,omitempty"`

		EpDn string `json:"EpDn,omitempty"`

		ExtEthInterfaceID string `json:"ExtEthInterfaceId,omitempty"`

		InterfaceType string `json:"InterfaceType,omitempty"`

		MacAddress string `json:"MacAddress,omitempty"`

		OperState string `json:"OperState,omitempty"`

		PeerDn string `json:"PeerDn,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.AdapterUnit = m.AdapterUnit

	dataAO1.AdminState = m.AdminState

	dataAO1.EpDn = m.EpDn

	dataAO1.ExtEthInterfaceID = m.ExtEthInterfaceID

	dataAO1.InterfaceType = m.InterfaceType

	dataAO1.MacAddress = m.MacAddress

	dataAO1.OperState = m.OperState

	dataAO1.PeerDn = m.PeerDn

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this adapter ext eth interface
func (m *AdapterExtEthInterface) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdapterUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdapterExtEthInterface) validateAdapterUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.AdapterUnit) { // not required
		return nil
	}

	if m.AdapterUnit != nil {
		if err := m.AdapterUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AdapterUnit")
			}
			return err
		}
	}

	return nil
}

func (m *AdapterExtEthInterface) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdapterExtEthInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdapterExtEthInterface) UnmarshalBinary(b []byte) error {
	var res AdapterExtEthInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
