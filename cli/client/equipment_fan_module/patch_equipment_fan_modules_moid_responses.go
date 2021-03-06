// Code generated by go-swagger; DO NOT EDIT.

package equipment_fan_module

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchEquipmentFanModulesMoidReader is a Reader for the PatchEquipmentFanModulesMoid structure.
type PatchEquipmentFanModulesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEquipmentFanModulesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchEquipmentFanModulesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchEquipmentFanModulesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEquipmentFanModulesMoidCreated creates a PatchEquipmentFanModulesMoidCreated with default headers values
func NewPatchEquipmentFanModulesMoidCreated() *PatchEquipmentFanModulesMoidCreated {
	return &PatchEquipmentFanModulesMoidCreated{}
}

/*PatchEquipmentFanModulesMoidCreated handles this case with default header values.

Null response
*/
type PatchEquipmentFanModulesMoidCreated struct {
}

func (o *PatchEquipmentFanModulesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /equipment/FanModules/{moid}][%d] patchEquipmentFanModulesMoidCreated ", 201)
}

func (o *PatchEquipmentFanModulesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEquipmentFanModulesMoidDefault creates a PatchEquipmentFanModulesMoidDefault with default headers values
func NewPatchEquipmentFanModulesMoidDefault(code int) *PatchEquipmentFanModulesMoidDefault {
	return &PatchEquipmentFanModulesMoidDefault{
		_statusCode: code,
	}
}

/*PatchEquipmentFanModulesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchEquipmentFanModulesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch equipment fan modules moid default response
func (o *PatchEquipmentFanModulesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchEquipmentFanModulesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /equipment/FanModules/{moid}][%d] PatchEquipmentFanModulesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEquipmentFanModulesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchEquipmentFanModulesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
