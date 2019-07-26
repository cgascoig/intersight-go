// Code generated by go-swagger; DO NOT EDIT.

package fault_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostFaultInstancesMoidReader is a Reader for the PostFaultInstancesMoid structure.
type PostFaultInstancesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFaultInstancesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFaultInstancesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostFaultInstancesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostFaultInstancesMoidCreated creates a PostFaultInstancesMoidCreated with default headers values
func NewPostFaultInstancesMoidCreated() *PostFaultInstancesMoidCreated {
	return &PostFaultInstancesMoidCreated{}
}

/*PostFaultInstancesMoidCreated handles this case with default header values.

Null response
*/
type PostFaultInstancesMoidCreated struct {
}

func (o *PostFaultInstancesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /fault/Instances/{moid}][%d] postFaultInstancesMoidCreated ", 201)
}

func (o *PostFaultInstancesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFaultInstancesMoidDefault creates a PostFaultInstancesMoidDefault with default headers values
func NewPostFaultInstancesMoidDefault(code int) *PostFaultInstancesMoidDefault {
	return &PostFaultInstancesMoidDefault{
		_statusCode: code,
	}
}

/*PostFaultInstancesMoidDefault handles this case with default header values.

unexpected error
*/
type PostFaultInstancesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post fault instances moid default response
func (o *PostFaultInstancesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostFaultInstancesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /fault/Instances/{moid}][%d] PostFaultInstancesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostFaultInstancesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFaultInstancesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}