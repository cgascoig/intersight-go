// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_volume_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureVolumeSnapshotsMoidReader is a Reader for the GetStoragePureVolumeSnapshotsMoid structure.
type GetStoragePureVolumeSnapshotsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureVolumeSnapshotsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureVolumeSnapshotsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePureVolumeSnapshotsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePureVolumeSnapshotsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureVolumeSnapshotsMoidOK creates a GetStoragePureVolumeSnapshotsMoidOK with default headers values
func NewGetStoragePureVolumeSnapshotsMoidOK() *GetStoragePureVolumeSnapshotsMoidOK {
	return &GetStoragePureVolumeSnapshotsMoidOK{}
}

/*GetStoragePureVolumeSnapshotsMoidOK handles this case with default header values.

An instance of storagePureVolumeSnapshot
*/
type GetStoragePureVolumeSnapshotsMoidOK struct {
	Payload *models.StoragePureVolumeSnapshot
}

func (o *GetStoragePureVolumeSnapshotsMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureVolumeSnapshots/{moid}][%d] getStoragePureVolumeSnapshotsMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureVolumeSnapshotsMoidOK) GetPayload() *models.StoragePureVolumeSnapshot {
	return o.Payload
}

func (o *GetStoragePureVolumeSnapshotsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureVolumeSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureVolumeSnapshotsMoidNotFound creates a GetStoragePureVolumeSnapshotsMoidNotFound with default headers values
func NewGetStoragePureVolumeSnapshotsMoidNotFound() *GetStoragePureVolumeSnapshotsMoidNotFound {
	return &GetStoragePureVolumeSnapshotsMoidNotFound{}
}

/*GetStoragePureVolumeSnapshotsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePureVolumeSnapshotsMoidNotFound struct {
}

func (o *GetStoragePureVolumeSnapshotsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PureVolumeSnapshots/{moid}][%d] getStoragePureVolumeSnapshotsMoidNotFound ", 404)
}

func (o *GetStoragePureVolumeSnapshotsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePureVolumeSnapshotsMoidDefault creates a GetStoragePureVolumeSnapshotsMoidDefault with default headers values
func NewGetStoragePureVolumeSnapshotsMoidDefault(code int) *GetStoragePureVolumeSnapshotsMoidDefault {
	return &GetStoragePureVolumeSnapshotsMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePureVolumeSnapshotsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureVolumeSnapshotsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure volume snapshots moid default response
func (o *GetStoragePureVolumeSnapshotsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureVolumeSnapshotsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureVolumeSnapshots/{moid}][%d] GetStoragePureVolumeSnapshotsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureVolumeSnapshotsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureVolumeSnapshotsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
