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

// PatchCvdDeploymentTasksMoidReader is a Reader for the PatchCvdDeploymentTasksMoid structure.
type PatchCvdDeploymentTasksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCvdDeploymentTasksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchCvdDeploymentTasksMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchCvdDeploymentTasksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCvdDeploymentTasksMoidCreated creates a PatchCvdDeploymentTasksMoidCreated with default headers values
func NewPatchCvdDeploymentTasksMoidCreated() *PatchCvdDeploymentTasksMoidCreated {
	return &PatchCvdDeploymentTasksMoidCreated{}
}

/*PatchCvdDeploymentTasksMoidCreated handles this case with default header values.

Null response
*/
type PatchCvdDeploymentTasksMoidCreated struct {
}

func (o *PatchCvdDeploymentTasksMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /cvd/DeploymentTasks/{moid}][%d] patchCvdDeploymentTasksMoidCreated ", 201)
}

func (o *PatchCvdDeploymentTasksMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCvdDeploymentTasksMoidDefault creates a PatchCvdDeploymentTasksMoidDefault with default headers values
func NewPatchCvdDeploymentTasksMoidDefault(code int) *PatchCvdDeploymentTasksMoidDefault {
	return &PatchCvdDeploymentTasksMoidDefault{
		_statusCode: code,
	}
}

/*PatchCvdDeploymentTasksMoidDefault handles this case with default header values.

unexpected error
*/
type PatchCvdDeploymentTasksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch cvd deployment tasks moid default response
func (o *PatchCvdDeploymentTasksMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchCvdDeploymentTasksMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /cvd/DeploymentTasks/{moid}][%d] PatchCvdDeploymentTasksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCvdDeploymentTasksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchCvdDeploymentTasksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}