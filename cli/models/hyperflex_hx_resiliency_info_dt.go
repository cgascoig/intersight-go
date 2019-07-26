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

// HyperflexHxResiliencyInfoDt Hyperflex:Hx Resiliency Info Dt
// swagger:model hyperflexHxResiliencyInfoDt
type HyperflexHxResiliencyInfoDt struct {

	// data replication factor
	// Read Only: true
	// Enum: [ONE_COPY TWO_COPIES THREE_COPIES FOUR_COPIES SIX_COPIES]
	DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

	// hdd failures tolerable
	// Read Only: true
	HddFailuresTolerable int64 `json:"HddFailuresTolerable,omitempty"`

	// messages
	// Read Only: true
	Messages []string `json:"Messages"`

	// node failures tolerable
	// Read Only: true
	NodeFailuresTolerable int64 `json:"NodeFailuresTolerable,omitempty"`

	// policy compliance
	// Read Only: true
	// Enum: [UNKNOWN COMPLIANT NON_COMPLIANT]
	PolicyCompliance string `json:"PolicyCompliance,omitempty"`

	// resiliency state
	// Read Only: true
	// Enum: [UNKNOWN HEALTHY WARNING OFFLINE]
	ResiliencyState string `json:"ResiliencyState,omitempty"`

	// ssd failures tolerable
	// Read Only: true
	SsdFailuresTolerable int64 `json:"SsdFailuresTolerable,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxResiliencyInfoDt) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

		HddFailuresTolerable int64 `json:"HddFailuresTolerable,omitempty"`

		Messages []string `json:"Messages"`

		NodeFailuresTolerable int64 `json:"NodeFailuresTolerable,omitempty"`

		PolicyCompliance string `json:"PolicyCompliance,omitempty"`

		ResiliencyState string `json:"ResiliencyState,omitempty"`

		SsdFailuresTolerable int64 `json:"SsdFailuresTolerable,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.DataReplicationFactor = dataAO0.DataReplicationFactor

	m.HddFailuresTolerable = dataAO0.HddFailuresTolerable

	m.Messages = dataAO0.Messages

	m.NodeFailuresTolerable = dataAO0.NodeFailuresTolerable

	m.PolicyCompliance = dataAO0.PolicyCompliance

	m.ResiliencyState = dataAO0.ResiliencyState

	m.SsdFailuresTolerable = dataAO0.SsdFailuresTolerable

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		DataReplicationFactor string `json:"DataReplicationFactor,omitempty"`

		HddFailuresTolerable int64 `json:"HddFailuresTolerable,omitempty"`

		Messages []string `json:"Messages"`

		NodeFailuresTolerable int64 `json:"NodeFailuresTolerable,omitempty"`

		PolicyCompliance string `json:"PolicyCompliance,omitempty"`

		ResiliencyState string `json:"ResiliencyState,omitempty"`

		SsdFailuresTolerable int64 `json:"SsdFailuresTolerable,omitempty"`
	}

	dataAO0.DataReplicationFactor = m.DataReplicationFactor

	dataAO0.HddFailuresTolerable = m.HddFailuresTolerable

	dataAO0.Messages = m.Messages

	dataAO0.NodeFailuresTolerable = m.NodeFailuresTolerable

	dataAO0.PolicyCompliance = m.PolicyCompliance

	dataAO0.ResiliencyState = m.ResiliencyState

	dataAO0.SsdFailuresTolerable = m.SsdFailuresTolerable

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hx resiliency info dt
func (m *HyperflexHxResiliencyInfoDt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataReplicationFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hyperflexHxResiliencyInfoDtTypeDataReplicationFactorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONE_COPY","TWO_COPIES","THREE_COPIES","FOUR_COPIES","SIX_COPIES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHxResiliencyInfoDtTypeDataReplicationFactorPropEnum = append(hyperflexHxResiliencyInfoDtTypeDataReplicationFactorPropEnum, v)
	}
}

// property enum
func (m *HyperflexHxResiliencyInfoDt) validateDataReplicationFactorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHxResiliencyInfoDtTypeDataReplicationFactorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHxResiliencyInfoDt) validateDataReplicationFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.DataReplicationFactor) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataReplicationFactorEnum("DataReplicationFactor", "body", m.DataReplicationFactor); err != nil {
		return err
	}

	return nil
}

var hyperflexHxResiliencyInfoDtTypePolicyCompliancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","COMPLIANT","NON_COMPLIANT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHxResiliencyInfoDtTypePolicyCompliancePropEnum = append(hyperflexHxResiliencyInfoDtTypePolicyCompliancePropEnum, v)
	}
}

// property enum
func (m *HyperflexHxResiliencyInfoDt) validatePolicyComplianceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHxResiliencyInfoDtTypePolicyCompliancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHxResiliencyInfoDt) validatePolicyCompliance(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyCompliance) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyComplianceEnum("PolicyCompliance", "body", m.PolicyCompliance); err != nil {
		return err
	}

	return nil
}

var hyperflexHxResiliencyInfoDtTypeResiliencyStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","HEALTHY","WARNING","OFFLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHxResiliencyInfoDtTypeResiliencyStatePropEnum = append(hyperflexHxResiliencyInfoDtTypeResiliencyStatePropEnum, v)
	}
}

// property enum
func (m *HyperflexHxResiliencyInfoDt) validateResiliencyStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHxResiliencyInfoDtTypeResiliencyStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHxResiliencyInfoDt) validateResiliencyState(formats strfmt.Registry) error {

	if swag.IsZero(m.ResiliencyState) { // not required
		return nil
	}

	// value enum
	if err := m.validateResiliencyStateEnum("ResiliencyState", "body", m.ResiliencyState); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxResiliencyInfoDt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxResiliencyInfoDt) UnmarshalBinary(b []byte) error {
	var res HyperflexHxResiliencyInfoDt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}