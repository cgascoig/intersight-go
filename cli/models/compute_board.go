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

// ComputeBoard Compute:Board
//
// Mother board of a server.
//
// swagger:model computeBoard
type ComputeBoard struct {
	EquipmentBase

	// board Id
	// Read Only: true
	BoardID int64 `json:"BoardId,omitempty"`

	// A collection of references to the [compute.Blade](mo://compute.Blade) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.Blade](mo://compute.Blade) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

	// A collection of references to the [compute.RackUnit](mo://compute.RackUnit) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.RackUnit](mo://compute.RackUnit) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

	// Cpu type controller
	// Read Only: true
	CPUTypeController string `json:"CpuTypeController,omitempty"`

	// equipment tpms
	// Read Only: true
	EquipmentTpms []*EquipmentTpmRef `json:"EquipmentTpms"`

	// It shows Graphics cards present in a server.
	//
	// Read Only: true
	GraphicsCards []*GraphicsCardRef `json:"GraphicsCards"`

	// memory arrays
	// Read Only: true
	MemoryArrays []*MemoryArrayRef `json:"MemoryArrays"`

	// oper power state
	// Read Only: true
	OperPowerState string `json:"OperPowerState,omitempty"`

	// It shows PCI CoprocessorCard present in a server.
	//
	// Read Only: true
	PciCoprocessorCards []*PciCoprocessorCardRef `json:"PciCoprocessorCards"`

	// It shows PCI Switches presen in a server.
	//
	// Read Only: true
	PciSwitch []*PciSwitchRef `json:"PciSwitch"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// processors
	// Read Only: true
	Processors []*ProcessorUnitRef `json:"Processors"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// security units
	// Read Only: true
	SecurityUnits []*SecurityUnitRef `json:"SecurityUnits"`

	// storage controllers
	// Read Only: true
	StorageControllers []*StorageControllerRef `json:"StorageControllers"`

	// storage flex flash controllers
	// Read Only: true
	StorageFlexFlashControllers []*StorageFlexFlashControllerRef `json:"StorageFlexFlashControllers"`

	// storage flex util controllers
	// Read Only: true
	StorageFlexUtilControllers []*StorageFlexUtilControllerRef `json:"StorageFlexUtilControllers"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ComputeBoard) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		BoardID int64 `json:"BoardId,omitempty"`

		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		CPUTypeController string `json:"CpuTypeController,omitempty"`

		EquipmentTpms []*EquipmentTpmRef `json:"EquipmentTpms"`

		GraphicsCards []*GraphicsCardRef `json:"GraphicsCards"`

		MemoryArrays []*MemoryArrayRef `json:"MemoryArrays"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		PciCoprocessorCards []*PciCoprocessorCardRef `json:"PciCoprocessorCards"`

		PciSwitch []*PciSwitchRef `json:"PciSwitch"`

		Presence string `json:"Presence,omitempty"`

		Processors []*ProcessorUnitRef `json:"Processors"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SecurityUnits []*SecurityUnitRef `json:"SecurityUnits"`

		StorageControllers []*StorageControllerRef `json:"StorageControllers"`

		StorageFlexFlashControllers []*StorageFlexFlashControllerRef `json:"StorageFlexFlashControllers"`

		StorageFlexUtilControllers []*StorageFlexUtilControllerRef `json:"StorageFlexUtilControllers"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BoardID = dataAO1.BoardID

	m.ComputeBlade = dataAO1.ComputeBlade

	m.ComputeRackUnit = dataAO1.ComputeRackUnit

	m.CPUTypeController = dataAO1.CPUTypeController

	m.EquipmentTpms = dataAO1.EquipmentTpms

	m.GraphicsCards = dataAO1.GraphicsCards

	m.MemoryArrays = dataAO1.MemoryArrays

	m.OperPowerState = dataAO1.OperPowerState

	m.PciCoprocessorCards = dataAO1.PciCoprocessorCards

	m.PciSwitch = dataAO1.PciSwitch

	m.Presence = dataAO1.Presence

	m.Processors = dataAO1.Processors

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.SecurityUnits = dataAO1.SecurityUnits

	m.StorageControllers = dataAO1.StorageControllers

	m.StorageFlexFlashControllers = dataAO1.StorageFlexFlashControllers

	m.StorageFlexUtilControllers = dataAO1.StorageFlexUtilControllers

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ComputeBoard) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BoardID int64 `json:"BoardId,omitempty"`

		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		CPUTypeController string `json:"CpuTypeController,omitempty"`

		EquipmentTpms []*EquipmentTpmRef `json:"EquipmentTpms"`

		GraphicsCards []*GraphicsCardRef `json:"GraphicsCards"`

		MemoryArrays []*MemoryArrayRef `json:"MemoryArrays"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		PciCoprocessorCards []*PciCoprocessorCardRef `json:"PciCoprocessorCards"`

		PciSwitch []*PciSwitchRef `json:"PciSwitch"`

		Presence string `json:"Presence,omitempty"`

		Processors []*ProcessorUnitRef `json:"Processors"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SecurityUnits []*SecurityUnitRef `json:"SecurityUnits"`

		StorageControllers []*StorageControllerRef `json:"StorageControllers"`

		StorageFlexFlashControllers []*StorageFlexFlashControllerRef `json:"StorageFlexFlashControllers"`

		StorageFlexUtilControllers []*StorageFlexUtilControllerRef `json:"StorageFlexUtilControllers"`
	}

	dataAO1.BoardID = m.BoardID

	dataAO1.ComputeBlade = m.ComputeBlade

	dataAO1.ComputeRackUnit = m.ComputeRackUnit

	dataAO1.CPUTypeController = m.CPUTypeController

	dataAO1.EquipmentTpms = m.EquipmentTpms

	dataAO1.GraphicsCards = m.GraphicsCards

	dataAO1.MemoryArrays = m.MemoryArrays

	dataAO1.OperPowerState = m.OperPowerState

	dataAO1.PciCoprocessorCards = m.PciCoprocessorCards

	dataAO1.PciSwitch = m.PciSwitch

	dataAO1.Presence = m.Presence

	dataAO1.Processors = m.Processors

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.SecurityUnits = m.SecurityUnits

	dataAO1.StorageControllers = m.StorageControllers

	dataAO1.StorageFlexFlashControllers = m.StorageFlexFlashControllers

	dataAO1.StorageFlexUtilControllers = m.StorageFlexUtilControllers

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this compute board
func (m *ComputeBoard) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBlade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeRackUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEquipmentTpms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraphicsCards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryArrays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePciCoprocessorCards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePciSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityUnits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageControllers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageFlexFlashControllers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageFlexUtilControllers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeBoard) validateComputeBlade(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBlade) { // not required
		return nil
	}

	if m.ComputeBlade != nil {
		if err := m.ComputeBlade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBlade")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeBoard) validateComputeRackUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeRackUnit) { // not required
		return nil
	}

	if m.ComputeRackUnit != nil {
		if err := m.ComputeRackUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeRackUnit")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeBoard) validateEquipmentTpms(formats strfmt.Registry) error {

	if swag.IsZero(m.EquipmentTpms) { // not required
		return nil
	}

	for i := 0; i < len(m.EquipmentTpms); i++ {
		if swag.IsZero(m.EquipmentTpms[i]) { // not required
			continue
		}

		if m.EquipmentTpms[i] != nil {
			if err := m.EquipmentTpms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EquipmentTpms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateGraphicsCards(formats strfmt.Registry) error {

	if swag.IsZero(m.GraphicsCards) { // not required
		return nil
	}

	for i := 0; i < len(m.GraphicsCards); i++ {
		if swag.IsZero(m.GraphicsCards[i]) { // not required
			continue
		}

		if m.GraphicsCards[i] != nil {
			if err := m.GraphicsCards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GraphicsCards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateMemoryArrays(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryArrays) { // not required
		return nil
	}

	for i := 0; i < len(m.MemoryArrays); i++ {
		if swag.IsZero(m.MemoryArrays[i]) { // not required
			continue
		}

		if m.MemoryArrays[i] != nil {
			if err := m.MemoryArrays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("MemoryArrays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validatePciCoprocessorCards(formats strfmt.Registry) error {

	if swag.IsZero(m.PciCoprocessorCards) { // not required
		return nil
	}

	for i := 0; i < len(m.PciCoprocessorCards); i++ {
		if swag.IsZero(m.PciCoprocessorCards[i]) { // not required
			continue
		}

		if m.PciCoprocessorCards[i] != nil {
			if err := m.PciCoprocessorCards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PciCoprocessorCards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validatePciSwitch(formats strfmt.Registry) error {

	if swag.IsZero(m.PciSwitch) { // not required
		return nil
	}

	for i := 0; i < len(m.PciSwitch); i++ {
		if swag.IsZero(m.PciSwitch[i]) { // not required
			continue
		}

		if m.PciSwitch[i] != nil {
			if err := m.PciSwitch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PciSwitch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateProcessors(formats strfmt.Registry) error {

	if swag.IsZero(m.Processors) { // not required
		return nil
	}

	for i := 0; i < len(m.Processors); i++ {
		if swag.IsZero(m.Processors[i]) { // not required
			continue
		}

		if m.Processors[i] != nil {
			if err := m.Processors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Processors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *ComputeBoard) validateSecurityUnits(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityUnits) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityUnits); i++ {
		if swag.IsZero(m.SecurityUnits[i]) { // not required
			continue
		}

		if m.SecurityUnits[i] != nil {
			if err := m.SecurityUnits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SecurityUnits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateStorageControllers(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageControllers) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageControllers); i++ {
		if swag.IsZero(m.StorageControllers[i]) { // not required
			continue
		}

		if m.StorageControllers[i] != nil {
			if err := m.StorageControllers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StorageControllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateStorageFlexFlashControllers(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageFlexFlashControllers) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageFlexFlashControllers); i++ {
		if swag.IsZero(m.StorageFlexFlashControllers[i]) { // not required
			continue
		}

		if m.StorageFlexFlashControllers[i] != nil {
			if err := m.StorageFlexFlashControllers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StorageFlexFlashControllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeBoard) validateStorageFlexUtilControllers(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageFlexUtilControllers) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageFlexUtilControllers); i++ {
		if swag.IsZero(m.StorageFlexUtilControllers[i]) { // not required
			continue
		}

		if m.StorageFlexUtilControllers[i] != nil {
			if err := m.StorageFlexUtilControllers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StorageFlexUtilControllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeBoard) UnmarshalBinary(b []byte) error {
	var res ComputeBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
