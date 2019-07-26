// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_if

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteVnicFcIfsMoidReader is a Reader for the DeleteVnicFcIfsMoid structure.
type DeleteVnicFcIfsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVnicFcIfsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVnicFcIfsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteVnicFcIfsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteVnicFcIfsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVnicFcIfsMoidOK creates a DeleteVnicFcIfsMoidOK with default headers values
func NewDeleteVnicFcIfsMoidOK() *DeleteVnicFcIfsMoidOK {
	return &DeleteVnicFcIfsMoidOK{}
}

/*DeleteVnicFcIfsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteVnicFcIfsMoidOK struct {
}

func (o *DeleteVnicFcIfsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /vnic/FcIfs/{moid}][%d] deleteVnicFcIfsMoidOK ", 200)
}

func (o *DeleteVnicFcIfsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVnicFcIfsMoidNotFound creates a DeleteVnicFcIfsMoidNotFound with default headers values
func NewDeleteVnicFcIfsMoidNotFound() *DeleteVnicFcIfsMoidNotFound {
	return &DeleteVnicFcIfsMoidNotFound{}
}

/*DeleteVnicFcIfsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteVnicFcIfsMoidNotFound struct {
}

func (o *DeleteVnicFcIfsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vnic/FcIfs/{moid}][%d] deleteVnicFcIfsMoidNotFound ", 404)
}

func (o *DeleteVnicFcIfsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVnicFcIfsMoidDefault creates a DeleteVnicFcIfsMoidDefault with default headers values
func NewDeleteVnicFcIfsMoidDefault(code int) *DeleteVnicFcIfsMoidDefault {
	return &DeleteVnicFcIfsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteVnicFcIfsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteVnicFcIfsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete vnic fc ifs moid default response
func (o *DeleteVnicFcIfsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVnicFcIfsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /vnic/FcIfs/{moid}][%d] DeleteVnicFcIfsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVnicFcIfsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVnicFcIfsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
