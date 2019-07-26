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

// UcsdconnectorTelemetryMo Ucsdconnector:Telemetry Mo
//
// It represents MO name, Field names and corresponding queries.
//
// swagger:model ucsdconnectorTelemetryMo
type UcsdconnectorTelemetryMo struct {

	// Its a collection of FieldQuery objects.
	//
	FieldQueries []*UcsdconnectorFieldQuery `json:"FieldQueries"`

	// MoName corresponds to Fanwood telemetry MO.
	//
	MoName string `json:"MoName,omitempty"`

	// It contains a single SQL Query that collects all the properties of a Mo.
	//
	ObjectQuery *UcsdconnectorSQLQuery `json:"ObjectQuery,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UcsdconnectorTelemetryMo) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		FieldQueries []*UcsdconnectorFieldQuery `json:"FieldQueries"`

		MoName string `json:"MoName,omitempty"`

		ObjectQuery *UcsdconnectorSQLQuery `json:"ObjectQuery,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.FieldQueries = dataAO0.FieldQueries

	m.MoName = dataAO0.MoName

	m.ObjectQuery = dataAO0.ObjectQuery

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UcsdconnectorTelemetryMo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		FieldQueries []*UcsdconnectorFieldQuery `json:"FieldQueries"`

		MoName string `json:"MoName,omitempty"`

		ObjectQuery *UcsdconnectorSQLQuery `json:"ObjectQuery,omitempty"`
	}

	dataAO0.FieldQueries = m.FieldQueries

	dataAO0.MoName = m.MoName

	dataAO0.ObjectQuery = m.ObjectQuery

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ucsdconnector telemetry mo
func (m *UcsdconnectorTelemetryMo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldQueries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UcsdconnectorTelemetryMo) validateFieldQueries(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldQueries) { // not required
		return nil
	}

	for i := 0; i < len(m.FieldQueries); i++ {
		if swag.IsZero(m.FieldQueries[i]) { // not required
			continue
		}

		if m.FieldQueries[i] != nil {
			if err := m.FieldQueries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FieldQueries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UcsdconnectorTelemetryMo) validateObjectQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectQuery) { // not required
		return nil
	}

	if m.ObjectQuery != nil {
		if err := m.ObjectQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectQuery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UcsdconnectorTelemetryMo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UcsdconnectorTelemetryMo) UnmarshalBinary(b []byte) error {
	var res UcsdconnectorTelemetryMo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}