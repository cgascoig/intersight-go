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

// DeleteWorkflowBatchAPIExecutorsMoidReader is a Reader for the DeleteWorkflowBatchAPIExecutorsMoid structure.
type DeleteWorkflowBatchAPIExecutorsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowBatchAPIExecutorsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWorkflowBatchAPIExecutorsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteWorkflowBatchAPIExecutorsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteWorkflowBatchAPIExecutorsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWorkflowBatchAPIExecutorsMoidOK creates a DeleteWorkflowBatchAPIExecutorsMoidOK with default headers values
func NewDeleteWorkflowBatchAPIExecutorsMoidOK() *DeleteWorkflowBatchAPIExecutorsMoidOK {
	return &DeleteWorkflowBatchAPIExecutorsMoidOK{}
}

/*DeleteWorkflowBatchAPIExecutorsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteWorkflowBatchAPIExecutorsMoidOK struct {
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /workflow/BatchApiExecutors/{moid}][%d] deleteWorkflowBatchApiExecutorsMoidOK ", 200)
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowBatchAPIExecutorsMoidNotFound creates a DeleteWorkflowBatchAPIExecutorsMoidNotFound with default headers values
func NewDeleteWorkflowBatchAPIExecutorsMoidNotFound() *DeleteWorkflowBatchAPIExecutorsMoidNotFound {
	return &DeleteWorkflowBatchAPIExecutorsMoidNotFound{}
}

/*DeleteWorkflowBatchAPIExecutorsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteWorkflowBatchAPIExecutorsMoidNotFound struct {
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflow/BatchApiExecutors/{moid}][%d] deleteWorkflowBatchApiExecutorsMoidNotFound ", 404)
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowBatchAPIExecutorsMoidDefault creates a DeleteWorkflowBatchAPIExecutorsMoidDefault with default headers values
func NewDeleteWorkflowBatchAPIExecutorsMoidDefault(code int) *DeleteWorkflowBatchAPIExecutorsMoidDefault {
	return &DeleteWorkflowBatchAPIExecutorsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteWorkflowBatchAPIExecutorsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteWorkflowBatchAPIExecutorsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete workflow batch API executors moid default response
func (o *DeleteWorkflowBatchAPIExecutorsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /workflow/BatchApiExecutors/{moid}][%d] DeleteWorkflowBatchAPIExecutorsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteWorkflowBatchAPIExecutorsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}