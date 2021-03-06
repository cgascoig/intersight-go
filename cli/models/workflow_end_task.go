// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowEndTask Workflow:End Task
//
// An EndTask denotes the completion of a workflow.
//
// swagger:model workflowEndTask
type WorkflowEndTask struct {
	WorkflowControlTask

	WorkflowEndTaskAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowEndTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowControlTask
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowControlTask = aO0

	// AO1
	var aO1 WorkflowEndTaskAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowEndTaskAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowEndTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowControlTask)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowEndTaskAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow end task
func (m *WorkflowEndTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowControlTask
	if err := m.WorkflowControlTask.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowEndTaskAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowEndTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowEndTask) UnmarshalBinary(b []byte) error {
	var res WorkflowEndTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowEndTaskAllOf1 workflow end task all of1
// swagger:model WorkflowEndTaskAllOf1
type WorkflowEndTaskAllOf1 interface{}
