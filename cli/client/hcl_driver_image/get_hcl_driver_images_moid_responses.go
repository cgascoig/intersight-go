// Code generated by go-swagger; DO NOT EDIT.

package hcl_driver_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHclDriverImagesMoidReader is a Reader for the GetHclDriverImagesMoid structure.
type GetHclDriverImagesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHclDriverImagesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHclDriverImagesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHclDriverImagesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHclDriverImagesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHclDriverImagesMoidOK creates a GetHclDriverImagesMoidOK with default headers values
func NewGetHclDriverImagesMoidOK() *GetHclDriverImagesMoidOK {
	return &GetHclDriverImagesMoidOK{}
}

/*GetHclDriverImagesMoidOK handles this case with default header values.

An instance of hclDriverImage
*/
type GetHclDriverImagesMoidOK struct {
	Payload *models.HclDriverImage
}

func (o *GetHclDriverImagesMoidOK) Error() string {
	return fmt.Sprintf("[GET /hcl/DriverImages/{moid}][%d] getHclDriverImagesMoidOK  %+v", 200, o.Payload)
}

func (o *GetHclDriverImagesMoidOK) GetPayload() *models.HclDriverImage {
	return o.Payload
}

func (o *GetHclDriverImagesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HclDriverImage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHclDriverImagesMoidNotFound creates a GetHclDriverImagesMoidNotFound with default headers values
func NewGetHclDriverImagesMoidNotFound() *GetHclDriverImagesMoidNotFound {
	return &GetHclDriverImagesMoidNotFound{}
}

/*GetHclDriverImagesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHclDriverImagesMoidNotFound struct {
}

func (o *GetHclDriverImagesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hcl/DriverImages/{moid}][%d] getHclDriverImagesMoidNotFound ", 404)
}

func (o *GetHclDriverImagesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHclDriverImagesMoidDefault creates a GetHclDriverImagesMoidDefault with default headers values
func NewGetHclDriverImagesMoidDefault(code int) *GetHclDriverImagesMoidDefault {
	return &GetHclDriverImagesMoidDefault{
		_statusCode: code,
	}
}

/*GetHclDriverImagesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHclDriverImagesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hcl driver images moid default response
func (o *GetHclDriverImagesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHclDriverImagesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hcl/DriverImages/{moid}][%d] GetHclDriverImagesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHclDriverImagesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHclDriverImagesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}