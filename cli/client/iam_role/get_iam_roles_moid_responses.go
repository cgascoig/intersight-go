// Code generated by go-swagger; DO NOT EDIT.

package iam_role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamRolesMoidReader is a Reader for the GetIamRolesMoid structure.
type GetIamRolesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamRolesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamRolesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIamRolesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIamRolesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamRolesMoidOK creates a GetIamRolesMoidOK with default headers values
func NewGetIamRolesMoidOK() *GetIamRolesMoidOK {
	return &GetIamRolesMoidOK{}
}

/*GetIamRolesMoidOK handles this case with default header values.

An instance of iamRole
*/
type GetIamRolesMoidOK struct {
	Payload *models.IamRole
}

func (o *GetIamRolesMoidOK) Error() string {
	return fmt.Sprintf("[GET /iam/Roles/{moid}][%d] getIamRolesMoidOK  %+v", 200, o.Payload)
}

func (o *GetIamRolesMoidOK) GetPayload() *models.IamRole {
	return o.Payload
}

func (o *GetIamRolesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamRolesMoidNotFound creates a GetIamRolesMoidNotFound with default headers values
func NewGetIamRolesMoidNotFound() *GetIamRolesMoidNotFound {
	return &GetIamRolesMoidNotFound{}
}

/*GetIamRolesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIamRolesMoidNotFound struct {
}

func (o *GetIamRolesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/Roles/{moid}][%d] getIamRolesMoidNotFound ", 404)
}

func (o *GetIamRolesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIamRolesMoidDefault creates a GetIamRolesMoidDefault with default headers values
func NewGetIamRolesMoidDefault(code int) *GetIamRolesMoidDefault {
	return &GetIamRolesMoidDefault{
		_statusCode: code,
	}
}

/*GetIamRolesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIamRolesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam roles moid default response
func (o *GetIamRolesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIamRolesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iam/Roles/{moid}][%d] GetIamRolesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamRolesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamRolesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}