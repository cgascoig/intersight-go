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

// PostIamLdapProvidersMoidReader is a Reader for the PostIamLdapProvidersMoid structure.
type PostIamLdapProvidersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamLdapProvidersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamLdapProvidersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamLdapProvidersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamLdapProvidersMoidCreated creates a PostIamLdapProvidersMoidCreated with default headers values
func NewPostIamLdapProvidersMoidCreated() *PostIamLdapProvidersMoidCreated {
	return &PostIamLdapProvidersMoidCreated{}
}

/*PostIamLdapProvidersMoidCreated handles this case with default header values.

Null response
*/
type PostIamLdapProvidersMoidCreated struct {
}

func (o *PostIamLdapProvidersMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/LdapProviders/{moid}][%d] postIamLdapProvidersMoidCreated ", 201)
}

func (o *PostIamLdapProvidersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamLdapProvidersMoidDefault creates a PostIamLdapProvidersMoidDefault with default headers values
func NewPostIamLdapProvidersMoidDefault(code int) *PostIamLdapProvidersMoidDefault {
	return &PostIamLdapProvidersMoidDefault{
		_statusCode: code,
	}
}

/*PostIamLdapProvidersMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamLdapProvidersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam ldap providers moid default response
func (o *PostIamLdapProvidersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamLdapProvidersMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/LdapProviders/{moid}][%d] PostIamLdapProvidersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamLdapProvidersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamLdapProvidersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
