// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexIPAddrRange Hyperflex:Ip Addr Range
//
// A range of IPv4 addresses.
//
// The range is inclusive and comprised of a start IP address, an end IP address, netmask, and default gateway.
//
// swagger:model hyperflexIpAddrRange
type HyperflexIPAddrRange struct {

	// The end IPv4 address of the range.
	//
	EndAddr string `json:"EndAddr,omitempty"`

	// The default gateway for the start and end IPv4 addresses.
	//
	Gateway string `json:"Gateway,omitempty"`

	// The netmask specified in dot decimal notation.
	//
	// The start address, end address, and gateway must all be within the network specified by this netmask.
	//
	//
	Netmask string `json:"Netmask,omitempty"`

	// The start IPv4 address of the range.
	//
	StartAddr string `json:"StartAddr,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexIPAddrRange) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		EndAddr string `json:"EndAddr,omitempty"`

		Gateway string `json:"Gateway,omitempty"`

		Netmask string `json:"Netmask,omitempty"`

		StartAddr string `json:"StartAddr,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.EndAddr = dataAO0.EndAddr

	m.Gateway = dataAO0.Gateway

	m.Netmask = dataAO0.Netmask

	m.StartAddr = dataAO0.StartAddr

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexIPAddrRange) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		EndAddr string `json:"EndAddr,omitempty"`

		Gateway string `json:"Gateway,omitempty"`

		Netmask string `json:"Netmask,omitempty"`

		StartAddr string `json:"StartAddr,omitempty"`
	}

	dataAO0.EndAddr = m.EndAddr

	dataAO0.Gateway = m.Gateway

	dataAO0.Netmask = m.Netmask

	dataAO0.StartAddr = m.StartAddr

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex Ip addr range
func (m *HyperflexIPAddrRange) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexIPAddrRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexIPAddrRange) UnmarshalBinary(b []byte) error {
	var res HyperflexIPAddrRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
