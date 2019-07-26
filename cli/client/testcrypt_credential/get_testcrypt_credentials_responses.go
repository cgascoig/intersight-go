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

// GetTestcryptCredentialsReader is a Reader for the GetTestcryptCredentials structure.
type GetTestcryptCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestcryptCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestcryptCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTestcryptCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTestcryptCredentialsOK creates a GetTestcryptCredentialsOK with default headers values
func NewGetTestcryptCredentialsOK() *GetTestcryptCredentialsOK {
	return &GetTestcryptCredentialsOK{}
}

/*GetTestcryptCredentialsOK handles this case with default header values.

List of testcryptCredentials for the given filter criteria
*/
type GetTestcryptCredentialsOK struct {
	Payload *models.TestcryptCredentialList
}

func (o *GetTestcryptCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Credentials][%d] getTestcryptCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetTestcryptCredentialsOK) GetPayload() *models.TestcryptCredentialList {
	return o.Payload
}

func (o *GetTestcryptCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestcryptCredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestcryptCredentialsDefault creates a GetTestcryptCredentialsDefault with default headers values
func NewGetTestcryptCredentialsDefault(code int) *GetTestcryptCredentialsDefault {
	return &GetTestcryptCredentialsDefault{
		_statusCode: code,
	}
}

/*GetTestcryptCredentialsDefault handles this case with default header values.

Unexpected error
*/
type GetTestcryptCredentialsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get testcrypt credentials default response
func (o *GetTestcryptCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *GetTestcryptCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Credentials][%d] GetTestcryptCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *GetTestcryptCredentialsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTestcryptCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
