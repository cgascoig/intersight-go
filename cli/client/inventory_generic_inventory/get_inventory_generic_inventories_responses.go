// Code generated by go-swagger; DO NOT EDIT.

package inventory_generic_inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetInventoryGenericInventoriesReader is a Reader for the GetInventoryGenericInventories structure.
type GetInventoryGenericInventoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryGenericInventoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryGenericInventoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetInventoryGenericInventoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInventoryGenericInventoriesOK creates a GetInventoryGenericInventoriesOK with default headers values
func NewGetInventoryGenericInventoriesOK() *GetInventoryGenericInventoriesOK {
	return &GetInventoryGenericInventoriesOK{}
}

/*GetInventoryGenericInventoriesOK handles this case with default header values.

List of inventoryGenericInventories for the given filter criteria
*/
type GetInventoryGenericInventoriesOK struct {
	Payload *models.InventoryGenericInventoryList
}

func (o *GetInventoryGenericInventoriesOK) Error() string {
	return fmt.Sprintf("[GET /inventory/GenericInventories][%d] getInventoryGenericInventoriesOK  %+v", 200, o.Payload)
}

func (o *GetInventoryGenericInventoriesOK) GetPayload() *models.InventoryGenericInventoryList {
	return o.Payload
}

func (o *GetInventoryGenericInventoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryGenericInventoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryGenericInventoriesDefault creates a GetInventoryGenericInventoriesDefault with default headers values
func NewGetInventoryGenericInventoriesDefault(code int) *GetInventoryGenericInventoriesDefault {
	return &GetInventoryGenericInventoriesDefault{
		_statusCode: code,
	}
}

/*GetInventoryGenericInventoriesDefault handles this case with default header values.

Unexpected error
*/
type GetInventoryGenericInventoriesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get inventory generic inventories default response
func (o *GetInventoryGenericInventoriesDefault) Code() int {
	return o._statusCode
}

func (o *GetInventoryGenericInventoriesDefault) Error() string {
	return fmt.Sprintf("[GET /inventory/GenericInventories][%d] GetInventoryGenericInventories default  %+v", o._statusCode, o.Payload)
}

func (o *GetInventoryGenericInventoriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInventoryGenericInventoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
