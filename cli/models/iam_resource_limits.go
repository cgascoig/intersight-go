// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IamResourceLimits Iam:Resource Limits
//
// The resource limits used to limit resources such as User accounts.
//
// swagger:model iamResourceLimits
type IamResourceLimits struct {
	MoBaseMo

	// A collection of references to the [iam.Account](mo://iam.Account) Managed Object.
	//
	// When this managed object is deleted, the referenced [iam.Account](mo://iam.Account) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// The maximum number of users allowed in an account. The default value is 200.
	//
	// Read Only: true
	PerAccountUserLimit int64 `json:"PerAccountUserLimit,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamResourceLimits) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		PerAccountUserLimit int64 `json:"PerAccountUserLimit,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.PerAccountUserLimit = dataAO1.PerAccountUserLimit

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamResourceLimits) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		PerAccountUserLimit int64 `json:"PerAccountUserLimit,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.PerAccountUserLimit = m.PerAccountUserLimit

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam resource limits
func (m *IamResourceLimits) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamResourceLimits) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamResourceLimits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamResourceLimits) UnmarshalBinary(b []byte) error {
	var res IamResourceLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
