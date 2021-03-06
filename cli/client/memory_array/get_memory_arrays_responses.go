// Code generated by go-swagger; DO NOT EDIT.

package memory_array

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetMemoryArraysReader is a Reader for the GetMemoryArrays structure.
type GetMemoryArraysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMemoryArraysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMemoryArraysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMemoryArraysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMemoryArraysOK creates a GetMemoryArraysOK with default headers values
func NewGetMemoryArraysOK() *GetMemoryArraysOK {
	return &GetMemoryArraysOK{}
}

/*GetMemoryArraysOK handles this case with default header values.

List of memoryArrays for the given filter criteria
*/
type GetMemoryArraysOK struct {
	Payload *models.MemoryArrayList
}

func (o *GetMemoryArraysOK) Error() string {
	return fmt.Sprintf("[GET /memory/Arrays][%d] getMemoryArraysOK  %+v", 200, o.Payload)
}

func (o *GetMemoryArraysOK) GetPayload() *models.MemoryArrayList {
	return o.Payload
}

func (o *GetMemoryArraysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemoryArrayList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemoryArraysDefault creates a GetMemoryArraysDefault with default headers values
func NewGetMemoryArraysDefault(code int) *GetMemoryArraysDefault {
	return &GetMemoryArraysDefault{
		_statusCode: code,
	}
}

/*GetMemoryArraysDefault handles this case with default header values.

Unexpected error
*/
type GetMemoryArraysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get memory arrays default response
func (o *GetMemoryArraysDefault) Code() int {
	return o._statusCode
}

func (o *GetMemoryArraysDefault) Error() string {
	return fmt.Sprintf("[GET /memory/Arrays][%d] GetMemoryArrays default  %+v", o._statusCode, o.Payload)
}

func (o *GetMemoryArraysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMemoryArraysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
