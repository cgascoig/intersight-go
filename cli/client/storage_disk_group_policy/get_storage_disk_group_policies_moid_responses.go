// Code generated by go-swagger; DO NOT EDIT.

package storage_disk_group_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageDiskGroupPoliciesMoidReader is a Reader for the GetStorageDiskGroupPoliciesMoid structure.
type GetStorageDiskGroupPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageDiskGroupPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageDiskGroupPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStorageDiskGroupPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStorageDiskGroupPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageDiskGroupPoliciesMoidOK creates a GetStorageDiskGroupPoliciesMoidOK with default headers values
func NewGetStorageDiskGroupPoliciesMoidOK() *GetStorageDiskGroupPoliciesMoidOK {
	return &GetStorageDiskGroupPoliciesMoidOK{}
}

/*GetStorageDiskGroupPoliciesMoidOK handles this case with default header values.

An instance of storageDiskGroupPolicy
*/
type GetStorageDiskGroupPoliciesMoidOK struct {
	Payload *models.StorageDiskGroupPolicy
}

func (o *GetStorageDiskGroupPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/DiskGroupPolicies/{moid}][%d] getStorageDiskGroupPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetStorageDiskGroupPoliciesMoidOK) GetPayload() *models.StorageDiskGroupPolicy {
	return o.Payload
}

func (o *GetStorageDiskGroupPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageDiskGroupPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageDiskGroupPoliciesMoidNotFound creates a GetStorageDiskGroupPoliciesMoidNotFound with default headers values
func NewGetStorageDiskGroupPoliciesMoidNotFound() *GetStorageDiskGroupPoliciesMoidNotFound {
	return &GetStorageDiskGroupPoliciesMoidNotFound{}
}

/*GetStorageDiskGroupPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStorageDiskGroupPoliciesMoidNotFound struct {
}

func (o *GetStorageDiskGroupPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/DiskGroupPolicies/{moid}][%d] getStorageDiskGroupPoliciesMoidNotFound ", 404)
}

func (o *GetStorageDiskGroupPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageDiskGroupPoliciesMoidDefault creates a GetStorageDiskGroupPoliciesMoidDefault with default headers values
func NewGetStorageDiskGroupPoliciesMoidDefault(code int) *GetStorageDiskGroupPoliciesMoidDefault {
	return &GetStorageDiskGroupPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetStorageDiskGroupPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStorageDiskGroupPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage disk group policies moid default response
func (o *GetStorageDiskGroupPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageDiskGroupPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/DiskGroupPolicies/{moid}][%d] GetStorageDiskGroupPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageDiskGroupPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageDiskGroupPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
