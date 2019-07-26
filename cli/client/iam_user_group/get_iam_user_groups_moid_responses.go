// Code generated by go-swagger; DO NOT EDIT.

package iam_user_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamUserGroupsMoidReader is a Reader for the GetIamUserGroupsMoid structure.
type GetIamUserGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamUserGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamUserGroupsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIamUserGroupsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIamUserGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamUserGroupsMoidOK creates a GetIamUserGroupsMoidOK with default headers values
func NewGetIamUserGroupsMoidOK() *GetIamUserGroupsMoidOK {
	return &GetIamUserGroupsMoidOK{}
}

/*GetIamUserGroupsMoidOK handles this case with default header values.

An instance of iamUserGroup
*/
type GetIamUserGroupsMoidOK struct {
	Payload *models.IamUserGroup
}

func (o *GetIamUserGroupsMoidOK) Error() string {
	return fmt.Sprintf("[GET /iam/UserGroups/{moid}][%d] getIamUserGroupsMoidOK  %+v", 200, o.Payload)
}

func (o *GetIamUserGroupsMoidOK) GetPayload() *models.IamUserGroup {
	return o.Payload
}

func (o *GetIamUserGroupsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamUserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamUserGroupsMoidNotFound creates a GetIamUserGroupsMoidNotFound with default headers values
func NewGetIamUserGroupsMoidNotFound() *GetIamUserGroupsMoidNotFound {
	return &GetIamUserGroupsMoidNotFound{}
}

/*GetIamUserGroupsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIamUserGroupsMoidNotFound struct {
}

func (o *GetIamUserGroupsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/UserGroups/{moid}][%d] getIamUserGroupsMoidNotFound ", 404)
}

func (o *GetIamUserGroupsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIamUserGroupsMoidDefault creates a GetIamUserGroupsMoidDefault with default headers values
func NewGetIamUserGroupsMoidDefault(code int) *GetIamUserGroupsMoidDefault {
	return &GetIamUserGroupsMoidDefault{
		_statusCode: code,
	}
}

/*GetIamUserGroupsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIamUserGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam user groups moid default response
func (o *GetIamUserGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIamUserGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iam/UserGroups/{moid}][%d] GetIamUserGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamUserGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamUserGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
