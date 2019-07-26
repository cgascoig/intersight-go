// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_snapshot_schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureSnapshotSchedulesMoidReader is a Reader for the GetStoragePureSnapshotSchedulesMoid structure.
type GetStoragePureSnapshotSchedulesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureSnapshotSchedulesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureSnapshotSchedulesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePureSnapshotSchedulesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePureSnapshotSchedulesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureSnapshotSchedulesMoidOK creates a GetStoragePureSnapshotSchedulesMoidOK with default headers values
func NewGetStoragePureSnapshotSchedulesMoidOK() *GetStoragePureSnapshotSchedulesMoidOK {
	return &GetStoragePureSnapshotSchedulesMoidOK{}
}

/*GetStoragePureSnapshotSchedulesMoidOK handles this case with default header values.

An instance of storagePureSnapshotSchedule
*/
type GetStoragePureSnapshotSchedulesMoidOK struct {
	Payload *models.StoragePureSnapshotSchedule
}

func (o *GetStoragePureSnapshotSchedulesMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureSnapshotSchedules/{moid}][%d] getStoragePureSnapshotSchedulesMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureSnapshotSchedulesMoidOK) GetPayload() *models.StoragePureSnapshotSchedule {
	return o.Payload
}

func (o *GetStoragePureSnapshotSchedulesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureSnapshotSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureSnapshotSchedulesMoidNotFound creates a GetStoragePureSnapshotSchedulesMoidNotFound with default headers values
func NewGetStoragePureSnapshotSchedulesMoidNotFound() *GetStoragePureSnapshotSchedulesMoidNotFound {
	return &GetStoragePureSnapshotSchedulesMoidNotFound{}
}

/*GetStoragePureSnapshotSchedulesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePureSnapshotSchedulesMoidNotFound struct {
}

func (o *GetStoragePureSnapshotSchedulesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PureSnapshotSchedules/{moid}][%d] getStoragePureSnapshotSchedulesMoidNotFound ", 404)
}

func (o *GetStoragePureSnapshotSchedulesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePureSnapshotSchedulesMoidDefault creates a GetStoragePureSnapshotSchedulesMoidDefault with default headers values
func NewGetStoragePureSnapshotSchedulesMoidDefault(code int) *GetStoragePureSnapshotSchedulesMoidDefault {
	return &GetStoragePureSnapshotSchedulesMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePureSnapshotSchedulesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureSnapshotSchedulesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure snapshot schedules moid default response
func (o *GetStoragePureSnapshotSchedulesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureSnapshotSchedulesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureSnapshotSchedules/{moid}][%d] GetStoragePureSnapshotSchedulesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureSnapshotSchedulesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureSnapshotSchedulesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}