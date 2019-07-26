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

// PostSyslogPoliciesReader is a Reader for the PostSyslogPolicies structure.
type PostSyslogPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSyslogPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSyslogPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostSyslogPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSyslogPoliciesCreated creates a PostSyslogPoliciesCreated with default headers values
func NewPostSyslogPoliciesCreated() *PostSyslogPoliciesCreated {
	return &PostSyslogPoliciesCreated{}
}

/*PostSyslogPoliciesCreated handles this case with default header values.

Null response
*/
type PostSyslogPoliciesCreated struct {
}

func (o *PostSyslogPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /syslog/Policies][%d] postSyslogPoliciesCreated ", 201)
}

func (o *PostSyslogPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSyslogPoliciesDefault creates a PostSyslogPoliciesDefault with default headers values
func NewPostSyslogPoliciesDefault(code int) *PostSyslogPoliciesDefault {
	return &PostSyslogPoliciesDefault{
		_statusCode: code,
	}
}

/*PostSyslogPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostSyslogPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post syslog policies default response
func (o *PostSyslogPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostSyslogPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /syslog/Policies][%d] PostSyslogPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostSyslogPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSyslogPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}