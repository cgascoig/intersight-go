// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageHostUtilization Storage:Host Utilization
//
// Storage space utilization of Pure FlashArray host entity.
//
// swagger:model storageHostUtilization
type StorageHostUtilization struct {
	StorageStorageUtilization

	StorageHostUtilizationAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageHostUtilization) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageStorageUtilization
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageStorageUtilization = aO0

	// AO1
	var aO1 StorageHostUtilizationAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.StorageHostUtilizationAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageHostUtilization) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageStorageUtilization)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.StorageHostUtilizationAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage host utilization
func (m *StorageHostUtilization) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageStorageUtilization
	if err := m.StorageStorageUtilization.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with StorageHostUtilizationAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageHostUtilization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageHostUtilization) UnmarshalBinary(b []byte) error {
	var res StorageHostUtilization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StorageHostUtilizationAllOf1 storage host utilization all of1
// swagger:model StorageHostUtilizationAllOf1
type StorageHostUtilizationAllOf1 interface{}
