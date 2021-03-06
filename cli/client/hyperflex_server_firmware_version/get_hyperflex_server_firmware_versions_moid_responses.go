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

// GetHyperflexServerFirmwareVersionsMoidReader is a Reader for the GetHyperflexServerFirmwareVersionsMoid structure.
type GetHyperflexServerFirmwareVersionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexServerFirmwareVersionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexServerFirmwareVersionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexServerFirmwareVersionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexServerFirmwareVersionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexServerFirmwareVersionsMoidOK creates a GetHyperflexServerFirmwareVersionsMoidOK with default headers values
func NewGetHyperflexServerFirmwareVersionsMoidOK() *GetHyperflexServerFirmwareVersionsMoidOK {
	return &GetHyperflexServerFirmwareVersionsMoidOK{}
}

/*GetHyperflexServerFirmwareVersionsMoidOK handles this case with default header values.

An instance of hyperflexServerFirmwareVersion
*/
type GetHyperflexServerFirmwareVersionsMoidOK struct {
	Payload *models.HyperflexServerFirmwareVersion
}

func (o *GetHyperflexServerFirmwareVersionsMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ServerFirmwareVersions/{moid}][%d] getHyperflexServerFirmwareVersionsMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexServerFirmwareVersionsMoidOK) GetPayload() *models.HyperflexServerFirmwareVersion {
	return o.Payload
}

func (o *GetHyperflexServerFirmwareVersionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexServerFirmwareVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexServerFirmwareVersionsMoidNotFound creates a GetHyperflexServerFirmwareVersionsMoidNotFound with default headers values
func NewGetHyperflexServerFirmwareVersionsMoidNotFound() *GetHyperflexServerFirmwareVersionsMoidNotFound {
	return &GetHyperflexServerFirmwareVersionsMoidNotFound{}
}

/*GetHyperflexServerFirmwareVersionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexServerFirmwareVersionsMoidNotFound struct {
}

func (o *GetHyperflexServerFirmwareVersionsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ServerFirmwareVersions/{moid}][%d] getHyperflexServerFirmwareVersionsMoidNotFound ", 404)
}

func (o *GetHyperflexServerFirmwareVersionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexServerFirmwareVersionsMoidDefault creates a GetHyperflexServerFirmwareVersionsMoidDefault with default headers values
func NewGetHyperflexServerFirmwareVersionsMoidDefault(code int) *GetHyperflexServerFirmwareVersionsMoidDefault {
	return &GetHyperflexServerFirmwareVersionsMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexServerFirmwareVersionsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexServerFirmwareVersionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex server firmware versions moid default response
func (o *GetHyperflexServerFirmwareVersionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexServerFirmwareVersionsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ServerFirmwareVersions/{moid}][%d] GetHyperflexServerFirmwareVersionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexServerFirmwareVersionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexServerFirmwareVersionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
