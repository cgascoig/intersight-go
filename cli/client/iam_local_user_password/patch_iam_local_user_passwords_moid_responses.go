// Code generated by go-swagger; DO NOT EDIT.

package iam_local_user_password

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchIamLocalUserPasswordsMoidReader is a Reader for the PatchIamLocalUserPasswordsMoid structure.
type PatchIamLocalUserPasswordsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIamLocalUserPasswordsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchIamLocalUserPasswordsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchIamLocalUserPasswordsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchIamLocalUserPasswordsMoidCreated creates a PatchIamLocalUserPasswordsMoidCreated with default headers values
func NewPatchIamLocalUserPasswordsMoidCreated() *PatchIamLocalUserPasswordsMoidCreated {
	return &PatchIamLocalUserPasswordsMoidCreated{}
}

/*PatchIamLocalUserPasswordsMoidCreated handles this case with default header values.

Null response
*/
type PatchIamLocalUserPasswordsMoidCreated struct {
}

func (o *PatchIamLocalUserPasswordsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /iam/LocalUserPasswords/{moid}][%d] patchIamLocalUserPasswordsMoidCreated ", 201)
}

func (o *PatchIamLocalUserPasswordsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchIamLocalUserPasswordsMoidDefault creates a PatchIamLocalUserPasswordsMoidDefault with default headers values
func NewPatchIamLocalUserPasswordsMoidDefault(code int) *PatchIamLocalUserPasswordsMoidDefault {
	return &PatchIamLocalUserPasswordsMoidDefault{
		_statusCode: code,
	}
}

/*PatchIamLocalUserPasswordsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchIamLocalUserPasswordsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch iam local user passwords moid default response
func (o *PatchIamLocalUserPasswordsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchIamLocalUserPasswordsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /iam/LocalUserPasswords/{moid}][%d] PatchIamLocalUserPasswordsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchIamLocalUserPasswordsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchIamLocalUserPasswordsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}