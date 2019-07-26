// Code generated by go-swagger; DO NOT EDIT.

package ether_physical_port

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchEtherPhysicalPortsMoidReader is a Reader for the PatchEtherPhysicalPortsMoid structure.
type PatchEtherPhysicalPortsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEtherPhysicalPortsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchEtherPhysicalPortsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchEtherPhysicalPortsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEtherPhysicalPortsMoidCreated creates a PatchEtherPhysicalPortsMoidCreated with default headers values
func NewPatchEtherPhysicalPortsMoidCreated() *PatchEtherPhysicalPortsMoidCreated {
	return &PatchEtherPhysicalPortsMoidCreated{}
}

/*PatchEtherPhysicalPortsMoidCreated handles this case with default header values.

Null response
*/
type PatchEtherPhysicalPortsMoidCreated struct {
}

func (o *PatchEtherPhysicalPortsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /ether/PhysicalPorts/{moid}][%d] patchEtherPhysicalPortsMoidCreated ", 201)
}

func (o *PatchEtherPhysicalPortsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEtherPhysicalPortsMoidDefault creates a PatchEtherPhysicalPortsMoidDefault with default headers values
func NewPatchEtherPhysicalPortsMoidDefault(code int) *PatchEtherPhysicalPortsMoidDefault {
	return &PatchEtherPhysicalPortsMoidDefault{
		_statusCode: code,
	}
}

/*PatchEtherPhysicalPortsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchEtherPhysicalPortsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch ether physical ports moid default response
func (o *PatchEtherPhysicalPortsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchEtherPhysicalPortsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /ether/PhysicalPorts/{moid}][%d] PatchEtherPhysicalPortsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEtherPhysicalPortsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchEtherPhysicalPortsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}