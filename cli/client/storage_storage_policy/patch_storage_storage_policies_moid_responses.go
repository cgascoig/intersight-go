// Code generated by go-swagger; DO NOT EDIT.

package storage_storage_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchStorageStoragePoliciesMoidReader is a Reader for the PatchStorageStoragePoliciesMoid structure.
type PatchStorageStoragePoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageStoragePoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchStorageStoragePoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchStorageStoragePoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchStorageStoragePoliciesMoidCreated creates a PatchStorageStoragePoliciesMoidCreated with default headers values
func NewPatchStorageStoragePoliciesMoidCreated() *PatchStorageStoragePoliciesMoidCreated {
	return &PatchStorageStoragePoliciesMoidCreated{}
}

/*PatchStorageStoragePoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchStorageStoragePoliciesMoidCreated struct {
}

func (o *PatchStorageStoragePoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /storage/StoragePolicies/{moid}][%d] patchStorageStoragePoliciesMoidCreated ", 201)
}

func (o *PatchStorageStoragePoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchStorageStoragePoliciesMoidDefault creates a PatchStorageStoragePoliciesMoidDefault with default headers values
func NewPatchStorageStoragePoliciesMoidDefault(code int) *PatchStorageStoragePoliciesMoidDefault {
	return &PatchStorageStoragePoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchStorageStoragePoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchStorageStoragePoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch storage storage policies moid default response
func (o *PatchStorageStoragePoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchStorageStoragePoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/StoragePolicies/{moid}][%d] PatchStorageStoragePoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchStorageStoragePoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchStorageStoragePoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
