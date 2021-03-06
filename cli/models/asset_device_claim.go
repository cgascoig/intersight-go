// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AssetDeviceClaim Asset:Device Claim
//
// DeviceClaim captures the intent to claim a device to an Intersight account. A device can be unclaimed by performing a DELETE on a DeviceClaim instance. When performing a claim, a secret passphrase must be obtained from the device connector UI/API by a sufficiently privileged user. The passphrase is timebound and proves that the user currently has privileged administrative access to the device being claimed.
//
// swagger:model assetDeviceClaim
type AssetDeviceClaim struct {
	MoBaseMo

	// The account of the user that has claimed the device.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// A collection of references to the [asset.DeviceRegistration](mo://asset.DeviceRegistration) Managed Object.
	//
	// When this managed object is deleted, the referenced [asset.DeviceRegistration](mo://asset.DeviceRegistration) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	Device *AssetDeviceRegistrationRef `json:"Device,omitempty"`

	// The list of devices that underwent change during the claim process.
	//
	// Read Only: true
	DeviceUpdates []*AssetConnectionControlMessage `json:"DeviceUpdates"`

	// Obtained from the device connector management UI or API (REST endpoint /connector/SecurityTokens).
	//
	SecurityToken string `json:"SecurityToken,omitempty"`

	// Obtained from the device connector management UI or API (REST endpoint /connector/DeviceIdentifiers).
	//
	SerialNumber string `json:"SerialNumber,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetDeviceClaim) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Device *AssetDeviceRegistrationRef `json:"Device,omitempty"`

		DeviceUpdates []*AssetConnectionControlMessage `json:"DeviceUpdates"`

		SecurityToken string `json:"SecurityToken,omitempty"`

		SerialNumber string `json:"SerialNumber,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.Device = dataAO1.Device

	m.DeviceUpdates = dataAO1.DeviceUpdates

	m.SecurityToken = dataAO1.SecurityToken

	m.SerialNumber = dataAO1.SerialNumber

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetDeviceClaim) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Device *AssetDeviceRegistrationRef `json:"Device,omitempty"`

		DeviceUpdates []*AssetConnectionControlMessage `json:"DeviceUpdates"`

		SecurityToken string `json:"SecurityToken,omitempty"`

		SerialNumber string `json:"SerialNumber,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.Device = m.Device

	dataAO1.DeviceUpdates = m.DeviceUpdates

	dataAO1.SecurityToken = m.SecurityToken

	dataAO1.SerialNumber = m.SerialNumber

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset device claim
func (m *AssetDeviceClaim) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetDeviceClaim) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

func (m *AssetDeviceClaim) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Device")
			}
			return err
		}
	}

	return nil
}

func (m *AssetDeviceClaim) validateDeviceUpdates(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceUpdates) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceUpdates); i++ {
		if swag.IsZero(m.DeviceUpdates[i]) { // not required
			continue
		}

		if m.DeviceUpdates[i] != nil {
			if err := m.DeviceUpdates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeviceUpdates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetDeviceClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetDeviceClaim) UnmarshalBinary(b []byte) error {
	var res AssetDeviceClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
