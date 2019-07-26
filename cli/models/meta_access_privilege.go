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

// MetaAccessPrivilege Meta:Access Privilege
//
// The required access privilege for a given Managed Object and CRUD operation.
//
// swagger:model metaAccessPrivilege
type MetaAccessPrivilege struct {

	// The type of CRUD operation (create, read, update, delete) for which an access privilege is required.
	//
	// Read Only: true
	// Enum: [Update Create Read Delete]
	Method string `json:"Method,omitempty"`

	// The name of the privilege which is required to invoke the specified CRUD method.
	//
	// Read Only: true
	Privilege string `json:"Privilege,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MetaAccessPrivilege) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Method string `json:"Method,omitempty"`

		Privilege string `json:"Privilege,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Method = dataAO0.Method

	m.Privilege = dataAO0.Privilege

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MetaAccessPrivilege) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Method string `json:"Method,omitempty"`

		Privilege string `json:"Privilege,omitempty"`
	}

	dataAO0.Method = m.Method

	dataAO0.Privilege = m.Privilege

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this meta access privilege
func (m *MetaAccessPrivilege) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metaAccessPrivilegeTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Update","Create","Read","Delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metaAccessPrivilegeTypeMethodPropEnum = append(metaAccessPrivilegeTypeMethodPropEnum, v)
	}
}

// property enum
func (m *MetaAccessPrivilege) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metaAccessPrivilegeTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetaAccessPrivilege) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("Method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetaAccessPrivilege) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetaAccessPrivilege) UnmarshalBinary(b []byte) error {
	var res MetaAccessPrivilege
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}