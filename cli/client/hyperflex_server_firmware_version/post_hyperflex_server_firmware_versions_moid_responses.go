// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_firmware_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexServerFirmwareVersionsMoidReader is a Reader for the PostHyperflexServerFirmwareVersionsMoid structure.
type PostHyperflexServerFirmwareVersionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexServerFirmwareVersionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexServerFirmwareVersionsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexServerFirmwareVersionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexServerFirmwareVersionsMoidCreated creates a PostHyperflexServerFirmwareVersionsMoidCreated with default headers values
func NewPostHyperflexServerFirmwareVersionsMoidCreated() *PostHyperflexServerFirmwareVersionsMoidCreated {
	return &PostHyperflexServerFirmwareVersionsMoidCreated{}
}

/*PostHyperflexServerFirmwareVersionsMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexServerFirmwareVersionsMoidCreated struct {
}

func (o *PostHyperflexServerFirmwareVersionsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ServerFirmwareVersions/{moid}][%d] postHyperflexServerFirmwareVersionsMoidCreated ", 201)
}

func (o *PostHyperflexServerFirmwareVersionsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexServerFirmwareVersionsMoidDefault creates a PostHyperflexServerFirmwareVersionsMoidDefault with default headers values
func NewPostHyperflexServerFirmwareVersionsMoidDefault(code int) *PostHyperflexServerFirmwareVersionsMoidDefault {
	return &PostHyperflexServerFirmwareVersionsMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexServerFirmwareVersionsMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexServerFirmwareVersionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex server firmware versions moid default response
func (o *PostHyperflexServerFirmwareVersionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexServerFirmwareVersionsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ServerFirmwareVersions/{moid}][%d] PostHyperflexServerFirmwareVersionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexServerFirmwareVersionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexServerFirmwareVersionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
