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

// DeleteWorkflowTaskDefinitionsMoidReader is a Reader for the DeleteWorkflowTaskDefinitionsMoid structure.
type DeleteWorkflowTaskDefinitionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowTaskDefinitionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWorkflowTaskDefinitionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteWorkflowTaskDefinitionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteWorkflowTaskDefinitionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWorkflowTaskDefinitionsMoidOK creates a DeleteWorkflowTaskDefinitionsMoidOK with default headers values
func NewDeleteWorkflowTaskDefinitionsMoidOK() *DeleteWorkflowTaskDefinitionsMoidOK {
	return &DeleteWorkflowTaskDefinitionsMoidOK{}
}

/*DeleteWorkflowTaskDefinitionsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteWorkflowTaskDefinitionsMoidOK struct {
}

func (o *DeleteWorkflowTaskDefinitionsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /workflow/TaskDefinitions/{moid}][%d] deleteWorkflowTaskDefinitionsMoidOK ", 200)
}

func (o *DeleteWorkflowTaskDefinitionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowTaskDefinitionsMoidNotFound creates a DeleteWorkflowTaskDefinitionsMoidNotFound with default headers values
func NewDeleteWorkflowTaskDefinitionsMoidNotFound() *DeleteWorkflowTaskDefinitionsMoidNotFound {
	return &DeleteWorkflowTaskDefinitionsMoidNotFound{}
}

/*DeleteWorkflowTaskDefinitionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteWorkflowTaskDefinitionsMoidNotFound struct {
}

func (o *DeleteWorkflowTaskDefinitionsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflow/TaskDefinitions/{moid}][%d] deleteWorkflowTaskDefinitionsMoidNotFound ", 404)
}

func (o *DeleteWorkflowTaskDefinitionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowTaskDefinitionsMoidDefault creates a DeleteWorkflowTaskDefinitionsMoidDefault with default headers values
func NewDeleteWorkflowTaskDefinitionsMoidDefault(code int) *DeleteWorkflowTaskDefinitionsMoidDefault {
	return &DeleteWorkflowTaskDefinitionsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteWorkflowTaskDefinitionsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteWorkflowTaskDefinitionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete workflow task definitions moid default response
func (o *DeleteWorkflowTaskDefinitionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWorkflowTaskDefinitionsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /workflow/TaskDefinitions/{moid}][%d] DeleteWorkflowTaskDefinitionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteWorkflowTaskDefinitionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteWorkflowTaskDefinitionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
