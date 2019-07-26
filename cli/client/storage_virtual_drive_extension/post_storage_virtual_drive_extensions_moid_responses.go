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

// PostStorageVirtualDriveExtensionsMoidReader is a Reader for the PostStorageVirtualDriveExtensionsMoid structure.
type PostStorageVirtualDriveExtensionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageVirtualDriveExtensionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStorageVirtualDriveExtensionsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStorageVirtualDriveExtensionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStorageVirtualDriveExtensionsMoidCreated creates a PostStorageVirtualDriveExtensionsMoidCreated with default headers values
func NewPostStorageVirtualDriveExtensionsMoidCreated() *PostStorageVirtualDriveExtensionsMoidCreated {
	return &PostStorageVirtualDriveExtensionsMoidCreated{}
}

/*PostStorageVirtualDriveExtensionsMoidCreated handles this case with default header values.

Null response
*/
type PostStorageVirtualDriveExtensionsMoidCreated struct {
}

func (o *PostStorageVirtualDriveExtensionsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/VirtualDriveExtensions/{moid}][%d] postStorageVirtualDriveExtensionsMoidCreated ", 201)
}

func (o *PostStorageVirtualDriveExtensionsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStorageVirtualDriveExtensionsMoidDefault creates a PostStorageVirtualDriveExtensionsMoidDefault with default headers values
func NewPostStorageVirtualDriveExtensionsMoidDefault(code int) *PostStorageVirtualDriveExtensionsMoidDefault {
	return &PostStorageVirtualDriveExtensionsMoidDefault{
		_statusCode: code,
	}
}

/*PostStorageVirtualDriveExtensionsMoidDefault handles this case with default header values.

unexpected error
*/
type PostStorageVirtualDriveExtensionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage virtual drive extensions moid default response
func (o *PostStorageVirtualDriveExtensionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStorageVirtualDriveExtensionsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/VirtualDriveExtensions/{moid}][%d] PostStorageVirtualDriveExtensionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStorageVirtualDriveExtensionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStorageVirtualDriveExtensionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}