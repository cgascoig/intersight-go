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

// DeleteMetaDefinitionsMoidReader is a Reader for the DeleteMetaDefinitionsMoid structure.
type DeleteMetaDefinitionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMetaDefinitionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMetaDefinitionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteMetaDefinitionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMetaDefinitionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMetaDefinitionsMoidOK creates a DeleteMetaDefinitionsMoidOK with default headers values
func NewDeleteMetaDefinitionsMoidOK() *DeleteMetaDefinitionsMoidOK {
	return &DeleteMetaDefinitionsMoidOK{}
}

/*DeleteMetaDefinitionsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteMetaDefinitionsMoidOK struct {
}

func (o *DeleteMetaDefinitionsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /meta/Definitions/{moid}][%d] deleteMetaDefinitionsMoidOK ", 200)
}

func (o *DeleteMetaDefinitionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMetaDefinitionsMoidNotFound creates a DeleteMetaDefinitionsMoidNotFound with default headers values
func NewDeleteMetaDefinitionsMoidNotFound() *DeleteMetaDefinitionsMoidNotFound {
	return &DeleteMetaDefinitionsMoidNotFound{}
}

/*DeleteMetaDefinitionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteMetaDefinitionsMoidNotFound struct {
}

func (o *DeleteMetaDefinitionsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /meta/Definitions/{moid}][%d] deleteMetaDefinitionsMoidNotFound ", 404)
}

func (o *DeleteMetaDefinitionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMetaDefinitionsMoidDefault creates a DeleteMetaDefinitionsMoidDefault with default headers values
func NewDeleteMetaDefinitionsMoidDefault(code int) *DeleteMetaDefinitionsMoidDefault {
	return &DeleteMetaDefinitionsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteMetaDefinitionsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteMetaDefinitionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete meta definitions moid default response
func (o *DeleteMetaDefinitionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMetaDefinitionsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /meta/Definitions/{moid}][%d] DeleteMetaDefinitionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMetaDefinitionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMetaDefinitionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
