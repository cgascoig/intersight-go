// Code generated by go-swagger; DO NOT EDIT.

package pci_switch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetPciSwitchesReader is a Reader for the GetPciSwitches structure.
type GetPciSwitchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPciSwitchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPciSwitchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPciSwitchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPciSwitchesOK creates a GetPciSwitchesOK with default headers values
func NewGetPciSwitchesOK() *GetPciSwitchesOK {
	return &GetPciSwitchesOK{}
}

/*GetPciSwitchesOK handles this case with default header values.

List of pciSwitches for the given filter criteria
*/
type GetPciSwitchesOK struct {
	Payload *models.PciSwitchList
}

func (o *GetPciSwitchesOK) Error() string {
	return fmt.Sprintf("[GET /pci/Switches][%d] getPciSwitchesOK  %+v", 200, o.Payload)
}

func (o *GetPciSwitchesOK) GetPayload() *models.PciSwitchList {
	return o.Payload
}

func (o *GetPciSwitchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PciSwitchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPciSwitchesDefault creates a GetPciSwitchesDefault with default headers values
func NewGetPciSwitchesDefault(code int) *GetPciSwitchesDefault {
	return &GetPciSwitchesDefault{
		_statusCode: code,
	}
}

/*GetPciSwitchesDefault handles this case with default header values.

Unexpected error
*/
type GetPciSwitchesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pci switches default response
func (o *GetPciSwitchesDefault) Code() int {
	return o._statusCode
}

func (o *GetPciSwitchesDefault) Error() string {
	return fmt.Sprintf("[GET /pci/Switches][%d] GetPciSwitches default  %+v", o._statusCode, o.Payload)
}

func (o *GetPciSwitchesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPciSwitchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
