// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_token_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostTestcryptTokenApisReader is a Reader for the PostTestcryptTokenApis structure.
type PostTestcryptTokenApisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTestcryptTokenApisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostTestcryptTokenApisCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostTestcryptTokenApisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTestcryptTokenApisCreated creates a PostTestcryptTokenApisCreated with default headers values
func NewPostTestcryptTokenApisCreated() *PostTestcryptTokenApisCreated {
	return &PostTestcryptTokenApisCreated{}
}

/*PostTestcryptTokenApisCreated handles this case with default header values.

Null response
*/
type PostTestcryptTokenApisCreated struct {
}

func (o *PostTestcryptTokenApisCreated) Error() string {
	return fmt.Sprintf("[POST /testcrypt/TokenApis][%d] postTestcryptTokenApisCreated ", 201)
}

func (o *PostTestcryptTokenApisCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTestcryptTokenApisDefault creates a PostTestcryptTokenApisDefault with default headers values
func NewPostTestcryptTokenApisDefault(code int) *PostTestcryptTokenApisDefault {
	return &PostTestcryptTokenApisDefault{
		_statusCode: code,
	}
}

/*PostTestcryptTokenApisDefault handles this case with default header values.

unexpected error
*/
type PostTestcryptTokenApisDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post testcrypt token apis default response
func (o *PostTestcryptTokenApisDefault) Code() int {
	return o._statusCode
}

func (o *PostTestcryptTokenApisDefault) Error() string {
	return fmt.Sprintf("[POST /testcrypt/TokenApis][%d] PostTestcryptTokenApis default  %+v", o._statusCode, o.Payload)
}

func (o *PostTestcryptTokenApisDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTestcryptTokenApisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}