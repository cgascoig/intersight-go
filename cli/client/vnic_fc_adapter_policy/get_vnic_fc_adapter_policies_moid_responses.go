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

// GetVnicFcAdapterPoliciesMoidReader is a Reader for the GetVnicFcAdapterPoliciesMoid structure.
type GetVnicFcAdapterPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVnicFcAdapterPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVnicFcAdapterPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetVnicFcAdapterPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetVnicFcAdapterPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVnicFcAdapterPoliciesMoidOK creates a GetVnicFcAdapterPoliciesMoidOK with default headers values
func NewGetVnicFcAdapterPoliciesMoidOK() *GetVnicFcAdapterPoliciesMoidOK {
	return &GetVnicFcAdapterPoliciesMoidOK{}
}

/*GetVnicFcAdapterPoliciesMoidOK handles this case with default header values.

An instance of vnicFcAdapterPolicy
*/
type GetVnicFcAdapterPoliciesMoidOK struct {
	Payload *models.VnicFcAdapterPolicy
}

func (o *GetVnicFcAdapterPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /vnic/FcAdapterPolicies/{moid}][%d] getVnicFcAdapterPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetVnicFcAdapterPoliciesMoidOK) GetPayload() *models.VnicFcAdapterPolicy {
	return o.Payload
}

func (o *GetVnicFcAdapterPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VnicFcAdapterPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVnicFcAdapterPoliciesMoidNotFound creates a GetVnicFcAdapterPoliciesMoidNotFound with default headers values
func NewGetVnicFcAdapterPoliciesMoidNotFound() *GetVnicFcAdapterPoliciesMoidNotFound {
	return &GetVnicFcAdapterPoliciesMoidNotFound{}
}

/*GetVnicFcAdapterPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetVnicFcAdapterPoliciesMoidNotFound struct {
}

func (o *GetVnicFcAdapterPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /vnic/FcAdapterPolicies/{moid}][%d] getVnicFcAdapterPoliciesMoidNotFound ", 404)
}

func (o *GetVnicFcAdapterPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVnicFcAdapterPoliciesMoidDefault creates a GetVnicFcAdapterPoliciesMoidDefault with default headers values
func NewGetVnicFcAdapterPoliciesMoidDefault(code int) *GetVnicFcAdapterPoliciesMoidDefault {
	return &GetVnicFcAdapterPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetVnicFcAdapterPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetVnicFcAdapterPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get vnic fc adapter policies moid default response
func (o *GetVnicFcAdapterPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetVnicFcAdapterPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /vnic/FcAdapterPolicies/{moid}][%d] GetVnicFcAdapterPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetVnicFcAdapterPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVnicFcAdapterPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}