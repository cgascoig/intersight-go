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

// PatchWorkflowTaskInfosMoidReader is a Reader for the PatchWorkflowTaskInfosMoid structure.
type PatchWorkflowTaskInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkflowTaskInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchWorkflowTaskInfosMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchWorkflowTaskInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWorkflowTaskInfosMoidCreated creates a PatchWorkflowTaskInfosMoidCreated with default headers values
func NewPatchWorkflowTaskInfosMoidCreated() *PatchWorkflowTaskInfosMoidCreated {
	return &PatchWorkflowTaskInfosMoidCreated{}
}

/*PatchWorkflowTaskInfosMoidCreated handles this case with default header values.

Null response
*/
type PatchWorkflowTaskInfosMoidCreated struct {
}

func (o *PatchWorkflowTaskInfosMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /workflow/TaskInfos/{moid}][%d] patchWorkflowTaskInfosMoidCreated ", 201)
}

func (o *PatchWorkflowTaskInfosMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchWorkflowTaskInfosMoidDefault creates a PatchWorkflowTaskInfosMoidDefault with default headers values
func NewPatchWorkflowTaskInfosMoidDefault(code int) *PatchWorkflowTaskInfosMoidDefault {
	return &PatchWorkflowTaskInfosMoidDefault{
		_statusCode: code,
	}
}

/*PatchWorkflowTaskInfosMoidDefault handles this case with default header values.

unexpected error
*/
type PatchWorkflowTaskInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch workflow task infos moid default response
func (o *PatchWorkflowTaskInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchWorkflowTaskInfosMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /workflow/TaskInfos/{moid}][%d] PatchWorkflowTaskInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchWorkflowTaskInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchWorkflowTaskInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
