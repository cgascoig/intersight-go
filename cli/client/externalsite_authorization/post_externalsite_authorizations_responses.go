// Code generated by go-swagger; DO NOT EDIT.

package externalsite_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostExternalsiteAuthorizationsReader is a Reader for the PostExternalsiteAuthorizations structure.
type PostExternalsiteAuthorizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalsiteAuthorizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostExternalsiteAuthorizationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostExternalsiteAuthorizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostExternalsiteAuthorizationsCreated creates a PostExternalsiteAuthorizationsCreated with default headers values
func NewPostExternalsiteAuthorizationsCreated() *PostExternalsiteAuthorizationsCreated {
	return &PostExternalsiteAuthorizationsCreated{}
}

/*PostExternalsiteAuthorizationsCreated handles this case with default header values.

Null response
*/
type PostExternalsiteAuthorizationsCreated struct {
}

func (o *PostExternalsiteAuthorizationsCreated) Error() string {
	return fmt.Sprintf("[POST /externalsite/Authorizations][%d] postExternalsiteAuthorizationsCreated ", 201)
}

func (o *PostExternalsiteAuthorizationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostExternalsiteAuthorizationsDefault creates a PostExternalsiteAuthorizationsDefault with default headers values
func NewPostExternalsiteAuthorizationsDefault(code int) *PostExternalsiteAuthorizationsDefault {
	return &PostExternalsiteAuthorizationsDefault{
		_statusCode: code,
	}
}

/*PostExternalsiteAuthorizationsDefault handles this case with default header values.

unexpected error
*/
type PostExternalsiteAuthorizationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post externalsite authorizations default response
func (o *PostExternalsiteAuthorizationsDefault) Code() int {
	return o._statusCode
}

func (o *PostExternalsiteAuthorizationsDefault) Error() string {
	return fmt.Sprintf("[POST /externalsite/Authorizations][%d] PostExternalsiteAuthorizations default  %+v", o._statusCode, o.Payload)
}

func (o *PostExternalsiteAuthorizationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostExternalsiteAuthorizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
