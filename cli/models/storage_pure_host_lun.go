// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StoragePureHostLun Storage:Pure Host Lun
//
// A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
//
// swagger:model storagePureHostLun
type StoragePureHostLun struct {
	StorageHostLun

	// Host group relationship object associated with LUN. It is empty, if volume has direct connection to host.
	//
	// Read Only: true
	HostGroup *StorageHostRef `json:"HostGroup,omitempty"`

	// Name of the host group associated with LUN.
	//
	// Read Only: true
	HostGroupName string `json:"HostGroupName,omitempty"`

	// Device registration managed object that represents this storage array connection to Intersight.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection.
	//
	// Read Only: true
	Shared *bool `json:"Shared,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureHostLun) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageHostLun
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageHostLun = aO0

	// AO1
	var dataAO1 struct {
		HostGroup *StorageHostRef `json:"HostGroup,omitempty"`

		HostGroupName string `json:"HostGroupName,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Shared *bool `json:"Shared,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.HostGroup = dataAO1.HostGroup

	m.HostGroupName = dataAO1.HostGroupName

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Shared = dataAO1.Shared

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureHostLun) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageHostLun)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		HostGroup *StorageHostRef `json:"HostGroup,omitempty"`

		HostGroupName string `json:"HostGroupName,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Shared *bool `json:"Shared,omitempty"`
	}

	dataAO1.HostGroup = m.HostGroup

	dataAO1.HostGroupName = m.HostGroupName

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Shared = m.Shared

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure host lun
func (m *StoragePureHostLun) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageHostLun
	if err := m.StorageHostLun.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostGroup(formats); err != nil {
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

func (m *StoragePureHostLun) validateHostGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroup) { // not required
		return nil
	}

	if m.HostGroup != nil {
		if err := m.HostGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostGroup")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePureHostLun) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StoragePureHostLun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureHostLun) UnmarshalBinary(b []byte) error {
	var res StoragePureHostLun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
