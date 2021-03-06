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

// GetIamIdpReferencesReader is a Reader for the GetIamIdpReferences structure.
type GetIamIdpReferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamIdpReferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamIdpReferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamIdpReferencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamIdpReferencesOK creates a GetIamIdpReferencesOK with default headers values
func NewGetIamIdpReferencesOK() *GetIamIdpReferencesOK {
	return &GetIamIdpReferencesOK{}
}

/*GetIamIdpReferencesOK handles this case with default header values.

List of iamIdpReferences for the given filter criteria
*/
type GetIamIdpReferencesOK struct {
	Payload *models.IamIdpReferenceList
}

func (o *GetIamIdpReferencesOK) Error() string {
	return fmt.Sprintf("[GET /iam/IdpReferences][%d] getIamIdpReferencesOK  %+v", 200, o.Payload)
}

func (o *GetIamIdpReferencesOK) GetPayload() *models.IamIdpReferenceList {
	return o.Payload
}

func (o *GetIamIdpReferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamIdpReferenceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamIdpReferencesDefault creates a GetIamIdpReferencesDefault with default headers values
func NewGetIamIdpReferencesDefault(code int) *GetIamIdpReferencesDefault {
	return &GetIamIdpReferencesDefault{
		_statusCode: code,
	}
}

/*GetIamIdpReferencesDefault handles this case with default header values.

Unexpected error
*/
type GetIamIdpReferencesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam idp references default response
func (o *GetIamIdpReferencesDefault) Code() int {
	return o._statusCode
}

func (o *GetIamIdpReferencesDefault) Error() string {
	return fmt.Sprintf("[GET /iam/IdpReferences][%d] GetIamIdpReferences default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamIdpReferencesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamIdpReferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
