// Code generated by go-swagger; DO NOT EDIT.

package bios_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostBiosUnitsMoidReader is a Reader for the PostBiosUnitsMoid structure.
type PostBiosUnitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBiosUnitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostBiosUnitsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostBiosUnitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostBiosUnitsMoidCreated creates a PostBiosUnitsMoidCreated with default headers values
func NewPostBiosUnitsMoidCreated() *PostBiosUnitsMoidCreated {
	return &PostBiosUnitsMoidCreated{}
}

/*PostBiosUnitsMoidCreated handles this case with default header values.

Null response
*/
type PostBiosUnitsMoidCreated struct {
}

func (o *PostBiosUnitsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /bios/Units/{moid}][%d] postBiosUnitsMoidCreated ", 201)
}

func (o *PostBiosUnitsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBiosUnitsMoidDefault creates a PostBiosUnitsMoidDefault with default headers values
func NewPostBiosUnitsMoidDefault(code int) *PostBiosUnitsMoidDefault {
	return &PostBiosUnitsMoidDefault{
		_statusCode: code,
	}
}

/*PostBiosUnitsMoidDefault handles this case with default header values.

unexpected error
*/
type PostBiosUnitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post bios units moid default response
func (o *PostBiosUnitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostBiosUnitsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /bios/Units/{moid}][%d] PostBiosUnitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostBiosUnitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostBiosUnitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}