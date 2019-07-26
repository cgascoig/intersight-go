// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageLocalDisk Storage:Local Disk
//
// This models disks in the disk group.
//
// swagger:model storageLocalDisk
type StorageLocalDisk struct {

	// The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server.
	//
	SlotNumber int64 `json:"SlotNumber,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageLocalDisk) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		SlotNumber int64 `json:"SlotNumber,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.SlotNumber = dataAO0.SlotNumber

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageLocalDisk) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		SlotNumber int64 `json:"SlotNumber,omitempty"`
	}

	dataAO0.SlotNumber = m.SlotNumber

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage local disk
func (m *StorageLocalDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageLocalDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageLocalDisk) UnmarshalBinary(b []byte) error {
	var res StorageLocalDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
