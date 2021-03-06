// Code generated by go-swagger; DO NOT EDIT.

package equipment_psu

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetEquipmentPsusMoidReader is a Reader for the GetEquipmentPsusMoid structure.
type GetEquipmentPsusMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentPsusMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentPsusMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentPsusMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentPsusMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentPsusMoidOK creates a GetEquipmentPsusMoidOK with default headers values
func NewGetEquipmentPsusMoidOK() *GetEquipmentPsusMoidOK {
	return &GetEquipmentPsusMoidOK{}
}

/*GetEquipmentPsusMoidOK handles this case with default header values.

An instance of equipmentPsu
*/
type GetEquipmentPsusMoidOK struct {
	Payload *models.EquipmentPsu
}

func (o *GetEquipmentPsusMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/Psus/{moid}][%d] getEquipmentPsusMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentPsusMoidOK) GetPayload() *models.EquipmentPsu {
	return o.Payload
}

func (o *GetEquipmentPsusMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentPsu)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentPsusMoidNotFound creates a GetEquipmentPsusMoidNotFound with default headers values
func NewGetEquipmentPsusMoidNotFound() *GetEquipmentPsusMoidNotFound {
	return &GetEquipmentPsusMoidNotFound{}
}

/*GetEquipmentPsusMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentPsusMoidNotFound struct {
}

func (o *GetEquipmentPsusMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/Psus/{moid}][%d] getEquipmentPsusMoidNotFound ", 404)
}

func (o *GetEquipmentPsusMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentPsusMoidDefault creates a GetEquipmentPsusMoidDefault with default headers values
func NewGetEquipmentPsusMoidDefault(code int) *GetEquipmentPsusMoidDefault {
	return &GetEquipmentPsusMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentPsusMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentPsusMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment psus moid default response
func (o *GetEquipmentPsusMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentPsusMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/Psus/{moid}][%d] GetEquipmentPsusMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentPsusMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentPsusMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
