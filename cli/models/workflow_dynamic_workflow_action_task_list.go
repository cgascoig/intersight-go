// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowDynamicWorkflowActionTaskList Workflow:Dynamic Workflow Action Task List
//
// Keeps a list of the tasks to add for a corresponding workflow action when PendingDynamicWorkflowInfo is built.
//
// swagger:model workflowDynamicWorkflowActionTaskList
type WorkflowDynamicWorkflowActionTaskList struct {

	// The action of the Dynamic Workflow.
	//
	Action string `json:"Action,omitempty"`

	// The task list that has precedence which dictates how the workflow should be constructed.
	//
	Tasks interface{} `json:"Tasks,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowDynamicWorkflowActionTaskList) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Action string `json:"Action,omitempty"`

		Tasks interface{} `json:"Tasks,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Action = dataAO0.Action

	m.Tasks = dataAO0.Tasks

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowDynamicWorkflowActionTaskList) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Action string `json:"Action,omitempty"`

		Tasks interface{} `json:"Tasks,omitempty"`
	}

	dataAO0.Action = m.Action

	dataAO0.Tasks = m.Tasks

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow dynamic workflow action task list
func (m *WorkflowDynamicWorkflowActionTaskList) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowDynamicWorkflowActionTaskList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowDynamicWorkflowActionTaskList) UnmarshalBinary(b []byte) error {
	var res WorkflowDynamicWorkflowActionTaskList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
