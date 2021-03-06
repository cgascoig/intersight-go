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

// StorageFlexUtilController Storage:Flex Util Controller
//
// Storage Flex Util Adapter.
//
// swagger:model storageFlexUtilController
type StorageFlexUtilController struct {
	InventoryBase

	// A collection of references to the [compute.Board](mo://compute.Board) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.Board](mo://compute.Board) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

	// controller name
	ControllerName string `json:"ControllerName,omitempty"`

	// controller status
	ControllerStatus string `json:"ControllerStatus,omitempty"`

	// ff controller Id
	FfControllerID string `json:"FfControllerId,omitempty"`

	// flex util physical drives
	// Read Only: true
	FlexUtilPhysicalDrives []*StorageFlexUtilPhysicalDriveRef `json:"FlexUtilPhysicalDrives"`

	// flex util virtual drives
	// Read Only: true
	FlexUtilVirtualDrives []*StorageFlexUtilVirtualDriveRef `json:"FlexUtilVirtualDrives"`

	// internal state
	InternalState string `json:"InternalState,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageFlexUtilController) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		ControllerName string `json:"ControllerName,omitempty"`

		ControllerStatus string `json:"ControllerStatus,omitempty"`

		FfControllerID string `json:"FfControllerId,omitempty"`

		FlexUtilPhysicalDrives []*StorageFlexUtilPhysicalDriveRef `json:"FlexUtilPhysicalDrives"`

		FlexUtilVirtualDrives []*StorageFlexUtilVirtualDriveRef `json:"FlexUtilVirtualDrives"`

		InternalState string `json:"InternalState,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ComputeBoard = dataAO1.ComputeBoard

	m.ControllerName = dataAO1.ControllerName

	m.ControllerStatus = dataAO1.ControllerStatus

	m.FfControllerID = dataAO1.FfControllerID

	m.FlexUtilPhysicalDrives = dataAO1.FlexUtilPhysicalDrives

	m.FlexUtilVirtualDrives = dataAO1.FlexUtilVirtualDrives

	m.InternalState = dataAO1.InternalState

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageFlexUtilController) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		ControllerName string `json:"ControllerName,omitempty"`

		ControllerStatus string `json:"ControllerStatus,omitempty"`

		FfControllerID string `json:"FfControllerId,omitempty"`

		FlexUtilPhysicalDrives []*StorageFlexUtilPhysicalDriveRef `json:"FlexUtilPhysicalDrives"`

		FlexUtilVirtualDrives []*StorageFlexUtilVirtualDriveRef `json:"FlexUtilVirtualDrives"`

		InternalState string `json:"InternalState,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.ComputeBoard = m.ComputeBoard

	dataAO1.ControllerName = m.ControllerName

	dataAO1.ControllerStatus = m.ControllerStatus

	dataAO1.FfControllerID = m.FfControllerID

	dataAO1.FlexUtilPhysicalDrives = m.FlexUtilPhysicalDrives

	dataAO1.FlexUtilVirtualDrives = m.FlexUtilVirtualDrives

	dataAO1.InternalState = m.InternalState

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage flex util controller
func (m *StorageFlexUtilController) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBoard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlexUtilPhysicalDrives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlexUtilVirtualDrives(formats); err != nil {
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

func (m *StorageFlexUtilController) validateComputeBoard(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBoard) { // not required
		return nil
	}

	if m.ComputeBoard != nil {
		if err := m.ComputeBoard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBoard")
			}
			return err
		}
	}

	return nil
}

func (m *StorageFlexUtilController) validateFlexUtilPhysicalDrives(formats strfmt.Registry) error {

	if swag.IsZero(m.FlexUtilPhysicalDrives) { // not required
		return nil
	}

	for i := 0; i < len(m.FlexUtilPhysicalDrives); i++ {
		if swag.IsZero(m.FlexUtilPhysicalDrives[i]) { // not required
			continue
		}

		if m.FlexUtilPhysicalDrives[i] != nil {
			if err := m.FlexUtilPhysicalDrives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FlexUtilPhysicalDrives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageFlexUtilController) validateFlexUtilVirtualDrives(formats strfmt.Registry) error {

	if swag.IsZero(m.FlexUtilVirtualDrives) { // not required
		return nil
	}

	for i := 0; i < len(m.FlexUtilVirtualDrives); i++ {
		if swag.IsZero(m.FlexUtilVirtualDrives[i]) { // not required
			continue
		}

		if m.FlexUtilVirtualDrives[i] != nil {
			if err := m.FlexUtilVirtualDrives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FlexUtilVirtualDrives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageFlexUtilController) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StorageFlexUtilController) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageFlexUtilController) UnmarshalBinary(b []byte) error {
	var res StorageFlexUtilController
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
