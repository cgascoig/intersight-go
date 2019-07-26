// Code generated by go-swagger; DO NOT EDIT.

package equipment_system_io_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetEquipmentSystemIoControllersMoidReader is a Reader for the GetEquipmentSystemIoControllersMoid structure.
type GetEquipmentSystemIoControllersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentSystemIoControllersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentSystemIoControllersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentSystemIoControllersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentSystemIoControllersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentSystemIoControllersMoidOK creates a GetEquipmentSystemIoControllersMoidOK with default headers values
func NewGetEquipmentSystemIoControllersMoidOK() *GetEquipmentSystemIoControllersMoidOK {
	return &GetEquipmentSystemIoControllersMoidOK{}
}

/*GetEquipmentSystemIoControllersMoidOK handles this case with default header values.

An instance of equipmentSystemIoController
*/
type GetEquipmentSystemIoControllersMoidOK struct {
	Payload *models.EquipmentSystemIoController
}

func (o *GetEquipmentSystemIoControllersMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/SystemIoControllers/{moid}][%d] getEquipmentSystemIoControllersMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentSystemIoControllersMoidOK) GetPayload() *models.EquipmentSystemIoController {
	return o.Payload
}

func (o *GetEquipmentSystemIoControllersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentSystemIoController)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentSystemIoControllersMoidNotFound creates a GetEquipmentSystemIoControllersMoidNotFound with default headers values
func NewGetEquipmentSystemIoControllersMoidNotFound() *GetEquipmentSystemIoControllersMoidNotFound {
	return &GetEquipmentSystemIoControllersMoidNotFound{}
}

/*GetEquipmentSystemIoControllersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentSystemIoControllersMoidNotFound struct {
}

func (o *GetEquipmentSystemIoControllersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/SystemIoControllers/{moid}][%d] getEquipmentSystemIoControllersMoidNotFound ", 404)
}

func (o *GetEquipmentSystemIoControllersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentSystemIoControllersMoidDefault creates a GetEquipmentSystemIoControllersMoidDefault with default headers values
func NewGetEquipmentSystemIoControllersMoidDefault(code int) *GetEquipmentSystemIoControllersMoidDefault {
	return &GetEquipmentSystemIoControllersMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentSystemIoControllersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentSystemIoControllersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment system io controllers moid default response
func (o *GetEquipmentSystemIoControllersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentSystemIoControllersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/SystemIoControllers/{moid}][%d] GetEquipmentSystemIoControllersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentSystemIoControllersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentSystemIoControllersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
