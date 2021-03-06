// Code generated by go-swagger; DO NOT EDIT.

package iam_session_limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamSessionLimitsReader is a Reader for the GetIamSessionLimits structure.
type GetIamSessionLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamSessionLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamSessionLimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamSessionLimitsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamSessionLimitsOK creates a GetIamSessionLimitsOK with default headers values
func NewGetIamSessionLimitsOK() *GetIamSessionLimitsOK {
	return &GetIamSessionLimitsOK{}
}

/*GetIamSessionLimitsOK handles this case with default header values.

List of iamSessionLimits for the given filter criteria
*/
type GetIamSessionLimitsOK struct {
	Payload *models.IamSessionLimitsList
}

func (o *GetIamSessionLimitsOK) Error() string {
	return fmt.Sprintf("[GET /iam/SessionLimits][%d] getIamSessionLimitsOK  %+v", 200, o.Payload)
}

func (o *GetIamSessionLimitsOK) GetPayload() *models.IamSessionLimitsList {
	return o.Payload
}

func (o *GetIamSessionLimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamSessionLimitsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamSessionLimitsDefault creates a GetIamSessionLimitsDefault with default headers values
func NewGetIamSessionLimitsDefault(code int) *GetIamSessionLimitsDefault {
	return &GetIamSessionLimitsDefault{
		_statusCode: code,
	}
}

/*GetIamSessionLimitsDefault handles this case with default header values.

Unexpected error
*/
type GetIamSessionLimitsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam session limits default response
func (o *GetIamSessionLimitsDefault) Code() int {
	return o._statusCode
}

func (o *GetIamSessionLimitsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/SessionLimits][%d] GetIamSessionLimits default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamSessionLimitsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamSessionLimitsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
