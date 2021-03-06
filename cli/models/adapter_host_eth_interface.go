// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdapterHostEthInterface Adapter:Host Eth Interface
//
// Physical / Virtual port of an adapter as seen by the host.
//
// swagger:model adapterHostEthInterface
type AdapterHostEthInterface struct {
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

	// host eth interface Id
	// Read Only: true
	HostEthInterfaceID int64 `json:"HostEthInterfaceId,omitempty"`

	// interface type
	// Read Only: true
	InterfaceType string `json:"InterfaceType,omitempty"`

	// mac address
	// Read Only: true
	MacAddress string `json:"MacAddress,omitempty"`

	// name
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// operability
	// Read Only: true
	Operability string `json:"Operability,omitempty"`

	// original mac address
	// Read Only: true
	OriginalMacAddress string `json:"OriginalMacAddress,omitempty"`

	// pci addr
	// Read Only: true
	PciAddr string `json:"PciAddr,omitempty"`

	// peer dn
	// Read Only: true
	PeerDn string `json:"PeerDn,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// virtualization preference
	// Read Only: true
	VirtualizationPreference string `json:"VirtualizationPreference,omitempty"`

	// vnic dn
	// Read Only: true
	VnicDn string `json:"VnicDn,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AdapterHostEthInterface) UnmarshalJSON(raw []byte) error {
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

		HostEthInterfaceID int64 `json:"HostEthInterfaceId,omitempty"`

		InterfaceType string `json:"InterfaceType,omitempty"`

		MacAddress string `json:"MacAddress,omitempty"`

		Name string `json:"Name,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		OriginalMacAddress string `json:"OriginalMacAddress,omitempty"`

		PciAddr string `json:"PciAddr,omitempty"`

		PeerDn string `json:"PeerDn,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		VirtualizationPreference string `json:"VirtualizationPreference,omitempty"`

		VnicDn string `json:"VnicDn,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdapterUnit = dataAO1.AdapterUnit

	m.AdminState = dataAO1.AdminState

	m.EpDn = dataAO1.EpDn

	m.HostEthInterfaceID = dataAO1.HostEthInterfaceID

	m.InterfaceType = dataAO1.InterfaceType

	m.MacAddress = dataAO1.MacAddress

	m.Name = dataAO1.Name

	m.OperState = dataAO1.OperState

	m.Operability = dataAO1.Operability

	m.OriginalMacAddress = dataAO1.OriginalMacAddress

	m.PciAddr = dataAO1.PciAddr

	m.PeerDn = dataAO1.PeerDn

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.VirtualizationPreference = dataAO1.VirtualizationPreference

	m.VnicDn = dataAO1.VnicDn

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AdapterHostEthInterface) MarshalJSON() ([]byte, error) {
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

		HostEthInterfaceID int64 `json:"HostEthInterfaceId,omitempty"`

		InterfaceType string `json:"InterfaceType,omitempty"`

		MacAddress string `json:"MacAddress,omitempty"`

		Name string `json:"Name,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		OriginalMacAddress string `json:"OriginalMacAddress,omitempty"`

		PciAddr string `json:"PciAddr,omitempty"`

		PeerDn string `json:"PeerDn,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		VirtualizationPreference string `json:"VirtualizationPreference,omitempty"`

		VnicDn string `json:"VnicDn,omitempty"`
	}

	dataAO1.AdapterUnit = m.AdapterUnit

	dataAO1.AdminState = m.AdminState

	dataAO1.EpDn = m.EpDn

	dataAO1.HostEthInterfaceID = m.HostEthInterfaceID

	dataAO1.InterfaceType = m.InterfaceType

	dataAO1.MacAddress = m.MacAddress

	dataAO1.Name = m.Name

	dataAO1.OperState = m.OperState

	dataAO1.Operability = m.Operability

	dataAO1.OriginalMacAddress = m.OriginalMacAddress

	dataAO1.PciAddr = m.PciAddr

	dataAO1.PeerDn = m.PeerDn

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.VirtualizationPreference = m.VirtualizationPreference

	dataAO1.VnicDn = m.VnicDn

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this adapter host eth interface
func (m *AdapterHostEthInterface) Validate(formats strfmt.Registry) error {
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

func (m *AdapterHostEthInterface) validateAdapterUnit(formats strfmt.Registry) error {

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

func (m *AdapterHostEthInterface) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *AdapterHostEthInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdapterHostEthInterface) UnmarshalBinary(b []byte) error {
	var res AdapterHostEthInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
