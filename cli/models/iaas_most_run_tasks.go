// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IaasMostRunTasks Iaas:Most Run Tasks
//
// Describes most run workflow tasks within UCSD.
//
// swagger:model iaasMostRunTasks
type IaasMostRunTasks struct {
	MoBaseMo

	// A collection of references to the [iaas.UcsdInfo](mo://iaas.UcsdInfo) Managed Object.
	//
	// When this managed object is deleted, the referenced [iaas.UcsdInfo](mo://iaas.UcsdInfo) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	GUID *IaasUcsdInfoRef `json:"Guid,omitempty"`

	// A functional area to which a task belongs to.
	//
	// Read Only: true
	TaskCategory string `json:"TaskCategory,omitempty"`

	// Number of times this task has executed.
	//
	// Read Only: true
	TaskExecutionCount int64 `json:"TaskExecutionCount,omitempty"`

	// Name of the task executed in UCSD.
	//
	// Read Only: true
	TaskName string `json:"TaskName,omitempty"`

	// Type of the task whether it is system task or custom task.
	//
	// Read Only: true
	TaskType string `json:"TaskType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IaasMostRunTasks) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		GUID *IaasUcsdInfoRef `json:"Guid,omitempty"`

		TaskCategory string `json:"TaskCategory,omitempty"`

		TaskExecutionCount int64 `json:"TaskExecutionCount,omitempty"`

		TaskName string `json:"TaskName,omitempty"`

		TaskType string `json:"TaskType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.GUID = dataAO1.GUID

	m.TaskCategory = dataAO1.TaskCategory

	m.TaskExecutionCount = dataAO1.TaskExecutionCount

	m.TaskName = dataAO1.TaskName

	m.TaskType = dataAO1.TaskType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IaasMostRunTasks) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		GUID *IaasUcsdInfoRef `json:"Guid,omitempty"`

		TaskCategory string `json:"TaskCategory,omitempty"`

		TaskExecutionCount int64 `json:"TaskExecutionCount,omitempty"`

		TaskName string `json:"TaskName,omitempty"`

		TaskType string `json:"TaskType,omitempty"`
	}

	dataAO1.GUID = m.GUID

	dataAO1.TaskCategory = m.TaskCategory

	dataAO1.TaskExecutionCount = m.TaskExecutionCount

	dataAO1.TaskName = m.TaskName

	dataAO1.TaskType = m.TaskType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iaas most run tasks
func (m *IaasMostRunTasks) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IaasMostRunTasks) validateGUID(formats strfmt.Registry) error {

	if swag.IsZero(m.GUID) { // not required
		return nil
	}

	if m.GUID != nil {
		if err := m.GUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Guid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IaasMostRunTasks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasMostRunTasks) UnmarshalBinary(b []byte) error {
	var res IaasMostRunTasks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
