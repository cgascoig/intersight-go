// Code generated by go-swagger; DO NOT EDIT.

package kvm_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostKvmPoliciesReader is a Reader for the PostKvmPolicies structure.
type PostKvmPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKvmPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostKvmPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostKvmPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostKvmPoliciesCreated creates a PostKvmPoliciesCreated with default headers values
func NewPostKvmPoliciesCreated() *PostKvmPoliciesCreated {
	return &PostKvmPoliciesCreated{}
}

/*PostKvmPoliciesCreated handles this case with default header values.

Null response
*/
type PostKvmPoliciesCreated struct {
}

func (o *PostKvmPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /kvm/Policies][%d] postKvmPoliciesCreated ", 201)
}

func (o *PostKvmPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostKvmPoliciesDefault creates a PostKvmPoliciesDefault with default headers values
func NewPostKvmPoliciesDefault(code int) *PostKvmPoliciesDefault {
	return &PostKvmPoliciesDefault{
		_statusCode: code,
	}
}

/*PostKvmPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostKvmPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post kvm policies default response
func (o *PostKvmPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostKvmPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /kvm/Policies][%d] PostKvmPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostKvmPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostKvmPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
