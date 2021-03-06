// Code generated by go-swagger; DO NOT EDIT.

package storage_enclosure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchStorageEnclosuresMoidReader is a Reader for the PatchStorageEnclosuresMoid structure.
type PatchStorageEnclosuresMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageEnclosuresMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchStorageEnclosuresMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchStorageEnclosuresMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchStorageEnclosuresMoidCreated creates a PatchStorageEnclosuresMoidCreated with default headers values
func NewPatchStorageEnclosuresMoidCreated() *PatchStorageEnclosuresMoidCreated {
	return &PatchStorageEnclosuresMoidCreated{}
}

/*PatchStorageEnclosuresMoidCreated handles this case with default header values.

Null response
*/
type PatchStorageEnclosuresMoidCreated struct {
}

func (o *PatchStorageEnclosuresMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /storage/Enclosures/{moid}][%d] patchStorageEnclosuresMoidCreated ", 201)
}

func (o *PatchStorageEnclosuresMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchStorageEnclosuresMoidDefault creates a PatchStorageEnclosuresMoidDefault with default headers values
func NewPatchStorageEnclosuresMoidDefault(code int) *PatchStorageEnclosuresMoidDefault {
	return &PatchStorageEnclosuresMoidDefault{
		_statusCode: code,
	}
}

/*PatchStorageEnclosuresMoidDefault handles this case with default header values.

unexpected error
*/
type PatchStorageEnclosuresMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch storage enclosures moid default response
func (o *PatchStorageEnclosuresMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchStorageEnclosuresMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/Enclosures/{moid}][%d] PatchStorageEnclosuresMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchStorageEnclosuresMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchStorageEnclosuresMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
