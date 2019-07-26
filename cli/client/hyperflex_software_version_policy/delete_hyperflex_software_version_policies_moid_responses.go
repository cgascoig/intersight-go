// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_software_version_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteHyperflexSoftwareVersionPoliciesMoidReader is a Reader for the DeleteHyperflexSoftwareVersionPoliciesMoid structure.
type DeleteHyperflexSoftwareVersionPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHyperflexSoftwareVersionPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHyperflexSoftwareVersionPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHyperflexSoftwareVersionPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHyperflexSoftwareVersionPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHyperflexSoftwareVersionPoliciesMoidOK creates a DeleteHyperflexSoftwareVersionPoliciesMoidOK with default headers values
func NewDeleteHyperflexSoftwareVersionPoliciesMoidOK() *DeleteHyperflexSoftwareVersionPoliciesMoidOK {
	return &DeleteHyperflexSoftwareVersionPoliciesMoidOK{}
}

/*DeleteHyperflexSoftwareVersionPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteHyperflexSoftwareVersionPoliciesMoidOK struct {
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/SoftwareVersionPolicies/{moid}][%d] deleteHyperflexSoftwareVersionPoliciesMoidOK ", 200)
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexSoftwareVersionPoliciesMoidNotFound creates a DeleteHyperflexSoftwareVersionPoliciesMoidNotFound with default headers values
func NewDeleteHyperflexSoftwareVersionPoliciesMoidNotFound() *DeleteHyperflexSoftwareVersionPoliciesMoidNotFound {
	return &DeleteHyperflexSoftwareVersionPoliciesMoidNotFound{}
}

/*DeleteHyperflexSoftwareVersionPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteHyperflexSoftwareVersionPoliciesMoidNotFound struct {
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/SoftwareVersionPolicies/{moid}][%d] deleteHyperflexSoftwareVersionPoliciesMoidNotFound ", 404)
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexSoftwareVersionPoliciesMoidDefault creates a DeleteHyperflexSoftwareVersionPoliciesMoidDefault with default headers values
func NewDeleteHyperflexSoftwareVersionPoliciesMoidDefault(code int) *DeleteHyperflexSoftwareVersionPoliciesMoidDefault {
	return &DeleteHyperflexSoftwareVersionPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteHyperflexSoftwareVersionPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteHyperflexSoftwareVersionPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hyperflex software version policies moid default response
func (o *DeleteHyperflexSoftwareVersionPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/SoftwareVersionPolicies/{moid}][%d] DeleteHyperflexSoftwareVersionPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHyperflexSoftwareVersionPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
