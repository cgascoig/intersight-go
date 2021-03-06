// Code generated by go-swagger; DO NOT EDIT.

package server_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteServerProfilesMoidReader is a Reader for the DeleteServerProfilesMoid structure.
type DeleteServerProfilesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServerProfilesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServerProfilesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteServerProfilesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServerProfilesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServerProfilesMoidOK creates a DeleteServerProfilesMoidOK with default headers values
func NewDeleteServerProfilesMoidOK() *DeleteServerProfilesMoidOK {
	return &DeleteServerProfilesMoidOK{}
}

/*DeleteServerProfilesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteServerProfilesMoidOK struct {
}

func (o *DeleteServerProfilesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /server/Profiles/{moid}][%d] deleteServerProfilesMoidOK ", 200)
}

func (o *DeleteServerProfilesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerProfilesMoidNotFound creates a DeleteServerProfilesMoidNotFound with default headers values
func NewDeleteServerProfilesMoidNotFound() *DeleteServerProfilesMoidNotFound {
	return &DeleteServerProfilesMoidNotFound{}
}

/*DeleteServerProfilesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteServerProfilesMoidNotFound struct {
}

func (o *DeleteServerProfilesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /server/Profiles/{moid}][%d] deleteServerProfilesMoidNotFound ", 404)
}

func (o *DeleteServerProfilesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerProfilesMoidDefault creates a DeleteServerProfilesMoidDefault with default headers values
func NewDeleteServerProfilesMoidDefault(code int) *DeleteServerProfilesMoidDefault {
	return &DeleteServerProfilesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteServerProfilesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteServerProfilesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete server profiles moid default response
func (o *DeleteServerProfilesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServerProfilesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /server/Profiles/{moid}][%d] DeleteServerProfilesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServerProfilesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerProfilesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
