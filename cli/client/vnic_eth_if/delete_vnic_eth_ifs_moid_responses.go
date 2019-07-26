// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_if

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteVnicEthIfsMoidReader is a Reader for the DeleteVnicEthIfsMoid structure.
type DeleteVnicEthIfsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVnicEthIfsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVnicEthIfsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteVnicEthIfsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteVnicEthIfsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVnicEthIfsMoidOK creates a DeleteVnicEthIfsMoidOK with default headers values
func NewDeleteVnicEthIfsMoidOK() *DeleteVnicEthIfsMoidOK {
	return &DeleteVnicEthIfsMoidOK{}
}

/*DeleteVnicEthIfsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteVnicEthIfsMoidOK struct {
}

func (o *DeleteVnicEthIfsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /vnic/EthIfs/{moid}][%d] deleteVnicEthIfsMoidOK ", 200)
}

func (o *DeleteVnicEthIfsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVnicEthIfsMoidNotFound creates a DeleteVnicEthIfsMoidNotFound with default headers values
func NewDeleteVnicEthIfsMoidNotFound() *DeleteVnicEthIfsMoidNotFound {
	return &DeleteVnicEthIfsMoidNotFound{}
}

/*DeleteVnicEthIfsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteVnicEthIfsMoidNotFound struct {
}

func (o *DeleteVnicEthIfsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vnic/EthIfs/{moid}][%d] deleteVnicEthIfsMoidNotFound ", 404)
}

func (o *DeleteVnicEthIfsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVnicEthIfsMoidDefault creates a DeleteVnicEthIfsMoidDefault with default headers values
func NewDeleteVnicEthIfsMoidDefault(code int) *DeleteVnicEthIfsMoidDefault {
	return &DeleteVnicEthIfsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteVnicEthIfsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteVnicEthIfsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete vnic eth ifs moid default response
func (o *DeleteVnicEthIfsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVnicEthIfsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /vnic/EthIfs/{moid}][%d] DeleteVnicEthIfsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVnicEthIfsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVnicEthIfsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
