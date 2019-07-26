// Code generated by go-swagger; DO NOT EDIT.

package smtp_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostSMTPPoliciesMoidReader is a Reader for the PostSMTPPoliciesMoid structure.
type PostSMTPPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSMTPPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSMTPPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostSMTPPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSMTPPoliciesMoidCreated creates a PostSMTPPoliciesMoidCreated with default headers values
func NewPostSMTPPoliciesMoidCreated() *PostSMTPPoliciesMoidCreated {
	return &PostSMTPPoliciesMoidCreated{}
}

/*PostSMTPPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostSMTPPoliciesMoidCreated struct {
}

func (o *PostSMTPPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /smtp/Policies/{moid}][%d] postSmtpPoliciesMoidCreated ", 201)
}

func (o *PostSMTPPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSMTPPoliciesMoidDefault creates a PostSMTPPoliciesMoidDefault with default headers values
func NewPostSMTPPoliciesMoidDefault(code int) *PostSMTPPoliciesMoidDefault {
	return &PostSMTPPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostSMTPPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostSMTPPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post SMTP policies moid default response
func (o *PostSMTPPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostSMTPPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /smtp/Policies/{moid}][%d] PostSMTPPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostSMTPPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSMTPPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
