// Code generated by go-swagger; DO NOT EDIT.

package cvd_deployment_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetCvdDeploymentTasksMoidReader is a Reader for the GetCvdDeploymentTasksMoid structure.
type GetCvdDeploymentTasksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCvdDeploymentTasksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCvdDeploymentTasksMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCvdDeploymentTasksMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCvdDeploymentTasksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCvdDeploymentTasksMoidOK creates a GetCvdDeploymentTasksMoidOK with default headers values
func NewGetCvdDeploymentTasksMoidOK() *GetCvdDeploymentTasksMoidOK {
	return &GetCvdDeploymentTasksMoidOK{}
}

/*GetCvdDeploymentTasksMoidOK handles this case with default header values.

An instance of cvdDeploymentTask
*/
type GetCvdDeploymentTasksMoidOK struct {
	Payload *models.CvdDeploymentTask
}

func (o *GetCvdDeploymentTasksMoidOK) Error() string {
	return fmt.Sprintf("[GET /cvd/DeploymentTasks/{moid}][%d] getCvdDeploymentTasksMoidOK  %+v", 200, o.Payload)
}

func (o *GetCvdDeploymentTasksMoidOK) GetPayload() *models.CvdDeploymentTask {
	return o.Payload
}

func (o *GetCvdDeploymentTasksMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CvdDeploymentTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCvdDeploymentTasksMoidNotFound creates a GetCvdDeploymentTasksMoidNotFound with default headers values
func NewGetCvdDeploymentTasksMoidNotFound() *GetCvdDeploymentTasksMoidNotFound {
	return &GetCvdDeploymentTasksMoidNotFound{}
}

/*GetCvdDeploymentTasksMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetCvdDeploymentTasksMoidNotFound struct {
}

func (o *GetCvdDeploymentTasksMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /cvd/DeploymentTasks/{moid}][%d] getCvdDeploymentTasksMoidNotFound ", 404)
}

func (o *GetCvdDeploymentTasksMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCvdDeploymentTasksMoidDefault creates a GetCvdDeploymentTasksMoidDefault with default headers values
func NewGetCvdDeploymentTasksMoidDefault(code int) *GetCvdDeploymentTasksMoidDefault {
	return &GetCvdDeploymentTasksMoidDefault{
		_statusCode: code,
	}
}

/*GetCvdDeploymentTasksMoidDefault handles this case with default header values.

Unexpected error
*/
type GetCvdDeploymentTasksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cvd deployment tasks moid default response
func (o *GetCvdDeploymentTasksMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetCvdDeploymentTasksMoidDefault) Error() string {
	return fmt.Sprintf("[GET /cvd/DeploymentTasks/{moid}][%d] GetCvdDeploymentTasksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCvdDeploymentTasksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCvdDeploymentTasksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}