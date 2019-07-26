// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoragePureVolume Storage:Pure Volume
//
// A volume entity in PureStorage FlashArray.
//
// swagger:model storagePureVolume
type StoragePureVolume struct {
	StorageVolume

	// Creation time of the volume.
	//
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// Device registration managed object that represents this storage array connection to Intersight.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Serial number of the volume.
	//
	// Read Only: true
	Serial string `json:"Serial,omitempty"`

	// Source from which the volume is created. Applicable only if the volume is cloned from other volume or snapshot.
	//
	// Read Only: true
	Source string `json:"Source,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureVolume) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageVolume
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageVolume = aO0

	// AO1
	var dataAO1 struct {
		Created strfmt.DateTime `json:"Created,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Serial string `json:"Serial,omitempty"`

		Source string `json:"Source,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Created = dataAO1.Created

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Serial = dataAO1.Serial

	m.Source = dataAO1.Source

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureVolume) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageVolume)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Created strfmt.DateTime `json:"Created,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Serial string `json:"Serial,omitempty"`

		Source string `json:"Source,omitempty"`
	}

	dataAO1.Created = m.Created

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Serial = m.Serial

	dataAO1.Source = m.Source

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure volume
func (m *StoragePureVolume) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageVolume
	if err := m.StorageVolume.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
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

func (m *StoragePureVolume) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StoragePureVolume) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StoragePureVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureVolume) UnmarshalBinary(b []byte) error {
	var res StoragePureVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
