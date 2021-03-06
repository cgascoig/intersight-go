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

// DeleteIamLdapProvidersMoidReader is a Reader for the DeleteIamLdapProvidersMoid structure.
type DeleteIamLdapProvidersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIamLdapProvidersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIamLdapProvidersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIamLdapProvidersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIamLdapProvidersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIamLdapProvidersMoidOK creates a DeleteIamLdapProvidersMoidOK with default headers values
func NewDeleteIamLdapProvidersMoidOK() *DeleteIamLdapProvidersMoidOK {
	return &DeleteIamLdapProvidersMoidOK{}
}

/*DeleteIamLdapProvidersMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteIamLdapProvidersMoidOK struct {
}

func (o *DeleteIamLdapProvidersMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapProviders/{moid}][%d] deleteIamLdapProvidersMoidOK ", 200)
}

func (o *DeleteIamLdapProvidersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamLdapProvidersMoidNotFound creates a DeleteIamLdapProvidersMoidNotFound with default headers values
func NewDeleteIamLdapProvidersMoidNotFound() *DeleteIamLdapProvidersMoidNotFound {
	return &DeleteIamLdapProvidersMoidNotFound{}
}

/*DeleteIamLdapProvidersMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteIamLdapProvidersMoidNotFound struct {
}

func (o *DeleteIamLdapProvidersMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapProviders/{moid}][%d] deleteIamLdapProvidersMoidNotFound ", 404)
}

func (o *DeleteIamLdapProvidersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamLdapProvidersMoidDefault creates a DeleteIamLdapProvidersMoidDefault with default headers values
func NewDeleteIamLdapProvidersMoidDefault(code int) *DeleteIamLdapProvidersMoidDefault {
	return &DeleteIamLdapProvidersMoidDefault{
		_statusCode: code,
	}
}

/*DeleteIamLdapProvidersMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteIamLdapProvidersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete iam ldap providers moid default response
func (o *DeleteIamLdapProvidersMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIamLdapProvidersMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapProviders/{moid}][%d] DeleteIamLdapProvidersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIamLdapProvidersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIamLdapProvidersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
