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

// StorageDiskGroupPolicyList storage disk group policy list
// swagger:model storageDiskGroupPolicyList
type StorageDiskGroupPolicyList struct {

	// The number of storageDiskGroupPolicies matching your request in total for all pages.
	Count int32 `json:"Count,omitempty"`

	// The array of storageDiskGroupPolicies matching your request.
	Results []*StorageDiskGroupPolicy `json:"Results"`
}

// Validate validates this storage disk group policy list
func (m *StorageDiskGroupPolicyList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageDiskGroupPolicyList) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageDiskGroupPolicyList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageDiskGroupPolicyList) UnmarshalBinary(b []byte) error {
	var res StorageDiskGroupPolicyList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
