// Code generated by go-swagger; DO NOT EDIT.

package fc_physical_port

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchFcPhysicalPortsMoidReader is a Reader for the PatchFcPhysicalPortsMoid structure.
type PatchFcPhysicalPortsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchFcPhysicalPortsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchFcPhysicalPortsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchFcPhysicalPortsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchFcPhysicalPortsMoidCreated creates a PatchFcPhysicalPortsMoidCreated with default headers values
func NewPatchFcPhysicalPortsMoidCreated() *PatchFcPhysicalPortsMoidCreated {
	return &PatchFcPhysicalPortsMoidCreated{}
}

/*PatchFcPhysicalPortsMoidCreated handles this case with default header values.

Null response
*/
type PatchFcPhysicalPortsMoidCreated struct {
}

func (o *PatchFcPhysicalPortsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /fc/PhysicalPorts/{moid}][%d] patchFcPhysicalPortsMoidCreated ", 201)
}

func (o *PatchFcPhysicalPortsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchFcPhysicalPortsMoidDefault creates a PatchFcPhysicalPortsMoidDefault with default headers values
func NewPatchFcPhysicalPortsMoidDefault(code int) *PatchFcPhysicalPortsMoidDefault {
	return &PatchFcPhysicalPortsMoidDefault{
		_statusCode: code,
	}
}

/*PatchFcPhysicalPortsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchFcPhysicalPortsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch fc physical ports moid default response
func (o *PatchFcPhysicalPortsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchFcPhysicalPortsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /fc/PhysicalPorts/{moid}][%d] PatchFcPhysicalPortsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchFcPhysicalPortsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchFcPhysicalPortsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}