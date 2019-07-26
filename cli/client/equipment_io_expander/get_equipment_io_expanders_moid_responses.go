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

// GetEquipmentIoExpandersMoidReader is a Reader for the GetEquipmentIoExpandersMoid structure.
type GetEquipmentIoExpandersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentIoExpandersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentIoExpandersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentIoExpandersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentIoExpandersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentIoExpandersMoidOK creates a GetEquipmentIoExpandersMoidOK with default headers values
func NewGetEquipmentIoExpandersMoidOK() *GetEquipmentIoExpandersMoidOK {
	return &GetEquipmentIoExpandersMoidOK{}
}

/*GetEquipmentIoExpandersMoidOK handles this case with default header values.

An instance of equipmentIoExpander
*/
type GetEquipmentIoExpandersMoidOK struct {
	Payload *models.EquipmentIoExpander
}

func (o *GetEquipmentIoExpandersMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/IoExpanders/{moid}][%d] getEquipmentIoExpandersMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentIoExpandersMoidOK) GetPayload() *models.EquipmentIoExpander {
	return o.Payload
}

func (o *GetEquipmentIoExpandersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentIoExpander)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentIoExpandersMoidNotFound creates a GetEquipmentIoExpandersMoidNotFound with default headers values
func NewGetEquipmentIoExpandersMoidNotFound() *GetEquipmentIoExpandersMoidNotFound {
	return &GetEquipmentIoExpandersMoidNotFound{}
}

/*GetEquipmentIoExpandersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentIoExpandersMoidNotFound struct {
}

func (o *GetEquipmentIoExpandersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/IoExpanders/{moid}][%d] getEquipmentIoExpandersMoidNotFound ", 404)
}

func (o *GetEquipmentIoExpandersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentIoExpandersMoidDefault creates a GetEquipmentIoExpandersMoidDefault with default headers values
func NewGetEquipmentIoExpandersMoidDefault(code int) *GetEquipmentIoExpandersMoidDefault {
	return &GetEquipmentIoExpandersMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentIoExpandersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentIoExpandersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment io expanders moid default response
func (o *GetEquipmentIoExpandersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentIoExpandersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/IoExpanders/{moid}][%d] GetEquipmentIoExpandersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentIoExpandersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentIoExpandersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}