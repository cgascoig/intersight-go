// Code generated by go-swagger; DO NOT EDIT.

package port_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostPortGroupsMoidReader is a Reader for the PostPortGroupsMoid structure.
type PostPortGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPortGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPortGroupsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostPortGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostPortGroupsMoidCreated creates a PostPortGroupsMoidCreated with default headers values
func NewPostPortGroupsMoidCreated() *PostPortGroupsMoidCreated {
	return &PostPortGroupsMoidCreated{}
}

/*PostPortGroupsMoidCreated handles this case with default header values.

Null response
*/
type PostPortGroupsMoidCreated struct {
}

func (o *PostPortGroupsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /port/Groups/{moid}][%d] postPortGroupsMoidCreated ", 201)
}

func (o *PostPortGroupsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPortGroupsMoidDefault creates a PostPortGroupsMoidDefault with default headers values
func NewPostPortGroupsMoidDefault(code int) *PostPortGroupsMoidDefault {
	return &PostPortGroupsMoidDefault{
		_statusCode: code,
	}
}

/*PostPortGroupsMoidDefault handles this case with default header values.

unexpected error
*/
type PostPortGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post port groups moid default response
func (o *PostPortGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostPortGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /port/Groups/{moid}][%d] PostPortGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostPortGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostPortGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}