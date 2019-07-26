// Code generated by go-swagger; DO NOT EDIT.

package workflow_build_task_meta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetWorkflowBuildTaskMetaMoidReader is a Reader for the GetWorkflowBuildTaskMetaMoid structure.
type GetWorkflowBuildTaskMetaMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowBuildTaskMetaMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowBuildTaskMetaMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWorkflowBuildTaskMetaMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkflowBuildTaskMetaMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowBuildTaskMetaMoidOK creates a GetWorkflowBuildTaskMetaMoidOK with default headers values
func NewGetWorkflowBuildTaskMetaMoidOK() *GetWorkflowBuildTaskMetaMoidOK {
	return &GetWorkflowBuildTaskMetaMoidOK{}
}

/*GetWorkflowBuildTaskMetaMoidOK handles this case with default header values.

An instance of workflowBuildTaskMeta
*/
type GetWorkflowBuildTaskMetaMoidOK struct {
	Payload *models.WorkflowBuildTaskMeta
}

func (o *GetWorkflowBuildTaskMetaMoidOK) Error() string {
	return fmt.Sprintf("[GET /workflow/BuildTaskMeta/{moid}][%d] getWorkflowBuildTaskMetaMoidOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowBuildTaskMetaMoidOK) GetPayload() *models.WorkflowBuildTaskMeta {
	return o.Payload
}

func (o *GetWorkflowBuildTaskMetaMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowBuildTaskMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowBuildTaskMetaMoidNotFound creates a GetWorkflowBuildTaskMetaMoidNotFound with default headers values
func NewGetWorkflowBuildTaskMetaMoidNotFound() *GetWorkflowBuildTaskMetaMoidNotFound {
	return &GetWorkflowBuildTaskMetaMoidNotFound{}
}

/*GetWorkflowBuildTaskMetaMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetWorkflowBuildTaskMetaMoidNotFound struct {
}

func (o *GetWorkflowBuildTaskMetaMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /workflow/BuildTaskMeta/{moid}][%d] getWorkflowBuildTaskMetaMoidNotFound ", 404)
}

func (o *GetWorkflowBuildTaskMetaMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowBuildTaskMetaMoidDefault creates a GetWorkflowBuildTaskMetaMoidDefault with default headers values
func NewGetWorkflowBuildTaskMetaMoidDefault(code int) *GetWorkflowBuildTaskMetaMoidDefault {
	return &GetWorkflowBuildTaskMetaMoidDefault{
		_statusCode: code,
	}
}

/*GetWorkflowBuildTaskMetaMoidDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowBuildTaskMetaMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflow build task meta moid default response
func (o *GetWorkflowBuildTaskMetaMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowBuildTaskMetaMoidDefault) Error() string {
	return fmt.Sprintf("[GET /workflow/BuildTaskMeta/{moid}][%d] GetWorkflowBuildTaskMetaMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowBuildTaskMetaMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkflowBuildTaskMetaMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
