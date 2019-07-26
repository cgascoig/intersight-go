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

// FeedbackFeedbackData Feedback:Feedback Data
//
// Feedback data that collected on the UI from user.
//
// swagger:model feedbackFeedbackData
type FeedbackFeedbackData struct {

	// Text of the feedback as provided by the user, if it is a bug or a comment.
	//
	Comment string `json:"Comment,omitempty"`

	// User's email address details.
	//
	Email string `json:"Email,omitempty"`

	// Evalation rating as provided by the user to capture user sentiment regarding the issue.
	//
	// Enum: [Excellent Poor Fair Good]
	Evaluation *string `json:"Evaluation,omitempty"`

	// If a user is open for follow-up or not.
	//
	FollowUp bool `json:"FollowUp,omitempty"`

	// Bunch of last traceId for reproducing user last activity.
	//
	TraceIds interface{} `json:"TraceIds,omitempty"`

	// Type of the feedback from user.
	//
	// Enum: [Evaluation Bug]
	Type *string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FeedbackFeedbackData) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Comment string `json:"Comment,omitempty"`

		Email string `json:"Email,omitempty"`

		Evaluation *string `json:"Evaluation,omitempty"`

		FollowUp bool `json:"FollowUp,omitempty"`

		TraceIds interface{} `json:"TraceIds,omitempty"`

		Type *string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Comment = dataAO0.Comment

	m.Email = dataAO0.Email

	m.Evaluation = dataAO0.Evaluation

	m.FollowUp = dataAO0.FollowUp

	m.TraceIds = dataAO0.TraceIds

	m.Type = dataAO0.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FeedbackFeedbackData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Comment string `json:"Comment,omitempty"`

		Email string `json:"Email,omitempty"`

		Evaluation *string `json:"Evaluation,omitempty"`

		FollowUp bool `json:"FollowUp,omitempty"`

		TraceIds interface{} `json:"TraceIds,omitempty"`

		Type *string `json:"Type,omitempty"`
	}

	dataAO0.Comment = m.Comment

	dataAO0.Email = m.Email

	dataAO0.Evaluation = m.Evaluation

	dataAO0.FollowUp = m.FollowUp

	dataAO0.TraceIds = m.TraceIds

	dataAO0.Type = m.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this feedback feedback data
func (m *FeedbackFeedbackData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var feedbackFeedbackDataTypeEvaluationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Excellent","Poor","Fair","Good"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedbackFeedbackDataTypeEvaluationPropEnum = append(feedbackFeedbackDataTypeEvaluationPropEnum, v)
	}
}

// property enum
func (m *FeedbackFeedbackData) validateEvaluationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, feedbackFeedbackDataTypeEvaluationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FeedbackFeedbackData) validateEvaluation(formats strfmt.Registry) error {

	if swag.IsZero(m.Evaluation) { // not required
		return nil
	}

	// value enum
	if err := m.validateEvaluationEnum("Evaluation", "body", *m.Evaluation); err != nil {
		return err
	}

	return nil
}

var feedbackFeedbackDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Evaluation","Bug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedbackFeedbackDataTypeTypePropEnum = append(feedbackFeedbackDataTypeTypePropEnum, v)
	}
}

// property enum
func (m *FeedbackFeedbackData) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, feedbackFeedbackDataTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FeedbackFeedbackData) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedbackFeedbackData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedbackFeedbackData) UnmarshalBinary(b []byte) error {
	var res FeedbackFeedbackData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}