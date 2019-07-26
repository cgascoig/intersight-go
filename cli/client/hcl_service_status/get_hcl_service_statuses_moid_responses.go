// Code generated by go-swagger; DO NOT EDIT.

package hcl_service_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHclServiceStatusesMoidReader is a Reader for the GetHclServiceStatusesMoid structure.
type GetHclServiceStatusesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHclServiceStatusesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHclServiceStatusesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHclServiceStatusesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHclServiceStatusesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHclServiceStatusesMoidOK creates a GetHclServiceStatusesMoidOK with default headers values
func NewGetHclServiceStatusesMoidOK() *GetHclServiceStatusesMoidOK {
	return &GetHclServiceStatusesMoidOK{}
}

/*GetHclServiceStatusesMoidOK handles this case with default header values.

An instance of hclServiceStatus
*/
type GetHclServiceStatusesMoidOK struct {
	Payload *models.HclServiceStatus
}

func (o *GetHclServiceStatusesMoidOK) Error() string {
	return fmt.Sprintf("[GET /hcl/ServiceStatuses/{moid}][%d] getHclServiceStatusesMoidOK  %+v", 200, o.Payload)
}

func (o *GetHclServiceStatusesMoidOK) GetPayload() *models.HclServiceStatus {
	return o.Payload
}

func (o *GetHclServiceStatusesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HclServiceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHclServiceStatusesMoidNotFound creates a GetHclServiceStatusesMoidNotFound with default headers values
func NewGetHclServiceStatusesMoidNotFound() *GetHclServiceStatusesMoidNotFound {
	return &GetHclServiceStatusesMoidNotFound{}
}

/*GetHclServiceStatusesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHclServiceStatusesMoidNotFound struct {
}

func (o *GetHclServiceStatusesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hcl/ServiceStatuses/{moid}][%d] getHclServiceStatusesMoidNotFound ", 404)
}

func (o *GetHclServiceStatusesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHclServiceStatusesMoidDefault creates a GetHclServiceStatusesMoidDefault with default headers values
func NewGetHclServiceStatusesMoidDefault(code int) *GetHclServiceStatusesMoidDefault {
	return &GetHclServiceStatusesMoidDefault{
		_statusCode: code,
	}
}

/*GetHclServiceStatusesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHclServiceStatusesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hcl service statuses moid default response
func (o *GetHclServiceStatusesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHclServiceStatusesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hcl/ServiceStatuses/{moid}][%d] GetHclServiceStatusesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHclServiceStatusesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHclServiceStatusesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}