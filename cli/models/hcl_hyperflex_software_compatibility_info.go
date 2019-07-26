// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HclHyperflexSoftwareCompatibilityInfo Hyperflex Software Compatibility Information
//
// Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
//
// swagger:model hclHyperflexSoftwareCompatibilityInfo
type HclHyperflexSoftwareCompatibilityInfo struct {
	MoBaseMo

	// A collection of references to the [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) MO unsets its reference to this deleted MO.
	//
	AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

	// Constraint for the matching software compatibility info, in case the match is applicable for certain cases only. For example a combination of (HyperFlex Data Platform, serverFw and hypervisor) versions can be applicable only for a HyperFlex Cluster UPGRADE operation, so a constraint of "supportedOperations=upgrade" can be added to the matching row.
	//
	Constraints []*HclConstraint `json:"Constraints"`

	// HXDP component software version.
	//
	HxdpVersion string `json:"HxdpVersion,omitempty"`

	// Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.
	//
	// Enum: [ESXi]
	HypervisorType *string `json:"HypervisorType,omitempty"`

	// Hypervisor component software version.
	//
	HypervisorVersion string `json:"HypervisorVersion,omitempty"`

	// UCS Server Firmware component software version.
	//
	ServerFwVersion string `json:"ServerFwVersion,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HclHyperflexSoftwareCompatibilityInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		Constraints []*HclConstraint `json:"Constraints"`

		HxdpVersion string `json:"HxdpVersion,omitempty"`

		HypervisorType *string `json:"HypervisorType,omitempty"`

		HypervisorVersion string `json:"HypervisorVersion,omitempty"`

		ServerFwVersion string `json:"ServerFwVersion,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppCatalog = dataAO1.AppCatalog

	m.Constraints = dataAO1.Constraints

	m.HxdpVersion = dataAO1.HxdpVersion

	m.HypervisorType = dataAO1.HypervisorType

	m.HypervisorVersion = dataAO1.HypervisorVersion

	m.ServerFwVersion = dataAO1.ServerFwVersion

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HclHyperflexSoftwareCompatibilityInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		Constraints []*HclConstraint `json:"Constraints"`

		HxdpVersion string `json:"HxdpVersion,omitempty"`

		HypervisorType *string `json:"HypervisorType,omitempty"`

		HypervisorVersion string `json:"HypervisorVersion,omitempty"`

		ServerFwVersion string `json:"ServerFwVersion,omitempty"`
	}

	dataAO1.AppCatalog = m.AppCatalog

	dataAO1.Constraints = m.Constraints

	dataAO1.HxdpVersion = m.HxdpVersion

	dataAO1.HypervisorType = m.HypervisorType

	dataAO1.HypervisorVersion = m.HypervisorVersion

	dataAO1.ServerFwVersion = m.ServerFwVersion

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hcl hyperflex software compatibility info
func (m *HclHyperflexSoftwareCompatibilityInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervisorType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HclHyperflexSoftwareCompatibilityInfo) validateAppCatalog(formats strfmt.Registry) error {

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

func (m *HclHyperflexSoftwareCompatibilityInfo) validateConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var hclHyperflexSoftwareCompatibilityInfoTypeHypervisorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ESXi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hclHyperflexSoftwareCompatibilityInfoTypeHypervisorTypePropEnum = append(hclHyperflexSoftwareCompatibilityInfoTypeHypervisorTypePropEnum, v)
	}
}

// property enum
func (m *HclHyperflexSoftwareCompatibilityInfo) validateHypervisorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hclHyperflexSoftwareCompatibilityInfoTypeHypervisorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HclHyperflexSoftwareCompatibilityInfo) validateHypervisorType(formats strfmt.Registry) error {

	if swag.IsZero(m.HypervisorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateHypervisorTypeEnum("HypervisorType", "body", *m.HypervisorType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HclHyperflexSoftwareCompatibilityInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HclHyperflexSoftwareCompatibilityInfo) UnmarshalBinary(b []byte) error {
	var res HclHyperflexSoftwareCompatibilityInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
