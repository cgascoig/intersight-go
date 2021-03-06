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

// PatchStorageFlexFlashControllerPropsMoidReader is a Reader for the PatchStorageFlexFlashControllerPropsMoid structure.
type PatchStorageFlexFlashControllerPropsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageFlexFlashControllerPropsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchStorageFlexFlashControllerPropsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchStorageFlexFlashControllerPropsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchStorageFlexFlashControllerPropsMoidCreated creates a PatchStorageFlexFlashControllerPropsMoidCreated with default headers values
func NewPatchStorageFlexFlashControllerPropsMoidCreated() *PatchStorageFlexFlashControllerPropsMoidCreated {
	return &PatchStorageFlexFlashControllerPropsMoidCreated{}
}

/*PatchStorageFlexFlashControllerPropsMoidCreated handles this case with default header values.

Null response
*/
type PatchStorageFlexFlashControllerPropsMoidCreated struct {
}

func (o *PatchStorageFlexFlashControllerPropsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /storage/FlexFlashControllerProps/{moid}][%d] patchStorageFlexFlashControllerPropsMoidCreated ", 201)
}

func (o *PatchStorageFlexFlashControllerPropsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchStorageFlexFlashControllerPropsMoidDefault creates a PatchStorageFlexFlashControllerPropsMoidDefault with default headers values
func NewPatchStorageFlexFlashControllerPropsMoidDefault(code int) *PatchStorageFlexFlashControllerPropsMoidDefault {
	return &PatchStorageFlexFlashControllerPropsMoidDefault{
		_statusCode: code,
	}
}

/*PatchStorageFlexFlashControllerPropsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchStorageFlexFlashControllerPropsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch storage flex flash controller props moid default response
func (o *PatchStorageFlexFlashControllerPropsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchStorageFlexFlashControllerPropsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/FlexFlashControllerProps/{moid}][%d] PatchStorageFlexFlashControllerPropsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchStorageFlexFlashControllerPropsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchStorageFlexFlashControllerPropsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
