// Code generated by go-swagger; DO NOT EDIT.

package iam_api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamAPIKeysMoidReader is a Reader for the GetIamAPIKeysMoid structure.
type GetIamAPIKeysMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamAPIKeysMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamAPIKeysMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIamAPIKeysMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIamAPIKeysMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamAPIKeysMoidOK creates a GetIamAPIKeysMoidOK with default headers values
func NewGetIamAPIKeysMoidOK() *GetIamAPIKeysMoidOK {
	return &GetIamAPIKeysMoidOK{}
}

/*GetIamAPIKeysMoidOK handles this case with default header values.

An instance of iamApiKey
*/
type GetIamAPIKeysMoidOK struct {
	Payload *models.IamAPIKey
}

func (o *GetIamAPIKeysMoidOK) Error() string {
	return fmt.Sprintf("[GET /iam/ApiKeys/{moid}][%d] getIamApiKeysMoidOK  %+v", 200, o.Payload)
}

func (o *GetIamAPIKeysMoidOK) GetPayload() *models.IamAPIKey {
	return o.Payload
}

func (o *GetIamAPIKeysMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamAPIKeysMoidNotFound creates a GetIamAPIKeysMoidNotFound with default headers values
func NewGetIamAPIKeysMoidNotFound() *GetIamAPIKeysMoidNotFound {
	return &GetIamAPIKeysMoidNotFound{}
}

/*GetIamAPIKeysMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIamAPIKeysMoidNotFound struct {
}

func (o *GetIamAPIKeysMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/ApiKeys/{moid}][%d] getIamApiKeysMoidNotFound ", 404)
}

func (o *GetIamAPIKeysMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIamAPIKeysMoidDefault creates a GetIamAPIKeysMoidDefault with default headers values
func NewGetIamAPIKeysMoidDefault(code int) *GetIamAPIKeysMoidDefault {
	return &GetIamAPIKeysMoidDefault{
		_statusCode: code,
	}
}

/*GetIamAPIKeysMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIamAPIKeysMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam API keys moid default response
func (o *GetIamAPIKeysMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIamAPIKeysMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iam/ApiKeys/{moid}][%d] GetIamAPIKeysMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamAPIKeysMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamAPIKeysMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}