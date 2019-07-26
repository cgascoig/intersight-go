// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerConfigResult Server:Config Result
//
// The profile configuration (deploy, validation) results with the overall state and detailed result messages.
//
// swagger:model serverConfigResult
type ServerConfigResult struct {
	PolicyAbstractConfigResult

	// A collection of references to the [server.Profile](mo://server.Profile) Managed Object.
	//
	// When this managed object is deleted, the referenced [server.Profile](mo://server.Profile) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	Profile *ServerProfileRef `json:"Profile,omitempty"`

	// Detailed result entries for both validation & configration. Each result entry can be error/warning/info messages and the context.
	//
	ResultEntries []*ServerConfigResultEntryRef `json:"ResultEntries"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServerConfigResult) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractConfigResult
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractConfigResult = aO0

	// AO1
	var dataAO1 struct {
		Profile *ServerProfileRef `json:"Profile,omitempty"`

		ResultEntries []*ServerConfigResultEntryRef `json:"ResultEntries"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Profile = dataAO1.Profile

	m.ResultEntries = dataAO1.ResultEntries

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServerConfigResult) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractConfigResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Profile *ServerProfileRef `json:"Profile,omitempty"`

		ResultEntries []*ServerConfigResultEntryRef `json:"ResultEntries"`
	}

	dataAO1.Profile = m.Profile

	dataAO1.ResultEntries = m.ResultEntries

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this server config result
func (m *ServerConfigResult) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractConfigResult
	if err := m.PolicyAbstractConfigResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerConfigResult) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Profile")
			}
			return err
		}
	}

	return nil
}

func (m *ServerConfigResult) validateResultEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.ResultEntries); i++ {
		if swag.IsZero(m.ResultEntries[i]) { // not required
			continue
		}

		if m.ResultEntries[i] != nil {
			if err := m.ResultEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ResultEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerConfigResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerConfigResult) UnmarshalBinary(b []byte) error {
	var res ServerConfigResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
