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

// GetStorageVirtualDrivesMoidReader is a Reader for the GetStorageVirtualDrivesMoid structure.
type GetStorageVirtualDrivesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageVirtualDrivesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageVirtualDrivesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStorageVirtualDrivesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStorageVirtualDrivesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageVirtualDrivesMoidOK creates a GetStorageVirtualDrivesMoidOK with default headers values
func NewGetStorageVirtualDrivesMoidOK() *GetStorageVirtualDrivesMoidOK {
	return &GetStorageVirtualDrivesMoidOK{}
}

/*GetStorageVirtualDrivesMoidOK handles this case with default header values.

An instance of storageVirtualDrive
*/
type GetStorageVirtualDrivesMoidOK struct {
	Payload *models.StorageVirtualDrive
}

func (o *GetStorageVirtualDrivesMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/VirtualDrives/{moid}][%d] getStorageVirtualDrivesMoidOK  %+v", 200, o.Payload)
}

func (o *GetStorageVirtualDrivesMoidOK) GetPayload() *models.StorageVirtualDrive {
	return o.Payload
}

func (o *GetStorageVirtualDrivesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageVirtualDrive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageVirtualDrivesMoidNotFound creates a GetStorageVirtualDrivesMoidNotFound with default headers values
func NewGetStorageVirtualDrivesMoidNotFound() *GetStorageVirtualDrivesMoidNotFound {
	return &GetStorageVirtualDrivesMoidNotFound{}
}

/*GetStorageVirtualDrivesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStorageVirtualDrivesMoidNotFound struct {
}

func (o *GetStorageVirtualDrivesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/VirtualDrives/{moid}][%d] getStorageVirtualDrivesMoidNotFound ", 404)
}

func (o *GetStorageVirtualDrivesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageVirtualDrivesMoidDefault creates a GetStorageVirtualDrivesMoidDefault with default headers values
func NewGetStorageVirtualDrivesMoidDefault(code int) *GetStorageVirtualDrivesMoidDefault {
	return &GetStorageVirtualDrivesMoidDefault{
		_statusCode: code,
	}
}

/*GetStorageVirtualDrivesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStorageVirtualDrivesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage virtual drives moid default response
func (o *GetStorageVirtualDrivesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageVirtualDrivesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/VirtualDrives/{moid}][%d] GetStorageVirtualDrivesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageVirtualDrivesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageVirtualDrivesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
