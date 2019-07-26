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

// PostIamUserGroupsMoidReader is a Reader for the PostIamUserGroupsMoid structure.
type PostIamUserGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamUserGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamUserGroupsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamUserGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamUserGroupsMoidCreated creates a PostIamUserGroupsMoidCreated with default headers values
func NewPostIamUserGroupsMoidCreated() *PostIamUserGroupsMoidCreated {
	return &PostIamUserGroupsMoidCreated{}
}

/*PostIamUserGroupsMoidCreated handles this case with default header values.

Null response
*/
type PostIamUserGroupsMoidCreated struct {
}

func (o *PostIamUserGroupsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/UserGroups/{moid}][%d] postIamUserGroupsMoidCreated ", 201)
}

func (o *PostIamUserGroupsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamUserGroupsMoidDefault creates a PostIamUserGroupsMoidDefault with default headers values
func NewPostIamUserGroupsMoidDefault(code int) *PostIamUserGroupsMoidDefault {
	return &PostIamUserGroupsMoidDefault{
		_statusCode: code,
	}
}

/*PostIamUserGroupsMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamUserGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam user groups moid default response
func (o *PostIamUserGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamUserGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/UserGroups/{moid}][%d] PostIamUserGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamUserGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamUserGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}