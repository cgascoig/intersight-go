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

// GetEquipmentPsusReader is a Reader for the GetEquipmentPsus structure.
type GetEquipmentPsusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentPsusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentPsusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEquipmentPsusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentPsusOK creates a GetEquipmentPsusOK with default headers values
func NewGetEquipmentPsusOK() *GetEquipmentPsusOK {
	return &GetEquipmentPsusOK{}
}

/*GetEquipmentPsusOK handles this case with default header values.

List of equipmentPsus for the given filter criteria
*/
type GetEquipmentPsusOK struct {
	Payload *models.EquipmentPsuList
}

func (o *GetEquipmentPsusOK) Error() string {
	return fmt.Sprintf("[GET /equipment/Psus][%d] getEquipmentPsusOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentPsusOK) GetPayload() *models.EquipmentPsuList {
	return o.Payload
}

func (o *GetEquipmentPsusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentPsuList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentPsusDefault creates a GetEquipmentPsusDefault with default headers values
func NewGetEquipmentPsusDefault(code int) *GetEquipmentPsusDefault {
	return &GetEquipmentPsusDefault{
		_statusCode: code,
	}
}

/*GetEquipmentPsusDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentPsusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment psus default response
func (o *GetEquipmentPsusDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentPsusDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/Psus][%d] GetEquipmentPsus default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentPsusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentPsusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
