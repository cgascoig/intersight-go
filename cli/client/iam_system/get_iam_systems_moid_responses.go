// Code generated by go-swagger; DO NOT EDIT.

package iam_system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamSystemsMoidReader is a Reader for the GetIamSystemsMoid structure.
type GetIamSystemsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamSystemsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamSystemsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIamSystemsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIamSystemsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamSystemsMoidOK creates a GetIamSystemsMoidOK with default headers values
func NewGetIamSystemsMoidOK() *GetIamSystemsMoidOK {
	return &GetIamSystemsMoidOK{}
}

/*GetIamSystemsMoidOK handles this case with default header values.

An instance of iamSystem
*/
type GetIamSystemsMoidOK struct {
	Payload *models.IamSystem
}

func (o *GetIamSystemsMoidOK) Error() string {
	return fmt.Sprintf("[GET /iam/Systems/{moid}][%d] getIamSystemsMoidOK  %+v", 200, o.Payload)
}

func (o *GetIamSystemsMoidOK) GetPayload() *models.IamSystem {
	return o.Payload
}

func (o *GetIamSystemsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamSystem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamSystemsMoidNotFound creates a GetIamSystemsMoidNotFound with default headers values
func NewGetIamSystemsMoidNotFound() *GetIamSystemsMoidNotFound {
	return &GetIamSystemsMoidNotFound{}
}

/*GetIamSystemsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIamSystemsMoidNotFound struct {
}

func (o *GetIamSystemsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/Systems/{moid}][%d] getIamSystemsMoidNotFound ", 404)
}

func (o *GetIamSystemsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIamSystemsMoidDefault creates a GetIamSystemsMoidDefault with default headers values
func NewGetIamSystemsMoidDefault(code int) *GetIamSystemsMoidDefault {
	return &GetIamSystemsMoidDefault{
		_statusCode: code,
	}
}

/*GetIamSystemsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIamSystemsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam systems moid default response
func (o *GetIamSystemsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIamSystemsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iam/Systems/{moid}][%d] GetIamSystemsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamSystemsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamSystemsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
