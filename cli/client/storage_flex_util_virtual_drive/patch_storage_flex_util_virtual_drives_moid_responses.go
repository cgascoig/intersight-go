// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_util_virtual_drive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchStorageFlexUtilVirtualDrivesMoidReader is a Reader for the PatchStorageFlexUtilVirtualDrivesMoid structure.
type PatchStorageFlexUtilVirtualDrivesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageFlexUtilVirtualDrivesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchStorageFlexUtilVirtualDrivesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchStorageFlexUtilVirtualDrivesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchStorageFlexUtilVirtualDrivesMoidCreated creates a PatchStorageFlexUtilVirtualDrivesMoidCreated with default headers values
func NewPatchStorageFlexUtilVirtualDrivesMoidCreated() *PatchStorageFlexUtilVirtualDrivesMoidCreated {
	return &PatchStorageFlexUtilVirtualDrivesMoidCreated{}
}

/*PatchStorageFlexUtilVirtualDrivesMoidCreated handles this case with default header values.

Null response
*/
type PatchStorageFlexUtilVirtualDrivesMoidCreated struct {
}

func (o *PatchStorageFlexUtilVirtualDrivesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /storage/FlexUtilVirtualDrives/{moid}][%d] patchStorageFlexUtilVirtualDrivesMoidCreated ", 201)
}

func (o *PatchStorageFlexUtilVirtualDrivesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchStorageFlexUtilVirtualDrivesMoidDefault creates a PatchStorageFlexUtilVirtualDrivesMoidDefault with default headers values
func NewPatchStorageFlexUtilVirtualDrivesMoidDefault(code int) *PatchStorageFlexUtilVirtualDrivesMoidDefault {
	return &PatchStorageFlexUtilVirtualDrivesMoidDefault{
		_statusCode: code,
	}
}

/*PatchStorageFlexUtilVirtualDrivesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchStorageFlexUtilVirtualDrivesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch storage flex util virtual drives moid default response
func (o *PatchStorageFlexUtilVirtualDrivesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchStorageFlexUtilVirtualDrivesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/FlexUtilVirtualDrives/{moid}][%d] PatchStorageFlexUtilVirtualDrivesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchStorageFlexUtilVirtualDrivesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchStorageFlexUtilVirtualDrivesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
