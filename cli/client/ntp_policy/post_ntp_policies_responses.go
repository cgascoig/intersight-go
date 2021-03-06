// Code generated by go-swagger; DO NOT EDIT.

package ntp_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostNtpPoliciesReader is a Reader for the PostNtpPolicies structure.
type PostNtpPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNtpPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNtpPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNtpPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNtpPoliciesCreated creates a PostNtpPoliciesCreated with default headers values
func NewPostNtpPoliciesCreated() *PostNtpPoliciesCreated {
	return &PostNtpPoliciesCreated{}
}

/*PostNtpPoliciesCreated handles this case with default header values.

Null response
*/
type PostNtpPoliciesCreated struct {
}

func (o *PostNtpPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /ntp/Policies][%d] postNtpPoliciesCreated ", 201)
}

func (o *PostNtpPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNtpPoliciesDefault creates a PostNtpPoliciesDefault with default headers values
func NewPostNtpPoliciesDefault(code int) *PostNtpPoliciesDefault {
	return &PostNtpPoliciesDefault{
		_statusCode: code,
	}
}

/*PostNtpPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostNtpPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post ntp policies default response
func (o *PostNtpPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostNtpPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /ntp/Policies][%d] PostNtpPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostNtpPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNtpPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
