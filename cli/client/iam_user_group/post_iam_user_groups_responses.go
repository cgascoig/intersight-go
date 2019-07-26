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

// PostIamUserGroupsReader is a Reader for the PostIamUserGroups structure.
type PostIamUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamUserGroupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamUserGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamUserGroupsCreated creates a PostIamUserGroupsCreated with default headers values
func NewPostIamUserGroupsCreated() *PostIamUserGroupsCreated {
	return &PostIamUserGroupsCreated{}
}

/*PostIamUserGroupsCreated handles this case with default header values.

Null response
*/
type PostIamUserGroupsCreated struct {
}

func (o *PostIamUserGroupsCreated) Error() string {
	return fmt.Sprintf("[POST /iam/UserGroups][%d] postIamUserGroupsCreated ", 201)
}

func (o *PostIamUserGroupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamUserGroupsDefault creates a PostIamUserGroupsDefault with default headers values
func NewPostIamUserGroupsDefault(code int) *PostIamUserGroupsDefault {
	return &PostIamUserGroupsDefault{
		_statusCode: code,
	}
}

/*PostIamUserGroupsDefault handles this case with default header values.

unexpected error
*/
type PostIamUserGroupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam user groups default response
func (o *PostIamUserGroupsDefault) Code() int {
	return o._statusCode
}

func (o *PostIamUserGroupsDefault) Error() string {
	return fmt.Sprintf("[POST /iam/UserGroups][%d] PostIamUserGroups default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamUserGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamUserGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}