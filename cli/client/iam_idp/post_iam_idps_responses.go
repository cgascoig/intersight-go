// Code generated by go-swagger; DO NOT EDIT.

package iam_idp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamIdpsReader is a Reader for the PostIamIdps structure.
type PostIamIdpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamIdpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamIdpsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamIdpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamIdpsCreated creates a PostIamIdpsCreated with default headers values
func NewPostIamIdpsCreated() *PostIamIdpsCreated {
	return &PostIamIdpsCreated{}
}

/*PostIamIdpsCreated handles this case with default header values.

Null response
*/
type PostIamIdpsCreated struct {
}

func (o *PostIamIdpsCreated) Error() string {
	return fmt.Sprintf("[POST /iam/Idps][%d] postIamIdpsCreated ", 201)
}

func (o *PostIamIdpsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamIdpsDefault creates a PostIamIdpsDefault with default headers values
func NewPostIamIdpsDefault(code int) *PostIamIdpsDefault {
	return &PostIamIdpsDefault{
		_statusCode: code,
	}
}

/*PostIamIdpsDefault handles this case with default header values.

unexpected error
*/
type PostIamIdpsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam idps default response
func (o *PostIamIdpsDefault) Code() int {
	return o._statusCode
}

func (o *PostIamIdpsDefault) Error() string {
	return fmt.Sprintf("[POST /iam/Idps][%d] PostIamIdps default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamIdpsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamIdpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
