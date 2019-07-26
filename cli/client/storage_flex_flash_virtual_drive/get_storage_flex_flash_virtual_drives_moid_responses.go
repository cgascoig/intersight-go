// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_flash_virtual_drive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageFlexFlashVirtualDrivesMoidReader is a Reader for the GetStorageFlexFlashVirtualDrivesMoid structure.
type GetStorageFlexFlashVirtualDrivesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageFlexFlashVirtualDrivesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageFlexFlashVirtualDrivesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStorageFlexFlashVirtualDrivesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStorageFlexFlashVirtualDrivesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageFlexFlashVirtualDrivesMoidOK creates a GetStorageFlexFlashVirtualDrivesMoidOK with default headers values
func NewGetStorageFlexFlashVirtualDrivesMoidOK() *GetStorageFlexFlashVirtualDrivesMoidOK {
	return &GetStorageFlexFlashVirtualDrivesMoidOK{}
}

/*GetStorageFlexFlashVirtualDrivesMoidOK handles this case with default header values.

An instance of storageFlexFlashVirtualDrive
*/
type GetStorageFlexFlashVirtualDrivesMoidOK struct {
	Payload *models.StorageFlexFlashVirtualDrive
}

func (o *GetStorageFlexFlashVirtualDrivesMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/FlexFlashVirtualDrives/{moid}][%d] getStorageFlexFlashVirtualDrivesMoidOK  %+v", 200, o.Payload)
}

func (o *GetStorageFlexFlashVirtualDrivesMoidOK) GetPayload() *models.StorageFlexFlashVirtualDrive {
	return o.Payload
}

func (o *GetStorageFlexFlashVirtualDrivesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageFlexFlashVirtualDrive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageFlexFlashVirtualDrivesMoidNotFound creates a GetStorageFlexFlashVirtualDrivesMoidNotFound with default headers values
func NewGetStorageFlexFlashVirtualDrivesMoidNotFound() *GetStorageFlexFlashVirtualDrivesMoidNotFound {
	return &GetStorageFlexFlashVirtualDrivesMoidNotFound{}
}

/*GetStorageFlexFlashVirtualDrivesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStorageFlexFlashVirtualDrivesMoidNotFound struct {
}

func (o *GetStorageFlexFlashVirtualDrivesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/FlexFlashVirtualDrives/{moid}][%d] getStorageFlexFlashVirtualDrivesMoidNotFound ", 404)
}

func (o *GetStorageFlexFlashVirtualDrivesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageFlexFlashVirtualDrivesMoidDefault creates a GetStorageFlexFlashVirtualDrivesMoidDefault with default headers values
func NewGetStorageFlexFlashVirtualDrivesMoidDefault(code int) *GetStorageFlexFlashVirtualDrivesMoidDefault {
	return &GetStorageFlexFlashVirtualDrivesMoidDefault{
		_statusCode: code,
	}
}

/*GetStorageFlexFlashVirtualDrivesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStorageFlexFlashVirtualDrivesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage flex flash virtual drives moid default response
func (o *GetStorageFlexFlashVirtualDrivesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageFlexFlashVirtualDrivesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/FlexFlashVirtualDrives/{moid}][%d] GetStorageFlexFlashVirtualDrivesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageFlexFlashVirtualDrivesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageFlexFlashVirtualDrivesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
