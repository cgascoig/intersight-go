// Code generated by go-swagger; DO NOT EDIT.

package equipment_rack_enclosure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetEquipmentRackEnclosuresMoidReader is a Reader for the GetEquipmentRackEnclosuresMoid structure.
type GetEquipmentRackEnclosuresMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentRackEnclosuresMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentRackEnclosuresMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentRackEnclosuresMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentRackEnclosuresMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentRackEnclosuresMoidOK creates a GetEquipmentRackEnclosuresMoidOK with default headers values
func NewGetEquipmentRackEnclosuresMoidOK() *GetEquipmentRackEnclosuresMoidOK {
	return &GetEquipmentRackEnclosuresMoidOK{}
}

/*GetEquipmentRackEnclosuresMoidOK handles this case with default header values.

An instance of equipmentRackEnclosure
*/
type GetEquipmentRackEnclosuresMoidOK struct {
	Payload *models.EquipmentRackEnclosure
}

func (o *GetEquipmentRackEnclosuresMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/RackEnclosures/{moid}][%d] getEquipmentRackEnclosuresMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentRackEnclosuresMoidOK) GetPayload() *models.EquipmentRackEnclosure {
	return o.Payload
}

func (o *GetEquipmentRackEnclosuresMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentRackEnclosure)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentRackEnclosuresMoidNotFound creates a GetEquipmentRackEnclosuresMoidNotFound with default headers values
func NewGetEquipmentRackEnclosuresMoidNotFound() *GetEquipmentRackEnclosuresMoidNotFound {
	return &GetEquipmentRackEnclosuresMoidNotFound{}
}

/*GetEquipmentRackEnclosuresMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentRackEnclosuresMoidNotFound struct {
}

func (o *GetEquipmentRackEnclosuresMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/RackEnclosures/{moid}][%d] getEquipmentRackEnclosuresMoidNotFound ", 404)
}

func (o *GetEquipmentRackEnclosuresMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentRackEnclosuresMoidDefault creates a GetEquipmentRackEnclosuresMoidDefault with default headers values
func NewGetEquipmentRackEnclosuresMoidDefault(code int) *GetEquipmentRackEnclosuresMoidDefault {
	return &GetEquipmentRackEnclosuresMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentRackEnclosuresMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentRackEnclosuresMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment rack enclosures moid default response
func (o *GetEquipmentRackEnclosuresMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentRackEnclosuresMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/RackEnclosures/{moid}][%d] GetEquipmentRackEnclosuresMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentRackEnclosuresMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentRackEnclosuresMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
