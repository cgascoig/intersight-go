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

// ContentGrammar Grammar Specification
//
// Content handler framework supports extraction of required values from API/device
// responses. These responses may be of various content types such as XML, JSON, etc.
// The values of importance are modeled as parameters in the content handler framework.
//
// The parameters can be of a scalar value type or a collection of values. A group
// of related parameters can be modeled as a single complex type parameter. These
// complex types will be very useful to extract a set of repeating group of related
// parameters.
//
// A grammar specification defines the set of parameters that need to be extracted
// from the content. The grammar specification allows complex type definitions to be
// defined for any complex parameters.
//
// swagger:model contentGrammar
type ContentGrammar struct {

	// The list of parameter definitions, if found in a given API/device response,
	// makes the content handlers to treat the response as error response.
	//
	// This is optional parameter.
	//
	//
	ErrorParameters []*ContentParameter `json:"ErrorParameters"`

	// The list of parameter definitions that mark the parameters to be
	// extracted using this grammar specification.
	//
	//
	Parameters []*ContentParameter `json:"Parameters"`

	// The collection of complex types definitions used in this grammar
	// specification.
	//
	// This is required only if any of the parameters provided in this grammar
	// is of complex type.
	//
	//
	Types []*ContentComplexType `json:"Types"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ContentGrammar) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ErrorParameters []*ContentParameter `json:"ErrorParameters"`

		Parameters []*ContentParameter `json:"Parameters"`

		Types []*ContentComplexType `json:"Types"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ErrorParameters = dataAO0.ErrorParameters

	m.Parameters = dataAO0.Parameters

	m.Types = dataAO0.Types

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ContentGrammar) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		ErrorParameters []*ContentParameter `json:"ErrorParameters"`

		Parameters []*ContentParameter `json:"Parameters"`

		Types []*ContentComplexType `json:"Types"`
	}

	dataAO0.ErrorParameters = m.ErrorParameters

	dataAO0.Parameters = m.Parameters

	dataAO0.Types = m.Types

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this content grammar
func (m *ContentGrammar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentGrammar) validateErrorParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.ErrorParameters); i++ {
		if swag.IsZero(m.ErrorParameters[i]) { // not required
			continue
		}

		if m.ErrorParameters[i] != nil {
			if err := m.ErrorParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ErrorParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentGrammar) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentGrammar) validateTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.Types) { // not required
		return nil
	}

	for i := 0; i < len(m.Types); i++ {
		if swag.IsZero(m.Types[i]) { // not required
			continue
		}

		if m.Types[i] != nil {
			if err := m.Types[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentGrammar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentGrammar) UnmarshalBinary(b []byte) error {
	var res ContentGrammar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
