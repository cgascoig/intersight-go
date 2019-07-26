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

// PostWorkflowBatchAPIExecutorsReader is a Reader for the PostWorkflowBatchAPIExecutors structure.
type PostWorkflowBatchAPIExecutorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkflowBatchAPIExecutorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWorkflowBatchAPIExecutorsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWorkflowBatchAPIExecutorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWorkflowBatchAPIExecutorsCreated creates a PostWorkflowBatchAPIExecutorsCreated with default headers values
func NewPostWorkflowBatchAPIExecutorsCreated() *PostWorkflowBatchAPIExecutorsCreated {
	return &PostWorkflowBatchAPIExecutorsCreated{}
}

/*PostWorkflowBatchAPIExecutorsCreated handles this case with default header values.

Null response
*/
type PostWorkflowBatchAPIExecutorsCreated struct {
}

func (o *PostWorkflowBatchAPIExecutorsCreated) Error() string {
	return fmt.Sprintf("[POST /workflow/BatchApiExecutors][%d] postWorkflowBatchApiExecutorsCreated ", 201)
}

func (o *PostWorkflowBatchAPIExecutorsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWorkflowBatchAPIExecutorsDefault creates a PostWorkflowBatchAPIExecutorsDefault with default headers values
func NewPostWorkflowBatchAPIExecutorsDefault(code int) *PostWorkflowBatchAPIExecutorsDefault {
	return &PostWorkflowBatchAPIExecutorsDefault{
		_statusCode: code,
	}
}

/*PostWorkflowBatchAPIExecutorsDefault handles this case with default header values.

unexpected error
*/
type PostWorkflowBatchAPIExecutorsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post workflow batch API executors default response
func (o *PostWorkflowBatchAPIExecutorsDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkflowBatchAPIExecutorsDefault) Error() string {
	return fmt.Sprintf("[POST /workflow/BatchApiExecutors][%d] PostWorkflowBatchAPIExecutors default  %+v", o._statusCode, o.Payload)
}

func (o *PostWorkflowBatchAPIExecutorsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWorkflowBatchAPIExecutorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
