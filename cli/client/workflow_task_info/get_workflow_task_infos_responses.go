// Code generated by go-swagger; DO NOT EDIT.

package workflow_task_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetWorkflowTaskInfosReader is a Reader for the GetWorkflowTaskInfos structure.
type GetWorkflowTaskInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowTaskInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowTaskInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWorkflowTaskInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowTaskInfosOK creates a GetWorkflowTaskInfosOK with default headers values
func NewGetWorkflowTaskInfosOK() *GetWorkflowTaskInfosOK {
	return &GetWorkflowTaskInfosOK{}
}

/*GetWorkflowTaskInfosOK handles this case with default header values.

List of workflowTaskInfos for the given filter criteria
*/
type GetWorkflowTaskInfosOK struct {
	Payload *models.WorkflowTaskInfoList
}

func (o *GetWorkflowTaskInfosOK) Error() string {
	return fmt.Sprintf("[GET /workflow/TaskInfos][%d] getWorkflowTaskInfosOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowTaskInfosOK) GetPayload() *models.WorkflowTaskInfoList {
	return o.Payload
}

func (o *GetWorkflowTaskInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowTaskInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowTaskInfosDefault creates a GetWorkflowTaskInfosDefault with default headers values
func NewGetWorkflowTaskInfosDefault(code int) *GetWorkflowTaskInfosDefault {
	return &GetWorkflowTaskInfosDefault{
		_statusCode: code,
	}
}

/*GetWorkflowTaskInfosDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowTaskInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflow task infos default response
func (o *GetWorkflowTaskInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowTaskInfosDefault) Error() string {
	return fmt.Sprintf("[GET /workflow/TaskInfos][%d] GetWorkflowTaskInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowTaskInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkflowTaskInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}