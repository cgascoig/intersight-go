// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageSnapshotSchedule Storage:Snapshot Schedule
//
// Configuration parameters for snapshot creation at source array.
//
// swagger:model storageSnapshotSchedule
type StorageSnapshotSchedule struct {
	MoBaseMo

	// Snapshot frequency. It is an interval on which snapshot is set to trigger on source array.
	// Examples:
	//     PT2H, Snapshot is performed for every 2 hours.
	//     P4D, Snapshot is scheduled for every 4 days.
	//     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	//
	//
	// Read Only: true
	Frequency string `json:"Frequency,omitempty"`

	// Name of the snapshot schedule.
	//
	Name string `json:"Name,omitempty"`

	// Protection group relationship object.
	//
	// Read Only: true
	ProtectionGroup *StorageProtectionGroupRef `json:"ProtectionGroup,omitempty"`

	// Duration to keep the snapshots on the source array.
	// Once the period expires, system deletes the snapshot automatically from source array.
	// Examples:
	// P200D,  200 days.
	// PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	//
	//
	// Read Only: true
	RetentionTime string `json:"RetentionTime,omitempty"`

	// Preferred time of the day to capture the snapshot. it is applicable only if the frequency is set for a day or more.
	// Format: hh:mm:ss
	// Example: 08:30:00, Snapshot is set for 08:30 AM.
	//
	//
	// Read Only: true
	SnapshotTime string `json:"SnapshotTime,omitempty"`

	// Storage array managed object.
	//
	// Read Only: true
	StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageSnapshotSchedule) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Frequency string `json:"Frequency,omitempty"`

		Name string `json:"Name,omitempty"`

		ProtectionGroup *StorageProtectionGroupRef `json:"ProtectionGroup,omitempty"`

		RetentionTime string `json:"RetentionTime,omitempty"`

		SnapshotTime string `json:"SnapshotTime,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Frequency = dataAO1.Frequency

	m.Name = dataAO1.Name

	m.ProtectionGroup = dataAO1.ProtectionGroup

	m.RetentionTime = dataAO1.RetentionTime

	m.SnapshotTime = dataAO1.SnapshotTime

	m.StorageArray = dataAO1.StorageArray

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageSnapshotSchedule) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Frequency string `json:"Frequency,omitempty"`

		Name string `json:"Name,omitempty"`

		ProtectionGroup *StorageProtectionGroupRef `json:"ProtectionGroup,omitempty"`

		RetentionTime string `json:"RetentionTime,omitempty"`

		SnapshotTime string `json:"SnapshotTime,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`
	}

	dataAO1.Frequency = m.Frequency

	dataAO1.Name = m.Name

	dataAO1.ProtectionGroup = m.ProtectionGroup

	dataAO1.RetentionTime = m.RetentionTime

	dataAO1.SnapshotTime = m.SnapshotTime

	dataAO1.StorageArray = m.StorageArray

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage snapshot schedule
func (m *StorageSnapshotSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageSnapshotSchedule) validateProtectionGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtectionGroup) { // not required
		return nil
	}

	if m.ProtectionGroup != nil {
		if err := m.ProtectionGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProtectionGroup")
			}
			return err
		}
	}

	return nil
}

func (m *StorageSnapshotSchedule) validateStorageArray(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageArray) { // not required
		return nil
	}

	if m.StorageArray != nil {
		if err := m.StorageArray.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageArray")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageSnapshotSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSnapshotSchedule) UnmarshalBinary(b []byte) error {
	var res StorageSnapshotSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
