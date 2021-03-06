// Code generated by go-swagger; DO NOT EDIT.

package server_config_result_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetServerConfigResultEntriesMoidReader is a Reader for the GetServerConfigResultEntriesMoid structure.
type GetServerConfigResultEntriesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerConfigResultEntriesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerConfigResultEntriesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServerConfigResultEntriesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServerConfigResultEntriesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerConfigResultEntriesMoidOK creates a GetServerConfigResultEntriesMoidOK with default headers values
func NewGetServerConfigResultEntriesMoidOK() *GetServerConfigResultEntriesMoidOK {
	return &GetServerConfigResultEntriesMoidOK{}
}

/*GetServerConfigResultEntriesMoidOK handles this case with default header values.

An instance of serverConfigResultEntry
*/
type GetServerConfigResultEntriesMoidOK struct {
	Payload *models.ServerConfigResultEntry
}

func (o *GetServerConfigResultEntriesMoidOK) Error() string {
	return fmt.Sprintf("[GET /server/ConfigResultEntries/{moid}][%d] getServerConfigResultEntriesMoidOK  %+v", 200, o.Payload)
}

func (o *GetServerConfigResultEntriesMoidOK) GetPayload() *models.ServerConfigResultEntry {
	return o.Payload
}

func (o *GetServerConfigResultEntriesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerConfigResultEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerConfigResultEntriesMoidNotFound creates a GetServerConfigResultEntriesMoidNotFound with default headers values
func NewGetServerConfigResultEntriesMoidNotFound() *GetServerConfigResultEntriesMoidNotFound {
	return &GetServerConfigResultEntriesMoidNotFound{}
}

/*GetServerConfigResultEntriesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetServerConfigResultEntriesMoidNotFound struct {
}

func (o *GetServerConfigResultEntriesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /server/ConfigResultEntries/{moid}][%d] getServerConfigResultEntriesMoidNotFound ", 404)
}

func (o *GetServerConfigResultEntriesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServerConfigResultEntriesMoidDefault creates a GetServerConfigResultEntriesMoidDefault with default headers values
func NewGetServerConfigResultEntriesMoidDefault(code int) *GetServerConfigResultEntriesMoidDefault {
	return &GetServerConfigResultEntriesMoidDefault{
		_statusCode: code,
	}
}

/*GetServerConfigResultEntriesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetServerConfigResultEntriesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get server config result entries moid default response
func (o *GetServerConfigResultEntriesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetServerConfigResultEntriesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /server/ConfigResultEntries/{moid}][%d] GetServerConfigResultEntriesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerConfigResultEntriesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerConfigResultEntriesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
