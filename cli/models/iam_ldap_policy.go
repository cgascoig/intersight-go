// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IamLdapPolicy LDAP
//
// LDAP Policy configurations.
//
// swagger:model iamLdapPolicy
type IamLdapPolicy struct {
	PolicyAbstractPolicy

	// Base settings of LDAP required while configuring LDAP policy.
	//
	BaseProperties *IamLdapBaseProperties `json:"BaseProperties,omitempty"`

	// Configuration settings to resolve LDAP servers, when DNS is enabled.
	//
	DNSParameters *IamLdapDNSParameters `json:"DnsParameters,omitempty"`

	// Enables DNS to access LDAP servers.
	//
	EnableDNS bool `json:"EnableDns,omitempty"`

	// LDAP server performs authentication.
	//
	Enabled bool `json:"Enabled,omitempty"`

	// Relationship to collection of LDAP Groups.
	//
	Groups []*IamLdapGroupRef `json:"Groups"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Relationship to the profile object.
	//
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

	// Relationship to collection of LDAP Providers.
	//
	Providers []*IamLdapProviderRef `json:"Providers"`

	// Search precedence between local user database and LDAP user database.
	//
	// Enum: [LocalUserDb LDAPUserDb]
	UserSearchPrecedence *string `json:"UserSearchPrecedence,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamLdapPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		BaseProperties *IamLdapBaseProperties `json:"BaseProperties,omitempty"`

		DNSParameters *IamLdapDNSParameters `json:"DnsParameters,omitempty"`

		EnableDNS bool `json:"EnableDns,omitempty"`

		Enabled bool `json:"Enabled,omitempty"`

		Groups []*IamLdapGroupRef `json:"Groups"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		Providers []*IamLdapProviderRef `json:"Providers"`

		UserSearchPrecedence *string `json:"UserSearchPrecedence,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BaseProperties = dataAO1.BaseProperties

	m.DNSParameters = dataAO1.DNSParameters

	m.EnableDNS = dataAO1.EnableDNS

	m.Enabled = dataAO1.Enabled

	m.Groups = dataAO1.Groups

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	m.Providers = dataAO1.Providers

	m.UserSearchPrecedence = dataAO1.UserSearchPrecedence

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamLdapPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BaseProperties *IamLdapBaseProperties `json:"BaseProperties,omitempty"`

		DNSParameters *IamLdapDNSParameters `json:"DnsParameters,omitempty"`

		EnableDNS bool `json:"EnableDns,omitempty"`

		Enabled bool `json:"Enabled,omitempty"`

		Groups []*IamLdapGroupRef `json:"Groups"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		Providers []*IamLdapProviderRef `json:"Providers"`

		UserSearchPrecedence *string `json:"UserSearchPrecedence,omitempty"`
	}

	dataAO1.BaseProperties = m.BaseProperties

	dataAO1.DNSParameters = m.DNSParameters

	dataAO1.EnableDNS = m.EnableDNS

	dataAO1.Enabled = m.Enabled

	dataAO1.Groups = m.Groups

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	dataAO1.Providers = m.Providers

	dataAO1.UserSearchPrecedence = m.UserSearchPrecedence

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam ldap policy
func (m *IamLdapPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchPrecedence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamLdapPolicy) validateBaseProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseProperties) { // not required
		return nil
	}

	if m.BaseProperties != nil {
		if err := m.BaseProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseProperties")
			}
			return err
		}
	}

	return nil
}

func (m *IamLdapPolicy) validateDNSParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSParameters) { // not required
		return nil
	}

	if m.DNSParameters != nil {
		if err := m.DNSParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DnsParameters")
			}
			return err
		}
	}

	return nil
}

func (m *IamLdapPolicy) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamLdapPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *IamLdapPolicy) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamLdapPolicy) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {
		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {
			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var iamLdapPolicyTypeUserSearchPrecedencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LocalUserDb","LDAPUserDb"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iamLdapPolicyTypeUserSearchPrecedencePropEnum = append(iamLdapPolicyTypeUserSearchPrecedencePropEnum, v)
	}
}

// property enum
func (m *IamLdapPolicy) validateUserSearchPrecedenceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iamLdapPolicyTypeUserSearchPrecedencePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IamLdapPolicy) validateUserSearchPrecedence(formats strfmt.Registry) error {

	if swag.IsZero(m.UserSearchPrecedence) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserSearchPrecedenceEnum("UserSearchPrecedence", "body", *m.UserSearchPrecedence); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamLdapPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLdapPolicy) UnmarshalBinary(b []byte) error {
	var res IamLdapPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
