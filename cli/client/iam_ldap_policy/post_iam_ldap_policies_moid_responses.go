// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamLdapPoliciesMoidReader is a Reader for the PostIamLdapPoliciesMoid structure.
type PostIamLdapPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamLdapPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamLdapPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamLdapPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamLdapPoliciesMoidCreated creates a PostIamLdapPoliciesMoidCreated with default headers values
func NewPostIamLdapPoliciesMoidCreated() *PostIamLdapPoliciesMoidCreated {
	return &PostIamLdapPoliciesMoidCreated{}
}

/*PostIamLdapPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostIamLdapPoliciesMoidCreated struct {
}

func (o *PostIamLdapPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/LdapPolicies/{moid}][%d] postIamLdapPoliciesMoidCreated ", 201)
}

func (o *PostIamLdapPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamLdapPoliciesMoidDefault creates a PostIamLdapPoliciesMoidDefault with default headers values
func NewPostIamLdapPoliciesMoidDefault(code int) *PostIamLdapPoliciesMoidDefault {
	return &PostIamLdapPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostIamLdapPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamLdapPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam ldap policies moid default response
func (o *PostIamLdapPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamLdapPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/LdapPolicies/{moid}][%d] PostIamLdapPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamLdapPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamLdapPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}