// Code generated by go-swagger; DO NOT EDIT.

package iam_local_user_password

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamLocalUserPasswordsMoidReader is a Reader for the PostIamLocalUserPasswordsMoid structure.
type PostIamLocalUserPasswordsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamLocalUserPasswordsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamLocalUserPasswordsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamLocalUserPasswordsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamLocalUserPasswordsMoidCreated creates a PostIamLocalUserPasswordsMoidCreated with default headers values
func NewPostIamLocalUserPasswordsMoidCreated() *PostIamLocalUserPasswordsMoidCreated {
	return &PostIamLocalUserPasswordsMoidCreated{}
}

/*PostIamLocalUserPasswordsMoidCreated handles this case with default header values.

Null response
*/
type PostIamLocalUserPasswordsMoidCreated struct {
}

func (o *PostIamLocalUserPasswordsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/LocalUserPasswords/{moid}][%d] postIamLocalUserPasswordsMoidCreated ", 201)
}

func (o *PostIamLocalUserPasswordsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamLocalUserPasswordsMoidDefault creates a PostIamLocalUserPasswordsMoidDefault with default headers values
func NewPostIamLocalUserPasswordsMoidDefault(code int) *PostIamLocalUserPasswordsMoidDefault {
	return &PostIamLocalUserPasswordsMoidDefault{
		_statusCode: code,
	}
}

/*PostIamLocalUserPasswordsMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamLocalUserPasswordsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam local user passwords moid default response
func (o *PostIamLocalUserPasswordsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamLocalUserPasswordsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/LocalUserPasswords/{moid}][%d] PostIamLocalUserPasswordsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamLocalUserPasswordsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamLocalUserPasswordsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}