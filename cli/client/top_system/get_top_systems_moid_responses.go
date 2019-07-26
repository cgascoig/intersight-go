// Code generated by go-swagger; DO NOT EDIT.

package top_system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetTopSystemsMoidReader is a Reader for the GetTopSystemsMoid structure.
type GetTopSystemsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTopSystemsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTopSystemsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTopSystemsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTopSystemsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTopSystemsMoidOK creates a GetTopSystemsMoidOK with default headers values
func NewGetTopSystemsMoidOK() *GetTopSystemsMoidOK {
	return &GetTopSystemsMoidOK{}
}

/*GetTopSystemsMoidOK handles this case with default header values.

An instance of topSystem
*/
type GetTopSystemsMoidOK struct {
	Payload *models.TopSystem
}

func (o *GetTopSystemsMoidOK) Error() string {
	return fmt.Sprintf("[GET /top/Systems/{moid}][%d] getTopSystemsMoidOK  %+v", 200, o.Payload)
}

func (o *GetTopSystemsMoidOK) GetPayload() *models.TopSystem {
	return o.Payload
}

func (o *GetTopSystemsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopSystem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTopSystemsMoidNotFound creates a GetTopSystemsMoidNotFound with default headers values
func NewGetTopSystemsMoidNotFound() *GetTopSystemsMoidNotFound {
	return &GetTopSystemsMoidNotFound{}
}

/*GetTopSystemsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetTopSystemsMoidNotFound struct {
}

func (o *GetTopSystemsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /top/Systems/{moid}][%d] getTopSystemsMoidNotFound ", 404)
}

func (o *GetTopSystemsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTopSystemsMoidDefault creates a GetTopSystemsMoidDefault with default headers values
func NewGetTopSystemsMoidDefault(code int) *GetTopSystemsMoidDefault {
	return &GetTopSystemsMoidDefault{
		_statusCode: code,
	}
}

/*GetTopSystemsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetTopSystemsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get top systems moid default response
func (o *GetTopSystemsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetTopSystemsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /top/Systems/{moid}][%d] GetTopSystemsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetTopSystemsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTopSystemsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}