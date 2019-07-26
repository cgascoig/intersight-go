// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_user_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamEndPointUserPoliciesReader is a Reader for the PostIamEndPointUserPolicies structure.
type PostIamEndPointUserPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamEndPointUserPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamEndPointUserPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamEndPointUserPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamEndPointUserPoliciesCreated creates a PostIamEndPointUserPoliciesCreated with default headers values
func NewPostIamEndPointUserPoliciesCreated() *PostIamEndPointUserPoliciesCreated {
	return &PostIamEndPointUserPoliciesCreated{}
}

/*PostIamEndPointUserPoliciesCreated handles this case with default header values.

Null response
*/
type PostIamEndPointUserPoliciesCreated struct {
}

func (o *PostIamEndPointUserPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /iam/EndPointUserPolicies][%d] postIamEndPointUserPoliciesCreated ", 201)
}

func (o *PostIamEndPointUserPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamEndPointUserPoliciesDefault creates a PostIamEndPointUserPoliciesDefault with default headers values
func NewPostIamEndPointUserPoliciesDefault(code int) *PostIamEndPointUserPoliciesDefault {
	return &PostIamEndPointUserPoliciesDefault{
		_statusCode: code,
	}
}

/*PostIamEndPointUserPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostIamEndPointUserPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam end point user policies default response
func (o *PostIamEndPointUserPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostIamEndPointUserPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /iam/EndPointUserPolicies][%d] PostIamEndPointUserPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamEndPointUserPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamEndPointUserPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
