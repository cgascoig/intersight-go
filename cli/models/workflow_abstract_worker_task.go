// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowAbstractWorkerTask Workflow:Abstract Worker Task
//
// An AbstractWorkerTask is used to model a task that does some end-user work. This can be another task or it can be another workflow.
//
// swagger:model workflowAbstractWorkerTask
type WorkflowAbstractWorkerTask struct {
	WorkflowWorkflowTask

	// JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is '${Source.input/output.JsonPath}'. 'Source' can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source's input/output.
	//
	InputParameters interface{} `json:"InputParameters,omitempty"`

	// This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
	//
	OnFailure string `json:"OnFailure,omitempty"`

	// This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
	//
	OnSuccess string `json:"OnSuccess,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowAbstractWorkerTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowWorkflowTask
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowWorkflowTask = aO0

	// AO1
	var dataAO1 struct {
		InputParameters interface{} `json:"InputParameters,omitempty"`

		OnFailure string `json:"OnFailure,omitempty"`

		OnSuccess string `json:"OnSuccess,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.InputParameters = dataAO1.InputParameters

	m.OnFailure = dataAO1.OnFailure

	m.OnSuccess = dataAO1.OnSuccess

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowAbstractWorkerTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowWorkflowTask)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		InputParameters interface{} `json:"InputParameters,omitempty"`

		OnFailure string `json:"OnFailure,omitempty"`

		OnSuccess string `json:"OnSuccess,omitempty"`
	}

	dataAO1.InputParameters = m.InputParameters

	dataAO1.OnFailure = m.OnFailure

	dataAO1.OnSuccess = m.OnSuccess

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow abstract worker task
func (m *WorkflowAbstractWorkerTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowWorkflowTask
	if err := m.WorkflowWorkflowTask.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowAbstractWorkerTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowAbstractWorkerTask) UnmarshalBinary(b []byte) error {
	var res WorkflowAbstractWorkerTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
