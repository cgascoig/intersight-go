// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_firmware_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchHyperflexServerFirmwareVersionsMoidReader is a Reader for the PatchHyperflexServerFirmwareVersionsMoid structure.
type PatchHyperflexServerFirmwareVersionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchHyperflexServerFirmwareVersionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchHyperflexServerFirmwareVersionsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchHyperflexServerFirmwareVersionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchHyperflexServerFirmwareVersionsMoidCreated creates a PatchHyperflexServerFirmwareVersionsMoidCreated with default headers values
func NewPatchHyperflexServerFirmwareVersionsMoidCreated() *PatchHyperflexServerFirmwareVersionsMoidCreated {
	return &PatchHyperflexServerFirmwareVersionsMoidCreated{}
}

/*PatchHyperflexServerFirmwareVersionsMoidCreated handles this case with default header values.

Null response
*/
type PatchHyperflexServerFirmwareVersionsMoidCreated struct {
}

func (o *PatchHyperflexServerFirmwareVersionsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/ServerFirmwareVersions/{moid}][%d] patchHyperflexServerFirmwareVersionsMoidCreated ", 201)
}

func (o *PatchHyperflexServerFirmwareVersionsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchHyperflexServerFirmwareVersionsMoidDefault creates a PatchHyperflexServerFirmwareVersionsMoidDefault with default headers values
func NewPatchHyperflexServerFirmwareVersionsMoidDefault(code int) *PatchHyperflexServerFirmwareVersionsMoidDefault {
	return &PatchHyperflexServerFirmwareVersionsMoidDefault{
		_statusCode: code,
	}
}

/*PatchHyperflexServerFirmwareVersionsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchHyperflexServerFirmwareVersionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch hyperflex server firmware versions moid default response
func (o *PatchHyperflexServerFirmwareVersionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchHyperflexServerFirmwareVersionsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/ServerFirmwareVersions/{moid}][%d] PatchHyperflexServerFirmwareVersionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchHyperflexServerFirmwareVersionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchHyperflexServerFirmwareVersionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}