// Code generated by go-swagger; DO NOT EDIT.

package workflow_task_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostWorkflowTaskDefinitionsReader is a Reader for the PostWorkflowTaskDefinitions structure.
type PostWorkflowTaskDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkflowTaskDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWorkflowTaskDefinitionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWorkflowTaskDefinitionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWorkflowTaskDefinitionsCreated creates a PostWorkflowTaskDefinitionsCreated with default headers values
func NewPostWorkflowTaskDefinitionsCreated() *PostWorkflowTaskDefinitionsCreated {
	return &PostWorkflowTaskDefinitionsCreated{}
}

/*PostWorkflowTaskDefinitionsCreated handles this case with default header values.

Null response
*/
type PostWorkflowTaskDefinitionsCreated struct {
}

func (o *PostWorkflowTaskDefinitionsCreated) Error() string {
	return fmt.Sprintf("[POST /workflow/TaskDefinitions][%d] postWorkflowTaskDefinitionsCreated ", 201)
}

func (o *PostWorkflowTaskDefinitionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWorkflowTaskDefinitionsDefault creates a PostWorkflowTaskDefinitionsDefault with default headers values
func NewPostWorkflowTaskDefinitionsDefault(code int) *PostWorkflowTaskDefinitionsDefault {
	return &PostWorkflowTaskDefinitionsDefault{
		_statusCode: code,
	}
}

/*PostWorkflowTaskDefinitionsDefault handles this case with default header values.

unexpected error
*/
type PostWorkflowTaskDefinitionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post workflow task definitions default response
func (o *PostWorkflowTaskDefinitionsDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkflowTaskDefinitionsDefault) Error() string {
	return fmt.Sprintf("[POST /workflow/TaskDefinitions][%d] PostWorkflowTaskDefinitions default  %+v", o._statusCode, o.Payload)
}

func (o *PostWorkflowTaskDefinitionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWorkflowTaskDefinitionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
