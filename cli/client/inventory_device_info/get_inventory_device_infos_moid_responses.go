// Code generated by go-swagger; DO NOT EDIT.

package inventory_device_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetInventoryDeviceInfosMoidReader is a Reader for the GetInventoryDeviceInfosMoid structure.
type GetInventoryDeviceInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryDeviceInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryDeviceInfosMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetInventoryDeviceInfosMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetInventoryDeviceInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInventoryDeviceInfosMoidOK creates a GetInventoryDeviceInfosMoidOK with default headers values
func NewGetInventoryDeviceInfosMoidOK() *GetInventoryDeviceInfosMoidOK {
	return &GetInventoryDeviceInfosMoidOK{}
}

/*GetInventoryDeviceInfosMoidOK handles this case with default header values.

An instance of inventoryDeviceInfo
*/
type GetInventoryDeviceInfosMoidOK struct {
	Payload *models.InventoryDeviceInfo
}

func (o *GetInventoryDeviceInfosMoidOK) Error() string {
	return fmt.Sprintf("[GET /inventory/DeviceInfos/{moid}][%d] getInventoryDeviceInfosMoidOK  %+v", 200, o.Payload)
}

func (o *GetInventoryDeviceInfosMoidOK) GetPayload() *models.InventoryDeviceInfo {
	return o.Payload
}

func (o *GetInventoryDeviceInfosMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryDeviceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryDeviceInfosMoidNotFound creates a GetInventoryDeviceInfosMoidNotFound with default headers values
func NewGetInventoryDeviceInfosMoidNotFound() *GetInventoryDeviceInfosMoidNotFound {
	return &GetInventoryDeviceInfosMoidNotFound{}
}

/*GetInventoryDeviceInfosMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetInventoryDeviceInfosMoidNotFound struct {
}

func (o *GetInventoryDeviceInfosMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /inventory/DeviceInfos/{moid}][%d] getInventoryDeviceInfosMoidNotFound ", 404)
}

func (o *GetInventoryDeviceInfosMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryDeviceInfosMoidDefault creates a GetInventoryDeviceInfosMoidDefault with default headers values
func NewGetInventoryDeviceInfosMoidDefault(code int) *GetInventoryDeviceInfosMoidDefault {
	return &GetInventoryDeviceInfosMoidDefault{
		_statusCode: code,
	}
}

/*GetInventoryDeviceInfosMoidDefault handles this case with default header values.

Unexpected error
*/
type GetInventoryDeviceInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get inventory device infos moid default response
func (o *GetInventoryDeviceInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetInventoryDeviceInfosMoidDefault) Error() string {
	return fmt.Sprintf("[GET /inventory/DeviceInfos/{moid}][%d] GetInventoryDeviceInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetInventoryDeviceInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInventoryDeviceInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
