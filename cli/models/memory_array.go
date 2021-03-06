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

// MemoryArray Memory:Array
//
// Holder housing multiple memory units.
//
// swagger:model memoryArray
type MemoryArray struct {
	EquipmentBase

	// array Id
	// Read Only: true
	ArrayID int64 `json:"ArrayId,omitempty"`

	// A collection of references to the [compute.Board](mo://compute.Board) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.Board](mo://compute.Board) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

	// Cpu Id
	// Read Only: true
	CPUID int64 `json:"CpuId,omitempty"`

	// current capacity
	// Read Only: true
	CurrentCapacity string `json:"CurrentCapacity,omitempty"`

	// error correction
	// Read Only: true
	ErrorCorrection string `json:"ErrorCorrection,omitempty"`

	// max capacity
	// Read Only: true
	MaxCapacity string `json:"MaxCapacity,omitempty"`

	// max devices
	// Read Only: true
	MaxDevices string `json:"MaxDevices,omitempty"`

	// oper power state
	// Read Only: true
	OperPowerState string `json:"OperPowerState,omitempty"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// units
	// Read Only: true
	Units []*MemoryUnitRef `json:"Units"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemoryArray) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		ArrayID int64 `json:"ArrayId,omitempty"`

		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		CPUID int64 `json:"CpuId,omitempty"`

		CurrentCapacity string `json:"CurrentCapacity,omitempty"`

		ErrorCorrection string `json:"ErrorCorrection,omitempty"`

		MaxCapacity string `json:"MaxCapacity,omitempty"`

		MaxDevices string `json:"MaxDevices,omitempty"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Units []*MemoryUnitRef `json:"Units"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ArrayID = dataAO1.ArrayID

	m.ComputeBoard = dataAO1.ComputeBoard

	m.CPUID = dataAO1.CPUID

	m.CurrentCapacity = dataAO1.CurrentCapacity

	m.ErrorCorrection = dataAO1.ErrorCorrection

	m.MaxCapacity = dataAO1.MaxCapacity

	m.MaxDevices = dataAO1.MaxDevices

	m.OperPowerState = dataAO1.OperPowerState

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Units = dataAO1.Units

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemoryArray) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ArrayID int64 `json:"ArrayId,omitempty"`

		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		CPUID int64 `json:"CpuId,omitempty"`

		CurrentCapacity string `json:"CurrentCapacity,omitempty"`

		ErrorCorrection string `json:"ErrorCorrection,omitempty"`

		MaxCapacity string `json:"MaxCapacity,omitempty"`

		MaxDevices string `json:"MaxDevices,omitempty"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Units []*MemoryUnitRef `json:"Units"`
	}

	dataAO1.ArrayID = m.ArrayID

	dataAO1.ComputeBoard = m.ComputeBoard

	dataAO1.CPUID = m.CPUID

	dataAO1.CurrentCapacity = m.CurrentCapacity

	dataAO1.ErrorCorrection = m.ErrorCorrection

	dataAO1.MaxCapacity = m.MaxCapacity

	dataAO1.MaxDevices = m.MaxDevices

	dataAO1.OperPowerState = m.OperPowerState

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Units = m.Units

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this memory array
func (m *MemoryArray) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBoard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemoryArray) validateComputeBoard(formats strfmt.Registry) error {

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

func (m *MemoryArray) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *MemoryArray) validateUnits(formats strfmt.Registry) error {

	if swag.IsZero(m.Units) { // not required
		return nil
	}

	for i := 0; i < len(m.Units); i++ {
		if swag.IsZero(m.Units[i]) { // not required
			continue
		}

		if m.Units[i] != nil {
			if err := m.Units[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemoryArray) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryArray) UnmarshalBinary(b []byte) error {
	var res MemoryArray
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
