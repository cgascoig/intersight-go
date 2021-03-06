// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexNamedVsan Hyperflex:Named Vsan
//
// A VSAN with a name and ID.
//
// VSANs are used when defining Fibre Channel external storage policies for the cluster.
//
// swagger:model hyperflexNamedVsan
type HyperflexNamedVsan struct {

	// The name of the VSAN.
	//
	// The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
	//
	//
	Name string `json:"Name,omitempty"`

	// The ID of the named VSAN.
	//
	// The ID can be any number between 1 and 4093, inclusive.
	//
	//
	VsanID int64 `json:"VsanId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexNamedVsan) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name string `json:"Name,omitempty"`

		VsanID int64 `json:"VsanId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	m.VsanID = dataAO0.VsanID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexNamedVsan) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Name string `json:"Name,omitempty"`

		VsanID int64 `json:"VsanId,omitempty"`
	}

	dataAO0.Name = m.Name

	dataAO0.VsanID = m.VsanID

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex named vsan
func (m *HyperflexNamedVsan) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexNamedVsan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexNamedVsan) UnmarshalBinary(b []byte) error {
	var res HyperflexNamedVsan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
