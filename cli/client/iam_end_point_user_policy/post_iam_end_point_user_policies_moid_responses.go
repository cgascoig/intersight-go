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

// PostIamEndPointUserPoliciesMoidReader is a Reader for the PostIamEndPointUserPoliciesMoid structure.
type PostIamEndPointUserPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamEndPointUserPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamEndPointUserPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamEndPointUserPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamEndPointUserPoliciesMoidCreated creates a PostIamEndPointUserPoliciesMoidCreated with default headers values
func NewPostIamEndPointUserPoliciesMoidCreated() *PostIamEndPointUserPoliciesMoidCreated {
	return &PostIamEndPointUserPoliciesMoidCreated{}
}

/*PostIamEndPointUserPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostIamEndPointUserPoliciesMoidCreated struct {
}

func (o *PostIamEndPointUserPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/EndPointUserPolicies/{moid}][%d] postIamEndPointUserPoliciesMoidCreated ", 201)
}

func (o *PostIamEndPointUserPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamEndPointUserPoliciesMoidDefault creates a PostIamEndPointUserPoliciesMoidDefault with default headers values
func NewPostIamEndPointUserPoliciesMoidDefault(code int) *PostIamEndPointUserPoliciesMoidDefault {
	return &PostIamEndPointUserPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostIamEndPointUserPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamEndPointUserPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam end point user policies moid default response
func (o *PostIamEndPointUserPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamEndPointUserPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/EndPointUserPolicies/{moid}][%d] PostIamEndPointUserPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamEndPointUserPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamEndPointUserPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
