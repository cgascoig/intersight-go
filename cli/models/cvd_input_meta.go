// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CvdInputMeta Cvd:Input Meta
// swagger:model cvdInputMeta
type CvdInputMeta struct {

	// Input field filter(REST API path with filter which will return the list of applicable values for this field)
	//
	FieldFilter string `json:"FieldFilter,omitempty"`

	// Input field label(for GUI use)
	//
	FieldLabel string `json:"FieldLabel,omitempty"`

	// Input field name
	//
	FieldName string `json:"FieldName,omitempty"`

	// Input field value
	//
	FieldValue string `json:"FieldValue,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CvdInputMeta) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		FieldFilter string `json:"FieldFilter,omitempty"`

		FieldLabel string `json:"FieldLabel,omitempty"`

		FieldName string `json:"FieldName,omitempty"`

		FieldValue string `json:"FieldValue,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.FieldFilter = dataAO0.FieldFilter

	m.FieldLabel = dataAO0.FieldLabel

	m.FieldName = dataAO0.FieldName

	m.FieldValue = dataAO0.FieldValue

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CvdInputMeta) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		FieldFilter string `json:"FieldFilter,omitempty"`

		FieldLabel string `json:"FieldLabel,omitempty"`

		FieldName string `json:"FieldName,omitempty"`

		FieldValue string `json:"FieldValue,omitempty"`
	}

	dataAO0.FieldFilter = m.FieldFilter

	dataAO0.FieldLabel = m.FieldLabel

	dataAO0.FieldName = m.FieldName

	dataAO0.FieldValue = m.FieldValue

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cvd input meta
func (m *CvdInputMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CvdInputMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CvdInputMeta) UnmarshalBinary(b []byte) error {
	var res CvdInputMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
