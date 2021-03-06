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

// HyperflexAppCatalog HyperFlex application settings
//
// A catalog for managing application settings for HyperFlex cluster configuration service.
//
// swagger:model hyperflexAppCatalog
type HyperflexAppCatalog struct {
	MoBaseMo

	// The HyperFlex feature limits that are available to end users.
	//
	FeatureLimitExternal *HyperflexFeatureLimitExternalRef `json:"FeatureLimitExternal,omitempty"`

	// The HyperFlex feature limits for internal system use.
	//
	FeatureLimitInternal *HyperflexFeatureLimitInternalRef `json:"FeatureLimitInternal,omitempty"`

	// The list of supported HyperFlex Data Platform versions.
	//
	HxdpVersions []*HyperflexHxdpVersionRef `json:"HxdpVersions"`

	// Lists all supported HyperFlex feature capabilities and limitations.
	//
	HyperflexCapabilityInfos []*HyperflexCapabilityInfoRef `json:"HyperflexCapabilityInfos"`

	// Lists software compatibility information between different HyperFlex component versions like HXDP, Hypervisor, Drive Firmware, etc.
	//
	HyperflexSoftwareCompatibilityInfos []*HclHyperflexSoftwareCompatibilityInfoRef `json:"HyperflexSoftwareCompatibilityInfos"`

	// The supported server firmware bundle.
	//
	ServerFirmwareVersion *HyperflexServerFirmwareVersionRef `json:"ServerFirmwareVersion,omitempty"`

	// The supported server models in regex format.
	//
	ServerModel *HyperflexServerModelRef `json:"ServerModel,omitempty"`

	// The catalog version used in HyperFlex cluster configuration service.
	//
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexAppCatalog) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		FeatureLimitExternal *HyperflexFeatureLimitExternalRef `json:"FeatureLimitExternal,omitempty"`

		FeatureLimitInternal *HyperflexFeatureLimitInternalRef `json:"FeatureLimitInternal,omitempty"`

		HxdpVersions []*HyperflexHxdpVersionRef `json:"HxdpVersions"`

		HyperflexCapabilityInfos []*HyperflexCapabilityInfoRef `json:"HyperflexCapabilityInfos"`

		HyperflexSoftwareCompatibilityInfos []*HclHyperflexSoftwareCompatibilityInfoRef `json:"HyperflexSoftwareCompatibilityInfos"`

		ServerFirmwareVersion *HyperflexServerFirmwareVersionRef `json:"ServerFirmwareVersion,omitempty"`

		ServerModel *HyperflexServerModelRef `json:"ServerModel,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.FeatureLimitExternal = dataAO1.FeatureLimitExternal

	m.FeatureLimitInternal = dataAO1.FeatureLimitInternal

	m.HxdpVersions = dataAO1.HxdpVersions

	m.HyperflexCapabilityInfos = dataAO1.HyperflexCapabilityInfos

	m.HyperflexSoftwareCompatibilityInfos = dataAO1.HyperflexSoftwareCompatibilityInfos

	m.ServerFirmwareVersion = dataAO1.ServerFirmwareVersion

	m.ServerModel = dataAO1.ServerModel

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexAppCatalog) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		FeatureLimitExternal *HyperflexFeatureLimitExternalRef `json:"FeatureLimitExternal,omitempty"`

		FeatureLimitInternal *HyperflexFeatureLimitInternalRef `json:"FeatureLimitInternal,omitempty"`

		HxdpVersions []*HyperflexHxdpVersionRef `json:"HxdpVersions"`

		HyperflexCapabilityInfos []*HyperflexCapabilityInfoRef `json:"HyperflexCapabilityInfos"`

		HyperflexSoftwareCompatibilityInfos []*HclHyperflexSoftwareCompatibilityInfoRef `json:"HyperflexSoftwareCompatibilityInfos"`

		ServerFirmwareVersion *HyperflexServerFirmwareVersionRef `json:"ServerFirmwareVersion,omitempty"`

		ServerModel *HyperflexServerModelRef `json:"ServerModel,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.FeatureLimitExternal = m.FeatureLimitExternal

	dataAO1.FeatureLimitInternal = m.FeatureLimitInternal

	dataAO1.HxdpVersions = m.HxdpVersions

	dataAO1.HyperflexCapabilityInfos = m.HyperflexCapabilityInfos

	dataAO1.HyperflexSoftwareCompatibilityInfos = m.HyperflexSoftwareCompatibilityInfos

	dataAO1.ServerFirmwareVersion = m.ServerFirmwareVersion

	dataAO1.ServerModel = m.ServerModel

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex app catalog
func (m *HyperflexAppCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureLimitExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureLimitInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHxdpVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperflexCapabilityInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperflexSoftwareCompatibilityInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerFirmwareVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexAppCatalog) validateFeatureLimitExternal(formats strfmt.Registry) error {

	if swag.IsZero(m.FeatureLimitExternal) { // not required
		return nil
	}

	if m.FeatureLimitExternal != nil {
		if err := m.FeatureLimitExternal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeatureLimitExternal")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexAppCatalog) validateFeatureLimitInternal(formats strfmt.Registry) error {

	if swag.IsZero(m.FeatureLimitInternal) { // not required
		return nil
	}

	if m.FeatureLimitInternal != nil {
		if err := m.FeatureLimitInternal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeatureLimitInternal")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexAppCatalog) validateHxdpVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.HxdpVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.HxdpVersions); i++ {
		if swag.IsZero(m.HxdpVersions[i]) { // not required
			continue
		}

		if m.HxdpVersions[i] != nil {
			if err := m.HxdpVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("HxdpVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexAppCatalog) validateHyperflexCapabilityInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.HyperflexCapabilityInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.HyperflexCapabilityInfos); i++ {
		if swag.IsZero(m.HyperflexCapabilityInfos[i]) { // not required
			continue
		}

		if m.HyperflexCapabilityInfos[i] != nil {
			if err := m.HyperflexCapabilityInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("HyperflexCapabilityInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexAppCatalog) validateHyperflexSoftwareCompatibilityInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.HyperflexSoftwareCompatibilityInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.HyperflexSoftwareCompatibilityInfos); i++ {
		if swag.IsZero(m.HyperflexSoftwareCompatibilityInfos[i]) { // not required
			continue
		}

		if m.HyperflexSoftwareCompatibilityInfos[i] != nil {
			if err := m.HyperflexSoftwareCompatibilityInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("HyperflexSoftwareCompatibilityInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexAppCatalog) validateServerFirmwareVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerFirmwareVersion) { // not required
		return nil
	}

	if m.ServerFirmwareVersion != nil {
		if err := m.ServerFirmwareVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServerFirmwareVersion")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexAppCatalog) validateServerModel(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerModel) { // not required
		return nil
	}

	if m.ServerModel != nil {
		if err := m.ServerModel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServerModel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexAppCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexAppCatalog) UnmarshalBinary(b []byte) error {
	var res HyperflexAppCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
