// Code generated by go-swagger; DO NOT EDIT.

package tam_security_advisory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetTamSecurityAdvisoriesMoidReader is a Reader for the GetTamSecurityAdvisoriesMoid structure.
type GetTamSecurityAdvisoriesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTamSecurityAdvisoriesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTamSecurityAdvisoriesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTamSecurityAdvisoriesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTamSecurityAdvisoriesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTamSecurityAdvisoriesMoidOK creates a GetTamSecurityAdvisoriesMoidOK with default headers values
func NewGetTamSecurityAdvisoriesMoidOK() *GetTamSecurityAdvisoriesMoidOK {
	return &GetTamSecurityAdvisoriesMoidOK{}
}

/*GetTamSecurityAdvisoriesMoidOK handles this case with default header values.

An instance of tamSecurityAdvisory
*/
type GetTamSecurityAdvisoriesMoidOK struct {
	Payload *models.TamSecurityAdvisory
}

func (o *GetTamSecurityAdvisoriesMoidOK) Error() string {
	return fmt.Sprintf("[GET /tam/SecurityAdvisories/{moid}][%d] getTamSecurityAdvisoriesMoidOK  %+v", 200, o.Payload)
}

func (o *GetTamSecurityAdvisoriesMoidOK) GetPayload() *models.TamSecurityAdvisory {
	return o.Payload
}

func (o *GetTamSecurityAdvisoriesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TamSecurityAdvisory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTamSecurityAdvisoriesMoidNotFound creates a GetTamSecurityAdvisoriesMoidNotFound with default headers values
func NewGetTamSecurityAdvisoriesMoidNotFound() *GetTamSecurityAdvisoriesMoidNotFound {
	return &GetTamSecurityAdvisoriesMoidNotFound{}
}

/*GetTamSecurityAdvisoriesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetTamSecurityAdvisoriesMoidNotFound struct {
}

func (o *GetTamSecurityAdvisoriesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /tam/SecurityAdvisories/{moid}][%d] getTamSecurityAdvisoriesMoidNotFound ", 404)
}

func (o *GetTamSecurityAdvisoriesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTamSecurityAdvisoriesMoidDefault creates a GetTamSecurityAdvisoriesMoidDefault with default headers values
func NewGetTamSecurityAdvisoriesMoidDefault(code int) *GetTamSecurityAdvisoriesMoidDefault {
	return &GetTamSecurityAdvisoriesMoidDefault{
		_statusCode: code,
	}
}

/*GetTamSecurityAdvisoriesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetTamSecurityAdvisoriesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tam security advisories moid default response
func (o *GetTamSecurityAdvisoriesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetTamSecurityAdvisoriesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /tam/SecurityAdvisories/{moid}][%d] GetTamSecurityAdvisoriesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetTamSecurityAdvisoriesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTamSecurityAdvisoriesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
