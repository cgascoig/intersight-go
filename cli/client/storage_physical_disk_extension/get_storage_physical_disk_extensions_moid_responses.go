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

// GetStoragePhysicalDiskExtensionsMoidReader is a Reader for the GetStoragePhysicalDiskExtensionsMoid structure.
type GetStoragePhysicalDiskExtensionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePhysicalDiskExtensionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePhysicalDiskExtensionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePhysicalDiskExtensionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePhysicalDiskExtensionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePhysicalDiskExtensionsMoidOK creates a GetStoragePhysicalDiskExtensionsMoidOK with default headers values
func NewGetStoragePhysicalDiskExtensionsMoidOK() *GetStoragePhysicalDiskExtensionsMoidOK {
	return &GetStoragePhysicalDiskExtensionsMoidOK{}
}

/*GetStoragePhysicalDiskExtensionsMoidOK handles this case with default header values.

An instance of storagePhysicalDiskExtension
*/
type GetStoragePhysicalDiskExtensionsMoidOK struct {
	Payload *models.StoragePhysicalDiskExtension
}

func (o *GetStoragePhysicalDiskExtensionsMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDiskExtensions/{moid}][%d] getStoragePhysicalDiskExtensionsMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePhysicalDiskExtensionsMoidOK) GetPayload() *models.StoragePhysicalDiskExtension {
	return o.Payload
}

func (o *GetStoragePhysicalDiskExtensionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePhysicalDiskExtension)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePhysicalDiskExtensionsMoidNotFound creates a GetStoragePhysicalDiskExtensionsMoidNotFound with default headers values
func NewGetStoragePhysicalDiskExtensionsMoidNotFound() *GetStoragePhysicalDiskExtensionsMoidNotFound {
	return &GetStoragePhysicalDiskExtensionsMoidNotFound{}
}

/*GetStoragePhysicalDiskExtensionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePhysicalDiskExtensionsMoidNotFound struct {
}

func (o *GetStoragePhysicalDiskExtensionsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDiskExtensions/{moid}][%d] getStoragePhysicalDiskExtensionsMoidNotFound ", 404)
}

func (o *GetStoragePhysicalDiskExtensionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePhysicalDiskExtensionsMoidDefault creates a GetStoragePhysicalDiskExtensionsMoidDefault with default headers values
func NewGetStoragePhysicalDiskExtensionsMoidDefault(code int) *GetStoragePhysicalDiskExtensionsMoidDefault {
	return &GetStoragePhysicalDiskExtensionsMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePhysicalDiskExtensionsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePhysicalDiskExtensionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage physical disk extensions moid default response
func (o *GetStoragePhysicalDiskExtensionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePhysicalDiskExtensionsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDiskExtensions/{moid}][%d] GetStoragePhysicalDiskExtensionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePhysicalDiskExtensionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePhysicalDiskExtensionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
