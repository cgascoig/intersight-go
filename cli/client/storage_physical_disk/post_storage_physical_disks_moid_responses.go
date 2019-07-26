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

// PostStoragePhysicalDisksMoidReader is a Reader for the PostStoragePhysicalDisksMoid structure.
type PostStoragePhysicalDisksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStoragePhysicalDisksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStoragePhysicalDisksMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStoragePhysicalDisksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStoragePhysicalDisksMoidCreated creates a PostStoragePhysicalDisksMoidCreated with default headers values
func NewPostStoragePhysicalDisksMoidCreated() *PostStoragePhysicalDisksMoidCreated {
	return &PostStoragePhysicalDisksMoidCreated{}
}

/*PostStoragePhysicalDisksMoidCreated handles this case with default header values.

Null response
*/
type PostStoragePhysicalDisksMoidCreated struct {
}

func (o *PostStoragePhysicalDisksMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/PhysicalDisks/{moid}][%d] postStoragePhysicalDisksMoidCreated ", 201)
}

func (o *PostStoragePhysicalDisksMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStoragePhysicalDisksMoidDefault creates a PostStoragePhysicalDisksMoidDefault with default headers values
func NewPostStoragePhysicalDisksMoidDefault(code int) *PostStoragePhysicalDisksMoidDefault {
	return &PostStoragePhysicalDisksMoidDefault{
		_statusCode: code,
	}
}

/*PostStoragePhysicalDisksMoidDefault handles this case with default header values.

unexpected error
*/
type PostStoragePhysicalDisksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage physical disks moid default response
func (o *PostStoragePhysicalDisksMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStoragePhysicalDisksMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/PhysicalDisks/{moid}][%d] PostStoragePhysicalDisksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStoragePhysicalDisksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStoragePhysicalDisksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}