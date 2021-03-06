// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexStPlatformClusterHealingInfo Hyperflex:St Platform Cluster Healing Info
// swagger:model hyperflexStPlatformClusterHealingInfo
type HyperflexStPlatformClusterHealingInfo struct {

	// estimated completion time in seconds
	// Read Only: true
	EstimatedCompletionTimeInSeconds int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`

	// in progress
	// Read Only: true
	InProgress *bool `json:"InProgress,omitempty"`

	// messages
	// Read Only: true
	Messages []string `json:"Messages"`

	// messages iterator
	// Read Only: true
	MessagesIterator interface{} `json:"MessagesIterator,omitempty"`

	// messages size
	// Read Only: true
	MessagesSize int64 `json:"MessagesSize,omitempty"`

	// percent complete
	// Read Only: true
	PercentComplete int64 `json:"PercentComplete,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexStPlatformClusterHealingInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		EstimatedCompletionTimeInSeconds int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`

		InProgress *bool `json:"InProgress,omitempty"`

		Messages []string `json:"Messages"`

		MessagesIterator interface{} `json:"MessagesIterator,omitempty"`

		MessagesSize int64 `json:"MessagesSize,omitempty"`

		PercentComplete int64 `json:"PercentComplete,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.EstimatedCompletionTimeInSeconds = dataAO0.EstimatedCompletionTimeInSeconds

	m.InProgress = dataAO0.InProgress

	m.Messages = dataAO0.Messages

	m.MessagesIterator = dataAO0.MessagesIterator

	m.MessagesSize = dataAO0.MessagesSize

	m.PercentComplete = dataAO0.PercentComplete

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexStPlatformClusterHealingInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		EstimatedCompletionTimeInSeconds int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`

		InProgress *bool `json:"InProgress,omitempty"`

		Messages []string `json:"Messages"`

		MessagesIterator interface{} `json:"MessagesIterator,omitempty"`

		MessagesSize int64 `json:"MessagesSize,omitempty"`

		PercentComplete int64 `json:"PercentComplete,omitempty"`
	}

	dataAO0.EstimatedCompletionTimeInSeconds = m.EstimatedCompletionTimeInSeconds

	dataAO0.InProgress = m.InProgress

	dataAO0.Messages = m.Messages

	dataAO0.MessagesIterator = m.MessagesIterator

	dataAO0.MessagesSize = m.MessagesSize

	dataAO0.PercentComplete = m.PercentComplete

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex st platform cluster healing info
func (m *HyperflexStPlatformClusterHealingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexStPlatformClusterHealingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexStPlatformClusterHealingInfo) UnmarshalBinary(b []byte) error {
	var res HyperflexStPlatformClusterHealingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
