// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_user_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteIamEndPointUserPoliciesMoidReader is a Reader for the DeleteIamEndPointUserPoliciesMoid structure.
type DeleteIamEndPointUserPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIamEndPointUserPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIamEndPointUserPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIamEndPointUserPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteIamEndPointUserPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIamEndPointUserPoliciesMoidOK creates a DeleteIamEndPointUserPoliciesMoidOK with default headers values
func NewDeleteIamEndPointUserPoliciesMoidOK() *DeleteIamEndPointUserPoliciesMoidOK {
	return &DeleteIamEndPointUserPoliciesMoidOK{}
}

/*DeleteIamEndPointUserPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteIamEndPointUserPoliciesMoidOK struct {
}

func (o *DeleteIamEndPointUserPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/EndPointUserPolicies/{moid}][%d] deleteIamEndPointUserPoliciesMoidOK ", 200)
}

func (o *DeleteIamEndPointUserPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamEndPointUserPoliciesMoidNotFound creates a DeleteIamEndPointUserPoliciesMoidNotFound with default headers values
func NewDeleteIamEndPointUserPoliciesMoidNotFound() *DeleteIamEndPointUserPoliciesMoidNotFound {
	return &DeleteIamEndPointUserPoliciesMoidNotFound{}
}

/*DeleteIamEndPointUserPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteIamEndPointUserPoliciesMoidNotFound struct {
}

func (o *DeleteIamEndPointUserPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/EndPointUserPolicies/{moid}][%d] deleteIamEndPointUserPoliciesMoidNotFound ", 404)
}

func (o *DeleteIamEndPointUserPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIamEndPointUserPoliciesMoidDefault creates a DeleteIamEndPointUserPoliciesMoidDefault with default headers values
func NewDeleteIamEndPointUserPoliciesMoidDefault(code int) *DeleteIamEndPointUserPoliciesMoidDefault {
	return &DeleteIamEndPointUserPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteIamEndPointUserPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteIamEndPointUserPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete iam end point user policies moid default response
func (o *DeleteIamEndPointUserPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIamEndPointUserPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/EndPointUserPolicies/{moid}][%d] DeleteIamEndPointUserPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIamEndPointUserPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIamEndPointUserPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}