// Code generated by go-swagger; DO NOT EDIT.

package graphics_card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchGraphicsCardsMoidReader is a Reader for the PatchGraphicsCardsMoid structure.
type PatchGraphicsCardsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGraphicsCardsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchGraphicsCardsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchGraphicsCardsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGraphicsCardsMoidCreated creates a PatchGraphicsCardsMoidCreated with default headers values
func NewPatchGraphicsCardsMoidCreated() *PatchGraphicsCardsMoidCreated {
	return &PatchGraphicsCardsMoidCreated{}
}

/*PatchGraphicsCardsMoidCreated handles this case with default header values.

Null response
*/
type PatchGraphicsCardsMoidCreated struct {
}

func (o *PatchGraphicsCardsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /graphics/Cards/{moid}][%d] patchGraphicsCardsMoidCreated ", 201)
}

func (o *PatchGraphicsCardsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGraphicsCardsMoidDefault creates a PatchGraphicsCardsMoidDefault with default headers values
func NewPatchGraphicsCardsMoidDefault(code int) *PatchGraphicsCardsMoidDefault {
	return &PatchGraphicsCardsMoidDefault{
		_statusCode: code,
	}
}

/*PatchGraphicsCardsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchGraphicsCardsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch graphics cards moid default response
func (o *PatchGraphicsCardsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchGraphicsCardsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /graphics/Cards/{moid}][%d] PatchGraphicsCardsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchGraphicsCardsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGraphicsCardsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
