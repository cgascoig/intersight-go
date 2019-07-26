// Code generated by go-swagger; DO NOT EDIT.

package storage_virtual_drive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchStorageVirtualDrivesMoidReader is a Reader for the PatchStorageVirtualDrivesMoid structure.
type PatchStorageVirtualDrivesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageVirtualDrivesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchStorageVirtualDrivesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchStorageVirtualDrivesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchStorageVirtualDrivesMoidCreated creates a PatchStorageVirtualDrivesMoidCreated with default headers values
func NewPatchStorageVirtualDrivesMoidCreated() *PatchStorageVirtualDrivesMoidCreated {
	return &PatchStorageVirtualDrivesMoidCreated{}
}

/*PatchStorageVirtualDrivesMoidCreated handles this case with default header values.

Null response
*/
type PatchStorageVirtualDrivesMoidCreated struct {
}

func (o *PatchStorageVirtualDrivesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /storage/VirtualDrives/{moid}][%d] patchStorageVirtualDrivesMoidCreated ", 201)
}

func (o *PatchStorageVirtualDrivesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchStorageVirtualDrivesMoidDefault creates a PatchStorageVirtualDrivesMoidDefault with default headers values
func NewPatchStorageVirtualDrivesMoidDefault(code int) *PatchStorageVirtualDrivesMoidDefault {
	return &PatchStorageVirtualDrivesMoidDefault{
		_statusCode: code,
	}
}

/*PatchStorageVirtualDrivesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchStorageVirtualDrivesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch storage virtual drives moid default response
func (o *PatchStorageVirtualDrivesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchStorageVirtualDrivesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/VirtualDrives/{moid}][%d] PatchStorageVirtualDrivesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchStorageVirtualDrivesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchStorageVirtualDrivesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
