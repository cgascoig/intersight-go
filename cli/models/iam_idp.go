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

// IamIdp Iam:Idp
//
// The SAML identity provider such as Cisco, that has been used to log in to Intersight.
//
// swagger:model iamIdp
type IamIdp struct {
	MoBaseMo

	// A collection of references to the [iam.Account](mo://iam.Account) Managed Object.
	//
	// When this managed object is deleted, the referenced [iam.Account](mo://iam.Account) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	//
	DomainName string `json:"DomainName,omitempty"`

	// The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider.
	//
	// Read Only: true
	IdpEntityID string `json:"IdpEntityId,omitempty"`

	// SAML metadata of the IdP.
	//
	Metadata string `json:"Metadata,omitempty"`

	// The name of the Identity Provider, for example Cisco, Okta, or OneID.
	//
	Name string `json:"Name,omitempty"`

	// A collection of references to the [iam.System](mo://iam.System) Managed Object.
	//
	// When this managed object is deleted, the referenced [iam.System](mo://iam.System) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	System *IamSystemRef `json:"System,omitempty"`

	// Authentication protocol used by the IdP.
	//
	// Enum: [saml oidc local]
	Type *string `json:"Type,omitempty"`

	// The last login session details for each logged in user of this IdP.
	//
	// Read Only: true
	UserLoginTime []*IamUserLoginTimeRef `json:"UserLoginTime"`

	// The UI preference object for each user logged in through this IdP.
	//
	// Read Only: true
	UserPreferences []*IamUserPreferenceRef `json:"UserPreferences"`

	// User groups added in an IdP. User group provides a way to configure permission assignment for a group of users based on the IdP attributes received after authentication.
	//
	Usergroups []*IamUserGroupRef `json:"Usergroups"`

	// Added or logged in users of an IdP who can access an Intersight account.
	//
	Users []*IamUserRef `json:"Users"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamIdp) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		DomainName string `json:"DomainName,omitempty"`

		IdpEntityID string `json:"IdpEntityId,omitempty"`

		Metadata string `json:"Metadata,omitempty"`

		Name string `json:"Name,omitempty"`

		System *IamSystemRef `json:"System,omitempty"`

		Type *string `json:"Type,omitempty"`

		UserLoginTime []*IamUserLoginTimeRef `json:"UserLoginTime"`

		UserPreferences []*IamUserPreferenceRef `json:"UserPreferences"`

		Usergroups []*IamUserGroupRef `json:"Usergroups"`

		Users []*IamUserRef `json:"Users"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.DomainName = dataAO1.DomainName

	m.IdpEntityID = dataAO1.IdpEntityID

	m.Metadata = dataAO1.Metadata

	m.Name = dataAO1.Name

	m.System = dataAO1.System

	m.Type = dataAO1.Type

	m.UserLoginTime = dataAO1.UserLoginTime

	m.UserPreferences = dataAO1.UserPreferences

	m.Usergroups = dataAO1.Usergroups

	m.Users = dataAO1.Users

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamIdp) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		DomainName string `json:"DomainName,omitempty"`

		IdpEntityID string `json:"IdpEntityId,omitempty"`

		Metadata string `json:"Metadata,omitempty"`

		Name string `json:"Name,omitempty"`

		System *IamSystemRef `json:"System,omitempty"`

		Type *string `json:"Type,omitempty"`

		UserLoginTime []*IamUserLoginTimeRef `json:"UserLoginTime"`

		UserPreferences []*IamUserPreferenceRef `json:"UserPreferences"`

		Usergroups []*IamUserGroupRef `json:"Usergroups"`

		Users []*IamUserRef `json:"Users"`
	}

	dataAO1.Account = m.Account

	dataAO1.DomainName = m.DomainName

	dataAO1.IdpEntityID = m.IdpEntityID

	dataAO1.Metadata = m.Metadata

	dataAO1.Name = m.Name

	dataAO1.System = m.System

	dataAO1.Type = m.Type

	dataAO1.UserLoginTime = m.UserLoginTime

	dataAO1.UserPreferences = m.UserPreferences

	dataAO1.Usergroups = m.Usergroups

	dataAO1.Users = m.Users

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam idp
func (m *IamIdp) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserLoginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsergroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamIdp) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

func (m *IamIdp) validateSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.System) { // not required
		return nil
	}

	if m.System != nil {
		if err := m.System.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("System")
			}
			return err
		}
	}

	return nil
}

var iamIdpTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["saml","oidc","local"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iamIdpTypeTypePropEnum = append(iamIdpTypeTypePropEnum, v)
	}
}

// property enum
func (m *IamIdp) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iamIdpTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IamIdp) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *IamIdp) validateUserLoginTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UserLoginTime) { // not required
		return nil
	}

	for i := 0; i < len(m.UserLoginTime); i++ {
		if swag.IsZero(m.UserLoginTime[i]) { // not required
			continue
		}

		if m.UserLoginTime[i] != nil {
			if err := m.UserLoginTime[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UserLoginTime" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamIdp) validateUserPreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.UserPreferences) { // not required
		return nil
	}

	for i := 0; i < len(m.UserPreferences); i++ {
		if swag.IsZero(m.UserPreferences[i]) { // not required
			continue
		}

		if m.UserPreferences[i] != nil {
			if err := m.UserPreferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UserPreferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamIdp) validateUsergroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Usergroups) { // not required
		return nil
	}

	for i := 0; i < len(m.Usergroups); i++ {
		if swag.IsZero(m.Usergroups[i]) { // not required
			continue
		}

		if m.Usergroups[i] != nil {
			if err := m.Usergroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Usergroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamIdp) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamIdp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamIdp) UnmarshalBinary(b []byte) error {
	var res IamIdp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
