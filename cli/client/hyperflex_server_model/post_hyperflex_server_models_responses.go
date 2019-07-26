// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexServerModelsReader is a Reader for the PostHyperflexServerModels structure.
type PostHyperflexServerModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexServerModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexServerModelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexServerModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexServerModelsCreated creates a PostHyperflexServerModelsCreated with default headers values
func NewPostHyperflexServerModelsCreated() *PostHyperflexServerModelsCreated {
	return &PostHyperflexServerModelsCreated{}
}

/*PostHyperflexServerModelsCreated handles this case with default header values.

Null response
*/
type PostHyperflexServerModelsCreated struct {
}

func (o *PostHyperflexServerModelsCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ServerModels][%d] postHyperflexServerModelsCreated ", 201)
}

func (o *PostHyperflexServerModelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexServerModelsDefault creates a PostHyperflexServerModelsDefault with default headers values
func NewPostHyperflexServerModelsDefault(code int) *PostHyperflexServerModelsDefault {
	return &PostHyperflexServerModelsDefault{
		_statusCode: code,
	}
}

/*PostHyperflexServerModelsDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexServerModelsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex server models default response
func (o *PostHyperflexServerModelsDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexServerModelsDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ServerModels][%d] PostHyperflexServerModels default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexServerModelsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexServerModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}