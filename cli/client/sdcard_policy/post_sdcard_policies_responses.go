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

// PostSdcardPoliciesReader is a Reader for the PostSdcardPolicies structure.
type PostSdcardPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSdcardPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSdcardPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostSdcardPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSdcardPoliciesCreated creates a PostSdcardPoliciesCreated with default headers values
func NewPostSdcardPoliciesCreated() *PostSdcardPoliciesCreated {
	return &PostSdcardPoliciesCreated{}
}

/*PostSdcardPoliciesCreated handles this case with default header values.

Null response
*/
type PostSdcardPoliciesCreated struct {
}

func (o *PostSdcardPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /sdcard/Policies][%d] postSdcardPoliciesCreated ", 201)
}

func (o *PostSdcardPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSdcardPoliciesDefault creates a PostSdcardPoliciesDefault with default headers values
func NewPostSdcardPoliciesDefault(code int) *PostSdcardPoliciesDefault {
	return &PostSdcardPoliciesDefault{
		_statusCode: code,
	}
}

/*PostSdcardPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostSdcardPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post sdcard policies default response
func (o *PostSdcardPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostSdcardPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /sdcard/Policies][%d] PostSdcardPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostSdcardPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSdcardPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
