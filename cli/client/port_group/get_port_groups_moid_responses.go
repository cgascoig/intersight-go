// Code generated by go-swagger; DO NOT EDIT.

package port_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetPortGroupsMoidReader is a Reader for the GetPortGroupsMoid structure.
type GetPortGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPortGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPortGroupsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPortGroupsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPortGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPortGroupsMoidOK creates a GetPortGroupsMoidOK with default headers values
func NewGetPortGroupsMoidOK() *GetPortGroupsMoidOK {
	return &GetPortGroupsMoidOK{}
}

/*GetPortGroupsMoidOK handles this case with default header values.

An instance of portGroup
*/
type GetPortGroupsMoidOK struct {
	Payload *models.PortGroup
}

func (o *GetPortGroupsMoidOK) Error() string {
	return fmt.Sprintf("[GET /port/Groups/{moid}][%d] getPortGroupsMoidOK  %+v", 200, o.Payload)
}

func (o *GetPortGroupsMoidOK) GetPayload() *models.PortGroup {
	return o.Payload
}

func (o *GetPortGroupsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPortGroupsMoidNotFound creates a GetPortGroupsMoidNotFound with default headers values
func NewGetPortGroupsMoidNotFound() *GetPortGroupsMoidNotFound {
	return &GetPortGroupsMoidNotFound{}
}

/*GetPortGroupsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetPortGroupsMoidNotFound struct {
}

func (o *GetPortGroupsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /port/Groups/{moid}][%d] getPortGroupsMoidNotFound ", 404)
}

func (o *GetPortGroupsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPortGroupsMoidDefault creates a GetPortGroupsMoidDefault with default headers values
func NewGetPortGroupsMoidDefault(code int) *GetPortGroupsMoidDefault {
	return &GetPortGroupsMoidDefault{
		_statusCode: code,
	}
}

/*GetPortGroupsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetPortGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get port groups moid default response
func (o *GetPortGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetPortGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /port/Groups/{moid}][%d] GetPortGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetPortGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPortGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}