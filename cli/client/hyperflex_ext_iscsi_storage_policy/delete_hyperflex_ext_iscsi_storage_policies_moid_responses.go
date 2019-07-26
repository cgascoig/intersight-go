// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ext_iscsi_storage_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteHyperflexExtIscsiStoragePoliciesMoidReader is a Reader for the DeleteHyperflexExtIscsiStoragePoliciesMoid structure.
type DeleteHyperflexExtIscsiStoragePoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHyperflexExtIscsiStoragePoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHyperflexExtIscsiStoragePoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHyperflexExtIscsiStoragePoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidOK creates a DeleteHyperflexExtIscsiStoragePoliciesMoidOK with default headers values
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidOK() *DeleteHyperflexExtIscsiStoragePoliciesMoidOK {
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidOK{}
}

/*DeleteHyperflexExtIscsiStoragePoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteHyperflexExtIscsiStoragePoliciesMoidOK struct {
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ExtIscsiStoragePolicies/{moid}][%d] deleteHyperflexExtIscsiStoragePoliciesMoidOK ", 200)
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidNotFound creates a DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound with default headers values
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidNotFound() *DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound {
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound{}
}

/*DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound struct {
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ExtIscsiStoragePolicies/{moid}][%d] deleteHyperflexExtIscsiStoragePoliciesMoidNotFound ", 404)
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexExtIscsiStoragePoliciesMoidDefault creates a DeleteHyperflexExtIscsiStoragePoliciesMoidDefault with default headers values
func NewDeleteHyperflexExtIscsiStoragePoliciesMoidDefault(code int) *DeleteHyperflexExtIscsiStoragePoliciesMoidDefault {
	return &DeleteHyperflexExtIscsiStoragePoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteHyperflexExtIscsiStoragePoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteHyperflexExtIscsiStoragePoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hyperflex ext iscsi storage policies moid default response
func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ExtIscsiStoragePolicies/{moid}][%d] DeleteHyperflexExtIscsiStoragePoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHyperflexExtIscsiStoragePoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}