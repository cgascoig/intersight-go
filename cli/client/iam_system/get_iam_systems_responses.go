// Code generated by go-swagger; DO NOT EDIT.

package iam_system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamSystemsReader is a Reader for the GetIamSystems structure.
type GetIamSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamSystemsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamSystemsOK creates a GetIamSystemsOK with default headers values
func NewGetIamSystemsOK() *GetIamSystemsOK {
	return &GetIamSystemsOK{}
}

/*GetIamSystemsOK handles this case with default header values.

List of iamSystems for the given filter criteria
*/
type GetIamSystemsOK struct {
	Payload *models.IamSystemList
}

func (o *GetIamSystemsOK) Error() string {
	return fmt.Sprintf("[GET /iam/Systems][%d] getIamSystemsOK  %+v", 200, o.Payload)
}

func (o *GetIamSystemsOK) GetPayload() *models.IamSystemList {
	return o.Payload
}

func (o *GetIamSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamSystemList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamSystemsDefault creates a GetIamSystemsDefault with default headers values
func NewGetIamSystemsDefault(code int) *GetIamSystemsDefault {
	return &GetIamSystemsDefault{
		_statusCode: code,
	}
}

/*GetIamSystemsDefault handles this case with default header values.

Unexpected error
*/
type GetIamSystemsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam systems default response
func (o *GetIamSystemsDefault) Code() int {
	return o._statusCode
}

func (o *GetIamSystemsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/Systems][%d] GetIamSystems default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamSystemsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamSystemsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}