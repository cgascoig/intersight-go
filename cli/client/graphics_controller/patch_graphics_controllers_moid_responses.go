// Code generated by go-swagger; DO NOT EDIT.

package graphics_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchGraphicsControllersMoidReader is a Reader for the PatchGraphicsControllersMoid structure.
type PatchGraphicsControllersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGraphicsControllersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchGraphicsControllersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchGraphicsControllersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGraphicsControllersMoidCreated creates a PatchGraphicsControllersMoidCreated with default headers values
func NewPatchGraphicsControllersMoidCreated() *PatchGraphicsControllersMoidCreated {
	return &PatchGraphicsControllersMoidCreated{}
}

/*PatchGraphicsControllersMoidCreated handles this case with default header values.

Null response
*/
type PatchGraphicsControllersMoidCreated struct {
}

func (o *PatchGraphicsControllersMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /graphics/Controllers/{moid}][%d] patchGraphicsControllersMoidCreated ", 201)
}

func (o *PatchGraphicsControllersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGraphicsControllersMoidDefault creates a PatchGraphicsControllersMoidDefault with default headers values
func NewPatchGraphicsControllersMoidDefault(code int) *PatchGraphicsControllersMoidDefault {
	return &PatchGraphicsControllersMoidDefault{
		_statusCode: code,
	}
}

/*PatchGraphicsControllersMoidDefault handles this case with default header values.

unexpected error
*/
type PatchGraphicsControllersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch graphics controllers moid default response
func (o *PatchGraphicsControllersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchGraphicsControllersMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /graphics/Controllers/{moid}][%d] PatchGraphicsControllersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchGraphicsControllersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGraphicsControllersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
