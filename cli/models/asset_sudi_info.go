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

// AssetSudiInfo Asset:Sudi Info
//
// The SUDI is an X.509v3 certificate, which maintains the product identifier and serial number. The identity is implemented at manufacturing and chained to a publicly identifiable root certificate authority. It can be used as an unchangeable identity for configuration, security, auditing, and management. This strucure contains the SUDI information read from the device's Trust Anchor Module (TAM).
//
// swagger:model assetSudiInfo
type AssetSudiInfo struct {

	// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
	//
	Pid string `json:"Pid,omitempty"`

	// The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.
	//
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.
	//
	Signature string `json:"Signature,omitempty"`

	// The validation status of the device.
	//
	// Enum: [DeviceStatusUnknown Verified CertificateValidationFailed UnsupportedFirmware UnsupportedHardware DeviceNotResponding]
	Status *string `json:"Status,omitempty"`

	// The X.509 SUDI Leaf Certificate from the Trust Anchor Module. The certificate is serialized in PEM format (Base64 encoded DER certificate).
	//
	SudiCertificate *X509Certificate `json:"SudiCertificate,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetSudiInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Pid string `json:"Pid,omitempty"`

		SerialNumber string `json:"SerialNumber,omitempty"`

		Signature string `json:"Signature,omitempty"`

		Status *string `json:"Status,omitempty"`

		SudiCertificate *X509Certificate `json:"SudiCertificate,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Pid = dataAO0.Pid

	m.SerialNumber = dataAO0.SerialNumber

	m.Signature = dataAO0.Signature

	m.Status = dataAO0.Status

	m.SudiCertificate = dataAO0.SudiCertificate

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetSudiInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Pid string `json:"Pid,omitempty"`

		SerialNumber string `json:"SerialNumber,omitempty"`

		Signature string `json:"Signature,omitempty"`

		Status *string `json:"Status,omitempty"`

		SudiCertificate *X509Certificate `json:"SudiCertificate,omitempty"`
	}

	dataAO0.Pid = m.Pid

	dataAO0.SerialNumber = m.SerialNumber

	dataAO0.Signature = m.Signature

	dataAO0.Status = m.Status

	dataAO0.SudiCertificate = m.SudiCertificate

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset sudi info
func (m *AssetSudiInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSudiCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var assetSudiInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DeviceStatusUnknown","Verified","CertificateValidationFailed","UnsupportedFirmware","UnsupportedHardware","DeviceNotResponding"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetSudiInfoTypeStatusPropEnum = append(assetSudiInfoTypeStatusPropEnum, v)
	}
}

// property enum
func (m *AssetSudiInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetSudiInfoTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetSudiInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AssetSudiInfo) validateSudiCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.SudiCertificate) { // not required
		return nil
	}

	if m.SudiCertificate != nil {
		if err := m.SudiCertificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SudiCertificate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetSudiInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetSudiInfo) UnmarshalBinary(b []byte) error {
	var res AssetSudiInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
