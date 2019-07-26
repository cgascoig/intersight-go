// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicFcNetworkPolicy Fibre Channel Network
//
// A Fibre Channel Network policy governs the VSAN configuration for the virtual interfaces.
//
// swagger:model vnicFcNetworkPolicy
type VnicFcNetworkPolicy struct {
	PolicyAbstractPolicy

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// VSAN configuration for the virtual interface.
	//
	VsanSettings *VnicVsanSettings `json:"VsanSettings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicFcNetworkPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		Organization *IamAccountRef `json:"Organization,omitempty"`

		VsanSettings *VnicVsanSettings `json:"VsanSettings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Organization = dataAO1.Organization

	m.VsanSettings = dataAO1.VsanSettings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicFcNetworkPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Organization *IamAccountRef `json:"Organization,omitempty"`

		VsanSettings *VnicVsanSettings `json:"VsanSettings,omitempty"`
	}

	dataAO1.Organization = m.Organization

	dataAO1.VsanSettings = m.VsanSettings

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic fc network policy
func (m *VnicFcNetworkPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsanSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VnicFcNetworkPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcNetworkPolicy) validateVsanSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.VsanSettings) { // not required
		return nil
	}

	if m.VsanSettings != nil {
		if err := m.VsanSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VsanSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcNetworkPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcNetworkPolicy) UnmarshalBinary(b []byte) error {
	var res VnicFcNetworkPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
