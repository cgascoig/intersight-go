// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_administrator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchTestcryptAdministratorsMoidReader is a Reader for the PatchTestcryptAdministratorsMoid structure.
type PatchTestcryptAdministratorsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTestcryptAdministratorsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchTestcryptAdministratorsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchTestcryptAdministratorsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchTestcryptAdministratorsMoidCreated creates a PatchTestcryptAdministratorsMoidCreated with default headers values
func NewPatchTestcryptAdministratorsMoidCreated() *PatchTestcryptAdministratorsMoidCreated {
	return &PatchTestcryptAdministratorsMoidCreated{}
}

/*PatchTestcryptAdministratorsMoidCreated handles this case with default header values.

Null response
*/
type PatchTestcryptAdministratorsMoidCreated struct {
}

func (o *PatchTestcryptAdministratorsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /testcrypt/Administrators/{moid}][%d] patchTestcryptAdministratorsMoidCreated ", 201)
}

func (o *PatchTestcryptAdministratorsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchTestcryptAdministratorsMoidDefault creates a PatchTestcryptAdministratorsMoidDefault with default headers values
func NewPatchTestcryptAdministratorsMoidDefault(code int) *PatchTestcryptAdministratorsMoidDefault {
	return &PatchTestcryptAdministratorsMoidDefault{
		_statusCode: code,
	}
}

/*PatchTestcryptAdministratorsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchTestcryptAdministratorsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch testcrypt administrators moid default response
func (o *PatchTestcryptAdministratorsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchTestcryptAdministratorsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /testcrypt/Administrators/{moid}][%d] PatchTestcryptAdministratorsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchTestcryptAdministratorsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchTestcryptAdministratorsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}