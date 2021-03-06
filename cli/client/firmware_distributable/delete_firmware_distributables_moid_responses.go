// Code generated by go-swagger; DO NOT EDIT.

package firmware_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteFirmwareDistributablesMoidReader is a Reader for the DeleteFirmwareDistributablesMoid structure.
type DeleteFirmwareDistributablesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFirmwareDistributablesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFirmwareDistributablesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteFirmwareDistributablesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteFirmwareDistributablesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFirmwareDistributablesMoidOK creates a DeleteFirmwareDistributablesMoidOK with default headers values
func NewDeleteFirmwareDistributablesMoidOK() *DeleteFirmwareDistributablesMoidOK {
	return &DeleteFirmwareDistributablesMoidOK{}
}

/*DeleteFirmwareDistributablesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteFirmwareDistributablesMoidOK struct {
}

func (o *DeleteFirmwareDistributablesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /firmware/Distributables/{moid}][%d] deleteFirmwareDistributablesMoidOK ", 200)
}

func (o *DeleteFirmwareDistributablesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFirmwareDistributablesMoidNotFound creates a DeleteFirmwareDistributablesMoidNotFound with default headers values
func NewDeleteFirmwareDistributablesMoidNotFound() *DeleteFirmwareDistributablesMoidNotFound {
	return &DeleteFirmwareDistributablesMoidNotFound{}
}

/*DeleteFirmwareDistributablesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteFirmwareDistributablesMoidNotFound struct {
}

func (o *DeleteFirmwareDistributablesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /firmware/Distributables/{moid}][%d] deleteFirmwareDistributablesMoidNotFound ", 404)
}

func (o *DeleteFirmwareDistributablesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFirmwareDistributablesMoidDefault creates a DeleteFirmwareDistributablesMoidDefault with default headers values
func NewDeleteFirmwareDistributablesMoidDefault(code int) *DeleteFirmwareDistributablesMoidDefault {
	return &DeleteFirmwareDistributablesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteFirmwareDistributablesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteFirmwareDistributablesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete firmware distributables moid default response
func (o *DeleteFirmwareDistributablesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFirmwareDistributablesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /firmware/Distributables/{moid}][%d] DeleteFirmwareDistributablesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFirmwareDistributablesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFirmwareDistributablesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
