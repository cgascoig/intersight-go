// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_user_role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamEndPointUserRolesReader is a Reader for the PostIamEndPointUserRoles structure.
type PostIamEndPointUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamEndPointUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamEndPointUserRolesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamEndPointUserRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamEndPointUserRolesCreated creates a PostIamEndPointUserRolesCreated with default headers values
func NewPostIamEndPointUserRolesCreated() *PostIamEndPointUserRolesCreated {
	return &PostIamEndPointUserRolesCreated{}
}

/*PostIamEndPointUserRolesCreated handles this case with default header values.

Null response
*/
type PostIamEndPointUserRolesCreated struct {
}

func (o *PostIamEndPointUserRolesCreated) Error() string {
	return fmt.Sprintf("[POST /iam/EndPointUserRoles][%d] postIamEndPointUserRolesCreated ", 201)
}

func (o *PostIamEndPointUserRolesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamEndPointUserRolesDefault creates a PostIamEndPointUserRolesDefault with default headers values
func NewPostIamEndPointUserRolesDefault(code int) *PostIamEndPointUserRolesDefault {
	return &PostIamEndPointUserRolesDefault{
		_statusCode: code,
	}
}

/*PostIamEndPointUserRolesDefault handles this case with default header values.

unexpected error
*/
type PostIamEndPointUserRolesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam end point user roles default response
func (o *PostIamEndPointUserRolesDefault) Code() int {
	return o._statusCode
}

func (o *PostIamEndPointUserRolesDefault) Error() string {
	return fmt.Sprintf("[POST /iam/EndPointUserRoles][%d] PostIamEndPointUserRoles default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamEndPointUserRolesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamEndPointUserRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
