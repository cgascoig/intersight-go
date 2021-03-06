// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ManagementEntity Management:Entity
//
// Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
//
// swagger:model managementEntity
type ManagementEntity struct {
	InventoryBase

	// Identity of the Fabric Interconnect - A/B.
	//
	// Read Only: true
	EntityID string `json:"EntityId,omitempty"`

	// Role (Primary / Subordinate) of the Fabric Interconnect.
	//
	// Read Only: true
	Leadership string `json:"Leadership,omitempty"`

	// A collection of references to the [network.Element](mo://network.Element) Managed Object.
	//
	// When this managed object is deleted, the referenced [network.Element](mo://network.Element) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	NetworkElement *NetworkElementRef `json:"NetworkElement,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ManagementEntity) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		EntityID string `json:"EntityId,omitempty"`

		Leadership string `json:"Leadership,omitempty"`

		NetworkElement *NetworkElementRef `json:"NetworkElement,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EntityID = dataAO1.EntityID

	m.Leadership = dataAO1.Leadership

	m.NetworkElement = dataAO1.NetworkElement

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ManagementEntity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		EntityID string `json:"EntityId,omitempty"`

		Leadership string `json:"Leadership,omitempty"`

		NetworkElement *NetworkElementRef `json:"NetworkElement,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.EntityID = m.EntityID

	dataAO1.Leadership = m.Leadership

	dataAO1.NetworkElement = m.NetworkElement

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this management entity
func (m *ManagementEntity) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkElement(formats); err != nil {
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

func (m *ManagementEntity) validateNetworkElement(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkElement) { // not required
		return nil
	}

	if m.NetworkElement != nil {
		if err := m.NetworkElement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkElement")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementEntity) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *ManagementEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementEntity) UnmarshalBinary(b []byte) error {
	var res ManagementEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
