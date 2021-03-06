// Code generated by go-swagger; DO NOT EDIT.

package equipment_device_summary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetEquipmentDeviceSummariesMoidReader is a Reader for the GetEquipmentDeviceSummariesMoid structure.
type GetEquipmentDeviceSummariesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEquipmentDeviceSummariesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEquipmentDeviceSummariesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEquipmentDeviceSummariesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEquipmentDeviceSummariesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEquipmentDeviceSummariesMoidOK creates a GetEquipmentDeviceSummariesMoidOK with default headers values
func NewGetEquipmentDeviceSummariesMoidOK() *GetEquipmentDeviceSummariesMoidOK {
	return &GetEquipmentDeviceSummariesMoidOK{}
}

/*GetEquipmentDeviceSummariesMoidOK handles this case with default header values.

An instance of equipmentDeviceSummary
*/
type GetEquipmentDeviceSummariesMoidOK struct {
	Payload *models.EquipmentDeviceSummary
}

func (o *GetEquipmentDeviceSummariesMoidOK) Error() string {
	return fmt.Sprintf("[GET /equipment/DeviceSummaries/{moid}][%d] getEquipmentDeviceSummariesMoidOK  %+v", 200, o.Payload)
}

func (o *GetEquipmentDeviceSummariesMoidOK) GetPayload() *models.EquipmentDeviceSummary {
	return o.Payload
}

func (o *GetEquipmentDeviceSummariesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EquipmentDeviceSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEquipmentDeviceSummariesMoidNotFound creates a GetEquipmentDeviceSummariesMoidNotFound with default headers values
func NewGetEquipmentDeviceSummariesMoidNotFound() *GetEquipmentDeviceSummariesMoidNotFound {
	return &GetEquipmentDeviceSummariesMoidNotFound{}
}

/*GetEquipmentDeviceSummariesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetEquipmentDeviceSummariesMoidNotFound struct {
}

func (o *GetEquipmentDeviceSummariesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /equipment/DeviceSummaries/{moid}][%d] getEquipmentDeviceSummariesMoidNotFound ", 404)
}

func (o *GetEquipmentDeviceSummariesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEquipmentDeviceSummariesMoidDefault creates a GetEquipmentDeviceSummariesMoidDefault with default headers values
func NewGetEquipmentDeviceSummariesMoidDefault(code int) *GetEquipmentDeviceSummariesMoidDefault {
	return &GetEquipmentDeviceSummariesMoidDefault{
		_statusCode: code,
	}
}

/*GetEquipmentDeviceSummariesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetEquipmentDeviceSummariesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get equipment device summaries moid default response
func (o *GetEquipmentDeviceSummariesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetEquipmentDeviceSummariesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /equipment/DeviceSummaries/{moid}][%d] GetEquipmentDeviceSummariesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetEquipmentDeviceSummariesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEquipmentDeviceSummariesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
