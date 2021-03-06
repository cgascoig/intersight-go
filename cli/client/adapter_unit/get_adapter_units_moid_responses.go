// Code generated by go-swagger; DO NOT EDIT.

package adapter_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetAdapterUnitsMoidReader is a Reader for the GetAdapterUnitsMoid structure.
type GetAdapterUnitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdapterUnitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdapterUnitsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAdapterUnitsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAdapterUnitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdapterUnitsMoidOK creates a GetAdapterUnitsMoidOK with default headers values
func NewGetAdapterUnitsMoidOK() *GetAdapterUnitsMoidOK {
	return &GetAdapterUnitsMoidOK{}
}

/*GetAdapterUnitsMoidOK handles this case with default header values.

An instance of adapterUnit
*/
type GetAdapterUnitsMoidOK struct {
	Payload *models.AdapterUnit
}

func (o *GetAdapterUnitsMoidOK) Error() string {
	return fmt.Sprintf("[GET /adapter/Units/{moid}][%d] getAdapterUnitsMoidOK  %+v", 200, o.Payload)
}

func (o *GetAdapterUnitsMoidOK) GetPayload() *models.AdapterUnit {
	return o.Payload
}

func (o *GetAdapterUnitsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdapterUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdapterUnitsMoidNotFound creates a GetAdapterUnitsMoidNotFound with default headers values
func NewGetAdapterUnitsMoidNotFound() *GetAdapterUnitsMoidNotFound {
	return &GetAdapterUnitsMoidNotFound{}
}

/*GetAdapterUnitsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetAdapterUnitsMoidNotFound struct {
}

func (o *GetAdapterUnitsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /adapter/Units/{moid}][%d] getAdapterUnitsMoidNotFound ", 404)
}

func (o *GetAdapterUnitsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdapterUnitsMoidDefault creates a GetAdapterUnitsMoidDefault with default headers values
func NewGetAdapterUnitsMoidDefault(code int) *GetAdapterUnitsMoidDefault {
	return &GetAdapterUnitsMoidDefault{
		_statusCode: code,
	}
}

/*GetAdapterUnitsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetAdapterUnitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get adapter units moid default response
func (o *GetAdapterUnitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetAdapterUnitsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /adapter/Units/{moid}][%d] GetAdapterUnitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdapterUnitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAdapterUnitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
