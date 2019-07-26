// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_network_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteVnicFcNetworkPoliciesMoidReader is a Reader for the DeleteVnicFcNetworkPoliciesMoid structure.
type DeleteVnicFcNetworkPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVnicFcNetworkPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVnicFcNetworkPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteVnicFcNetworkPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteVnicFcNetworkPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVnicFcNetworkPoliciesMoidOK creates a DeleteVnicFcNetworkPoliciesMoidOK with default headers values
func NewDeleteVnicFcNetworkPoliciesMoidOK() *DeleteVnicFcNetworkPoliciesMoidOK {
	return &DeleteVnicFcNetworkPoliciesMoidOK{}
}

/*DeleteVnicFcNetworkPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteVnicFcNetworkPoliciesMoidOK struct {
}

func (o *DeleteVnicFcNetworkPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /vnic/FcNetworkPolicies/{moid}][%d] deleteVnicFcNetworkPoliciesMoidOK ", 200)
}

func (o *DeleteVnicFcNetworkPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVnicFcNetworkPoliciesMoidNotFound creates a DeleteVnicFcNetworkPoliciesMoidNotFound with default headers values
func NewDeleteVnicFcNetworkPoliciesMoidNotFound() *DeleteVnicFcNetworkPoliciesMoidNotFound {
	return &DeleteVnicFcNetworkPoliciesMoidNotFound{}
}

/*DeleteVnicFcNetworkPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteVnicFcNetworkPoliciesMoidNotFound struct {
}

func (o *DeleteVnicFcNetworkPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vnic/FcNetworkPolicies/{moid}][%d] deleteVnicFcNetworkPoliciesMoidNotFound ", 404)
}

func (o *DeleteVnicFcNetworkPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVnicFcNetworkPoliciesMoidDefault creates a DeleteVnicFcNetworkPoliciesMoidDefault with default headers values
func NewDeleteVnicFcNetworkPoliciesMoidDefault(code int) *DeleteVnicFcNetworkPoliciesMoidDefault {
	return &DeleteVnicFcNetworkPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteVnicFcNetworkPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteVnicFcNetworkPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete vnic fc network policies moid default response
func (o *DeleteVnicFcNetworkPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVnicFcNetworkPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /vnic/FcNetworkPolicies/{moid}][%d] DeleteVnicFcNetworkPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVnicFcNetworkPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVnicFcNetworkPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
