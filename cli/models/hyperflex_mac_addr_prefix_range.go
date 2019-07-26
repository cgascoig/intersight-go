// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HyperflexMacAddrPrefixRange Hyperflex:Mac Addr Prefix Range
//
// A MAC address prefix range.
//
// The range is inclusive and comprised of a start and end MAC addresses.
// A single address can be specified by setting it as the start and end of the range.
//
// swagger:model hyperflexMacAddrPrefixRange
type HyperflexMacAddrPrefixRange struct {

	// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	//
	EndAddr string `json:"EndAddr,omitempty"`

	// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	//
	StartAddr string `json:"StartAddr,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexMacAddrPrefixRange) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		EndAddr string `json:"EndAddr,omitempty"`

		StartAddr string `json:"StartAddr,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.EndAddr = dataAO0.EndAddr

	m.StartAddr = dataAO0.StartAddr

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexMacAddrPrefixRange) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		EndAddr string `json:"EndAddr,omitempty"`

		StartAddr string `json:"StartAddr,omitempty"`
	}

	dataAO0.EndAddr = m.EndAddr

	dataAO0.StartAddr = m.StartAddr

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex mac addr prefix range
func (m *HyperflexMacAddrPrefixRange) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexMacAddrPrefixRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexMacAddrPrefixRange) UnmarshalBinary(b []byte) error {
	var res HyperflexMacAddrPrefixRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}