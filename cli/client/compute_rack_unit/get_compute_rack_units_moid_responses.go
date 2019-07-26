// Code generated by go-swagger; DO NOT EDIT.

package compute_rack_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetComputeRackUnitsMoidReader is a Reader for the GetComputeRackUnitsMoid structure.
type GetComputeRackUnitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeRackUnitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeRackUnitsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetComputeRackUnitsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetComputeRackUnitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetComputeRackUnitsMoidOK creates a GetComputeRackUnitsMoidOK with default headers values
func NewGetComputeRackUnitsMoidOK() *GetComputeRackUnitsMoidOK {
	return &GetComputeRackUnitsMoidOK{}
}

/*GetComputeRackUnitsMoidOK handles this case with default header values.

An instance of computeRackUnit
*/
type GetComputeRackUnitsMoidOK struct {
	Payload *models.ComputeRackUnit
}

func (o *GetComputeRackUnitsMoidOK) Error() string {
	return fmt.Sprintf("[GET /compute/RackUnits/{moid}][%d] getComputeRackUnitsMoidOK  %+v", 200, o.Payload)
}

func (o *GetComputeRackUnitsMoidOK) GetPayload() *models.ComputeRackUnit {
	return o.Payload
}

func (o *GetComputeRackUnitsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeRackUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeRackUnitsMoidNotFound creates a GetComputeRackUnitsMoidNotFound with default headers values
func NewGetComputeRackUnitsMoidNotFound() *GetComputeRackUnitsMoidNotFound {
	return &GetComputeRackUnitsMoidNotFound{}
}

/*GetComputeRackUnitsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetComputeRackUnitsMoidNotFound struct {
}

func (o *GetComputeRackUnitsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /compute/RackUnits/{moid}][%d] getComputeRackUnitsMoidNotFound ", 404)
}

func (o *GetComputeRackUnitsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComputeRackUnitsMoidDefault creates a GetComputeRackUnitsMoidDefault with default headers values
func NewGetComputeRackUnitsMoidDefault(code int) *GetComputeRackUnitsMoidDefault {
	return &GetComputeRackUnitsMoidDefault{
		_statusCode: code,
	}
}

/*GetComputeRackUnitsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetComputeRackUnitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get compute rack units moid default response
func (o *GetComputeRackUnitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetComputeRackUnitsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /compute/RackUnits/{moid}][%d] GetComputeRackUnitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetComputeRackUnitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetComputeRackUnitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}