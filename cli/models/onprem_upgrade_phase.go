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

// OnpremUpgradePhase Onprem:Upgrade Phase
//
// UpgradePhase represents a phase of the Intersight Appliance software upgrade
// process. This data structure is shared by both the Intersight upgrade service
// and the Intersight Appliance's upgrade service.
//
// swagger:model onpremUpgradePhase
type OnpremUpgradePhase struct {

	// Name of the upgrade phase.
	//
	// Read Only: true
	// Enum: [init Prepare ServiceLoad UiLoad GenerateConfig DeployService Success Fail Cancel Telemetry]
	Name string `json:"Name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OnpremUpgradePhase) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name string `json:"Name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OnpremUpgradePhase) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Name string `json:"Name,omitempty"`
	}

	dataAO0.Name = m.Name

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this onprem upgrade phase
func (m *OnpremUpgradePhase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var onpremUpgradePhaseTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["init","Prepare","ServiceLoad","UiLoad","GenerateConfig","DeployService","Success","Fail","Cancel","Telemetry"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		onpremUpgradePhaseTypeNamePropEnum = append(onpremUpgradePhaseTypeNamePropEnum, v)
	}
}

// property enum
func (m *OnpremUpgradePhase) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, onpremUpgradePhaseTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OnpremUpgradePhase) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnpremUpgradePhase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnpremUpgradePhase) UnmarshalBinary(b []byte) error {
	var res OnpremUpgradePhase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}