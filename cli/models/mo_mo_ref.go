// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MoMoRef Mo:Mo Ref
//
// A reference to a REST resource, uniquely identified by object type and Moid.
//
// swagger:model moMoRef
type MoMoRef struct {

	// The Moid of the referenced REST resource.
	//
	// Read Only: true
	Moid string `json:"Moid,omitempty"`

	// The Object Type of the referenced REST resource.
	//
	// Read Only: true
	ObjectType string `json:"ObjectType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MoMoRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Moid string `json:"Moid,omitempty"`

		ObjectType string `json:"ObjectType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Moid = dataAO0.Moid

	m.ObjectType = dataAO0.ObjectType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MoMoRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Moid string `json:"Moid,omitempty"`

		ObjectType string `json:"ObjectType,omitempty"`
	}

	dataAO0.Moid = m.Moid

	dataAO0.ObjectType = m.ObjectType

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this mo mo ref
func (m *MoMoRef) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MoMoRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoMoRef) UnmarshalBinary(b []byte) error {
	var res MoMoRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}