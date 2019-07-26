// Code generated by go-swagger; DO NOT EDIT.

package storage_physical_disk_extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchStoragePhysicalDiskExtensionsMoidReader is a Reader for the PatchStoragePhysicalDiskExtensionsMoid structure.
type PatchStoragePhysicalDiskExtensionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStoragePhysicalDiskExtensionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchStoragePhysicalDiskExtensionsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchStoragePhysicalDiskExtensionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchStoragePhysicalDiskExtensionsMoidCreated creates a PatchStoragePhysicalDiskExtensionsMoidCreated with default headers values
func NewPatchStoragePhysicalDiskExtensionsMoidCreated() *PatchStoragePhysicalDiskExtensionsMoidCreated {
	return &PatchStoragePhysicalDiskExtensionsMoidCreated{}
}

/*PatchStoragePhysicalDiskExtensionsMoidCreated handles this case with default header values.

Null response
*/
type PatchStoragePhysicalDiskExtensionsMoidCreated struct {
}

func (o *PatchStoragePhysicalDiskExtensionsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /storage/PhysicalDiskExtensions/{moid}][%d] patchStoragePhysicalDiskExtensionsMoidCreated ", 201)
}

func (o *PatchStoragePhysicalDiskExtensionsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchStoragePhysicalDiskExtensionsMoidDefault creates a PatchStoragePhysicalDiskExtensionsMoidDefault with default headers values
func NewPatchStoragePhysicalDiskExtensionsMoidDefault(code int) *PatchStoragePhysicalDiskExtensionsMoidDefault {
	return &PatchStoragePhysicalDiskExtensionsMoidDefault{
		_statusCode: code,
	}
}

/*PatchStoragePhysicalDiskExtensionsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchStoragePhysicalDiskExtensionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch storage physical disk extensions moid default response
func (o *PatchStoragePhysicalDiskExtensionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchStoragePhysicalDiskExtensionsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/PhysicalDiskExtensions/{moid}][%d] PatchStoragePhysicalDiskExtensionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchStoragePhysicalDiskExtensionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchStoragePhysicalDiskExtensionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
