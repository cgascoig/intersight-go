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

// WorkflowConstraints Workflow:Constraints
//
// Captures the constraints for valid parameter values.
//
// swagger:model workflowConstraints
type WorkflowConstraints struct {

	// When the parameter is a enum then this list of enum entry is used to validate the input belongs to one of items in the list.
	//
	EnumList []*WorkflowEnumEntry `json:"EnumList"`

	// Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked.
	//
	//
	Max float64 `json:"Max,omitempty"`

	// Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked.
	//
	//
	Min float64 `json:"Min,omitempty"`

	// When the parameter is a string this regular expression is used to ensure the value is valid.
	//
	Regex string `json:"Regex,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowConstraints) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		EnumList []*WorkflowEnumEntry `json:"EnumList"`

		Max float64 `json:"Max,omitempty"`

		Min float64 `json:"Min,omitempty"`

		Regex string `json:"Regex,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.EnumList = dataAO0.EnumList

	m.Max = dataAO0.Max

	m.Min = dataAO0.Min

	m.Regex = dataAO0.Regex

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowConstraints) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		EnumList []*WorkflowEnumEntry `json:"EnumList"`

		Max float64 `json:"Max,omitempty"`

		Min float64 `json:"Min,omitempty"`

		Regex string `json:"Regex,omitempty"`
	}

	dataAO0.EnumList = m.EnumList

	dataAO0.Max = m.Max

	dataAO0.Min = m.Min

	dataAO0.Regex = m.Regex

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow constraints
func (m *WorkflowConstraints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnumList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowConstraints) validateEnumList(formats strfmt.Registry) error {

	if swag.IsZero(m.EnumList) { // not required
		return nil
	}

	for i := 0; i < len(m.EnumList); i++ {
		if swag.IsZero(m.EnumList[i]) { // not required
			continue
		}

		if m.EnumList[i] != nil {
			if err := m.EnumList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EnumList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowConstraints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowConstraints) UnmarshalBinary(b []byte) error {
	var res WorkflowConstraints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
