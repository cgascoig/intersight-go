// Code generated by go-swagger; DO NOT EDIT.

package memory_array

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostMemoryArraysMoidReader is a Reader for the PostMemoryArraysMoid structure.
type PostMemoryArraysMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMemoryArraysMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostMemoryArraysMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostMemoryArraysMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMemoryArraysMoidCreated creates a PostMemoryArraysMoidCreated with default headers values
func NewPostMemoryArraysMoidCreated() *PostMemoryArraysMoidCreated {
	return &PostMemoryArraysMoidCreated{}
}

/*PostMemoryArraysMoidCreated handles this case with default header values.

Null response
*/
type PostMemoryArraysMoidCreated struct {
}

func (o *PostMemoryArraysMoidCreated) Error() string {
	return fmt.Sprintf("[POST /memory/Arrays/{moid}][%d] postMemoryArraysMoidCreated ", 201)
}

func (o *PostMemoryArraysMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMemoryArraysMoidDefault creates a PostMemoryArraysMoidDefault with default headers values
func NewPostMemoryArraysMoidDefault(code int) *PostMemoryArraysMoidDefault {
	return &PostMemoryArraysMoidDefault{
		_statusCode: code,
	}
}

/*PostMemoryArraysMoidDefault handles this case with default header values.

unexpected error
*/
type PostMemoryArraysMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post memory arrays moid default response
func (o *PostMemoryArraysMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostMemoryArraysMoidDefault) Error() string {
	return fmt.Sprintf("[POST /memory/Arrays/{moid}][%d] PostMemoryArraysMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostMemoryArraysMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostMemoryArraysMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}