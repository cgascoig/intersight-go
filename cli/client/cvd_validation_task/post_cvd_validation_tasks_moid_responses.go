// Code generated by go-swagger; DO NOT EDIT.

package cvd_validation_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostCvdValidationTasksMoidReader is a Reader for the PostCvdValidationTasksMoid structure.
type PostCvdValidationTasksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCvdValidationTasksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCvdValidationTasksMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCvdValidationTasksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCvdValidationTasksMoidCreated creates a PostCvdValidationTasksMoidCreated with default headers values
func NewPostCvdValidationTasksMoidCreated() *PostCvdValidationTasksMoidCreated {
	return &PostCvdValidationTasksMoidCreated{}
}

/*PostCvdValidationTasksMoidCreated handles this case with default header values.

Null response
*/
type PostCvdValidationTasksMoidCreated struct {
}

func (o *PostCvdValidationTasksMoidCreated) Error() string {
	return fmt.Sprintf("[POST /cvd/ValidationTasks/{moid}][%d] postCvdValidationTasksMoidCreated ", 201)
}

func (o *PostCvdValidationTasksMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCvdValidationTasksMoidDefault creates a PostCvdValidationTasksMoidDefault with default headers values
func NewPostCvdValidationTasksMoidDefault(code int) *PostCvdValidationTasksMoidDefault {
	return &PostCvdValidationTasksMoidDefault{
		_statusCode: code,
	}
}

/*PostCvdValidationTasksMoidDefault handles this case with default header values.

unexpected error
*/
type PostCvdValidationTasksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post cvd validation tasks moid default response
func (o *PostCvdValidationTasksMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostCvdValidationTasksMoidDefault) Error() string {
	return fmt.Sprintf("[POST /cvd/ValidationTasks/{moid}][%d] PostCvdValidationTasksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostCvdValidationTasksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostCvdValidationTasksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
