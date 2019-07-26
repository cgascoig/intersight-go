// Code generated by go-swagger; DO NOT EDIT.

package workflow_task_meta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetWorkflowTaskMetaMoidReader is a Reader for the GetWorkflowTaskMetaMoid structure.
type GetWorkflowTaskMetaMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowTaskMetaMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowTaskMetaMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWorkflowTaskMetaMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkflowTaskMetaMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowTaskMetaMoidOK creates a GetWorkflowTaskMetaMoidOK with default headers values
func NewGetWorkflowTaskMetaMoidOK() *GetWorkflowTaskMetaMoidOK {
	return &GetWorkflowTaskMetaMoidOK{}
}

/*GetWorkflowTaskMetaMoidOK handles this case with default header values.

An instance of workflowTaskMeta
*/
type GetWorkflowTaskMetaMoidOK struct {
	Payload *models.WorkflowTaskMeta
}

func (o *GetWorkflowTaskMetaMoidOK) Error() string {
	return fmt.Sprintf("[GET /workflow/TaskMeta/{moid}][%d] getWorkflowTaskMetaMoidOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowTaskMetaMoidOK) GetPayload() *models.WorkflowTaskMeta {
	return o.Payload
}

func (o *GetWorkflowTaskMetaMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowTaskMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowTaskMetaMoidNotFound creates a GetWorkflowTaskMetaMoidNotFound with default headers values
func NewGetWorkflowTaskMetaMoidNotFound() *GetWorkflowTaskMetaMoidNotFound {
	return &GetWorkflowTaskMetaMoidNotFound{}
}

/*GetWorkflowTaskMetaMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetWorkflowTaskMetaMoidNotFound struct {
}

func (o *GetWorkflowTaskMetaMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /workflow/TaskMeta/{moid}][%d] getWorkflowTaskMetaMoidNotFound ", 404)
}

func (o *GetWorkflowTaskMetaMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowTaskMetaMoidDefault creates a GetWorkflowTaskMetaMoidDefault with default headers values
func NewGetWorkflowTaskMetaMoidDefault(code int) *GetWorkflowTaskMetaMoidDefault {
	return &GetWorkflowTaskMetaMoidDefault{
		_statusCode: code,
	}
}

/*GetWorkflowTaskMetaMoidDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowTaskMetaMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflow task meta moid default response
func (o *GetWorkflowTaskMetaMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowTaskMetaMoidDefault) Error() string {
	return fmt.Sprintf("[GET /workflow/TaskMeta/{moid}][%d] GetWorkflowTaskMetaMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowTaskMetaMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkflowTaskMetaMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}