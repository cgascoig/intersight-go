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

// HclHardwareCompatibilityProfile Hcl:Hardware Compatibility Profile
//
// Profile giving server hardware details, OS details and UCS Version details.
//
// swagger:model hclHardwareCompatibilityProfile
type HclHardwareCompatibilityProfile struct {

	// Url for the ISO with the drivers supported for the server.
	//
	DriverIsoURL string `json:"DriverIsoUrl,omitempty"`

	// Error code indicating the compatibility status.
	//
	// Read Only: true
	// Enum: [Success Unknown UnknownServer InvalidUcsVersion ProcessorNotSupported OSNotSupported OSUnknown UCSVersionNotSupported UcsVersionServerOSCombinationNotSupported ProductUnknown ProductNotSupported DriverNameNotSupported FirmwareVersionNotSupported DriverVersionNotSupported FirmwareVersionDriverVersionCombinationNotSupported FirmwareVersionAndDriverVersionNotSupported FirmwareVersionAndDriverNameNotSupported InternalError MarshallingError Exempted]
	ErrorCode string `json:"ErrorCode,omitempty"`

	// Identifier of the hardware compatibility profile.
	//
	ID string `json:"Id,omitempty"`

	// Vendor of the Operating System running on the server.
	//
	OsVendor string `json:"OsVendor,omitempty"`

	// Version of the Operating System running on the server.
	//
	OsVersion string `json:"OsVersion,omitempty"`

	// Model of the processor present in the server.
	//
	ProcessorModel string `json:"ProcessorModel,omitempty"`

	// List of the products (adapters/storage controllers) for which compatibility status needs to be checked.
	//
	Products []*HclProduct `json:"Products"`

	// Model of the server as returned by UCSM/CIMC XML API.
	//
	ServerModel string `json:"ServerModel,omitempty"`

	// Revision of the server model.
	//
	ServerRevision string `json:"ServerRevision,omitempty"`

	// Version of the UCS software.
	//
	UcsVersion string `json:"UcsVersion,omitempty"`

	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	//
	// Enum: [UCSM IMC]
	VersionType *string `json:"VersionType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HclHardwareCompatibilityProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		DriverIsoURL string `json:"DriverIsoUrl,omitempty"`

		ErrorCode string `json:"ErrorCode,omitempty"`

		ID string `json:"Id,omitempty"`

		OsVendor string `json:"OsVendor,omitempty"`

		OsVersion string `json:"OsVersion,omitempty"`

		ProcessorModel string `json:"ProcessorModel,omitempty"`

		Products []*HclProduct `json:"Products"`

		ServerModel string `json:"ServerModel,omitempty"`

		ServerRevision string `json:"ServerRevision,omitempty"`

		UcsVersion string `json:"UcsVersion,omitempty"`

		VersionType *string `json:"VersionType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.DriverIsoURL = dataAO0.DriverIsoURL

	m.ErrorCode = dataAO0.ErrorCode

	m.ID = dataAO0.ID

	m.OsVendor = dataAO0.OsVendor

	m.OsVersion = dataAO0.OsVersion

	m.ProcessorModel = dataAO0.ProcessorModel

	m.Products = dataAO0.Products

	m.ServerModel = dataAO0.ServerModel

	m.ServerRevision = dataAO0.ServerRevision

	m.UcsVersion = dataAO0.UcsVersion

	m.VersionType = dataAO0.VersionType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HclHardwareCompatibilityProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		DriverIsoURL string `json:"DriverIsoUrl,omitempty"`

		ErrorCode string `json:"ErrorCode,omitempty"`

		ID string `json:"Id,omitempty"`

		OsVendor string `json:"OsVendor,omitempty"`

		OsVersion string `json:"OsVersion,omitempty"`

		ProcessorModel string `json:"ProcessorModel,omitempty"`

		Products []*HclProduct `json:"Products"`

		ServerModel string `json:"ServerModel,omitempty"`

		ServerRevision string `json:"ServerRevision,omitempty"`

		UcsVersion string `json:"UcsVersion,omitempty"`

		VersionType *string `json:"VersionType,omitempty"`
	}

	dataAO0.DriverIsoURL = m.DriverIsoURL

	dataAO0.ErrorCode = m.ErrorCode

	dataAO0.ID = m.ID

	dataAO0.OsVendor = m.OsVendor

	dataAO0.OsVersion = m.OsVersion

	dataAO0.ProcessorModel = m.ProcessorModel

	dataAO0.Products = m.Products

	dataAO0.ServerModel = m.ServerModel

	dataAO0.ServerRevision = m.ServerRevision

	dataAO0.UcsVersion = m.UcsVersion

	dataAO0.VersionType = m.VersionType

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hcl hardware compatibility profile
func (m *HclHardwareCompatibilityProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hclHardwareCompatibilityProfileTypeErrorCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Success","Unknown","UnknownServer","InvalidUcsVersion","ProcessorNotSupported","OSNotSupported","OSUnknown","UCSVersionNotSupported","UcsVersionServerOSCombinationNotSupported","ProductUnknown","ProductNotSupported","DriverNameNotSupported","FirmwareVersionNotSupported","DriverVersionNotSupported","FirmwareVersionDriverVersionCombinationNotSupported","FirmwareVersionAndDriverVersionNotSupported","FirmwareVersionAndDriverNameNotSupported","InternalError","MarshallingError","Exempted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hclHardwareCompatibilityProfileTypeErrorCodePropEnum = append(hclHardwareCompatibilityProfileTypeErrorCodePropEnum, v)
	}
}

// property enum
func (m *HclHardwareCompatibilityProfile) validateErrorCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hclHardwareCompatibilityProfileTypeErrorCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HclHardwareCompatibilityProfile) validateErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorCodeEnum("ErrorCode", "body", m.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (m *HclHardwareCompatibilityProfile) validateProducts(formats strfmt.Registry) error {

	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var hclHardwareCompatibilityProfileTypeVersionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UCSM","IMC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hclHardwareCompatibilityProfileTypeVersionTypePropEnum = append(hclHardwareCompatibilityProfileTypeVersionTypePropEnum, v)
	}
}

// property enum
func (m *HclHardwareCompatibilityProfile) validateVersionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hclHardwareCompatibilityProfileTypeVersionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HclHardwareCompatibilityProfile) validateVersionType(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionTypeEnum("VersionType", "body", *m.VersionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HclHardwareCompatibilityProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HclHardwareCompatibilityProfile) UnmarshalBinary(b []byte) error {
	var res HclHardwareCompatibilityProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
