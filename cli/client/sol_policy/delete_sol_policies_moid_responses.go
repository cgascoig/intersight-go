// Code generated by go-swagger; DO NOT EDIT.

package sol_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteSolPoliciesMoidReader is a Reader for the DeleteSolPoliciesMoid structure.
type DeleteSolPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSolPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSolPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSolPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSolPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSolPoliciesMoidOK creates a DeleteSolPoliciesMoidOK with default headers values
func NewDeleteSolPoliciesMoidOK() *DeleteSolPoliciesMoidOK {
	return &DeleteSolPoliciesMoidOK{}
}

/*DeleteSolPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteSolPoliciesMoidOK struct {
}

func (o *DeleteSolPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /sol/Policies/{moid}][%d] deleteSolPoliciesMoidOK ", 200)
}

func (o *DeleteSolPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSolPoliciesMoidNotFound creates a DeleteSolPoliciesMoidNotFound with default headers values
func NewDeleteSolPoliciesMoidNotFound() *DeleteSolPoliciesMoidNotFound {
	return &DeleteSolPoliciesMoidNotFound{}
}

/*DeleteSolPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteSolPoliciesMoidNotFound struct {
}

func (o *DeleteSolPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sol/Policies/{moid}][%d] deleteSolPoliciesMoidNotFound ", 404)
}

func (o *DeleteSolPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSolPoliciesMoidDefault creates a DeleteSolPoliciesMoidDefault with default headers values
func NewDeleteSolPoliciesMoidDefault(code int) *DeleteSolPoliciesMoidDefault {
	return &DeleteSolPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteSolPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteSolPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete sol policies moid default response
func (o *DeleteSolPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSolPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /sol/Policies/{moid}][%d] DeleteSolPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSolPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSolPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
