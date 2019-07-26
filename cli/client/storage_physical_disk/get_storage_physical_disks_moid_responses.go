// Code generated by go-swagger; DO NOT EDIT.

package storage_physical_disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePhysicalDisksMoidReader is a Reader for the GetStoragePhysicalDisksMoid structure.
type GetStoragePhysicalDisksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePhysicalDisksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePhysicalDisksMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePhysicalDisksMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePhysicalDisksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePhysicalDisksMoidOK creates a GetStoragePhysicalDisksMoidOK with default headers values
func NewGetStoragePhysicalDisksMoidOK() *GetStoragePhysicalDisksMoidOK {
	return &GetStoragePhysicalDisksMoidOK{}
}

/*GetStoragePhysicalDisksMoidOK handles this case with default header values.

An instance of storagePhysicalDisk
*/
type GetStoragePhysicalDisksMoidOK struct {
	Payload *models.StoragePhysicalDisk
}

func (o *GetStoragePhysicalDisksMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDisks/{moid}][%d] getStoragePhysicalDisksMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePhysicalDisksMoidOK) GetPayload() *models.StoragePhysicalDisk {
	return o.Payload
}

func (o *GetStoragePhysicalDisksMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePhysicalDisk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePhysicalDisksMoidNotFound creates a GetStoragePhysicalDisksMoidNotFound with default headers values
func NewGetStoragePhysicalDisksMoidNotFound() *GetStoragePhysicalDisksMoidNotFound {
	return &GetStoragePhysicalDisksMoidNotFound{}
}

/*GetStoragePhysicalDisksMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePhysicalDisksMoidNotFound struct {
}

func (o *GetStoragePhysicalDisksMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDisks/{moid}][%d] getStoragePhysicalDisksMoidNotFound ", 404)
}

func (o *GetStoragePhysicalDisksMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePhysicalDisksMoidDefault creates a GetStoragePhysicalDisksMoidDefault with default headers values
func NewGetStoragePhysicalDisksMoidDefault(code int) *GetStoragePhysicalDisksMoidDefault {
	return &GetStoragePhysicalDisksMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePhysicalDisksMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePhysicalDisksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage physical disks moid default response
func (o *GetStoragePhysicalDisksMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePhysicalDisksMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDisks/{moid}][%d] GetStoragePhysicalDisksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePhysicalDisksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePhysicalDisksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}