// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_qos_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostVnicFcQosPoliciesReader is a Reader for the PostVnicFcQosPolicies structure.
type PostVnicFcQosPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVnicFcQosPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVnicFcQosPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVnicFcQosPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVnicFcQosPoliciesCreated creates a PostVnicFcQosPoliciesCreated with default headers values
func NewPostVnicFcQosPoliciesCreated() *PostVnicFcQosPoliciesCreated {
	return &PostVnicFcQosPoliciesCreated{}
}

/*PostVnicFcQosPoliciesCreated handles this case with default header values.

Null response
*/
type PostVnicFcQosPoliciesCreated struct {
}

func (o *PostVnicFcQosPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /vnic/FcQosPolicies][%d] postVnicFcQosPoliciesCreated ", 201)
}

func (o *PostVnicFcQosPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVnicFcQosPoliciesDefault creates a PostVnicFcQosPoliciesDefault with default headers values
func NewPostVnicFcQosPoliciesDefault(code int) *PostVnicFcQosPoliciesDefault {
	return &PostVnicFcQosPoliciesDefault{
		_statusCode: code,
	}
}

/*PostVnicFcQosPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostVnicFcQosPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post vnic fc qos policies default response
func (o *PostVnicFcQosPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostVnicFcQosPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /vnic/FcQosPolicies][%d] PostVnicFcQosPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostVnicFcQosPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVnicFcQosPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
