// Code generated by go-swagger; DO NOT EDIT.

package iam_trust_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteIamTrustPointsMoidReader is a Reader for the DeleteIamTrustPointsMoid structure.
type DeleteIamTrustPointsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIamTrustPointsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIamTrustPointsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIamTrustPointsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIamTrustPointsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIamTrustPointsMoidOK creates a DeleteIamTrustPointsMoidOK with default headers values
func NewDeleteIamTrustPointsMoidOK() *DeleteIamTrustPointsMoidOK {
	return &DeleteIamTrustPointsMoidOK{}
}

/*DeleteIamTrustPointsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteIamTrustPointsMoidOK struct {
}

func (o *DeleteIamTrustPointsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/TrustPoints/{moid}][%d] deleteIamTrustPointsMoidOK ", 200)
}

func (o *DeleteIamTrustPointsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamTrustPointsMoidNotFound creates a DeleteIamTrustPointsMoidNotFound with default headers values
func NewDeleteIamTrustPointsMoidNotFound() *DeleteIamTrustPointsMoidNotFound {
	return &DeleteIamTrustPointsMoidNotFound{}
}

/*DeleteIamTrustPointsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteIamTrustPointsMoidNotFound struct {
}

func (o *DeleteIamTrustPointsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/TrustPoints/{moid}][%d] deleteIamTrustPointsMoidNotFound ", 404)
}

func (o *DeleteIamTrustPointsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamTrustPointsMoidDefault creates a DeleteIamTrustPointsMoidDefault with default headers values
func NewDeleteIamTrustPointsMoidDefault(code int) *DeleteIamTrustPointsMoidDefault {
	return &DeleteIamTrustPointsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteIamTrustPointsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteIamTrustPointsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete iam trust points moid default response
func (o *DeleteIamTrustPointsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIamTrustPointsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/TrustPoints/{moid}][%d] DeleteIamTrustPointsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIamTrustPointsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIamTrustPointsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}