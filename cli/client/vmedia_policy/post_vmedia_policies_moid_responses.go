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

// PostVmediaPoliciesMoidReader is a Reader for the PostVmediaPoliciesMoid structure.
type PostVmediaPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVmediaPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVmediaPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVmediaPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVmediaPoliciesMoidCreated creates a PostVmediaPoliciesMoidCreated with default headers values
func NewPostVmediaPoliciesMoidCreated() *PostVmediaPoliciesMoidCreated {
	return &PostVmediaPoliciesMoidCreated{}
}

/*PostVmediaPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostVmediaPoliciesMoidCreated struct {
}

func (o *PostVmediaPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /vmedia/Policies/{moid}][%d] postVmediaPoliciesMoidCreated ", 201)
}

func (o *PostVmediaPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVmediaPoliciesMoidDefault creates a PostVmediaPoliciesMoidDefault with default headers values
func NewPostVmediaPoliciesMoidDefault(code int) *PostVmediaPoliciesMoidDefault {
	return &PostVmediaPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostVmediaPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostVmediaPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post vmedia policies moid default response
func (o *PostVmediaPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostVmediaPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /vmedia/Policies/{moid}][%d] PostVmediaPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostVmediaPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVmediaPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
