// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_administrator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetTestcryptAdministratorsReader is a Reader for the GetTestcryptAdministrators structure.
type GetTestcryptAdministratorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestcryptAdministratorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestcryptAdministratorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTestcryptAdministratorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTestcryptAdministratorsOK creates a GetTestcryptAdministratorsOK with default headers values
func NewGetTestcryptAdministratorsOK() *GetTestcryptAdministratorsOK {
	return &GetTestcryptAdministratorsOK{}
}

/*GetTestcryptAdministratorsOK handles this case with default header values.

List of testcryptAdministrators for the given filter criteria
*/
type GetTestcryptAdministratorsOK struct {
	Payload *models.TestcryptAdministratorList
}

func (o *GetTestcryptAdministratorsOK) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Administrators][%d] getTestcryptAdministratorsOK  %+v", 200, o.Payload)
}

func (o *GetTestcryptAdministratorsOK) GetPayload() *models.TestcryptAdministratorList {
	return o.Payload
}

func (o *GetTestcryptAdministratorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestcryptAdministratorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestcryptAdministratorsDefault creates a GetTestcryptAdministratorsDefault with default headers values
func NewGetTestcryptAdministratorsDefault(code int) *GetTestcryptAdministratorsDefault {
	return &GetTestcryptAdministratorsDefault{
		_statusCode: code,
	}
}

/*GetTestcryptAdministratorsDefault handles this case with default header values.

Unexpected error
*/
type GetTestcryptAdministratorsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get testcrypt administrators default response
func (o *GetTestcryptAdministratorsDefault) Code() int {
	return o._statusCode
}

func (o *GetTestcryptAdministratorsDefault) Error() string {
	return fmt.Sprintf("[GET /testcrypt/Administrators][%d] GetTestcryptAdministrators default  %+v", o._statusCode, o.Payload)
}

func (o *GetTestcryptAdministratorsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTestcryptAdministratorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
