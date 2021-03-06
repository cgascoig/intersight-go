// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowDecisionCase Workflow:Decision Case
//
// A Decision Case is a condition for a Decision Task. It is the equivalent of case statement in the programming language paradigm of switch and case.
//
// swagger:model workflowDecisionCase
type WorkflowDecisionCase struct {

	// Description of this decision case.
	//
	Description string `json:"Description,omitempty"`

	// The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node.
	//
	NextTask string `json:"NextTask,omitempty"`

	// Value for the decision case.
	//
	Value string `json:"Value,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowDecisionCase) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Description string `json:"Description,omitempty"`

		NextTask string `json:"NextTask,omitempty"`

		Value string `json:"Value,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Description = dataAO0.Description

	m.NextTask = dataAO0.NextTask

	m.Value = dataAO0.Value

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowDecisionCase) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Description string `json:"Description,omitempty"`

		NextTask string `json:"NextTask,omitempty"`

		Value string `json:"Value,omitempty"`
	}

	dataAO0.Description = m.Description

	dataAO0.NextTask = m.NextTask

	dataAO0.Value = m.Value

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow decision case
func (m *WorkflowDecisionCase) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowDecisionCase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowDecisionCase) UnmarshalBinary(b []byte) error {
	var res WorkflowDecisionCase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
