// Code generated by go-swagger; DO NOT EDIT.

package workflow_batch_api_executor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchWorkflowBatchAPIExecutorsMoidReader is a Reader for the PatchWorkflowBatchAPIExecutorsMoid structure.
type PatchWorkflowBatchAPIExecutorsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkflowBatchAPIExecutorsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchWorkflowBatchAPIExecutorsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchWorkflowBatchAPIExecutorsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWorkflowBatchAPIExecutorsMoidCreated creates a PatchWorkflowBatchAPIExecutorsMoidCreated with default headers values
func NewPatchWorkflowBatchAPIExecutorsMoidCreated() *PatchWorkflowBatchAPIExecutorsMoidCreated {
	return &PatchWorkflowBatchAPIExecutorsMoidCreated{}
}

/*PatchWorkflowBatchAPIExecutorsMoidCreated handles this case with default header values.

Null response
*/
type PatchWorkflowBatchAPIExecutorsMoidCreated struct {
}

func (o *PatchWorkflowBatchAPIExecutorsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /workflow/BatchApiExecutors/{moid}][%d] patchWorkflowBatchApiExecutorsMoidCreated ", 201)
}

func (o *PatchWorkflowBatchAPIExecutorsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchWorkflowBatchAPIExecutorsMoidDefault creates a PatchWorkflowBatchAPIExecutorsMoidDefault with default headers values
func NewPatchWorkflowBatchAPIExecutorsMoidDefault(code int) *PatchWorkflowBatchAPIExecutorsMoidDefault {
	return &PatchWorkflowBatchAPIExecutorsMoidDefault{
		_statusCode: code,
	}
}

/*PatchWorkflowBatchAPIExecutorsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchWorkflowBatchAPIExecutorsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch workflow batch API executors moid default response
func (o *PatchWorkflowBatchAPIExecutorsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchWorkflowBatchAPIExecutorsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /workflow/BatchApiExecutors/{moid}][%d] PatchWorkflowBatchAPIExecutorsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchWorkflowBatchAPIExecutorsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchWorkflowBatchAPIExecutorsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
