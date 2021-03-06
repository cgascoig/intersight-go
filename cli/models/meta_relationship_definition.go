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

// MetaRelationshipDefinition Meta:Relationship Definition
//
// Definitions for the relationship in a meta.
//
// swagger:model metaRelationshipDefinition
type MetaRelationshipDefinition struct {

	// API access definition for this relationship.
	//
	// Read Only: true
	// Enum: [NoAccess ReadOnly CreateOnly ReadWrite WriteOnly ReadOnCreate]
	APIAccess string `json:"ApiAccess,omitempty"`

	// Specifies whether the relationship is a collection.
	//
	// Read Only: true
	Collection *bool `json:"Collection,omitempty"`

	// The name of the relationship.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Fully qualified type of the foreign managed object.
	//
	// Read Only: true
	Type string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MetaRelationshipDefinition) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		APIAccess string `json:"ApiAccess,omitempty"`

		Collection *bool `json:"Collection,omitempty"`

		Name string `json:"Name,omitempty"`

		Type string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.APIAccess = dataAO0.APIAccess

	m.Collection = dataAO0.Collection

	m.Name = dataAO0.Name

	m.Type = dataAO0.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MetaRelationshipDefinition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		APIAccess string `json:"ApiAccess,omitempty"`

		Collection *bool `json:"Collection,omitempty"`

		Name string `json:"Name,omitempty"`

		Type string `json:"Type,omitempty"`
	}

	dataAO0.APIAccess = m.APIAccess

	dataAO0.Collection = m.Collection

	dataAO0.Name = m.Name

	dataAO0.Type = m.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this meta relationship definition
func (m *MetaRelationshipDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metaRelationshipDefinitionTypeAPIAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NoAccess","ReadOnly","CreateOnly","ReadWrite","WriteOnly","ReadOnCreate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metaRelationshipDefinitionTypeAPIAccessPropEnum = append(metaRelationshipDefinitionTypeAPIAccessPropEnum, v)
	}
}

// property enum
func (m *MetaRelationshipDefinition) validateAPIAccessEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metaRelationshipDefinitionTypeAPIAccessPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetaRelationshipDefinition) validateAPIAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.APIAccess) { // not required
		return nil
	}

	// value enum
	if err := m.validateAPIAccessEnum("ApiAccess", "body", m.APIAccess); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetaRelationshipDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetaRelationshipDefinition) UnmarshalBinary(b []byte) error {
	var res MetaRelationshipDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
