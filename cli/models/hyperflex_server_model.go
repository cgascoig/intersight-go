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

// HyperflexServerModel Hyperflex:Server Model
//
// A supported server model.
//
// swagger:model hyperflexServerModel
type HyperflexServerModel struct {
	MoBaseMo

	// A collection of references to the [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) MO unsets its reference to this deleted MO.
	//
	AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

	// The supported server models in regex format.
	//
	ServerModelEntries []*HyperflexServerModelEntry `json:"ServerModelEntries"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexServerModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		ServerModelEntries []*HyperflexServerModelEntry `json:"ServerModelEntries"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppCatalog = dataAO1.AppCatalog

	m.ServerModelEntries = dataAO1.ServerModelEntries

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexServerModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		ServerModelEntries []*HyperflexServerModelEntry `json:"ServerModelEntries"`
	}

	dataAO1.AppCatalog = m.AppCatalog

	dataAO1.ServerModelEntries = m.ServerModelEntries

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex server model
func (m *HyperflexServerModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerModelEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexServerModel) validateAppCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.AppCatalog) { // not required
		return nil
	}

	if m.AppCatalog != nil {
		if err := m.AppCatalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AppCatalog")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexServerModel) validateServerModelEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerModelEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerModelEntries); i++ {
		if swag.IsZero(m.ServerModelEntries[i]) { // not required
			continue
		}

		if m.ServerModelEntries[i] != nil {
			if err := m.ServerModelEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ServerModelEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexServerModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexServerModel) UnmarshalBinary(b []byte) error {
	var res HyperflexServerModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
