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

// GetEquipmentIoExpandersReader is a Reader for the GetEquipmentIoExpanders structure.
type GetEquipmentIoExpandersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentIoExpandersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentIoExpandersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEquipmentIoExpandersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentIoExpandersOK creates a GetEquipmentIoExpandersOK with default headers values
func NewGetEquipmentIoExpandersOK() *GetEquipmentIoExpandersOK {
	return &GetEquipmentIoExpandersOK{}
}

/*GetEquipmentIoExpandersOK handles this case with default header values.

List of equipmentIoExpanders for the given filter criteria
*/
type GetEquipmentIoExpandersOK struct {
	Payload *models.EquipmentIoExpanderList
}

func (o *GetEquipmentIoExpandersOK) Error() string {
	return fmt.Sprintf("[GET /equipment/IoExpanders][%d] getEquipmentIoExpandersOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentIoExpandersOK) GetPayload() *models.EquipmentIoExpanderList {
	return o.Payload
}

func (o *GetEquipmentIoExpandersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentIoExpanderList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentIoExpandersDefault creates a GetEquipmentIoExpandersDefault with default headers values
func NewGetEquipmentIoExpandersDefault(code int) *GetEquipmentIoExpandersDefault {
	return &GetEquipmentIoExpandersDefault{
		_statusCode: code,
	}
}

/*GetEquipmentIoExpandersDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentIoExpandersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment io expanders default response
func (o *GetEquipmentIoExpandersDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentIoExpandersDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/IoExpanders][%d] GetEquipmentIoExpanders default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentIoExpandersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentIoExpandersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
