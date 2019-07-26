// Code generated by go-swagger; DO NOT EDIT.

package sdcard_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostSdcardPoliciesMoidReader is a Reader for the PostSdcardPoliciesMoid structure.
type PostSdcardPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSdcardPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSdcardPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostSdcardPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSdcardPoliciesMoidCreated creates a PostSdcardPoliciesMoidCreated with default headers values
func NewPostSdcardPoliciesMoidCreated() *PostSdcardPoliciesMoidCreated {
	return &PostSdcardPoliciesMoidCreated{}
}

/*PostSdcardPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostSdcardPoliciesMoidCreated struct {
}

func (o *PostSdcardPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /sdcard/Policies/{moid}][%d] postSdcardPoliciesMoidCreated ", 201)
}

func (o *PostSdcardPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSdcardPoliciesMoidDefault creates a PostSdcardPoliciesMoidDefault with default headers values
func NewPostSdcardPoliciesMoidDefault(code int) *PostSdcardPoliciesMoidDefault {
	return &PostSdcardPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostSdcardPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostSdcardPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post sdcard policies moid default response
func (o *PostSdcardPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostSdcardPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /sdcard/Policies/{moid}][%d] PostSdcardPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostSdcardPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSdcardPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}