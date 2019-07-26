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

// DeleteIamLdapPoliciesMoidReader is a Reader for the DeleteIamLdapPoliciesMoid structure.
type DeleteIamLdapPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIamLdapPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIamLdapPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIamLdapPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIamLdapPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIamLdapPoliciesMoidOK creates a DeleteIamLdapPoliciesMoidOK with default headers values
func NewDeleteIamLdapPoliciesMoidOK() *DeleteIamLdapPoliciesMoidOK {
	return &DeleteIamLdapPoliciesMoidOK{}
}

/*DeleteIamLdapPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteIamLdapPoliciesMoidOK struct {
}

func (o *DeleteIamLdapPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapPolicies/{moid}][%d] deleteIamLdapPoliciesMoidOK ", 200)
}

func (o *DeleteIamLdapPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamLdapPoliciesMoidNotFound creates a DeleteIamLdapPoliciesMoidNotFound with default headers values
func NewDeleteIamLdapPoliciesMoidNotFound() *DeleteIamLdapPoliciesMoidNotFound {
	return &DeleteIamLdapPoliciesMoidNotFound{}
}

/*DeleteIamLdapPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteIamLdapPoliciesMoidNotFound struct {
}

func (o *DeleteIamLdapPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapPolicies/{moid}][%d] deleteIamLdapPoliciesMoidNotFound ", 404)
}

func (o *DeleteIamLdapPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamLdapPoliciesMoidDefault creates a DeleteIamLdapPoliciesMoidDefault with default headers values
func NewDeleteIamLdapPoliciesMoidDefault(code int) *DeleteIamLdapPoliciesMoidDefault {
	return &DeleteIamLdapPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteIamLdapPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteIamLdapPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete iam ldap policies moid default response
func (o *DeleteIamLdapPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIamLdapPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapPolicies/{moid}][%d] DeleteIamLdapPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIamLdapPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIamLdapPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
