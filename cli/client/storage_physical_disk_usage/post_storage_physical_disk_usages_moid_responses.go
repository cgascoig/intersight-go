// Code generated by go-swagger; DO NOT EDIT.

package storage_physical_disk_usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostStoragePhysicalDiskUsagesMoidReader is a Reader for the PostStoragePhysicalDiskUsagesMoid structure.
type PostStoragePhysicalDiskUsagesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStoragePhysicalDiskUsagesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStoragePhysicalDiskUsagesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStoragePhysicalDiskUsagesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStoragePhysicalDiskUsagesMoidCreated creates a PostStoragePhysicalDiskUsagesMoidCreated with default headers values
func NewPostStoragePhysicalDiskUsagesMoidCreated() *PostStoragePhysicalDiskUsagesMoidCreated {
	return &PostStoragePhysicalDiskUsagesMoidCreated{}
}

/*PostStoragePhysicalDiskUsagesMoidCreated handles this case with default header values.

Null response
*/
type PostStoragePhysicalDiskUsagesMoidCreated struct {
}

func (o *PostStoragePhysicalDiskUsagesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/PhysicalDiskUsages/{moid}][%d] postStoragePhysicalDiskUsagesMoidCreated ", 201)
}

func (o *PostStoragePhysicalDiskUsagesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStoragePhysicalDiskUsagesMoidDefault creates a PostStoragePhysicalDiskUsagesMoidDefault with default headers values
func NewPostStoragePhysicalDiskUsagesMoidDefault(code int) *PostStoragePhysicalDiskUsagesMoidDefault {
	return &PostStoragePhysicalDiskUsagesMoidDefault{
		_statusCode: code,
	}
}

/*PostStoragePhysicalDiskUsagesMoidDefault handles this case with default header values.

unexpected error
*/
type PostStoragePhysicalDiskUsagesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage physical disk usages moid default response
func (o *PostStoragePhysicalDiskUsagesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStoragePhysicalDiskUsagesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/PhysicalDiskUsages/{moid}][%d] PostStoragePhysicalDiskUsagesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStoragePhysicalDiskUsagesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStoragePhysicalDiskUsagesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
