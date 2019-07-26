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

// StorageSnapshot Storage:Snapshot
//
// Reference marker for volume at a particular point in time. It is a point-in-time copy of the volume.
//
// swagger:model storageSnapshot
type StorageSnapshot struct {
	MoBaseMo

	// Exact date and time on which snapshot is created.
	//
	// Read Only: true
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"CreatedTime,omitempty"`

	// Name of the snapshot which represents point in time copy of volume.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Name of the protection group to which the snapshot belongs. Value is emptry, if the snapshot is created directly on volume.
	//
	// Read Only: true
	ProtectionGroupName string `json:"ProtectionGroupName,omitempty"`

	// Protection group snapshot relationship object.
	//
	// Read Only: true
	ProtectionGroupSnapshot *StorageProtectionGroupSnapshotRef `json:"ProtectionGroupSnapshot,omitempty"`

	// Snapshot size represented in bytes.
	//
	// Read Only: true
	Size int64 `json:"Size,omitempty"`

	// Source object on which the snapshot is created. It is a name of the originating volume.
	//
	// Read Only: true
	Source string `json:"Source,omitempty"`

	// Storage array managed object.
	//
	// Read Only: true
	StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`

	// Volume managed object where the snapshot is created.
	//
	// Read Only: true
	Volume *StorageVolumeRef `json:"Volume,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageSnapshot) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		CreatedTime strfmt.DateTime `json:"CreatedTime,omitempty"`

		Name string `json:"Name,omitempty"`

		ProtectionGroupName string `json:"ProtectionGroupName,omitempty"`

		ProtectionGroupSnapshot *StorageProtectionGroupSnapshotRef `json:"ProtectionGroupSnapshot,omitempty"`

		Size int64 `json:"Size,omitempty"`

		Source string `json:"Source,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`

		Volume *StorageVolumeRef `json:"Volume,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CreatedTime = dataAO1.CreatedTime

	m.Name = dataAO1.Name

	m.ProtectionGroupName = dataAO1.ProtectionGroupName

	m.ProtectionGroupSnapshot = dataAO1.ProtectionGroupSnapshot

	m.Size = dataAO1.Size

	m.Source = dataAO1.Source

	m.StorageArray = dataAO1.StorageArray

	m.Volume = dataAO1.Volume

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageSnapshot) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		CreatedTime strfmt.DateTime `json:"CreatedTime,omitempty"`

		Name string `json:"Name,omitempty"`

		ProtectionGroupName string `json:"ProtectionGroupName,omitempty"`

		ProtectionGroupSnapshot *StorageProtectionGroupSnapshotRef `json:"ProtectionGroupSnapshot,omitempty"`

		Size int64 `json:"Size,omitempty"`

		Source string `json:"Source,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`

		Volume *StorageVolumeRef `json:"Volume,omitempty"`
	}

	dataAO1.CreatedTime = m.CreatedTime

	dataAO1.Name = m.Name

	dataAO1.ProtectionGroupName = m.ProtectionGroupName

	dataAO1.ProtectionGroupSnapshot = m.ProtectionGroupSnapshot

	dataAO1.Size = m.Size

	dataAO1.Source = m.Source

	dataAO1.StorageArray = m.StorageArray

	dataAO1.Volume = m.Volume

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage snapshot
func (m *StorageSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionGroupSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArray(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageSnapshot) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorageSnapshot) validateProtectionGroupSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtectionGroupSnapshot) { // not required
		return nil
	}

	if m.ProtectionGroupSnapshot != nil {
		if err := m.ProtectionGroupSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProtectionGroupSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *StorageSnapshot) validateStorageArray(formats strfmt.Registry) error {

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

func (m *StorageSnapshot) validateVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSnapshot) UnmarshalBinary(b []byte) error {
	var res StorageSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}