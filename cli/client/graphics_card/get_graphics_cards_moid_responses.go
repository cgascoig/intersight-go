// Code generated by go-swagger; DO NOT EDIT.

package graphics_card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetGraphicsCardsMoidReader is a Reader for the GetGraphicsCardsMoid structure.
type GetGraphicsCardsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGraphicsCardsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGraphicsCardsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetGraphicsCardsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetGraphicsCardsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGraphicsCardsMoidOK creates a GetGraphicsCardsMoidOK with default headers values
func NewGetGraphicsCardsMoidOK() *GetGraphicsCardsMoidOK {
	return &GetGraphicsCardsMoidOK{}
}

/*GetGraphicsCardsMoidOK handles this case with default header values.

An instance of graphicsCard
*/
type GetGraphicsCardsMoidOK struct {
	Payload *models.GraphicsCard
}

func (o *GetGraphicsCardsMoidOK) Error() string {
	return fmt.Sprintf("[GET /graphics/Cards/{moid}][%d] getGraphicsCardsMoidOK  %+v", 200, o.Payload)
}

func (o *GetGraphicsCardsMoidOK) GetPayload() *models.GraphicsCard {
	return o.Payload
}

func (o *GetGraphicsCardsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphicsCard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGraphicsCardsMoidNotFound creates a GetGraphicsCardsMoidNotFound with default headers values
func NewGetGraphicsCardsMoidNotFound() *GetGraphicsCardsMoidNotFound {
	return &GetGraphicsCardsMoidNotFound{}
}

/*GetGraphicsCardsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetGraphicsCardsMoidNotFound struct {
}

func (o *GetGraphicsCardsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /graphics/Cards/{moid}][%d] getGraphicsCardsMoidNotFound ", 404)
}

func (o *GetGraphicsCardsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGraphicsCardsMoidDefault creates a GetGraphicsCardsMoidDefault with default headers values
func NewGetGraphicsCardsMoidDefault(code int) *GetGraphicsCardsMoidDefault {
	return &GetGraphicsCardsMoidDefault{
		_statusCode: code,
	}
}

/*GetGraphicsCardsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetGraphicsCardsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get graphics cards moid default response
func (o *GetGraphicsCardsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetGraphicsCardsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /graphics/Cards/{moid}][%d] GetGraphicsCardsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetGraphicsCardsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGraphicsCardsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}