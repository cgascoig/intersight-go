// Code generated by go-swagger; DO NOT EDIT.

package kvm_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteKvmPoliciesMoidReader is a Reader for the DeleteKvmPoliciesMoid structure.
type DeleteKvmPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKvmPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteKvmPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteKvmPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteKvmPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteKvmPoliciesMoidOK creates a DeleteKvmPoliciesMoidOK with default headers values
func NewDeleteKvmPoliciesMoidOK() *DeleteKvmPoliciesMoidOK {
	return &DeleteKvmPoliciesMoidOK{}
}

/*DeleteKvmPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteKvmPoliciesMoidOK struct {
}

func (o *DeleteKvmPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /kvm/Policies/{moid}][%d] deleteKvmPoliciesMoidOK ", 200)
}

func (o *DeleteKvmPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKvmPoliciesMoidNotFound creates a DeleteKvmPoliciesMoidNotFound with default headers values
func NewDeleteKvmPoliciesMoidNotFound() *DeleteKvmPoliciesMoidNotFound {
	return &DeleteKvmPoliciesMoidNotFound{}
}

/*DeleteKvmPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteKvmPoliciesMoidNotFound struct {
}

func (o *DeleteKvmPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /kvm/Policies/{moid}][%d] deleteKvmPoliciesMoidNotFound ", 404)
}

func (o *DeleteKvmPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKvmPoliciesMoidDefault creates a DeleteKvmPoliciesMoidDefault with default headers values
func NewDeleteKvmPoliciesMoidDefault(code int) *DeleteKvmPoliciesMoidDefault {
	return &DeleteKvmPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteKvmPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteKvmPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete kvm policies moid default response
func (o *DeleteKvmPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteKvmPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /kvm/Policies/{moid}][%d] DeleteKvmPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteKvmPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteKvmPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
