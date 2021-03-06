// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteIamLdapGroupsMoidReader is a Reader for the DeleteIamLdapGroupsMoid structure.
type DeleteIamLdapGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIamLdapGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIamLdapGroupsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIamLdapGroupsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIamLdapGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIamLdapGroupsMoidOK creates a DeleteIamLdapGroupsMoidOK with default headers values
func NewDeleteIamLdapGroupsMoidOK() *DeleteIamLdapGroupsMoidOK {
	return &DeleteIamLdapGroupsMoidOK{}
}

/*DeleteIamLdapGroupsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteIamLdapGroupsMoidOK struct {
}

func (o *DeleteIamLdapGroupsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapGroups/{moid}][%d] deleteIamLdapGroupsMoidOK ", 200)
}

func (o *DeleteIamLdapGroupsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamLdapGroupsMoidNotFound creates a DeleteIamLdapGroupsMoidNotFound with default headers values
func NewDeleteIamLdapGroupsMoidNotFound() *DeleteIamLdapGroupsMoidNotFound {
	return &DeleteIamLdapGroupsMoidNotFound{}
}

/*DeleteIamLdapGroupsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteIamLdapGroupsMoidNotFound struct {
}

func (o *DeleteIamLdapGroupsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapGroups/{moid}][%d] deleteIamLdapGroupsMoidNotFound ", 404)
}

func (o *DeleteIamLdapGroupsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamLdapGroupsMoidDefault creates a DeleteIamLdapGroupsMoidDefault with default headers values
func NewDeleteIamLdapGroupsMoidDefault(code int) *DeleteIamLdapGroupsMoidDefault {
	return &DeleteIamLdapGroupsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteIamLdapGroupsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteIamLdapGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete iam ldap groups moid default response
func (o *DeleteIamLdapGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIamLdapGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/LdapGroups/{moid}][%d] DeleteIamLdapGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIamLdapGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIamLdapGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
