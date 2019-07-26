// Code generated by go-swagger; DO NOT EDIT.

package meta_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetMetaDefinitionsMoidReader is a Reader for the GetMetaDefinitionsMoid structure.
type GetMetaDefinitionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetaDefinitionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetaDefinitionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetMetaDefinitionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMetaDefinitionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetaDefinitionsMoidOK creates a GetMetaDefinitionsMoidOK with default headers values
func NewGetMetaDefinitionsMoidOK() *GetMetaDefinitionsMoidOK {
	return &GetMetaDefinitionsMoidOK{}
}

/*GetMetaDefinitionsMoidOK handles this case with default header values.

An instance of metaDefinition
*/
type GetMetaDefinitionsMoidOK struct {
	Payload *models.MetaDefinition
}

func (o *GetMetaDefinitionsMoidOK) Error() string {
	return fmt.Sprintf("[GET /meta/Definitions/{moid}][%d] getMetaDefinitionsMoidOK  %+v", 200, o.Payload)
}

func (o *GetMetaDefinitionsMoidOK) GetPayload() *models.MetaDefinition {
	return o.Payload
}

func (o *GetMetaDefinitionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetaDefinitionsMoidNotFound creates a GetMetaDefinitionsMoidNotFound with default headers values
func NewGetMetaDefinitionsMoidNotFound() *GetMetaDefinitionsMoidNotFound {
	return &GetMetaDefinitionsMoidNotFound{}
}

/*GetMetaDefinitionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetMetaDefinitionsMoidNotFound struct {
}

func (o *GetMetaDefinitionsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /meta/Definitions/{moid}][%d] getMetaDefinitionsMoidNotFound ", 404)
}

func (o *GetMetaDefinitionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMetaDefinitionsMoidDefault creates a GetMetaDefinitionsMoidDefault with default headers values
func NewGetMetaDefinitionsMoidDefault(code int) *GetMetaDefinitionsMoidDefault {
	return &GetMetaDefinitionsMoidDefault{
		_statusCode: code,
	}
}

/*GetMetaDefinitionsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetMetaDefinitionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get meta definitions moid default response
func (o *GetMetaDefinitionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetMetaDefinitionsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /meta/Definitions/{moid}][%d] GetMetaDefinitionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetaDefinitionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMetaDefinitionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}