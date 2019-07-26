// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageRemoteKeySetting Remote Key Management
//
// This type models the remote key configurarion required for disks encryptions. KMIP is the only remote key protocol supported in the policy.
//
// swagger:model storageRemoteKeySetting
type StorageRemoteKeySetting struct {

	// is password set
	IsPasswordSet bool `json:"IsPasswordSet,omitempty"`

	// This property is used to specify password for the KMIP server login.
	//
	Password string `json:"Password,omitempty"`

	// This property is used to port to which the KMIP client should connect.
	//
	Port int64 `json:"Port,omitempty"`

	// This property is used to store the address of the KMIP server. It could be an IPv4 address or an IPv6 address or hostname. Hostnames are valid only when Inband is configured for the CIMC address.
	//
	PrimaryServer string `json:"PrimaryServer,omitempty"`

	// This property is used to store the address of the KMIP server. It could be an IPv4 address or an IPv6 address or hostname. Hostnames are valid only when Inband is configured for the CIMC address.
	//
	SecondaryServer string `json:"SecondaryServer,omitempty"`

	// This property is used to store the certificate/ public key of the KMIP server This is required for initiating secure communication with the server.
	//
	ServerCertificate string `json:"ServerCertificate,omitempty"`

	// This property is used to specify user name for the KMIP server login.
	//
	Username string `json:"Username,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageRemoteKeySetting) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		IsPasswordSet bool `json:"IsPasswordSet,omitempty"`

		Password string `json:"Password,omitempty"`

		Port int64 `json:"Port,omitempty"`

		PrimaryServer string `json:"PrimaryServer,omitempty"`

		SecondaryServer string `json:"SecondaryServer,omitempty"`

		ServerCertificate string `json:"ServerCertificate,omitempty"`

		Username string `json:"Username,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.IsPasswordSet = dataAO0.IsPasswordSet

	m.Password = dataAO0.Password

	m.Port = dataAO0.Port

	m.PrimaryServer = dataAO0.PrimaryServer

	m.SecondaryServer = dataAO0.SecondaryServer

	m.ServerCertificate = dataAO0.ServerCertificate

	m.Username = dataAO0.Username

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageRemoteKeySetting) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		IsPasswordSet bool `json:"IsPasswordSet,omitempty"`

		Password string `json:"Password,omitempty"`

		Port int64 `json:"Port,omitempty"`

		PrimaryServer string `json:"PrimaryServer,omitempty"`

		SecondaryServer string `json:"SecondaryServer,omitempty"`

		ServerCertificate string `json:"ServerCertificate,omitempty"`

		Username string `json:"Username,omitempty"`
	}

	dataAO0.IsPasswordSet = m.IsPasswordSet

	dataAO0.Password = m.Password

	dataAO0.Port = m.Port

	dataAO0.PrimaryServer = m.PrimaryServer

	dataAO0.SecondaryServer = m.SecondaryServer

	dataAO0.ServerCertificate = m.ServerCertificate

	dataAO0.Username = m.Username

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage remote key setting
func (m *StorageRemoteKeySetting) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageRemoteKeySetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageRemoteKeySetting) UnmarshalBinary(b []byte) error {
	var res StorageRemoteKeySetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
