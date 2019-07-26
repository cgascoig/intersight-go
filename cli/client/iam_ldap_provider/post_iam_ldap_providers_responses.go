// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamLdapProvidersReader is a Reader for the PostIamLdapProviders structure.
type PostIamLdapProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamLdapProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamLdapProvidersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamLdapProvidersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamLdapProvidersCreated creates a PostIamLdapProvidersCreated with default headers values
func NewPostIamLdapProvidersCreated() *PostIamLdapProvidersCreated {
	return &PostIamLdapProvidersCreated{}
}

/*PostIamLdapProvidersCreated handles this case with default header values.

Null response
*/
type PostIamLdapProvidersCreated struct {
}

func (o *PostIamLdapProvidersCreated) Error() string {
	return fmt.Sprintf("[POST /iam/LdapProviders][%d] postIamLdapProvidersCreated ", 201)
}

func (o *PostIamLdapProvidersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamLdapProvidersDefault creates a PostIamLdapProvidersDefault with default headers values
func NewPostIamLdapProvidersDefault(code int) *PostIamLdapProvidersDefault {
	return &PostIamLdapProvidersDefault{
		_statusCode: code,
	}
}

/*PostIamLdapProvidersDefault handles this case with default header values.

unexpected error
*/
type PostIamLdapProvidersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam ldap providers default response
func (o *PostIamLdapProvidersDefault) Code() int {
	return o._statusCode
}

func (o *PostIamLdapProvidersDefault) Error() string {
	return fmt.Sprintf("[POST /iam/LdapProviders][%d] PostIamLdapProviders default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamLdapProvidersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamLdapProvidersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
