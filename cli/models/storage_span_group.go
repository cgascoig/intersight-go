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

// StorageSpanGroup Storage:Span Group
//
// This models a single span group of disks in a RAID group.
//
// swagger:model storageSpanGroup
type StorageSpanGroup struct {

	// Collection of local disks that are part of this span group. The minimum number of disks needed in a span group varies based on RAID level. Raid0 requires at least one disk, Raid1 and Raid10 requires at least 2 and in multiples of 2, Raid5 Raid50 Raid6 and Raid60 require at least 3 disks in a span group.
	//
	Disks []*StorageLocalDisk `json:"Disks"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageSpanGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Disks []*StorageLocalDisk `json:"Disks"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Disks = dataAO0.Disks

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageSpanGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Disks []*StorageLocalDisk `json:"Disks"`
	}

	dataAO0.Disks = m.Disks

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage span group
func (m *StorageSpanGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageSpanGroup) validateDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageSpanGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSpanGroup) UnmarshalBinary(b []byte) error {
	var res StorageSpanGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
