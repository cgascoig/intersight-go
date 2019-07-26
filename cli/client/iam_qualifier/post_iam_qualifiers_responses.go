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

// PostIamQualifiersReader is a Reader for the PostIamQualifiers structure.
type PostIamQualifiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIamQualifiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIamQualifiersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIamQualifiersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIamQualifiersCreated creates a PostIamQualifiersCreated with default headers values
func NewPostIamQualifiersCreated() *PostIamQualifiersCreated {
	return &PostIamQualifiersCreated{}
}

/*PostIamQualifiersCreated handles this case with default header values.

Null response
*/
type PostIamQualifiersCreated struct {
}

func (o *PostIamQualifiersCreated) Error() string {
	return fmt.Sprintf("[POST /iam/Qualifiers][%d] postIamQualifiersCreated ", 201)
}

func (o *PostIamQualifiersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIamQualifiersDefault creates a PostIamQualifiersDefault with default headers values
func NewPostIamQualifiersDefault(code int) *PostIamQualifiersDefault {
	return &PostIamQualifiersDefault{
		_statusCode: code,
	}
}

/*PostIamQualifiersDefault handles this case with default header values.

unexpected error
*/
type PostIamQualifiersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post iam qualifiers default response
func (o *PostIamQualifiersDefault) Code() int {
	return o._statusCode
}

func (o *PostIamQualifiersDefault) Error() string {
	return fmt.Sprintf("[POST /iam/Qualifiers][%d] PostIamQualifiers default  %+v", o._statusCode, o.Payload)
}

func (o *PostIamQualifiersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostIamQualifiersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
