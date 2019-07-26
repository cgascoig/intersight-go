// Code generated by go-swagger; DO NOT EDIT.

package iam_resource_limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamResourceLimitsMoidReader is a Reader for the GetIamResourceLimitsMoid structure.
type GetIamResourceLimitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamResourceLimitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamResourceLimitsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIamResourceLimitsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIamResourceLimitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamResourceLimitsMoidOK creates a GetIamResourceLimitsMoidOK with default headers values
func NewGetIamResourceLimitsMoidOK() *GetIamResourceLimitsMoidOK {
	return &GetIamResourceLimitsMoidOK{}
}

/*GetIamResourceLimitsMoidOK handles this case with default header values.

An instance of iamResourceLimits
*/
type GetIamResourceLimitsMoidOK struct {
	Payload *models.IamResourceLimits
}

func (o *GetIamResourceLimitsMoidOK) Error() string {
	return fmt.Sprintf("[GET /iam/ResourceLimits/{moid}][%d] getIamResourceLimitsMoidOK  %+v", 200, o.Payload)
}

func (o *GetIamResourceLimitsMoidOK) GetPayload() *models.IamResourceLimits {
	return o.Payload
}

func (o *GetIamResourceLimitsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamResourceLimits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamResourceLimitsMoidNotFound creates a GetIamResourceLimitsMoidNotFound with default headers values
func NewGetIamResourceLimitsMoidNotFound() *GetIamResourceLimitsMoidNotFound {
	return &GetIamResourceLimitsMoidNotFound{}
}

/*GetIamResourceLimitsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIamResourceLimitsMoidNotFound struct {
}

func (o *GetIamResourceLimitsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/ResourceLimits/{moid}][%d] getIamResourceLimitsMoidNotFound ", 404)
}

func (o *GetIamResourceLimitsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIamResourceLimitsMoidDefault creates a GetIamResourceLimitsMoidDefault with default headers values
func NewGetIamResourceLimitsMoidDefault(code int) *GetIamResourceLimitsMoidDefault {
	return &GetIamResourceLimitsMoidDefault{
		_statusCode: code,
	}
}

/*GetIamResourceLimitsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIamResourceLimitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam resource limits moid default response
func (o *GetIamResourceLimitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIamResourceLimitsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iam/ResourceLimits/{moid}][%d] GetIamResourceLimitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamResourceLimitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamResourceLimitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
