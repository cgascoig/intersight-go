// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetaPropDefinition Meta:Prop Definition
//
// Definitions for the properties in a meta.
//
// swagger:model metaPropDefinition
type MetaPropDefinition struct {

	// API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc.
	//
	// Read Only: true
	// Enum: [NoAccess ReadOnly CreateOnly ReadWrite WriteOnly ReadOnCreate]
	APIAccess string `json:"ApiAccess,omitempty"`

	// The name of the property.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// The data-at-rest security protection applied to this property when a Managed Object is persisted.
	//
	// For example, Encrypted or Cleartext.
	//
	//
	// Read Only: true
	// Enum: [ClearText Encrypted]
	OpSecurity string `json:"OpSecurity,omitempty"`

	// Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable.
	//
	// Read Only: true
	SearchWeight float32 `json:"SearchWeight,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MetaPropDefinition) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		APIAccess string `json:"ApiAccess,omitempty"`

		Name string `json:"Name,omitempty"`

		OpSecurity string `json:"OpSecurity,omitempty"`

		SearchWeight float32 `json:"SearchWeight,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.APIAccess = dataAO0.APIAccess

	m.Name = dataAO0.Name

	m.OpSecurity = dataAO0.OpSecurity

	m.SearchWeight = dataAO0.SearchWeight

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MetaPropDefinition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		APIAccess string `json:"ApiAccess,omitempty"`

		Name string `json:"Name,omitempty"`

		OpSecurity string `json:"OpSecurity,omitempty"`

		SearchWeight float32 `json:"SearchWeight,omitempty"`
	}

	dataAO0.APIAccess = m.APIAccess

	dataAO0.Name = m.Name

	dataAO0.OpSecurity = m.OpSecurity

	dataAO0.SearchWeight = m.SearchWeight

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this meta prop definition
func (m *MetaPropDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpSecurity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metaPropDefinitionTypeAPIAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NoAccess","ReadOnly","CreateOnly","ReadWrite","WriteOnly","ReadOnCreate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metaPropDefinitionTypeAPIAccessPropEnum = append(metaPropDefinitionTypeAPIAccessPropEnum, v)
	}
}

// property enum
func (m *MetaPropDefinition) validateAPIAccessEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metaPropDefinitionTypeAPIAccessPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetaPropDefinition) validateAPIAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.APIAccess) { // not required
		return nil
	}

	// value enum
	if err := m.validateAPIAccessEnum("ApiAccess", "body", m.APIAccess); err != nil {
		return err
	}

	return nil
}

var metaPropDefinitionTypeOpSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ClearText","Encrypted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metaPropDefinitionTypeOpSecurityPropEnum = append(metaPropDefinitionTypeOpSecurityPropEnum, v)
	}
}

// property enum
func (m *MetaPropDefinition) validateOpSecurityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metaPropDefinitionTypeOpSecurityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetaPropDefinition) validateOpSecurity(formats strfmt.Registry) error {

	if swag.IsZero(m.OpSecurity) { // not required
		return nil
	}

	// value enum
	if err := m.validateOpSecurityEnum("OpSecurity", "body", m.OpSecurity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetaPropDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetaPropDefinition) UnmarshalBinary(b []byte) error {
	var res MetaPropDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
