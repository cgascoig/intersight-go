// Code generated by go-swagger; DO NOT EDIT.

package equipment_chassis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetEquipmentChassesMoidReader is a Reader for the GetEquipmentChassesMoid structure.
type GetEquipmentChassesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentChassesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentChassesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentChassesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentChassesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentChassesMoidOK creates a GetEquipmentChassesMoidOK with default headers values
func NewGetEquipmentChassesMoidOK() *GetEquipmentChassesMoidOK {
	return &GetEquipmentChassesMoidOK{}
}

/*GetEquipmentChassesMoidOK handles this case with default header values.

An instance of equipmentChassis
*/
type GetEquipmentChassesMoidOK struct {
	Payload *models.EquipmentChassis
}

func (o *GetEquipmentChassesMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/Chasses/{moid}][%d] getEquipmentChassesMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentChassesMoidOK) GetPayload() *models.EquipmentChassis {
	return o.Payload
}

func (o *GetEquipmentChassesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentChassesMoidNotFound creates a GetEquipmentChassesMoidNotFound with default headers values
func NewGetEquipmentChassesMoidNotFound() *GetEquipmentChassesMoidNotFound {
	return &GetEquipmentChassesMoidNotFound{}
}

/*GetEquipmentChassesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentChassesMoidNotFound struct {
}

func (o *GetEquipmentChassesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/Chasses/{moid}][%d] getEquipmentChassesMoidNotFound ", 404)
}

func (o *GetEquipmentChassesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentChassesMoidDefault creates a GetEquipmentChassesMoidDefault with default headers values
func NewGetEquipmentChassesMoidDefault(code int) *GetEquipmentChassesMoidDefault {
	return &GetEquipmentChassesMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentChassesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentChassesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment chasses moid default response
func (o *GetEquipmentChassesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentChassesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/Chasses/{moid}][%d] GetEquipmentChassesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentChassesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentChassesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
