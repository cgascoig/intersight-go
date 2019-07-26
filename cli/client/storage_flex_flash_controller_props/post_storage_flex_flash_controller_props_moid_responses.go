// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_flash_controller_props

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostStorageFlexFlashControllerPropsMoidReader is a Reader for the PostStorageFlexFlashControllerPropsMoid structure.
type PostStorageFlexFlashControllerPropsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageFlexFlashControllerPropsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStorageFlexFlashControllerPropsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStorageFlexFlashControllerPropsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStorageFlexFlashControllerPropsMoidCreated creates a PostStorageFlexFlashControllerPropsMoidCreated with default headers values
func NewPostStorageFlexFlashControllerPropsMoidCreated() *PostStorageFlexFlashControllerPropsMoidCreated {
	return &PostStorageFlexFlashControllerPropsMoidCreated{}
}

/*PostStorageFlexFlashControllerPropsMoidCreated handles this case with default header values.

Null response
*/
type PostStorageFlexFlashControllerPropsMoidCreated struct {
}

func (o *PostStorageFlexFlashControllerPropsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/FlexFlashControllerProps/{moid}][%d] postStorageFlexFlashControllerPropsMoidCreated ", 201)
}

func (o *PostStorageFlexFlashControllerPropsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStorageFlexFlashControllerPropsMoidDefault creates a PostStorageFlexFlashControllerPropsMoidDefault with default headers values
func NewPostStorageFlexFlashControllerPropsMoidDefault(code int) *PostStorageFlexFlashControllerPropsMoidDefault {
	return &PostStorageFlexFlashControllerPropsMoidDefault{
		_statusCode: code,
	}
}

/*PostStorageFlexFlashControllerPropsMoidDefault handles this case with default header values.

unexpected error
*/
type PostStorageFlexFlashControllerPropsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage flex flash controller props moid default response
func (o *PostStorageFlexFlashControllerPropsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStorageFlexFlashControllerPropsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/FlexFlashControllerProps/{moid}][%d] PostStorageFlexFlashControllerPropsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStorageFlexFlashControllerPropsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStorageFlexFlashControllerPropsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}