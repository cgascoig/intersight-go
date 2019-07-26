// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_protection_group_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureProtectionGroupSnapshotsMoidReader is a Reader for the GetStoragePureProtectionGroupSnapshotsMoid structure.
type GetStoragePureProtectionGroupSnapshotsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureProtectionGroupSnapshotsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureProtectionGroupSnapshotsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePureProtectionGroupSnapshotsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePureProtectionGroupSnapshotsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureProtectionGroupSnapshotsMoidOK creates a GetStoragePureProtectionGroupSnapshotsMoidOK with default headers values
func NewGetStoragePureProtectionGroupSnapshotsMoidOK() *GetStoragePureProtectionGroupSnapshotsMoidOK {
	return &GetStoragePureProtectionGroupSnapshotsMoidOK{}
}

/*GetStoragePureProtectionGroupSnapshotsMoidOK handles this case with default header values.

An instance of storagePureProtectionGroupSnapshot
*/
type GetStoragePureProtectionGroupSnapshotsMoidOK struct {
	Payload *models.StoragePureProtectionGroupSnapshot
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroupSnapshots/{moid}][%d] getStoragePureProtectionGroupSnapshotsMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidOK) GetPayload() *models.StoragePureProtectionGroupSnapshot {
	return o.Payload
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureProtectionGroupSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureProtectionGroupSnapshotsMoidNotFound creates a GetStoragePureProtectionGroupSnapshotsMoidNotFound with default headers values
func NewGetStoragePureProtectionGroupSnapshotsMoidNotFound() *GetStoragePureProtectionGroupSnapshotsMoidNotFound {
	return &GetStoragePureProtectionGroupSnapshotsMoidNotFound{}
}

/*GetStoragePureProtectionGroupSnapshotsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePureProtectionGroupSnapshotsMoidNotFound struct {
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroupSnapshots/{moid}][%d] getStoragePureProtectionGroupSnapshotsMoidNotFound ", 404)
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePureProtectionGroupSnapshotsMoidDefault creates a GetStoragePureProtectionGroupSnapshotsMoidDefault with default headers values
func NewGetStoragePureProtectionGroupSnapshotsMoidDefault(code int) *GetStoragePureProtectionGroupSnapshotsMoidDefault {
	return &GetStoragePureProtectionGroupSnapshotsMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePureProtectionGroupSnapshotsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureProtectionGroupSnapshotsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure protection group snapshots moid default response
func (o *GetStoragePureProtectionGroupSnapshotsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroupSnapshots/{moid}][%d] GetStoragePureProtectionGroupSnapshotsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureProtectionGroupSnapshotsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
