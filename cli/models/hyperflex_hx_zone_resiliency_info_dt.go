// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexHxZoneResiliencyInfoDt Hyperflex:Hx Zone Resiliency Info Dt
// swagger:model hyperflexHxZoneResiliencyInfoDt
type HyperflexHxZoneResiliencyInfoDt struct {

	// name
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// resiliency info
	// Read Only: true
	ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxZoneResiliencyInfoDt) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name string `json:"Name,omitempty"`

		ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	m.ResiliencyInfo = dataAO0.ResiliencyInfo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxZoneResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Name string `json:"Name,omitempty"`

		ResiliencyInfo *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	}

	dataAO0.Name = m.Name

	dataAO0.ResiliencyInfo = m.ResiliencyInfo

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hx zone resiliency info dt
func (m *HyperflexHxZoneResiliencyInfoDt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResiliencyInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexHxZoneResiliencyInfoDt) validateResiliencyInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ResiliencyInfo) { // not required
		return nil
	}

	if m.ResiliencyInfo != nil {
		if err := m.ResiliencyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResiliencyInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxZoneResiliencyInfoDt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxZoneResiliencyInfoDt) UnmarshalBinary(b []byte) error {
	var res HyperflexHxZoneResiliencyInfoDt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}