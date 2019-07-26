// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommIPV6Interface Comm:Ip V6 Interface
//
// The configuration data of a single IPv6 interface, including IP address, IPv6 prefix and default gateway.
//
// swagger:model commIpV6Interface
type CommIPV6Interface struct {

	// The IPv6 address of the default gateway.
	//
	Gateway string `json:"Gateway,omitempty"`

	// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
	//
	IPAddress string `json:"IpAddress,omitempty"`

	// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
	//
	Prefix string `json:"Prefix,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommIPV6Interface) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Gateway string `json:"Gateway,omitempty"`

		IPAddress string `json:"IpAddress,omitempty"`

		Prefix string `json:"Prefix,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Gateway = dataAO0.Gateway

	m.IPAddress = dataAO0.IPAddress

	m.Prefix = dataAO0.Prefix

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommIPV6Interface) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Gateway string `json:"Gateway,omitempty"`

		IPAddress string `json:"IpAddress,omitempty"`

		Prefix string `json:"Prefix,omitempty"`
	}

	dataAO0.Gateway = m.Gateway

	dataAO0.IPAddress = m.IPAddress

	dataAO0.Prefix = m.Prefix

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comm Ip v6 interface
func (m *CommIPV6Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CommIPV6Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommIPV6Interface) UnmarshalBinary(b []byte) error {
	var res CommIPV6Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
