// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_hxdp_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteHyperflexHxdpVersionsMoidReader is a Reader for the DeleteHyperflexHxdpVersionsMoid structure.
type DeleteHyperflexHxdpVersionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHyperflexHxdpVersionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHyperflexHxdpVersionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHyperflexHxdpVersionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHyperflexHxdpVersionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHyperflexHxdpVersionsMoidOK creates a DeleteHyperflexHxdpVersionsMoidOK with default headers values
func NewDeleteHyperflexHxdpVersionsMoidOK() *DeleteHyperflexHxdpVersionsMoidOK {
	return &DeleteHyperflexHxdpVersionsMoidOK{}
}

/*DeleteHyperflexHxdpVersionsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteHyperflexHxdpVersionsMoidOK struct {
}

func (o *DeleteHyperflexHxdpVersionsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/HxdpVersions/{moid}][%d] deleteHyperflexHxdpVersionsMoidOK ", 200)
}

func (o *DeleteHyperflexHxdpVersionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexHxdpVersionsMoidNotFound creates a DeleteHyperflexHxdpVersionsMoidNotFound with default headers values
func NewDeleteHyperflexHxdpVersionsMoidNotFound() *DeleteHyperflexHxdpVersionsMoidNotFound {
	return &DeleteHyperflexHxdpVersionsMoidNotFound{}
}

/*DeleteHyperflexHxdpVersionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteHyperflexHxdpVersionsMoidNotFound struct {
}

func (o *DeleteHyperflexHxdpVersionsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/HxdpVersions/{moid}][%d] deleteHyperflexHxdpVersionsMoidNotFound ", 404)
}

func (o *DeleteHyperflexHxdpVersionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexHxdpVersionsMoidDefault creates a DeleteHyperflexHxdpVersionsMoidDefault with default headers values
func NewDeleteHyperflexHxdpVersionsMoidDefault(code int) *DeleteHyperflexHxdpVersionsMoidDefault {
	return &DeleteHyperflexHxdpVersionsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteHyperflexHxdpVersionsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteHyperflexHxdpVersionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hyperflex hxdp versions moid default response
func (o *DeleteHyperflexHxdpVersionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHyperflexHxdpVersionsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/HxdpVersions/{moid}][%d] DeleteHyperflexHxdpVersionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHyperflexHxdpVersionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHyperflexHxdpVersionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
