// Code generated by go-swagger; DO NOT EDIT.

package ssh_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchSSHPoliciesMoidReader is a Reader for the PatchSSHPoliciesMoid structure.
type PatchSSHPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSSHPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchSSHPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchSSHPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSSHPoliciesMoidCreated creates a PatchSSHPoliciesMoidCreated with default headers values
func NewPatchSSHPoliciesMoidCreated() *PatchSSHPoliciesMoidCreated {
	return &PatchSSHPoliciesMoidCreated{}
}

/*PatchSSHPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchSSHPoliciesMoidCreated struct {
}

func (o *PatchSSHPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /ssh/Policies/{moid}][%d] patchSshPoliciesMoidCreated ", 201)
}

func (o *PatchSSHPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSSHPoliciesMoidDefault creates a PatchSSHPoliciesMoidDefault with default headers values
func NewPatchSSHPoliciesMoidDefault(code int) *PatchSSHPoliciesMoidDefault {
	return &PatchSSHPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchSSHPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchSSHPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch SSH policies moid default response
func (o *PatchSSHPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchSSHPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /ssh/Policies/{moid}][%d] PatchSSHPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSSHPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSSHPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}