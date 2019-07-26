// Code generated by go-swagger; DO NOT EDIT.

package iam_qualifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteIamQualifiersMoidReader is a Reader for the DeleteIamQualifiersMoid structure.
type DeleteIamQualifiersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIamQualifiersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIamQualifiersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIamQualifiersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIamQualifiersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIamQualifiersMoidOK creates a DeleteIamQualifiersMoidOK with default headers values
func NewDeleteIamQualifiersMoidOK() *DeleteIamQualifiersMoidOK {
	return &DeleteIamQualifiersMoidOK{}
}

/*DeleteIamQualifiersMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteIamQualifiersMoidOK struct {
}

func (o *DeleteIamQualifiersMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/Qualifiers/{moid}][%d] deleteIamQualifiersMoidOK ", 200)
}

func (o *DeleteIamQualifiersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamQualifiersMoidNotFound creates a DeleteIamQualifiersMoidNotFound with default headers values
func NewDeleteIamQualifiersMoidNotFound() *DeleteIamQualifiersMoidNotFound {
	return &DeleteIamQualifiersMoidNotFound{}
}

/*DeleteIamQualifiersMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteIamQualifiersMoidNotFound struct {
}

func (o *DeleteIamQualifiersMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/Qualifiers/{moid}][%d] deleteIamQualifiersMoidNotFound ", 404)
}

func (o *DeleteIamQualifiersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamQualifiersMoidDefault creates a DeleteIamQualifiersMoidDefault with default headers values
func NewDeleteIamQualifiersMoidDefault(code int) *DeleteIamQualifiersMoidDefault {
	return &DeleteIamQualifiersMoidDefault{
		_statusCode: code,
	}
}

/*DeleteIamQualifiersMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteIamQualifiersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete iam qualifiers moid default response
func (o *DeleteIamQualifiersMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIamQualifiersMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/Qualifiers/{moid}][%d] DeleteIamQualifiersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIamQualifiersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIamQualifiersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
