// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_capability_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteHyperflexCapabilityInfosMoidReader is a Reader for the DeleteHyperflexCapabilityInfosMoid structure.
type DeleteHyperflexCapabilityInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHyperflexCapabilityInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHyperflexCapabilityInfosMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHyperflexCapabilityInfosMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHyperflexCapabilityInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHyperflexCapabilityInfosMoidOK creates a DeleteHyperflexCapabilityInfosMoidOK with default headers values
func NewDeleteHyperflexCapabilityInfosMoidOK() *DeleteHyperflexCapabilityInfosMoidOK {
	return &DeleteHyperflexCapabilityInfosMoidOK{}
}

/*DeleteHyperflexCapabilityInfosMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteHyperflexCapabilityInfosMoidOK struct {
}

func (o *DeleteHyperflexCapabilityInfosMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/CapabilityInfos/{moid}][%d] deleteHyperflexCapabilityInfosMoidOK ", 200)
}

func (o *DeleteHyperflexCapabilityInfosMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexCapabilityInfosMoidNotFound creates a DeleteHyperflexCapabilityInfosMoidNotFound with default headers values
func NewDeleteHyperflexCapabilityInfosMoidNotFound() *DeleteHyperflexCapabilityInfosMoidNotFound {
	return &DeleteHyperflexCapabilityInfosMoidNotFound{}
}

/*DeleteHyperflexCapabilityInfosMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteHyperflexCapabilityInfosMoidNotFound struct {
}

func (o *DeleteHyperflexCapabilityInfosMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/CapabilityInfos/{moid}][%d] deleteHyperflexCapabilityInfosMoidNotFound ", 404)
}

func (o *DeleteHyperflexCapabilityInfosMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexCapabilityInfosMoidDefault creates a DeleteHyperflexCapabilityInfosMoidDefault with default headers values
func NewDeleteHyperflexCapabilityInfosMoidDefault(code int) *DeleteHyperflexCapabilityInfosMoidDefault {
	return &DeleteHyperflexCapabilityInfosMoidDefault{
		_statusCode: code,
	}
}

/*DeleteHyperflexCapabilityInfosMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteHyperflexCapabilityInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hyperflex capability infos moid default response
func (o *DeleteHyperflexCapabilityInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHyperflexCapabilityInfosMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/CapabilityInfos/{moid}][%d] DeleteHyperflexCapabilityInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHyperflexCapabilityInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHyperflexCapabilityInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}