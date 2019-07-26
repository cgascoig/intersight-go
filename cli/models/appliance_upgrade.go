// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplianceUpgrade Appliance:Upgrade
//
// Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
// upgrade service automatically creates an Upgrade managed object when there is a
// pending software upgrade. User may modify an active Upgrade managed object to reset
// the software upgrade start time. However, user may not be able to postpone an upgrade
// beyond the limit enforced by the upgrade grace period set in the software manifest.
//
// swagger:model applianceUpgrade
type ApplianceUpgrade struct {
	MoBaseMo

	// Upgrade managed object to Account relationship.
	//
	Account *IamAccountRef `json:"Account,omitempty"`

	// Indicates if the software upgrade is active or not.
	//
	// Read Only: true
	Active *bool `json:"Active,omitempty"`

	// Indicates that the request was automatically created by the system.
	//
	// Read Only: true
	AutoCreated *bool `json:"AutoCreated,omitempty"`

	// Collection of the completed software upgrade phases.
	//
	// Read Only: true
	CompletedPhases []*OnpremUpgradePhase `json:"CompletedPhases"`

	// Current phase of the Intersight Appliance's software upgrade.
	//
	// Read Only: true
	CurrentPhase *OnpremUpgradePhase `json:"CurrentPhase,omitempty"`

	// Description of the software upgrade.
	//
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// Elapsed time in seconds during the software upgrade.
	//
	// Read Only: true
	ElapsedTime int64 `json:"ElapsedTime,omitempty"`

	// End date of the software upgrade.
	//
	// Read Only: true
	// Format: date-time
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	// Software upgrade manifest's fingerprint.
	//
	// Read Only: true
	Fingerprint string `json:"Fingerprint,omitempty"`

	// Upgrade managed object to ImageBundle relationship.
	//
	// Read Only: true
	ImageBundle *ApplianceImageBundleRef `json:"ImageBundle,omitempty"`

	// Messages generated during the software upgrade process.
	//
	// Read Only: true
	Messages []string `json:"Messages"`

	// Services that are upgraded during the software upgrade process. For example, if the software upgrade has updates for five Intersight micro-services, then this field will have the names of those five micro-services.
	//
	// Read Only: true
	Services []string `json:"Services"`

	// Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.
	//
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// Status of the Intersight Appliance's software upgrade.
	//
	// Read Only: true
	Status string `json:"Status,omitempty"`

	// Name of the UI packages that are upgraded. For example, if the software upgrade has updates for five Intersight micro-service UI packages, then this field will have the names of those five micro-services.
	//
	// Read Only: true
	UIPackages []string `json:"UiPackages"`

	// Software upgrade manifest's version.
	//
	// Read Only: true
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceUpgrade) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Active *bool `json:"Active,omitempty"`

		AutoCreated *bool `json:"AutoCreated,omitempty"`

		CompletedPhases []*OnpremUpgradePhase `json:"CompletedPhases"`

		CurrentPhase *OnpremUpgradePhase `json:"CurrentPhase,omitempty"`

		Description string `json:"Description,omitempty"`

		ElapsedTime int64 `json:"ElapsedTime,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		Fingerprint string `json:"Fingerprint,omitempty"`

		ImageBundle *ApplianceImageBundleRef `json:"ImageBundle,omitempty"`

		Messages []string `json:"Messages"`

		Services []string `json:"Services"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`

		UIPackages []string `json:"UiPackages"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.Active = dataAO1.Active

	m.AutoCreated = dataAO1.AutoCreated

	m.CompletedPhases = dataAO1.CompletedPhases

	m.CurrentPhase = dataAO1.CurrentPhase

	m.Description = dataAO1.Description

	m.ElapsedTime = dataAO1.ElapsedTime

	m.EndTime = dataAO1.EndTime

	m.Fingerprint = dataAO1.Fingerprint

	m.ImageBundle = dataAO1.ImageBundle

	m.Messages = dataAO1.Messages

	m.Services = dataAO1.Services

	m.StartTime = dataAO1.StartTime

	m.Status = dataAO1.Status

	m.UIPackages = dataAO1.UIPackages

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceUpgrade) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Active *bool `json:"Active,omitempty"`

		AutoCreated *bool `json:"AutoCreated,omitempty"`

		CompletedPhases []*OnpremUpgradePhase `json:"CompletedPhases"`

		CurrentPhase *OnpremUpgradePhase `json:"CurrentPhase,omitempty"`

		Description string `json:"Description,omitempty"`

		ElapsedTime int64 `json:"ElapsedTime,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		Fingerprint string `json:"Fingerprint,omitempty"`

		ImageBundle *ApplianceImageBundleRef `json:"ImageBundle,omitempty"`

		Messages []string `json:"Messages"`

		Services []string `json:"Services"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`

		UIPackages []string `json:"UiPackages"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.Active = m.Active

	dataAO1.AutoCreated = m.AutoCreated

	dataAO1.CompletedPhases = m.CompletedPhases

	dataAO1.CurrentPhase = m.CurrentPhase

	dataAO1.Description = m.Description

	dataAO1.ElapsedTime = m.ElapsedTime

	dataAO1.EndTime = m.EndTime

	dataAO1.Fingerprint = m.Fingerprint

	dataAO1.ImageBundle = m.ImageBundle

	dataAO1.Messages = m.Messages

	dataAO1.Services = m.Services

	dataAO1.StartTime = m.StartTime

	dataAO1.Status = m.Status

	dataAO1.UIPackages = m.UIPackages

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance upgrade
func (m *ApplianceUpgrade) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedPhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentPhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageBundle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplianceUpgrade) validateAccount(formats strfmt.Registry) error {

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

func (m *ApplianceUpgrade) validateCompletedPhases(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedPhases) { // not required
		return nil
	}

	for i := 0; i < len(m.CompletedPhases); i++ {
		if swag.IsZero(m.CompletedPhases[i]) { // not required
			continue
		}

		if m.CompletedPhases[i] != nil {
			if err := m.CompletedPhases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CompletedPhases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceUpgrade) validateCurrentPhase(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentPhase) { // not required
		return nil
	}

	if m.CurrentPhase != nil {
		if err := m.CurrentPhase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrentPhase")
			}
			return err
		}
	}

	return nil
}

func (m *ApplianceUpgrade) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApplianceUpgrade) validateImageBundle(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageBundle) { // not required
		return nil
	}

	if m.ImageBundle != nil {
		if err := m.ImageBundle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ImageBundle")
			}
			return err
		}
	}

	return nil
}

func (m *ApplianceUpgrade) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceUpgrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceUpgrade) UnmarshalBinary(b []byte) error {
	var res ApplianceUpgrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
