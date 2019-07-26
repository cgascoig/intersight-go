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

// GetIamIdpReferencesMoidReader is a Reader for the GetIamIdpReferencesMoid structure.
type GetIamIdpReferencesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamIdpReferencesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamIdpReferencesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIamIdpReferencesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIamIdpReferencesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamIdpReferencesMoidOK creates a GetIamIdpReferencesMoidOK with default headers values
func NewGetIamIdpReferencesMoidOK() *GetIamIdpReferencesMoidOK {
	return &GetIamIdpReferencesMoidOK{}
}

/*GetIamIdpReferencesMoidOK handles this case with default header values.

An instance of iamIdpReference
*/
type GetIamIdpReferencesMoidOK struct {
	Payload *models.IamIdpReference
}

func (o *GetIamIdpReferencesMoidOK) Error() string {
	return fmt.Sprintf("[GET /iam/IdpReferences/{moid}][%d] getIamIdpReferencesMoidOK  %+v", 200, o.Payload)
}

func (o *GetIamIdpReferencesMoidOK) GetPayload() *models.IamIdpReference {
	return o.Payload
}

func (o *GetIamIdpReferencesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamIdpReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamIdpReferencesMoidNotFound creates a GetIamIdpReferencesMoidNotFound with default headers values
func NewGetIamIdpReferencesMoidNotFound() *GetIamIdpReferencesMoidNotFound {
	return &GetIamIdpReferencesMoidNotFound{}
}

/*GetIamIdpReferencesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIamIdpReferencesMoidNotFound struct {
}

func (o *GetIamIdpReferencesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/IdpReferences/{moid}][%d] getIamIdpReferencesMoidNotFound ", 404)
}

func (o *GetIamIdpReferencesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIamIdpReferencesMoidDefault creates a GetIamIdpReferencesMoidDefault with default headers values
func NewGetIamIdpReferencesMoidDefault(code int) *GetIamIdpReferencesMoidDefault {
	return &GetIamIdpReferencesMoidDefault{
		_statusCode: code,
	}
}

/*GetIamIdpReferencesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIamIdpReferencesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam idp references moid default response
func (o *GetIamIdpReferencesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIamIdpReferencesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iam/IdpReferences/{moid}][%d] GetIamIdpReferencesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamIdpReferencesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamIdpReferencesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
