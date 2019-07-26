// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_util_physical_drive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageFlexUtilPhysicalDrivesMoidReader is a Reader for the GetStorageFlexUtilPhysicalDrivesMoid structure.
type GetStorageFlexUtilPhysicalDrivesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageFlexUtilPhysicalDrivesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageFlexUtilPhysicalDrivesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStorageFlexUtilPhysicalDrivesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStorageFlexUtilPhysicalDrivesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageFlexUtilPhysicalDrivesMoidOK creates a GetStorageFlexUtilPhysicalDrivesMoidOK with default headers values
func NewGetStorageFlexUtilPhysicalDrivesMoidOK() *GetStorageFlexUtilPhysicalDrivesMoidOK {
	return &GetStorageFlexUtilPhysicalDrivesMoidOK{}
}

/*GetStorageFlexUtilPhysicalDrivesMoidOK handles this case with default header values.

An instance of storageFlexUtilPhysicalDrive
*/
type GetStorageFlexUtilPhysicalDrivesMoidOK struct {
	Payload *models.StorageFlexUtilPhysicalDrive
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/FlexUtilPhysicalDrives/{moid}][%d] getStorageFlexUtilPhysicalDrivesMoidOK  %+v", 200, o.Payload)
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidOK) GetPayload() *models.StorageFlexUtilPhysicalDrive {
	return o.Payload
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageFlexUtilPhysicalDrive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageFlexUtilPhysicalDrivesMoidNotFound creates a GetStorageFlexUtilPhysicalDrivesMoidNotFound with default headers values
func NewGetStorageFlexUtilPhysicalDrivesMoidNotFound() *GetStorageFlexUtilPhysicalDrivesMoidNotFound {
	return &GetStorageFlexUtilPhysicalDrivesMoidNotFound{}
}

/*GetStorageFlexUtilPhysicalDrivesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStorageFlexUtilPhysicalDrivesMoidNotFound struct {
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/FlexUtilPhysicalDrives/{moid}][%d] getStorageFlexUtilPhysicalDrivesMoidNotFound ", 404)
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageFlexUtilPhysicalDrivesMoidDefault creates a GetStorageFlexUtilPhysicalDrivesMoidDefault with default headers values
func NewGetStorageFlexUtilPhysicalDrivesMoidDefault(code int) *GetStorageFlexUtilPhysicalDrivesMoidDefault {
	return &GetStorageFlexUtilPhysicalDrivesMoidDefault{
		_statusCode: code,
	}
}

/*GetStorageFlexUtilPhysicalDrivesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStorageFlexUtilPhysicalDrivesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage flex util physical drives moid default response
func (o *GetStorageFlexUtilPhysicalDrivesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/FlexUtilPhysicalDrives/{moid}][%d] GetStorageFlexUtilPhysicalDrivesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageFlexUtilPhysicalDrivesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
