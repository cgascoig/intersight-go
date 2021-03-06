// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StoragePureProtectionGroupSnapshot Storage:Pure Protection Group Snapshot
//
// Protection group snapshot entity in Pure protection group.
//
// swagger:model storagePureProtectionGroupSnapshot
type StoragePureProtectionGroupSnapshot struct {
	StorageProtectionGroupSnapshot

	// Device registration managed object that represents this storage array connection to Intersight.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureProtectionGroupSnapshot) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageProtectionGroupSnapshot
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageProtectionGroupSnapshot = aO0

	// AO1
	var dataAO1 struct {
		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureProtectionGroupSnapshot) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageProtectionGroupSnapshot)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure protection group snapshot
func (m *StoragePureProtectionGroupSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageProtectionGroupSnapshot
	if err := m.StorageProtectionGroupSnapshot.Validate(formats); err != nil {
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

func (m *StoragePureProtectionGroupSnapshot) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StoragePureProtectionGroupSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureProtectionGroupSnapshot) UnmarshalBinary(b []byte) error {
	var res StoragePureProtectionGroupSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
