// Code generated by go-swagger; DO NOT EDIT.

package processor_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetProcessorUnitsMoidReader is a Reader for the GetProcessorUnitsMoid structure.
type GetProcessorUnitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorUnitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessorUnitsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProcessorUnitsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProcessorUnitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProcessorUnitsMoidOK creates a GetProcessorUnitsMoidOK with default headers values
func NewGetProcessorUnitsMoidOK() *GetProcessorUnitsMoidOK {
	return &GetProcessorUnitsMoidOK{}
}

/*GetProcessorUnitsMoidOK handles this case with default header values.

An instance of processorUnit
*/
type GetProcessorUnitsMoidOK struct {
	Payload *models.ProcessorUnit
}

func (o *GetProcessorUnitsMoidOK) Error() string {
	return fmt.Sprintf("[GET /processor/Units/{moid}][%d] getProcessorUnitsMoidOK  %+v", 200, o.Payload)
}

func (o *GetProcessorUnitsMoidOK) GetPayload() *models.ProcessorUnit {
	return o.Payload
}

func (o *GetProcessorUnitsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorUnitsMoidNotFound creates a GetProcessorUnitsMoidNotFound with default headers values
func NewGetProcessorUnitsMoidNotFound() *GetProcessorUnitsMoidNotFound {
	return &GetProcessorUnitsMoidNotFound{}
}

/*GetProcessorUnitsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetProcessorUnitsMoidNotFound struct {
}

func (o *GetProcessorUnitsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /processor/Units/{moid}][%d] getProcessorUnitsMoidNotFound ", 404)
}

func (o *GetProcessorUnitsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorUnitsMoidDefault creates a GetProcessorUnitsMoidDefault with default headers values
func NewGetProcessorUnitsMoidDefault(code int) *GetProcessorUnitsMoidDefault {
	return &GetProcessorUnitsMoidDefault{
		_statusCode: code,
	}
}

/*GetProcessorUnitsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetProcessorUnitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get processor units moid default response
func (o *GetProcessorUnitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetProcessorUnitsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /processor/Units/{moid}][%d] GetProcessorUnitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetProcessorUnitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProcessorUnitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
