// Code generated by go-swagger; DO NOT EDIT.

package vmedia_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostVmediaPoliciesReader is a Reader for the PostVmediaPolicies structure.
type PostVmediaPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVmediaPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVmediaPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVmediaPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVmediaPoliciesCreated creates a PostVmediaPoliciesCreated with default headers values
func NewPostVmediaPoliciesCreated() *PostVmediaPoliciesCreated {
	return &PostVmediaPoliciesCreated{}
}

/*PostVmediaPoliciesCreated handles this case with default header values.

Null response
*/
type PostVmediaPoliciesCreated struct {
}

func (o *PostVmediaPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /vmedia/Policies][%d] postVmediaPoliciesCreated ", 201)
}

func (o *PostVmediaPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVmediaPoliciesDefault creates a PostVmediaPoliciesDefault with default headers values
func NewPostVmediaPoliciesDefault(code int) *PostVmediaPoliciesDefault {
	return &PostVmediaPoliciesDefault{
		_statusCode: code,
	}
}

/*PostVmediaPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostVmediaPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post vmedia policies default response
func (o *PostVmediaPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostVmediaPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /vmedia/Policies][%d] PostVmediaPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostVmediaPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVmediaPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
