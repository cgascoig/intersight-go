// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageCapacity Storage:Capacity
//
// Storage capacity information which includes, total capacity, available capacity, used capacity and free capacity.
//
// swagger:model storageCapacity
type StorageCapacity struct {

	// Total consumable storage capacity represented in bytes. System may reserve some space for internal purpose which is excluded from total capacity.
	//
	// Read Only: true
	Available int64 `json:"Available,omitempty"`

	// Unused space available for user to consume, represented in bytes.
	//
	// Read Only: true
	Free int64 `json:"Free,omitempty"`

	// Total storage capacity, represented in bytes. It is set by the component manufacture.
	//
	// Read Only: true
	Total int64 `json:"Total,omitempty"`

	// Used or consumed storage capacity, represented in bytes.
	//
	// Read Only: true
	Used int64 `json:"Used,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageCapacity) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Available int64 `json:"Available,omitempty"`

		Free int64 `json:"Free,omitempty"`

		Total int64 `json:"Total,omitempty"`

		Used int64 `json:"Used,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Available = dataAO0.Available

	m.Free = dataAO0.Free

	m.Total = dataAO0.Total

	m.Used = dataAO0.Used

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageCapacity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Available int64 `json:"Available,omitempty"`

		Free int64 `json:"Free,omitempty"`

		Total int64 `json:"Total,omitempty"`

		Used int64 `json:"Used,omitempty"`
	}

	dataAO0.Available = m.Available

	dataAO0.Free = m.Free

	dataAO0.Total = m.Total

	dataAO0.Used = m.Used

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage capacity
func (m *StorageCapacity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCapacity) UnmarshalBinary(b []byte) error {
	var res StorageCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
