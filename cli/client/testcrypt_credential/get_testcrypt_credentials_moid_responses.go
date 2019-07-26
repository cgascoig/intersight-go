// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetTestcryptCredentialsMoidReader is a Reader for the GetTestcryptCredentialsMoid structure.
type GetTestcryptCredentialsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestcryptCredentialsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestcryptCredentialsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTestcryptCredentialsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTestcryptCredentialsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTestcryptCredentialsMoidOK creates a GetTestcryptCredentialsMoidOK with default headers values
func NewGetTestcryptCredentialsMoidOK() *GetTestcryptCredentialsMoidOK {
	return &GetTestcryptCredentialsMoidOK{}
}

/*GetTestcryptCredentialsMoidOK handles this case with default header values.

An instance of testcryptCredential
*/
type GetTestcryptCredentialsMoidOK struct {
	Payload *models.TestcryptCredential
}

func (o *GetTestcryptCredentialsMoidOK) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Credentials/{moid}][%d] getTestcryptCredentialsMoidOK  %+v", 200, o.Payload)
}

func (o *GetTestcryptCredentialsMoidOK) GetPayload() *models.TestcryptCredential {
	return o.Payload
}

func (o *GetTestcryptCredentialsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestcryptCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestcryptCredentialsMoidNotFound creates a GetTestcryptCredentialsMoidNotFound with default headers values
func NewGetTestcryptCredentialsMoidNotFound() *GetTestcryptCredentialsMoidNotFound {
	return &GetTestcryptCredentialsMoidNotFound{}
}

/*GetTestcryptCredentialsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetTestcryptCredentialsMoidNotFound struct {
}

func (o *GetTestcryptCredentialsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Credentials/{moid}][%d] getTestcryptCredentialsMoidNotFound ", 404)
}

func (o *GetTestcryptCredentialsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTestcryptCredentialsMoidDefault creates a GetTestcryptCredentialsMoidDefault with default headers values
func NewGetTestcryptCredentialsMoidDefault(code int) *GetTestcryptCredentialsMoidDefault {
	return &GetTestcryptCredentialsMoidDefault{
		_statusCode: code,
	}
}

/*GetTestcryptCredentialsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetTestcryptCredentialsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get testcrypt credentials moid default response
func (o *GetTestcryptCredentialsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetTestcryptCredentialsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Credentials/{moid}][%d] GetTestcryptCredentialsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetTestcryptCredentialsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTestcryptCredentialsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
