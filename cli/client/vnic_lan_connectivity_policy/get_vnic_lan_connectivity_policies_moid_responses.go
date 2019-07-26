// Code generated by go-swagger; DO NOT EDIT.

package vnic_lan_connectivity_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetVnicLanConnectivityPoliciesMoidReader is a Reader for the GetVnicLanConnectivityPoliciesMoid structure.
type GetVnicLanConnectivityPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVnicLanConnectivityPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVnicLanConnectivityPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetVnicLanConnectivityPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetVnicLanConnectivityPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVnicLanConnectivityPoliciesMoidOK creates a GetVnicLanConnectivityPoliciesMoidOK with default headers values
func NewGetVnicLanConnectivityPoliciesMoidOK() *GetVnicLanConnectivityPoliciesMoidOK {
	return &GetVnicLanConnectivityPoliciesMoidOK{}
}

/*GetVnicLanConnectivityPoliciesMoidOK handles this case with default header values.

An instance of vnicLanConnectivityPolicy
*/
type GetVnicLanConnectivityPoliciesMoidOK struct {
	Payload *models.VnicLanConnectivityPolicy
}

func (o *GetVnicLanConnectivityPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /vnic/LanConnectivityPolicies/{moid}][%d] getVnicLanConnectivityPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetVnicLanConnectivityPoliciesMoidOK) GetPayload() *models.VnicLanConnectivityPolicy {
	return o.Payload
}

func (o *GetVnicLanConnectivityPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VnicLanConnectivityPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVnicLanConnectivityPoliciesMoidNotFound creates a GetVnicLanConnectivityPoliciesMoidNotFound with default headers values
func NewGetVnicLanConnectivityPoliciesMoidNotFound() *GetVnicLanConnectivityPoliciesMoidNotFound {
	return &GetVnicLanConnectivityPoliciesMoidNotFound{}
}

/*GetVnicLanConnectivityPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetVnicLanConnectivityPoliciesMoidNotFound struct {
}

func (o *GetVnicLanConnectivityPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /vnic/LanConnectivityPolicies/{moid}][%d] getVnicLanConnectivityPoliciesMoidNotFound ", 404)
}

func (o *GetVnicLanConnectivityPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVnicLanConnectivityPoliciesMoidDefault creates a GetVnicLanConnectivityPoliciesMoidDefault with default headers values
func NewGetVnicLanConnectivityPoliciesMoidDefault(code int) *GetVnicLanConnectivityPoliciesMoidDefault {
	return &GetVnicLanConnectivityPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetVnicLanConnectivityPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetVnicLanConnectivityPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get vnic lan connectivity policies moid default response
func (o *GetVnicLanConnectivityPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetVnicLanConnectivityPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /vnic/LanConnectivityPolicies/{moid}][%d] GetVnicLanConnectivityPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetVnicLanConnectivityPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVnicLanConnectivityPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}