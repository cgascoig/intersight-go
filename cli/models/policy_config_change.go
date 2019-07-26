// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PolicyConfigChange Policy:Config Change
//
// Defines the configuration changes at the summary level including configuration changes and disruptions.
//
// swagger:model policyConfigChange
type PolicyConfigChange struct {

	// Configuration changes at summary level.
	//
	Changes []string `json:"Changes"`

	// Configuration disruptions.
	//
	Disruptions []string `json:"Disruptions"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyConfigChange) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Changes []string `json:"Changes"`

		Disruptions []string `json:"Disruptions"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Changes = dataAO0.Changes

	m.Disruptions = dataAO0.Disruptions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyConfigChange) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Changes []string `json:"Changes"`

		Disruptions []string `json:"Disruptions"`
	}

	dataAO0.Changes = m.Changes

	dataAO0.Disruptions = m.Disruptions

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy config change
func (m *PolicyConfigChange) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConfigChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConfigChange) UnmarshalBinary(b []byte) error {
	var res PolicyConfigChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
