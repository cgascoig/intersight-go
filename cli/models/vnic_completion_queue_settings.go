// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicCompletionQueueSettings Completion Queue Settings
//
// Completion Queue resource settings.
//
// swagger:model vnicCompletionQueueSettings
type VnicCompletionQueueSettings struct {

	// The number of completion queue resources to allocate. In general, the number of completion queue resources you should allocate is equal to the number of transmit queue resources plus the number of receive queue resources.
	//
	Count int64 `json:"Count,omitempty"`

	// The number of descriptors in each completion queue.
	//
	// Read Only: true
	RingSize int64 `json:"RingSize,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicCompletionQueueSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Count int64 `json:"Count,omitempty"`

		RingSize int64 `json:"RingSize,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Count = dataAO0.Count

	m.RingSize = dataAO0.RingSize

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicCompletionQueueSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Count int64 `json:"Count,omitempty"`

		RingSize int64 `json:"RingSize,omitempty"`
	}

	dataAO0.Count = m.Count

	dataAO0.RingSize = m.RingSize

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic completion queue settings
func (m *VnicCompletionQueueSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicCompletionQueueSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicCompletionQueueSettings) UnmarshalBinary(b []byte) error {
	var res VnicCompletionQueueSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}