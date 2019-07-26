// Code generated by go-swagger; DO NOT EDIT.

package workflow_workflow_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetWorkflowWorkflowDefinitionsMoidReader is a Reader for the GetWorkflowWorkflowDefinitionsMoid structure.
type GetWorkflowWorkflowDefinitionsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowWorkflowDefinitionsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowWorkflowDefinitionsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWorkflowWorkflowDefinitionsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkflowWorkflowDefinitionsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowWorkflowDefinitionsMoidOK creates a GetWorkflowWorkflowDefinitionsMoidOK with default headers values
func NewGetWorkflowWorkflowDefinitionsMoidOK() *GetWorkflowWorkflowDefinitionsMoidOK {
	return &GetWorkflowWorkflowDefinitionsMoidOK{}
}

/*GetWorkflowWorkflowDefinitionsMoidOK handles this case with default header values.

An instance of workflowWorkflowDefinition
*/
type GetWorkflowWorkflowDefinitionsMoidOK struct {
	Payload *models.WorkflowWorkflowDefinition
}

func (o *GetWorkflowWorkflowDefinitionsMoidOK) Error() string {
	return fmt.Sprintf("[GET /workflow/WorkflowDefinitions/{moid}][%d] getWorkflowWorkflowDefinitionsMoidOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowWorkflowDefinitionsMoidOK) GetPayload() *models.WorkflowWorkflowDefinition {
	return o.Payload
}

func (o *GetWorkflowWorkflowDefinitionsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowWorkflowDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowWorkflowDefinitionsMoidNotFound creates a GetWorkflowWorkflowDefinitionsMoidNotFound with default headers values
func NewGetWorkflowWorkflowDefinitionsMoidNotFound() *GetWorkflowWorkflowDefinitionsMoidNotFound {
	return &GetWorkflowWorkflowDefinitionsMoidNotFound{}
}

/*GetWorkflowWorkflowDefinitionsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetWorkflowWorkflowDefinitionsMoidNotFound struct {
}

func (o *GetWorkflowWorkflowDefinitionsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /workflow/WorkflowDefinitions/{moid}][%d] getWorkflowWorkflowDefinitionsMoidNotFound ", 404)
}

func (o *GetWorkflowWorkflowDefinitionsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowWorkflowDefinitionsMoidDefault creates a GetWorkflowWorkflowDefinitionsMoidDefault with default headers values
func NewGetWorkflowWorkflowDefinitionsMoidDefault(code int) *GetWorkflowWorkflowDefinitionsMoidDefault {
	return &GetWorkflowWorkflowDefinitionsMoidDefault{
		_statusCode: code,
	}
}

/*GetWorkflowWorkflowDefinitionsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowWorkflowDefinitionsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflow workflow definitions moid default response
func (o *GetWorkflowWorkflowDefinitionsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowWorkflowDefinitionsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /workflow/WorkflowDefinitions/{moid}][%d] GetWorkflowWorkflowDefinitionsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowWorkflowDefinitionsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkflowWorkflowDefinitionsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}