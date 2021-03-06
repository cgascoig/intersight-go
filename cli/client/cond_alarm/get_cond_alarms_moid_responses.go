// Code generated by go-swagger; DO NOT EDIT.

package cond_alarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetCondAlarmsMoidReader is a Reader for the GetCondAlarmsMoid structure.
type GetCondAlarmsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCondAlarmsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCondAlarmsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCondAlarmsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCondAlarmsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCondAlarmsMoidOK creates a GetCondAlarmsMoidOK with default headers values
func NewGetCondAlarmsMoidOK() *GetCondAlarmsMoidOK {
	return &GetCondAlarmsMoidOK{}
}

/*GetCondAlarmsMoidOK handles this case with default header values.

An instance of condAlarm
*/
type GetCondAlarmsMoidOK struct {
	Payload *models.CondAlarm
}

func (o *GetCondAlarmsMoidOK) Error() string {
	return fmt.Sprintf("[GET /cond/Alarms/{moid}][%d] getCondAlarmsMoidOK  %+v", 200, o.Payload)
}

func (o *GetCondAlarmsMoidOK) GetPayload() *models.CondAlarm {
	return o.Payload
}

func (o *GetCondAlarmsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CondAlarm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCondAlarmsMoidNotFound creates a GetCondAlarmsMoidNotFound with default headers values
func NewGetCondAlarmsMoidNotFound() *GetCondAlarmsMoidNotFound {
	return &GetCondAlarmsMoidNotFound{}
}

/*GetCondAlarmsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetCondAlarmsMoidNotFound struct {
}

func (o *GetCondAlarmsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /cond/Alarms/{moid}][%d] getCondAlarmsMoidNotFound ", 404)
}

func (o *GetCondAlarmsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCondAlarmsMoidDefault creates a GetCondAlarmsMoidDefault with default headers values
func NewGetCondAlarmsMoidDefault(code int) *GetCondAlarmsMoidDefault {
	return &GetCondAlarmsMoidDefault{
		_statusCode: code,
	}
}

/*GetCondAlarmsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetCondAlarmsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cond alarms moid default response
func (o *GetCondAlarmsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetCondAlarmsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /cond/Alarms/{moid}][%d] GetCondAlarmsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCondAlarmsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCondAlarmsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
