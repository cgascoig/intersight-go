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

// GetPciLinksMoidReader is a Reader for the GetPciLinksMoid structure.
type GetPciLinksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPciLinksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPciLinksMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPciLinksMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPciLinksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPciLinksMoidOK creates a GetPciLinksMoidOK with default headers values
func NewGetPciLinksMoidOK() *GetPciLinksMoidOK {
	return &GetPciLinksMoidOK{}
}

/*GetPciLinksMoidOK handles this case with default header values.

An instance of pciLink
*/
type GetPciLinksMoidOK struct {
	Payload *models.PciLink
}

func (o *GetPciLinksMoidOK) Error() string {
	return fmt.Sprintf("[GET /pci/Links/{moid}][%d] getPciLinksMoidOK  %+v", 200, o.Payload)
}

func (o *GetPciLinksMoidOK) GetPayload() *models.PciLink {
	return o.Payload
}

func (o *GetPciLinksMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PciLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPciLinksMoidNotFound creates a GetPciLinksMoidNotFound with default headers values
func NewGetPciLinksMoidNotFound() *GetPciLinksMoidNotFound {
	return &GetPciLinksMoidNotFound{}
}

/*GetPciLinksMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetPciLinksMoidNotFound struct {
}

func (o *GetPciLinksMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /pci/Links/{moid}][%d] getPciLinksMoidNotFound ", 404)
}

func (o *GetPciLinksMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPciLinksMoidDefault creates a GetPciLinksMoidDefault with default headers values
func NewGetPciLinksMoidDefault(code int) *GetPciLinksMoidDefault {
	return &GetPciLinksMoidDefault{
		_statusCode: code,
	}
}

/*GetPciLinksMoidDefault handles this case with default header values.

Unexpected error
*/
type GetPciLinksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pci links moid default response
func (o *GetPciLinksMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetPciLinksMoidDefault) Error() string {
	return fmt.Sprintf("[GET /pci/Links/{moid}][%d] GetPciLinksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetPciLinksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPciLinksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
