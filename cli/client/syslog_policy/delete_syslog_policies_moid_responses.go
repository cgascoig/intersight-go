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

// DeleteSyslogPoliciesMoidReader is a Reader for the DeleteSyslogPoliciesMoid structure.
type DeleteSyslogPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSyslogPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSyslogPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSyslogPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSyslogPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSyslogPoliciesMoidOK creates a DeleteSyslogPoliciesMoidOK with default headers values
func NewDeleteSyslogPoliciesMoidOK() *DeleteSyslogPoliciesMoidOK {
	return &DeleteSyslogPoliciesMoidOK{}
}

/*DeleteSyslogPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteSyslogPoliciesMoidOK struct {
}

func (o *DeleteSyslogPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /syslog/Policies/{moid}][%d] deleteSyslogPoliciesMoidOK ", 200)
}

func (o *DeleteSyslogPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSyslogPoliciesMoidNotFound creates a DeleteSyslogPoliciesMoidNotFound with default headers values
func NewDeleteSyslogPoliciesMoidNotFound() *DeleteSyslogPoliciesMoidNotFound {
	return &DeleteSyslogPoliciesMoidNotFound{}
}

/*DeleteSyslogPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteSyslogPoliciesMoidNotFound struct {
}

func (o *DeleteSyslogPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /syslog/Policies/{moid}][%d] deleteSyslogPoliciesMoidNotFound ", 404)
}

func (o *DeleteSyslogPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSyslogPoliciesMoidDefault creates a DeleteSyslogPoliciesMoidDefault with default headers values
func NewDeleteSyslogPoliciesMoidDefault(code int) *DeleteSyslogPoliciesMoidDefault {
	return &DeleteSyslogPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteSyslogPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteSyslogPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete syslog policies moid default response
func (o *DeleteSyslogPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSyslogPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /syslog/Policies/{moid}][%d] DeleteSyslogPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyslogPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSyslogPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}