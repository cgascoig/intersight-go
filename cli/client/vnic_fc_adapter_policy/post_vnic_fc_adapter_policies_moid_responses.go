// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_adapter_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostVnicFcAdapterPoliciesMoidReader is a Reader for the PostVnicFcAdapterPoliciesMoid structure.
type PostVnicFcAdapterPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVnicFcAdapterPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVnicFcAdapterPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVnicFcAdapterPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVnicFcAdapterPoliciesMoidCreated creates a PostVnicFcAdapterPoliciesMoidCreated with default headers values
func NewPostVnicFcAdapterPoliciesMoidCreated() *PostVnicFcAdapterPoliciesMoidCreated {
	return &PostVnicFcAdapterPoliciesMoidCreated{}
}

/*PostVnicFcAdapterPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostVnicFcAdapterPoliciesMoidCreated struct {
}

func (o *PostVnicFcAdapterPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /vnic/FcAdapterPolicies/{moid}][%d] postVnicFcAdapterPoliciesMoidCreated ", 201)
}

func (o *PostVnicFcAdapterPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVnicFcAdapterPoliciesMoidDefault creates a PostVnicFcAdapterPoliciesMoidDefault with default headers values
func NewPostVnicFcAdapterPoliciesMoidDefault(code int) *PostVnicFcAdapterPoliciesMoidDefault {
	return &PostVnicFcAdapterPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostVnicFcAdapterPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostVnicFcAdapterPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post vnic fc adapter policies moid default response
func (o *PostVnicFcAdapterPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostVnicFcAdapterPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /vnic/FcAdapterPolicies/{moid}][%d] PostVnicFcAdapterPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostVnicFcAdapterPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVnicFcAdapterPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
