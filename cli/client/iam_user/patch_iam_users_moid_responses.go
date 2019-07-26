// Code generated by go-swagger; DO NOT EDIT.

package iam_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchIamUsersMoidReader is a Reader for the PatchIamUsersMoid structure.
type PatchIamUsersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIamUsersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchIamUsersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchIamUsersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchIamUsersMoidCreated creates a PatchIamUsersMoidCreated with default headers values
func NewPatchIamUsersMoidCreated() *PatchIamUsersMoidCreated {
	return &PatchIamUsersMoidCreated{}
}

/*PatchIamUsersMoidCreated handles this case with default header values.

Null response
*/
type PatchIamUsersMoidCreated struct {
}

func (o *PatchIamUsersMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /iam/Users/{moid}][%d] patchIamUsersMoidCreated ", 201)
}

func (o *PatchIamUsersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchIamUsersMoidDefault creates a PatchIamUsersMoidDefault with default headers values
func NewPatchIamUsersMoidDefault(code int) *PatchIamUsersMoidDefault {
	return &PatchIamUsersMoidDefault{
		_statusCode: code,
	}
}

/*PatchIamUsersMoidDefault handles this case with default header values.

unexpected error
*/
type PatchIamUsersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch iam users moid default response
func (o *PatchIamUsersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchIamUsersMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /iam/Users/{moid}][%d] PatchIamUsersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchIamUsersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchIamUsersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
