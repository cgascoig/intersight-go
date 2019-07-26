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

// GetTestcryptTokenApisReader is a Reader for the GetTestcryptTokenApis structure.
type GetTestcryptTokenApisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestcryptTokenApisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestcryptTokenApisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTestcryptTokenApisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTestcryptTokenApisOK creates a GetTestcryptTokenApisOK with default headers values
func NewGetTestcryptTokenApisOK() *GetTestcryptTokenApisOK {
	return &GetTestcryptTokenApisOK{}
}

/*GetTestcryptTokenApisOK handles this case with default header values.

List of testcryptTokenApis for the given filter criteria
*/
type GetTestcryptTokenApisOK struct {
	Payload *models.TestcryptTokenAPIList
}

func (o *GetTestcryptTokenApisOK) Error() string {
	return fmt.Sprintf("[GET /testcrypt/TokenApis][%d] getTestcryptTokenApisOK  %+v", 200, o.Payload)
}

func (o *GetTestcryptTokenApisOK) GetPayload() *models.TestcryptTokenAPIList {
	return o.Payload
}

func (o *GetTestcryptTokenApisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestcryptTokenAPIList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestcryptTokenApisDefault creates a GetTestcryptTokenApisDefault with default headers values
func NewGetTestcryptTokenApisDefault(code int) *GetTestcryptTokenApisDefault {
	return &GetTestcryptTokenApisDefault{
		_statusCode: code,
	}
}

/*GetTestcryptTokenApisDefault handles this case with default header values.

Unexpected error
*/
type GetTestcryptTokenApisDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get testcrypt token apis default response
func (o *GetTestcryptTokenApisDefault) Code() int {
	return o._statusCode
}

func (o *GetTestcryptTokenApisDefault) Error() string {
	return fmt.Sprintf("[GET /testcrypt/TokenApis][%d] GetTestcryptTokenApis default  %+v", o._statusCode, o.Payload)
}

func (o *GetTestcryptTokenApisDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTestcryptTokenApisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
