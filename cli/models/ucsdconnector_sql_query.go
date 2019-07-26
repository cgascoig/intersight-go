// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UcsdconnectorSQLQuery Ucsdconnector:Sql Query
//
// SqlQuery is an object of sql query string and its params.
//
// swagger:model ucsdconnectorSqlQuery
type UcsdconnectorSQLQuery struct {

	// Parameters in query string.
	//
	Params interface{} `json:"Params,omitempty"`

	// SQL query string to be executed on UCSD database.
	//
	Query string `json:"Query,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UcsdconnectorSQLQuery) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Params interface{} `json:"Params,omitempty"`

		Query string `json:"Query,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Params = dataAO0.Params

	m.Query = dataAO0.Query

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UcsdconnectorSQLQuery) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Params interface{} `json:"Params,omitempty"`

		Query string `json:"Query,omitempty"`
	}

	dataAO0.Params = m.Params

	dataAO0.Query = m.Query

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ucsdconnector Sql query
func (m *UcsdconnectorSQLQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UcsdconnectorSQLQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UcsdconnectorSQLQuery) UnmarshalBinary(b []byte) error {
	var res UcsdconnectorSQLQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
