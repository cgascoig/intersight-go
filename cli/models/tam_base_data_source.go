// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TamBaseDataSource Tam:Base Data Source
//
// Represents source for the data needed to analyze the alerts. this could one of several supported source types (textFsmTemplates/Intersight API/external API). TextFsmTemplates and Intersight API are the only ones supported currently.
//
// swagger:model tamBaseDataSource
type TamBaseDataSource struct {

	// Name is used to unique identify and refer a given data source in an alert definition.
	//
	Name string `json:"Name,omitempty"`

	// Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).
	//
	// Enum: [nxos intersightApi]
	Type *string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamBaseDataSource) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name string `json:"Name,omitempty"`

		Type *string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	m.Type = dataAO0.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamBaseDataSource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Name string `json:"Name,omitempty"`

		Type *string `json:"Type,omitempty"`
	}

	dataAO0.Name = m.Name

	dataAO0.Type = m.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam base data source
func (m *TamBaseDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tamBaseDataSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nxos","intersightApi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tamBaseDataSourceTypeTypePropEnum = append(tamBaseDataSourceTypeTypePropEnum, v)
	}
}

// property enum
func (m *TamBaseDataSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tamBaseDataSourceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TamBaseDataSource) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TamBaseDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamBaseDataSource) UnmarshalBinary(b []byte) error {
	var res TamBaseDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
