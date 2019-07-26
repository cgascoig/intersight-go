// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PolicyAbstractConfigProfile Policy:Abstract Config Profile
//
// AbstractConfigProfile is an abstract base type for all config actions on a profile.
//
// swagger:model policyAbstractConfigProfile
type PolicyAbstractConfigProfile struct {
	PolicyAbstractProfile

	// User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
	//
	Action string `json:"Action,omitempty"`

	// The configuration state and results of the last configuration operation.
	//
	ConfigContext *PolicyConfigContext `json:"ConfigContext,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyAbstractConfigProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractProfile
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractProfile = aO0

	// AO1
	var dataAO1 struct {
		Action string `json:"Action,omitempty"`

		ConfigContext *PolicyConfigContext `json:"ConfigContext,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Action = dataAO1.Action

	m.ConfigContext = dataAO1.ConfigContext

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyAbstractConfigProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractProfile)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Action string `json:"Action,omitempty"`

		ConfigContext *PolicyConfigContext `json:"ConfigContext,omitempty"`
	}

	dataAO1.Action = m.Action

	dataAO1.ConfigContext = m.ConfigContext

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy abstract config profile
func (m *PolicyAbstractConfigProfile) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractProfile
	if err := m.PolicyAbstractProfile.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyAbstractConfigProfile) validateConfigContext(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigContext) { // not required
		return nil
	}

	if m.ConfigContext != nil {
		if err := m.ConfigContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyAbstractConfigProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyAbstractConfigProfile) UnmarshalBinary(b []byte) error {
	var res PolicyAbstractConfigProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}