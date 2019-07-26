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

// PatchIamLdapGroupsMoidReader is a Reader for the PatchIamLdapGroupsMoid structure.
type PatchIamLdapGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIamLdapGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchIamLdapGroupsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchIamLdapGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchIamLdapGroupsMoidCreated creates a PatchIamLdapGroupsMoidCreated with default headers values
func NewPatchIamLdapGroupsMoidCreated() *PatchIamLdapGroupsMoidCreated {
	return &PatchIamLdapGroupsMoidCreated{}
}

/*PatchIamLdapGroupsMoidCreated handles this case with default header values.

Null response
*/
type PatchIamLdapGroupsMoidCreated struct {
}

func (o *PatchIamLdapGroupsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /iam/LdapGroups/{moid}][%d] patchIamLdapGroupsMoidCreated ", 201)
}

func (o *PatchIamLdapGroupsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchIamLdapGroupsMoidDefault creates a PatchIamLdapGroupsMoidDefault with default headers values
func NewPatchIamLdapGroupsMoidDefault(code int) *PatchIamLdapGroupsMoidDefault {
	return &PatchIamLdapGroupsMoidDefault{
		_statusCode: code,
	}
}

/*PatchIamLdapGroupsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchIamLdapGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch iam ldap groups moid default response
func (o *PatchIamLdapGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchIamLdapGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /iam/LdapGroups/{moid}][%d] PatchIamLdapGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchIamLdapGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchIamLdapGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
