// Code generated by go-swagger; DO NOT EDIT.

package iam_idp_reference

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostIamIdpReferencesMoidReader is a Reader for the PostIamIdpReferencesMoid structure.
type PostIamIdpReferencesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamIdpReferencesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamIdpReferencesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamIdpReferencesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamIdpReferencesMoidCreated creates a PostIamIdpReferencesMoidCreated with default headers values
func NewPostIamIdpReferencesMoidCreated() *PostIamIdpReferencesMoidCreated {
	return &PostIamIdpReferencesMoidCreated{}
}

/*PostIamIdpReferencesMoidCreated handles this case with default header values.

Null response
*/
type PostIamIdpReferencesMoidCreated struct {
}

func (o *PostIamIdpReferencesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /iam/IdpReferences/{moid}][%d] postIamIdpReferencesMoidCreated ", 201)
}

func (o *PostIamIdpReferencesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamIdpReferencesMoidDefault creates a PostIamIdpReferencesMoidDefault with default headers values
func NewPostIamIdpReferencesMoidDefault(code int) *PostIamIdpReferencesMoidDefault {
	return &PostIamIdpReferencesMoidDefault{
		_statusCode: code,
	}
}

/*PostIamIdpReferencesMoidDefault handles this case with default header values.

unexpected error
*/
type PostIamIdpReferencesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam idp references moid default response
func (o *PostIamIdpReferencesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostIamIdpReferencesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /iam/IdpReferences/{moid}][%d] PostIamIdpReferencesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamIdpReferencesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamIdpReferencesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
