// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_util_virtual_drive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostStorageFlexUtilVirtualDrivesMoidReader is a Reader for the PostStorageFlexUtilVirtualDrivesMoid structure.
type PostStorageFlexUtilVirtualDrivesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageFlexUtilVirtualDrivesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStorageFlexUtilVirtualDrivesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStorageFlexUtilVirtualDrivesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStorageFlexUtilVirtualDrivesMoidCreated creates a PostStorageFlexUtilVirtualDrivesMoidCreated with default headers values
func NewPostStorageFlexUtilVirtualDrivesMoidCreated() *PostStorageFlexUtilVirtualDrivesMoidCreated {
	return &PostStorageFlexUtilVirtualDrivesMoidCreated{}
}

/*PostStorageFlexUtilVirtualDrivesMoidCreated handles this case with default header values.

Null response
*/
type PostStorageFlexUtilVirtualDrivesMoidCreated struct {
}

func (o *PostStorageFlexUtilVirtualDrivesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/FlexUtilVirtualDrives/{moid}][%d] postStorageFlexUtilVirtualDrivesMoidCreated ", 201)
}

func (o *PostStorageFlexUtilVirtualDrivesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStorageFlexUtilVirtualDrivesMoidDefault creates a PostStorageFlexUtilVirtualDrivesMoidDefault with default headers values
func NewPostStorageFlexUtilVirtualDrivesMoidDefault(code int) *PostStorageFlexUtilVirtualDrivesMoidDefault {
	return &PostStorageFlexUtilVirtualDrivesMoidDefault{
		_statusCode: code,
	}
}

/*PostStorageFlexUtilVirtualDrivesMoidDefault handles this case with default header values.

unexpected error
*/
type PostStorageFlexUtilVirtualDrivesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage flex util virtual drives moid default response
func (o *PostStorageFlexUtilVirtualDrivesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStorageFlexUtilVirtualDrivesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/FlexUtilVirtualDrives/{moid}][%d] PostStorageFlexUtilVirtualDrivesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStorageFlexUtilVirtualDrivesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStorageFlexUtilVirtualDrivesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}