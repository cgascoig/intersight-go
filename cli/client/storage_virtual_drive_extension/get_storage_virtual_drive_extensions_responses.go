// Code generated by go-swagger; DO NOT EDIT.

package storage_virtual_drive_extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageVirtualDriveExtensionsReader is a Reader for the GetStorageVirtualDriveExtensions structure.
type GetStorageVirtualDriveExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageVirtualDriveExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageVirtualDriveExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStorageVirtualDriveExtensionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageVirtualDriveExtensionsOK creates a GetStorageVirtualDriveExtensionsOK with default headers values
func NewGetStorageVirtualDriveExtensionsOK() *GetStorageVirtualDriveExtensionsOK {
	return &GetStorageVirtualDriveExtensionsOK{}
}

/*GetStorageVirtualDriveExtensionsOK handles this case with default header values.

List of storageVirtualDriveExtensions for the given filter criteria
*/
type GetStorageVirtualDriveExtensionsOK struct {
	Payload *models.StorageVirtualDriveExtensionList
}

func (o *GetStorageVirtualDriveExtensionsOK) Error() string {
	return fmt.Sprintf("[GET /storage/VirtualDriveExtensions][%d] getStorageVirtualDriveExtensionsOK  %+v", 200, o.Payload)
}

func (o *GetStorageVirtualDriveExtensionsOK) GetPayload() *models.StorageVirtualDriveExtensionList {
	return o.Payload
}

func (o *GetStorageVirtualDriveExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageVirtualDriveExtensionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageVirtualDriveExtensionsDefault creates a GetStorageVirtualDriveExtensionsDefault with default headers values
func NewGetStorageVirtualDriveExtensionsDefault(code int) *GetStorageVirtualDriveExtensionsDefault {
	return &GetStorageVirtualDriveExtensionsDefault{
		_statusCode: code,
	}
}

/*GetStorageVirtualDriveExtensionsDefault handles this case with default header values.

Unexpected error
*/
type GetStorageVirtualDriveExtensionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage virtual drive extensions default response
func (o *GetStorageVirtualDriveExtensionsDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageVirtualDriveExtensionsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/VirtualDriveExtensions][%d] GetStorageVirtualDriveExtensions default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageVirtualDriveExtensionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageVirtualDriveExtensionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
