// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IamLdapProvider Iam:Ldap Provider
//
// LDAP Provider or LDAP Server for user authentication.
//
// swagger:model iamLdapProvider
type IamLdapProvider struct {
	MoBaseMo

	// A collection of references to the [iam.LdapPolicy](mo://iam.LdapPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [iam.LdapPolicy](mo://iam.LdapPolicy) MO unsets its reference to this deleted MO.
	//
	LdapPolicy *IamLdapPolicyRef `json:"LdapPolicy,omitempty"`

	// LDAP Server Port for connection establishment.
	//
	Port int64 `json:"Port,omitempty"`

	// LDAP Server Address, can be IP address or hostname.
	//
	Server string `json:"Server,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamLdapProvider) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		LdapPolicy *IamLdapPolicyRef `json:"LdapPolicy,omitempty"`

		Port int64 `json:"Port,omitempty"`

		Server string `json:"Server,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LdapPolicy = dataAO1.LdapPolicy

	m.Port = dataAO1.Port

	m.Server = dataAO1.Server

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamLdapProvider) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		LdapPolicy *IamLdapPolicyRef `json:"LdapPolicy,omitempty"`

		Port int64 `json:"Port,omitempty"`

		Server string `json:"Server,omitempty"`
	}

	dataAO1.LdapPolicy = m.LdapPolicy

	dataAO1.Port = m.Port

	dataAO1.Server = m.Server

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam ldap provider
func (m *IamLdapProvider) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdapPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamLdapProvider) validateLdapPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.LdapPolicy) { // not required
		return nil
	}

	if m.LdapPolicy != nil {
		if err := m.LdapPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LdapPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamLdapProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLdapProvider) UnmarshalBinary(b []byte) error {
	var res IamLdapProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
