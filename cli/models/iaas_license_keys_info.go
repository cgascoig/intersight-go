// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IaasLicenseKeysInfo Iaas:License Keys Info
//
// License list with the utilization info for UCSD.
//
// swagger:model iaasLicenseKeysInfo
type IaasLicenseKeysInfo struct {

	// Number of licenses available for the UCSD PID (Product ID).
	//
	// Read Only: true
	Count int64 `json:"Count,omitempty"`

	// Expiration date for the license.
	//
	// Read Only: true
	ExpirationDate string `json:"ExpirationDate,omitempty"`

	// UCS Director Unique license ID.
	//
	// Read Only: true
	LicenseID string `json:"LicenseId,omitempty"`

	// PID (Product ID) for UCSD License.
	//
	// Read Only: true
	Pid string `json:"Pid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IaasLicenseKeysInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Count int64 `json:"Count,omitempty"`

		ExpirationDate string `json:"ExpirationDate,omitempty"`

		LicenseID string `json:"LicenseId,omitempty"`

		Pid string `json:"Pid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Count = dataAO0.Count

	m.ExpirationDate = dataAO0.ExpirationDate

	m.LicenseID = dataAO0.LicenseID

	m.Pid = dataAO0.Pid

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IaasLicenseKeysInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Count int64 `json:"Count,omitempty"`

		ExpirationDate string `json:"ExpirationDate,omitempty"`

		LicenseID string `json:"LicenseId,omitempty"`

		Pid string `json:"Pid,omitempty"`
	}

	dataAO0.Count = m.Count

	dataAO0.ExpirationDate = m.ExpirationDate

	dataAO0.LicenseID = m.LicenseID

	dataAO0.Pid = m.Pid

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iaas license keys info
func (m *IaasLicenseKeysInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IaasLicenseKeysInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasLicenseKeysInfo) UnmarshalBinary(b []byte) error {
	var res IaasLicenseKeysInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}