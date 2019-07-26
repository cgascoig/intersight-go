// Code generated by go-swagger; DO NOT EDIT.

package vnic_san_connectivity_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostVnicSanConnectivityPoliciesMoidReader is a Reader for the PostVnicSanConnectivityPoliciesMoid structure.
type PostVnicSanConnectivityPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVnicSanConnectivityPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVnicSanConnectivityPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVnicSanConnectivityPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVnicSanConnectivityPoliciesMoidCreated creates a PostVnicSanConnectivityPoliciesMoidCreated with default headers values
func NewPostVnicSanConnectivityPoliciesMoidCreated() *PostVnicSanConnectivityPoliciesMoidCreated {
	return &PostVnicSanConnectivityPoliciesMoidCreated{}
}

/*PostVnicSanConnectivityPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostVnicSanConnectivityPoliciesMoidCreated struct {
}

func (o *PostVnicSanConnectivityPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /vnic/SanConnectivityPolicies/{moid}][%d] postVnicSanConnectivityPoliciesMoidCreated ", 201)
}

func (o *PostVnicSanConnectivityPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVnicSanConnectivityPoliciesMoidDefault creates a PostVnicSanConnectivityPoliciesMoidDefault with default headers values
func NewPostVnicSanConnectivityPoliciesMoidDefault(code int) *PostVnicSanConnectivityPoliciesMoidDefault {
	return &PostVnicSanConnectivityPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostVnicSanConnectivityPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostVnicSanConnectivityPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post vnic san connectivity policies moid default response
func (o *PostVnicSanConnectivityPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostVnicSanConnectivityPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /vnic/SanConnectivityPolicies/{moid}][%d] PostVnicSanConnectivityPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostVnicSanConnectivityPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVnicSanConnectivityPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
