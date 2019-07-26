// Code generated by go-swagger; DO NOT EDIT.

package bios_boot_mode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchBiosBootModesMoidReader is a Reader for the PatchBiosBootModesMoid structure.
type PatchBiosBootModesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBiosBootModesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchBiosBootModesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchBiosBootModesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchBiosBootModesMoidCreated creates a PatchBiosBootModesMoidCreated with default headers values
func NewPatchBiosBootModesMoidCreated() *PatchBiosBootModesMoidCreated {
	return &PatchBiosBootModesMoidCreated{}
}

/*PatchBiosBootModesMoidCreated handles this case with default header values.

Null response
*/
type PatchBiosBootModesMoidCreated struct {
}

func (o *PatchBiosBootModesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /bios/BootModes/{moid}][%d] patchBiosBootModesMoidCreated ", 201)
}

func (o *PatchBiosBootModesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchBiosBootModesMoidDefault creates a PatchBiosBootModesMoidDefault with default headers values
func NewPatchBiosBootModesMoidDefault(code int) *PatchBiosBootModesMoidDefault {
	return &PatchBiosBootModesMoidDefault{
		_statusCode: code,
	}
}

/*PatchBiosBootModesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchBiosBootModesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch bios boot modes moid default response
func (o *PatchBiosBootModesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchBiosBootModesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /bios/BootModes/{moid}][%d] PatchBiosBootModesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchBiosBootModesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBiosBootModesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
