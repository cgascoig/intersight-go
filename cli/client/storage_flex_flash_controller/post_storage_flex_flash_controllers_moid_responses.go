// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_flash_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostStorageFlexFlashControllersMoidReader is a Reader for the PostStorageFlexFlashControllersMoid structure.
type PostStorageFlexFlashControllersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageFlexFlashControllersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStorageFlexFlashControllersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStorageFlexFlashControllersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStorageFlexFlashControllersMoidCreated creates a PostStorageFlexFlashControllersMoidCreated with default headers values
func NewPostStorageFlexFlashControllersMoidCreated() *PostStorageFlexFlashControllersMoidCreated {
	return &PostStorageFlexFlashControllersMoidCreated{}
}

/*PostStorageFlexFlashControllersMoidCreated handles this case with default header values.

Null response
*/
type PostStorageFlexFlashControllersMoidCreated struct {
}

func (o *PostStorageFlexFlashControllersMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/FlexFlashControllers/{moid}][%d] postStorageFlexFlashControllersMoidCreated ", 201)
}

func (o *PostStorageFlexFlashControllersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStorageFlexFlashControllersMoidDefault creates a PostStorageFlexFlashControllersMoidDefault with default headers values
func NewPostStorageFlexFlashControllersMoidDefault(code int) *PostStorageFlexFlashControllersMoidDefault {
	return &PostStorageFlexFlashControllersMoidDefault{
		_statusCode: code,
	}
}

/*PostStorageFlexFlashControllersMoidDefault handles this case with default header values.

unexpected error
*/
type PostStorageFlexFlashControllersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage flex flash controllers moid default response
func (o *PostStorageFlexFlashControllersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStorageFlexFlashControllersMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/FlexFlashControllers/{moid}][%d] PostStorageFlexFlashControllersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStorageFlexFlashControllersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStorageFlexFlashControllersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
