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

// WorkflowBatchAPIExecutor Batch API
//
// Intersight allows generic API tasks to be created by taking the API request
// body and a response parser specification in the form of content.Grammar object.
//
// Batch API associates the list of API requests to be executed as part of single
// task execution. Each API request takes the request body and a response parser
// specification.
//
// swagger:model workflowBatchApiExecutor
type WorkflowBatchAPIExecutor struct {
	MoBaseMo

	// Intersight Orchestrator supports one or a batch of APIs to be executed as part of
	// a task execution.
	//
	// The batch cannot be empty.
	//
	//
	Batch []*WorkflowAPI `json:"Batch"`

	// Enter the constraints on when this task should match against the task definition.
	//
	//
	Constraints string `json:"Constraints,omitempty"`

	// A detailed description about the batch APIs.
	//
	//
	Description string `json:"Description,omitempty"`

	// Name for the batch API task.
	//
	//
	Name string `json:"Name,omitempty"`

	// Intersight Orchestrator allows the extraction of required values from API
	// responses using the API response grammar. These extracted values can be mapped
	// to task output parameters defined in task definition.
	//
	// The mapping of API output parameters to the task output parameters is provided
	// as JSON in this property.
	//
	//
	Output interface{} `json:"Output,omitempty"`

	// The skip expression, if provided, allows the batch API executor to skip the
	// task execution when the given expression evaluates to true.
	//
	// The expression is given as such a golang template that has to be
	// evaluated to a final content true/false. The expression is an optional and in
	// case not provided, the API will always be executed.
	//
	//
	SkipOnCondition string `json:"SkipOnCondition,omitempty"`

	// The interface task definition for which this batch API is one of the implementation.
	//
	TaskDefinition *WorkflowTaskDefinitionRef `json:"TaskDefinition,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowBatchAPIExecutor) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Batch []*WorkflowAPI `json:"Batch"`

		Constraints string `json:"Constraints,omitempty"`

		Description string `json:"Description,omitempty"`

		Name string `json:"Name,omitempty"`

		Output interface{} `json:"Output,omitempty"`

		SkipOnCondition string `json:"SkipOnCondition,omitempty"`

		TaskDefinition *WorkflowTaskDefinitionRef `json:"TaskDefinition,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Batch = dataAO1.Batch

	m.Constraints = dataAO1.Constraints

	m.Description = dataAO1.Description

	m.Name = dataAO1.Name

	m.Output = dataAO1.Output

	m.SkipOnCondition = dataAO1.SkipOnCondition

	m.TaskDefinition = dataAO1.TaskDefinition

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowBatchAPIExecutor) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Batch []*WorkflowAPI `json:"Batch"`

		Constraints string `json:"Constraints,omitempty"`

		Description string `json:"Description,omitempty"`

		Name string `json:"Name,omitempty"`

		Output interface{} `json:"Output,omitempty"`

		SkipOnCondition string `json:"SkipOnCondition,omitempty"`

		TaskDefinition *WorkflowTaskDefinitionRef `json:"TaskDefinition,omitempty"`
	}

	dataAO1.Batch = m.Batch

	dataAO1.Constraints = m.Constraints

	dataAO1.Description = m.Description

	dataAO1.Name = m.Name

	dataAO1.Output = m.Output

	dataAO1.SkipOnCondition = m.SkipOnCondition

	dataAO1.TaskDefinition = m.TaskDefinition

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow batch Api executor
func (m *WorkflowBatchAPIExecutor) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowBatchAPIExecutor) validateBatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Batch) { // not required
		return nil
	}

	for i := 0; i < len(m.Batch); i++ {
		if swag.IsZero(m.Batch[i]) { // not required
			continue
		}

		if m.Batch[i] != nil {
			if err := m.Batch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Batch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowBatchAPIExecutor) validateTaskDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskDefinition) { // not required
		return nil
	}

	if m.TaskDefinition != nil {
		if err := m.TaskDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaskDefinition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowBatchAPIExecutor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowBatchAPIExecutor) UnmarshalBinary(b []byte) error {
	var res WorkflowBatchAPIExecutor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
