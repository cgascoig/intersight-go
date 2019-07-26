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

// GetIamAPIKeysReader is a Reader for the GetIamAPIKeys structure.
type GetIamAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamAPIKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamAPIKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamAPIKeysOK creates a GetIamAPIKeysOK with default headers values
func NewGetIamAPIKeysOK() *GetIamAPIKeysOK {
	return &GetIamAPIKeysOK{}
}

/*GetIamAPIKeysOK handles this case with default header values.

List of iamApiKeys for the given filter criteria
*/
type GetIamAPIKeysOK struct {
	Payload *models.IamAPIKeyList
}

func (o *GetIamAPIKeysOK) Error() string {
	return fmt.Sprintf("[GET /iam/ApiKeys][%d] getIamApiKeysOK  %+v", 200, o.Payload)
}

func (o *GetIamAPIKeysOK) GetPayload() *models.IamAPIKeyList {
	return o.Payload
}

func (o *GetIamAPIKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamAPIKeyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamAPIKeysDefault creates a GetIamAPIKeysDefault with default headers values
func NewGetIamAPIKeysDefault(code int) *GetIamAPIKeysDefault {
	return &GetIamAPIKeysDefault{
		_statusCode: code,
	}
}

/*GetIamAPIKeysDefault handles this case with default header values.

Unexpected error
*/
type GetIamAPIKeysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam API keys default response
func (o *GetIamAPIKeysDefault) Code() int {
	return o._statusCode
}

func (o *GetIamAPIKeysDefault) Error() string {
	return fmt.Sprintf("[GET /iam/ApiKeys][%d] GetIamAPIKeys default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamAPIKeysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamAPIKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
