// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FeedbackFeedbackPost Feedback:Feedback Post
//
// Initial feedback submitted by the user from UI.
//
// swagger:model feedbackFeedbackPost
type FeedbackFeedbackPost struct {
	MoBaseMo

	// An Intersight Account. An account provides personalized access to Intersight.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// Feedback collected from the user and latest activity.
	//
	FeedbackData *FeedbackFeedbackData `json:"FeedbackData,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FeedbackFeedbackPost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		FeedbackData *FeedbackFeedbackData `json:"FeedbackData,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.FeedbackData = dataAO1.FeedbackData

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FeedbackFeedbackPost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		FeedbackData *FeedbackFeedbackData `json:"FeedbackData,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.FeedbackData = m.FeedbackData

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this feedback feedback post
func (m *FeedbackFeedbackPost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedbackData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeedbackFeedbackPost) validateAccount(formats strfmt.Registry) error {

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

func (m *FeedbackFeedbackPost) validateFeedbackData(formats strfmt.Registry) error {

	if swag.IsZero(m.FeedbackData) { // not required
		return nil
	}

	if m.FeedbackData != nil {
		if err := m.FeedbackData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeedbackData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedbackFeedbackPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedbackFeedbackPost) UnmarshalBinary(b []byte) error {
	var res FeedbackFeedbackPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
