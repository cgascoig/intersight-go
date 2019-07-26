// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_replication_schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureReplicationSchedulesMoidReader is a Reader for the GetStoragePureReplicationSchedulesMoid structure.
type GetStoragePureReplicationSchedulesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureReplicationSchedulesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureReplicationSchedulesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePureReplicationSchedulesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePureReplicationSchedulesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureReplicationSchedulesMoidOK creates a GetStoragePureReplicationSchedulesMoidOK with default headers values
func NewGetStoragePureReplicationSchedulesMoidOK() *GetStoragePureReplicationSchedulesMoidOK {
	return &GetStoragePureReplicationSchedulesMoidOK{}
}

/*GetStoragePureReplicationSchedulesMoidOK handles this case with default header values.

An instance of storagePureReplicationSchedule
*/
type GetStoragePureReplicationSchedulesMoidOK struct {
	Payload *models.StoragePureReplicationSchedule
}

func (o *GetStoragePureReplicationSchedulesMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureReplicationSchedules/{moid}][%d] getStoragePureReplicationSchedulesMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureReplicationSchedulesMoidOK) GetPayload() *models.StoragePureReplicationSchedule {
	return o.Payload
}

func (o *GetStoragePureReplicationSchedulesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureReplicationSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureReplicationSchedulesMoidNotFound creates a GetStoragePureReplicationSchedulesMoidNotFound with default headers values
func NewGetStoragePureReplicationSchedulesMoidNotFound() *GetStoragePureReplicationSchedulesMoidNotFound {
	return &GetStoragePureReplicationSchedulesMoidNotFound{}
}

/*GetStoragePureReplicationSchedulesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePureReplicationSchedulesMoidNotFound struct {
}

func (o *GetStoragePureReplicationSchedulesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PureReplicationSchedules/{moid}][%d] getStoragePureReplicationSchedulesMoidNotFound ", 404)
}

func (o *GetStoragePureReplicationSchedulesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePureReplicationSchedulesMoidDefault creates a GetStoragePureReplicationSchedulesMoidDefault with default headers values
func NewGetStoragePureReplicationSchedulesMoidDefault(code int) *GetStoragePureReplicationSchedulesMoidDefault {
	return &GetStoragePureReplicationSchedulesMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePureReplicationSchedulesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureReplicationSchedulesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure replication schedules moid default response
func (o *GetStoragePureReplicationSchedulesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureReplicationSchedulesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureReplicationSchedules/{moid}][%d] GetStoragePureReplicationSchedulesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureReplicationSchedulesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureReplicationSchedulesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
