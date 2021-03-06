// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CondHclStatusJob Cond:Hcl Status Job
//
// An HCLStatusJob is used to batch mo inventory notifications and process the evaluation of HCLStatus. When we receive a notification for an inventory MO, we will create a HCLStatusJob and inserted into the DB if it doesn't already exist. Then based on a timer we process the jobs in the DB and clear them.
//
// swagger:model condHclStatusJob
type CondHclStatusJob struct {
	MoBaseMo

	// The related managed object that will be re-evaluated when this job executes.
	//
	// Read Only: true
	ManagedObject *InventoryBaseRef `json:"ManagedObject,omitempty"`

	// The registered device associated with this job. We need this to correctly set permissions for the HCLStatus that will be re-evaluated when this job executes.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CondHclStatusJob) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ManagedObject *InventoryBaseRef `json:"ManagedObject,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ManagedObject = dataAO1.ManagedObject

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CondHclStatusJob) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ManagedObject *InventoryBaseRef `json:"ManagedObject,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.ManagedObject = m.ManagedObject

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cond hcl status job
func (m *CondHclStatusJob) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedObject(formats); err != nil {
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

func (m *CondHclStatusJob) validateManagedObject(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedObject) { // not required
		return nil
	}

	if m.ManagedObject != nil {
		if err := m.ManagedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ManagedObject")
			}
			return err
		}
	}

	return nil
}

func (m *CondHclStatusJob) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *CondHclStatusJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CondHclStatusJob) UnmarshalBinary(b []byte) error {
	var res CondHclStatusJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
