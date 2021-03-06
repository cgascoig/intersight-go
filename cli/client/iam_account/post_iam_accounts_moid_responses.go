// Code generated by go-swagger; DO NOT EDIT.

package iam_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamAccountsMoidReader is a Reader for the PostIamAccountsMoid structure.
type PostIamAccountsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamAccountsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamAccountsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamAccountsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamAccountsMoidCreated creates a PostIamAccountsMoidCreated with default headers values
func NewPostIamAccountsMoidCreated() *PostIamAccountsMoidCreated {
	return &PostIamAccountsMoidCreated{}
}

/*PostIamAccountsMoidCreated handles this case with default header values.

Null response
*/
type PostIamAccountsMoidCreated struct {
}

func (o *PostIamAccountsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/Accounts/{moid}][%d] postIamAccountsMoidCreated ", 201)
}

func (o *PostIamAccountsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamAccountsMoidDefault creates a PostIamAccountsMoidDefault with default headers values
func NewPostIamAccountsMoidDefault(code int) *PostIamAccountsMoidDefault {
	return &PostIamAccountsMoidDefault{
		_statusCode: code,
	}
}

/*PostIamAccountsMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamAccountsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam accounts moid default response
func (o *PostIamAccountsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamAccountsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/Accounts/{moid}][%d] PostIamAccountsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamAccountsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamAccountsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
