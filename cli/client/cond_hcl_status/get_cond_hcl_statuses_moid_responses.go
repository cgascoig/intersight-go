// Code generated by go-swagger; DO NOT EDIT.

package cond_hcl_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetCondHclStatusesMoidReader is a Reader for the GetCondHclStatusesMoid structure.
type GetCondHclStatusesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCondHclStatusesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCondHclStatusesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCondHclStatusesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCondHclStatusesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCondHclStatusesMoidOK creates a GetCondHclStatusesMoidOK with default headers values
func NewGetCondHclStatusesMoidOK() *GetCondHclStatusesMoidOK {
	return &GetCondHclStatusesMoidOK{}
}

/*GetCondHclStatusesMoidOK handles this case with default header values.

An instance of condHclStatus
*/
type GetCondHclStatusesMoidOK struct {
	Payload *models.CondHclStatus
}

func (o *GetCondHclStatusesMoidOK) Error() string {
	return fmt.Sprintf("[GET /cond/HclStatuses/{moid}][%d] getCondHclStatusesMoidOK  %+v", 200, o.Payload)
}

func (o *GetCondHclStatusesMoidOK) GetPayload() *models.CondHclStatus {
	return o.Payload
}

func (o *GetCondHclStatusesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CondHclStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCondHclStatusesMoidNotFound creates a GetCondHclStatusesMoidNotFound with default headers values
func NewGetCondHclStatusesMoidNotFound() *GetCondHclStatusesMoidNotFound {
	return &GetCondHclStatusesMoidNotFound{}
}

/*GetCondHclStatusesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetCondHclStatusesMoidNotFound struct {
}

func (o *GetCondHclStatusesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /cond/HclStatuses/{moid}][%d] getCondHclStatusesMoidNotFound ", 404)
}

func (o *GetCondHclStatusesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCondHclStatusesMoidDefault creates a GetCondHclStatusesMoidDefault with default headers values
func NewGetCondHclStatusesMoidDefault(code int) *GetCondHclStatusesMoidDefault {
	return &GetCondHclStatusesMoidDefault{
		_statusCode: code,
	}
}

/*GetCondHclStatusesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetCondHclStatusesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cond hcl statuses moid default response
func (o *GetCondHclStatusesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetCondHclStatusesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /cond/HclStatuses/{moid}][%d] GetCondHclStatusesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCondHclStatusesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCondHclStatusesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}