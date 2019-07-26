// Code generated by go-swagger; DO NOT EDIT.

package syslog_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetSyslogPoliciesMoidReader is a Reader for the GetSyslogPoliciesMoid structure.
type GetSyslogPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyslogPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyslogPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSyslogPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSyslogPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSyslogPoliciesMoidOK creates a GetSyslogPoliciesMoidOK with default headers values
func NewGetSyslogPoliciesMoidOK() *GetSyslogPoliciesMoidOK {
	return &GetSyslogPoliciesMoidOK{}
}

/*GetSyslogPoliciesMoidOK handles this case with default header values.

An instance of syslogPolicy
*/
type GetSyslogPoliciesMoidOK struct {
	Payload *models.SyslogPolicy
}

func (o *GetSyslogPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /syslog/Policies/{moid}][%d] getSyslogPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetSyslogPoliciesMoidOK) GetPayload() *models.SyslogPolicy {
	return o.Payload
}

func (o *GetSyslogPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyslogPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyslogPoliciesMoidNotFound creates a GetSyslogPoliciesMoidNotFound with default headers values
func NewGetSyslogPoliciesMoidNotFound() *GetSyslogPoliciesMoidNotFound {
	return &GetSyslogPoliciesMoidNotFound{}
}

/*GetSyslogPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetSyslogPoliciesMoidNotFound struct {
}

func (o *GetSyslogPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /syslog/Policies/{moid}][%d] getSyslogPoliciesMoidNotFound ", 404)
}

func (o *GetSyslogPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSyslogPoliciesMoidDefault creates a GetSyslogPoliciesMoidDefault with default headers values
func NewGetSyslogPoliciesMoidDefault(code int) *GetSyslogPoliciesMoidDefault {
	return &GetSyslogPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetSyslogPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetSyslogPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get syslog policies moid default response
func (o *GetSyslogPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetSyslogPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /syslog/Policies/{moid}][%d] GetSyslogPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetSyslogPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSyslogPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
