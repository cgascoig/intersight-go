// Code generated by go-swagger; DO NOT EDIT.

package compute_board

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetComputeBoardsMoidReader is a Reader for the GetComputeBoardsMoid structure.
type GetComputeBoardsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeBoardsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeBoardsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetComputeBoardsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetComputeBoardsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetComputeBoardsMoidOK creates a GetComputeBoardsMoidOK with default headers values
func NewGetComputeBoardsMoidOK() *GetComputeBoardsMoidOK {
	return &GetComputeBoardsMoidOK{}
}

/*GetComputeBoardsMoidOK handles this case with default header values.

An instance of computeBoard
*/
type GetComputeBoardsMoidOK struct {
	Payload *models.ComputeBoard
}

func (o *GetComputeBoardsMoidOK) Error() string {
	return fmt.Sprintf("[GET /compute/Boards/{moid}][%d] getComputeBoardsMoidOK  %+v", 200, o.Payload)
}

func (o *GetComputeBoardsMoidOK) GetPayload() *models.ComputeBoard {
	return o.Payload
}

func (o *GetComputeBoardsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeBoard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeBoardsMoidNotFound creates a GetComputeBoardsMoidNotFound with default headers values
func NewGetComputeBoardsMoidNotFound() *GetComputeBoardsMoidNotFound {
	return &GetComputeBoardsMoidNotFound{}
}

/*GetComputeBoardsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetComputeBoardsMoidNotFound struct {
}

func (o *GetComputeBoardsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /compute/Boards/{moid}][%d] getComputeBoardsMoidNotFound ", 404)
}

func (o *GetComputeBoardsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComputeBoardsMoidDefault creates a GetComputeBoardsMoidDefault with default headers values
func NewGetComputeBoardsMoidDefault(code int) *GetComputeBoardsMoidDefault {
	return &GetComputeBoardsMoidDefault{
		_statusCode: code,
	}
}

/*GetComputeBoardsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetComputeBoardsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get compute boards moid default response
func (o *GetComputeBoardsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetComputeBoardsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /compute/Boards/{moid}][%d] GetComputeBoardsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetComputeBoardsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetComputeBoardsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
