// Code generated by go-swagger; DO NOT EDIT.

package cond_hcl_status_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetCondHclStatusJobsReader is a Reader for the GetCondHclStatusJobs structure.
type GetCondHclStatusJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCondHclStatusJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCondHclStatusJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCondHclStatusJobsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCondHclStatusJobsOK creates a GetCondHclStatusJobsOK with default headers values
func NewGetCondHclStatusJobsOK() *GetCondHclStatusJobsOK {
	return &GetCondHclStatusJobsOK{}
}

/*GetCondHclStatusJobsOK handles this case with default header values.

List of condHclStatusJobs for the given filter criteria
*/
type GetCondHclStatusJobsOK struct {
	Payload *models.CondHclStatusJobList
}

func (o *GetCondHclStatusJobsOK) Error() string {
	return fmt.Sprintf("[GET /cond/HclStatusJobs][%d] getCondHclStatusJobsOK  %+v", 200, o.Payload)
}

func (o *GetCondHclStatusJobsOK) GetPayload() *models.CondHclStatusJobList {
	return o.Payload
}

func (o *GetCondHclStatusJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CondHclStatusJobList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCondHclStatusJobsDefault creates a GetCondHclStatusJobsDefault with default headers values
func NewGetCondHclStatusJobsDefault(code int) *GetCondHclStatusJobsDefault {
	return &GetCondHclStatusJobsDefault{
		_statusCode: code,
	}
}

/*GetCondHclStatusJobsDefault handles this case with default header values.

Unexpected error
*/
type GetCondHclStatusJobsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cond hcl status jobs default response
func (o *GetCondHclStatusJobsDefault) Code() int {
	return o._statusCode
}

func (o *GetCondHclStatusJobsDefault) Error() string {
	return fmt.Sprintf("[GET /cond/HclStatusJobs][%d] GetCondHclStatusJobs default  %+v", o._statusCode, o.Payload)
}

func (o *GetCondHclStatusJobsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCondHclStatusJobsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}