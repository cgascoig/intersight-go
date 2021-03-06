// Code generated by go-swagger; DO NOT EDIT.

package workflow_pending_dynamic_workflow_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetWorkflowPendingDynamicWorkflowInfosReader is a Reader for the GetWorkflowPendingDynamicWorkflowInfos structure.
type GetWorkflowPendingDynamicWorkflowInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowPendingDynamicWorkflowInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowPendingDynamicWorkflowInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWorkflowPendingDynamicWorkflowInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowPendingDynamicWorkflowInfosOK creates a GetWorkflowPendingDynamicWorkflowInfosOK with default headers values
func NewGetWorkflowPendingDynamicWorkflowInfosOK() *GetWorkflowPendingDynamicWorkflowInfosOK {
	return &GetWorkflowPendingDynamicWorkflowInfosOK{}
}

/*GetWorkflowPendingDynamicWorkflowInfosOK handles this case with default header values.

List of workflowPendingDynamicWorkflowInfos for the given filter criteria
*/
type GetWorkflowPendingDynamicWorkflowInfosOK struct {
	Payload *models.WorkflowPendingDynamicWorkflowInfoList
}

func (o *GetWorkflowPendingDynamicWorkflowInfosOK) Error() string {
	return fmt.Sprintf("[GET /workflow/PendingDynamicWorkflowInfos][%d] getWorkflowPendingDynamicWorkflowInfosOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowPendingDynamicWorkflowInfosOK) GetPayload() *models.WorkflowPendingDynamicWorkflowInfoList {
	return o.Payload
}

func (o *GetWorkflowPendingDynamicWorkflowInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowPendingDynamicWorkflowInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowPendingDynamicWorkflowInfosDefault creates a GetWorkflowPendingDynamicWorkflowInfosDefault with default headers values
func NewGetWorkflowPendingDynamicWorkflowInfosDefault(code int) *GetWorkflowPendingDynamicWorkflowInfosDefault {
	return &GetWorkflowPendingDynamicWorkflowInfosDefault{
		_statusCode: code,
	}
}

/*GetWorkflowPendingDynamicWorkflowInfosDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowPendingDynamicWorkflowInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflow pending dynamic workflow infos default response
func (o *GetWorkflowPendingDynamicWorkflowInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowPendingDynamicWorkflowInfosDefault) Error() string {
	return fmt.Sprintf("[GET /workflow/PendingDynamicWorkflowInfos][%d] GetWorkflowPendingDynamicWorkflowInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowPendingDynamicWorkflowInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkflowPendingDynamicWorkflowInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
