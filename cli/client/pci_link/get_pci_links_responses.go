// Code generated by go-swagger; DO NOT EDIT.

package pci_link

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetPciLinksReader is a Reader for the GetPciLinks structure.
type GetPciLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPciLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPciLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPciLinksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPciLinksOK creates a GetPciLinksOK with default headers values
func NewGetPciLinksOK() *GetPciLinksOK {
	return &GetPciLinksOK{}
}

/*GetPciLinksOK handles this case with default header values.

List of pciLinks for the given filter criteria
*/
type GetPciLinksOK struct {
	Payload *models.PciLinkList
}

func (o *GetPciLinksOK) Error() string {
	return fmt.Sprintf("[GET /pci/Links][%d] getPciLinksOK  %+v", 200, o.Payload)
}

func (o *GetPciLinksOK) GetPayload() *models.PciLinkList {
	return o.Payload
}

func (o *GetPciLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PciLinkList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPciLinksDefault creates a GetPciLinksDefault with default headers values
func NewGetPciLinksDefault(code int) *GetPciLinksDefault {
	return &GetPciLinksDefault{
		_statusCode: code,
	}
}

/*GetPciLinksDefault handles this case with default header values.

Unexpected error
*/
type GetPciLinksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pci links default response
func (o *GetPciLinksDefault) Code() int {
	return o._statusCode
}

func (o *GetPciLinksDefault) Error() string {
	return fmt.Sprintf("[GET /pci/Links][%d] GetPciLinks default  %+v", o._statusCode, o.Payload)
}

func (o *GetPciLinksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPciLinksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
