// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageReplicationBlackout Storage:Replication Blackout
//
// Range of time at which to suspend replication. System disables replication during this interval.
//
// swagger:model storageReplicationBlackout
type StorageReplicationBlackout struct {

	// The end time of day for replication blackout window.
	// Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
	//
	//
	// Read Only: true
	End string `json:"End,omitempty"`

	// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication.
	// Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds.
	// The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
	//
	//
	// Read Only: true
	Start string `json:"Start,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageReplicationBlackout) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		End string `json:"End,omitempty"`

		Start string `json:"Start,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.End = dataAO0.End

	m.Start = dataAO0.Start

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageReplicationBlackout) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		End string `json:"End,omitempty"`

		Start string `json:"Start,omitempty"`
	}

	dataAO0.End = m.End

	dataAO0.Start = m.Start

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage replication blackout
func (m *StorageReplicationBlackout) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageReplicationBlackout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageReplicationBlackout) UnmarshalBinary(b []byte) error {
	var res StorageReplicationBlackout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}