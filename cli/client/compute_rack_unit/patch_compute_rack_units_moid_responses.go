// Code generated by go-swagger; DO NOT EDIT.

package compute_rack_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchComputeRackUnitsMoidReader is a Reader for the PatchComputeRackUnitsMoid structure.
type PatchComputeRackUnitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchComputeRackUnitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchComputeRackUnitsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchComputeRackUnitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchComputeRackUnitsMoidCreated creates a PatchComputeRackUnitsMoidCreated with default headers values
func NewPatchComputeRackUnitsMoidCreated() *PatchComputeRackUnitsMoidCreated {
	return &PatchComputeRackUnitsMoidCreated{}
}

/*PatchComputeRackUnitsMoidCreated handles this case with default header values.

Null response
*/
type PatchComputeRackUnitsMoidCreated struct {
}

func (o *PatchComputeRackUnitsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /compute/RackUnits/{moid}][%d] patchComputeRackUnitsMoidCreated ", 201)
}

func (o *PatchComputeRackUnitsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchComputeRackUnitsMoidDefault creates a PatchComputeRackUnitsMoidDefault with default headers values
func NewPatchComputeRackUnitsMoidDefault(code int) *PatchComputeRackUnitsMoidDefault {
	return &PatchComputeRackUnitsMoidDefault{
		_statusCode: code,
	}
}

/*PatchComputeRackUnitsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchComputeRackUnitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch compute rack units moid default response
func (o *PatchComputeRackUnitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchComputeRackUnitsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /compute/RackUnits/{moid}][%d] PatchComputeRackUnitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchComputeRackUnitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchComputeRackUnitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
