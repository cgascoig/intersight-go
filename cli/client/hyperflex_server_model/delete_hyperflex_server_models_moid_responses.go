// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteHyperflexServerModelsMoidReader is a Reader for the DeleteHyperflexServerModelsMoid structure.
type DeleteHyperflexServerModelsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHyperflexServerModelsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHyperflexServerModelsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHyperflexServerModelsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHyperflexServerModelsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHyperflexServerModelsMoidOK creates a DeleteHyperflexServerModelsMoidOK with default headers values
func NewDeleteHyperflexServerModelsMoidOK() *DeleteHyperflexServerModelsMoidOK {
	return &DeleteHyperflexServerModelsMoidOK{}
}

/*DeleteHyperflexServerModelsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteHyperflexServerModelsMoidOK struct {
}

func (o *DeleteHyperflexServerModelsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ServerModels/{moid}][%d] deleteHyperflexServerModelsMoidOK ", 200)
}

func (o *DeleteHyperflexServerModelsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexServerModelsMoidNotFound creates a DeleteHyperflexServerModelsMoidNotFound with default headers values
func NewDeleteHyperflexServerModelsMoidNotFound() *DeleteHyperflexServerModelsMoidNotFound {
	return &DeleteHyperflexServerModelsMoidNotFound{}
}

/*DeleteHyperflexServerModelsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteHyperflexServerModelsMoidNotFound struct {
}

func (o *DeleteHyperflexServerModelsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ServerModels/{moid}][%d] deleteHyperflexServerModelsMoidNotFound ", 404)
}

func (o *DeleteHyperflexServerModelsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexServerModelsMoidDefault creates a DeleteHyperflexServerModelsMoidDefault with default headers values
func NewDeleteHyperflexServerModelsMoidDefault(code int) *DeleteHyperflexServerModelsMoidDefault {
	return &DeleteHyperflexServerModelsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteHyperflexServerModelsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteHyperflexServerModelsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hyperflex server models moid default response
func (o *DeleteHyperflexServerModelsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHyperflexServerModelsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ServerModels/{moid}][%d] DeleteHyperflexServerModelsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHyperflexServerModelsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHyperflexServerModelsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
