// Code generated by go-swagger; DO NOT EDIT.

package server_config_import

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostServerConfigImportsReader is a Reader for the PostServerConfigImports structure.
type PostServerConfigImportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServerConfigImportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostServerConfigImportsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostServerConfigImportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostServerConfigImportsCreated creates a PostServerConfigImportsCreated with default headers values
func NewPostServerConfigImportsCreated() *PostServerConfigImportsCreated {
	return &PostServerConfigImportsCreated{}
}

/*PostServerConfigImportsCreated handles this case with default header values.

Null response
*/
type PostServerConfigImportsCreated struct {
}

func (o *PostServerConfigImportsCreated) Error() string {
	return fmt.Sprintf("[POST /server/ConfigImports][%d] postServerConfigImportsCreated ", 201)
}

func (o *PostServerConfigImportsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServerConfigImportsDefault creates a PostServerConfigImportsDefault with default headers values
func NewPostServerConfigImportsDefault(code int) *PostServerConfigImportsDefault {
	return &PostServerConfigImportsDefault{
		_statusCode: code,
	}
}

/*PostServerConfigImportsDefault handles this case with default header values.

unexpected error
*/
type PostServerConfigImportsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post server config imports default response
func (o *PostServerConfigImportsDefault) Code() int {
	return o._statusCode
}

func (o *PostServerConfigImportsDefault) Error() string {
	return fmt.Sprintf("[POST /server/ConfigImports][%d] PostServerConfigImports default  %+v", o._statusCode, o.Payload)
}

func (o *PostServerConfigImportsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServerConfigImportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
