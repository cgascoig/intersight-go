// Code generated by go-swagger; DO NOT EDIT.

package firmware_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostFirmwareDistributablesMoidReader is a Reader for the PostFirmwareDistributablesMoid structure.
type PostFirmwareDistributablesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFirmwareDistributablesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFirmwareDistributablesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostFirmwareDistributablesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostFirmwareDistributablesMoidCreated creates a PostFirmwareDistributablesMoidCreated with default headers values
func NewPostFirmwareDistributablesMoidCreated() *PostFirmwareDistributablesMoidCreated {
	return &PostFirmwareDistributablesMoidCreated{}
}

/*PostFirmwareDistributablesMoidCreated handles this case with default header values.

Null response
*/
type PostFirmwareDistributablesMoidCreated struct {
}

func (o *PostFirmwareDistributablesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /firmware/Distributables/{moid}][%d] postFirmwareDistributablesMoidCreated ", 201)
}

func (o *PostFirmwareDistributablesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFirmwareDistributablesMoidDefault creates a PostFirmwareDistributablesMoidDefault with default headers values
func NewPostFirmwareDistributablesMoidDefault(code int) *PostFirmwareDistributablesMoidDefault {
	return &PostFirmwareDistributablesMoidDefault{
		_statusCode: code,
	}
}

/*PostFirmwareDistributablesMoidDefault handles this case with default header values.

unexpected error
*/
type PostFirmwareDistributablesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post firmware distributables moid default response
func (o *PostFirmwareDistributablesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostFirmwareDistributablesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /firmware/Distributables/{moid}][%d] PostFirmwareDistributablesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostFirmwareDistributablesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFirmwareDistributablesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
