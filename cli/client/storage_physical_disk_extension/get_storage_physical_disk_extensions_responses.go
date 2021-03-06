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

// GetStoragePhysicalDiskExtensionsReader is a Reader for the GetStoragePhysicalDiskExtensions structure.
type GetStoragePhysicalDiskExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePhysicalDiskExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePhysicalDiskExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStoragePhysicalDiskExtensionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePhysicalDiskExtensionsOK creates a GetStoragePhysicalDiskExtensionsOK with default headers values
func NewGetStoragePhysicalDiskExtensionsOK() *GetStoragePhysicalDiskExtensionsOK {
	return &GetStoragePhysicalDiskExtensionsOK{}
}

/*GetStoragePhysicalDiskExtensionsOK handles this case with default header values.

List of storagePhysicalDiskExtensions for the given filter criteria
*/
type GetStoragePhysicalDiskExtensionsOK struct {
	Payload *models.StoragePhysicalDiskExtensionList
}

func (o *GetStoragePhysicalDiskExtensionsOK) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDiskExtensions][%d] getStoragePhysicalDiskExtensionsOK  %+v", 200, o.Payload)
}

func (o *GetStoragePhysicalDiskExtensionsOK) GetPayload() *models.StoragePhysicalDiskExtensionList {
	return o.Payload
}

func (o *GetStoragePhysicalDiskExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePhysicalDiskExtensionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePhysicalDiskExtensionsDefault creates a GetStoragePhysicalDiskExtensionsDefault with default headers values
func NewGetStoragePhysicalDiskExtensionsDefault(code int) *GetStoragePhysicalDiskExtensionsDefault {
	return &GetStoragePhysicalDiskExtensionsDefault{
		_statusCode: code,
	}
}

/*GetStoragePhysicalDiskExtensionsDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePhysicalDiskExtensionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage physical disk extensions default response
func (o *GetStoragePhysicalDiskExtensionsDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePhysicalDiskExtensionsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PhysicalDiskExtensions][%d] GetStoragePhysicalDiskExtensions default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePhysicalDiskExtensionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePhysicalDiskExtensionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
