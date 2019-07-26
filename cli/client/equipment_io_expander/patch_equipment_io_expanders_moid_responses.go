// Code generated by go-swagger; DO NOT EDIT.

package equipment_io_expander

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchEquipmentIoExpandersMoidReader is a Reader for the PatchEquipmentIoExpandersMoid structure.
type PatchEquipmentIoExpandersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEquipmentIoExpandersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchEquipmentIoExpandersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchEquipmentIoExpandersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEquipmentIoExpandersMoidCreated creates a PatchEquipmentIoExpandersMoidCreated with default headers values
func NewPatchEquipmentIoExpandersMoidCreated() *PatchEquipmentIoExpandersMoidCreated {
	return &PatchEquipmentIoExpandersMoidCreated{}
}

/*PatchEquipmentIoExpandersMoidCreated handles this case with default header values.

Null response
*/
type PatchEquipmentIoExpandersMoidCreated struct {
}

func (o *PatchEquipmentIoExpandersMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /equipment/IoExpanders/{moid}][%d] patchEquipmentIoExpandersMoidCreated ", 201)
}

func (o *PatchEquipmentIoExpandersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEquipmentIoExpandersMoidDefault creates a PatchEquipmentIoExpandersMoidDefault with default headers values
func NewPatchEquipmentIoExpandersMoidDefault(code int) *PatchEquipmentIoExpandersMoidDefault {
	return &PatchEquipmentIoExpandersMoidDefault{
		_statusCode: code,
	}
}

/*PatchEquipmentIoExpandersMoidDefault handles this case with default header values.

unexpected error
*/
type PatchEquipmentIoExpandersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch equipment io expanders moid default response
func (o *PatchEquipmentIoExpandersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchEquipmentIoExpandersMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /equipment/IoExpanders/{moid}][%d] PatchEquipmentIoExpandersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEquipmentIoExpandersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchEquipmentIoExpandersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
