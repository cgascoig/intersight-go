// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ApplianceUpgradePolicy Appliance:Upgrade Policy
//
// UpgradePolicy stores the Intersight Appliance's software upgrade policy. UpgradePolicy
// is a sinlgeton managed object. A default upgrade policy is created during the Intersight
// Appliance setup, and it is configured with an automatic upgrade policy.
//
// Automatic upgrade policy lets the system start software upgrade after a pre-defined
// number of seconds set in the software manifest.
//
// Scheduled upgrade pilicy lets the user start software upgrade at a specified schedule.
// However, scheduled time cannot be beyond the time limit enforced by the upgrade grace
// period set in the software manifest.
//
// swagger:model applianceUpgradePolicy
type ApplianceUpgradePolicy struct {
	MoBaseMo

	// UpgradePolicy managed object to Account relationship.
	//
	Account *IamAccountRef `json:"Account,omitempty"`

	// Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored.
	//
	AutoUpgrade bool `json:"AutoUpgrade,omitempty"`

	// User defined schedule for upgrading the Intersight Appliance software.
	//
	Schedule *OnpremSchedule `json:"Schedule,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceUpgradePolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		AutoUpgrade bool `json:"AutoUpgrade,omitempty"`

		Schedule *OnpremSchedule `json:"Schedule,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.AutoUpgrade = dataAO1.AutoUpgrade

	m.Schedule = dataAO1.Schedule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceUpgradePolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		AutoUpgrade bool `json:"AutoUpgrade,omitempty"`

		Schedule *OnpremSchedule `json:"Schedule,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.AutoUpgrade = m.AutoUpgrade

	dataAO1.Schedule = m.Schedule

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance upgrade policy
func (m *ApplianceUpgradePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplianceUpgradePolicy) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

func (m *ApplianceUpgradePolicy) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceUpgradePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceUpgradePolicy) UnmarshalBinary(b []byte) error {
	var res ApplianceUpgradePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}