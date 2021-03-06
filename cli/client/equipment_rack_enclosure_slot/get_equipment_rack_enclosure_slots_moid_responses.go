// Code generated by go-swagger; DO NOT EDIT.

package equipment_rack_enclosure_slot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetEquipmentRackEnclosureSlotsMoidReader is a Reader for the GetEquipmentRackEnclosureSlotsMoid structure.
type GetEquipmentRackEnclosureSlotsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentRackEnclosureSlotsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentRackEnclosureSlotsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentRackEnclosureSlotsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentRackEnclosureSlotsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentRackEnclosureSlotsMoidOK creates a GetEquipmentRackEnclosureSlotsMoidOK with default headers values
func NewGetEquipmentRackEnclosureSlotsMoidOK() *GetEquipmentRackEnclosureSlotsMoidOK {
	return &GetEquipmentRackEnclosureSlotsMoidOK{}
}

/*GetEquipmentRackEnclosureSlotsMoidOK handles this case with default header values.

An instance of equipmentRackEnclosureSlot
*/
type GetEquipmentRackEnclosureSlotsMoidOK struct {
	Payload *models.EquipmentRackEnclosureSlot
}

func (o *GetEquipmentRackEnclosureSlotsMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/RackEnclosureSlots/{moid}][%d] getEquipmentRackEnclosureSlotsMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentRackEnclosureSlotsMoidOK) GetPayload() *models.EquipmentRackEnclosureSlot {
	return o.Payload
}

func (o *GetEquipmentRackEnclosureSlotsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentRackEnclosureSlot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentRackEnclosureSlotsMoidNotFound creates a GetEquipmentRackEnclosureSlotsMoidNotFound with default headers values
func NewGetEquipmentRackEnclosureSlotsMoidNotFound() *GetEquipmentRackEnclosureSlotsMoidNotFound {
	return &GetEquipmentRackEnclosureSlotsMoidNotFound{}
}

/*GetEquipmentRackEnclosureSlotsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentRackEnclosureSlotsMoidNotFound struct {
}

func (o *GetEquipmentRackEnclosureSlotsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/RackEnclosureSlots/{moid}][%d] getEquipmentRackEnclosureSlotsMoidNotFound ", 404)
}

func (o *GetEquipmentRackEnclosureSlotsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentRackEnclosureSlotsMoidDefault creates a GetEquipmentRackEnclosureSlotsMoidDefault with default headers values
func NewGetEquipmentRackEnclosureSlotsMoidDefault(code int) *GetEquipmentRackEnclosureSlotsMoidDefault {
	return &GetEquipmentRackEnclosureSlotsMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentRackEnclosureSlotsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentRackEnclosureSlotsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment rack enclosure slots moid default response
func (o *GetEquipmentRackEnclosureSlotsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentRackEnclosureSlotsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/RackEnclosureSlots/{moid}][%d] GetEquipmentRackEnclosureSlotsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentRackEnclosureSlotsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentRackEnclosureSlotsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
