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

// PatchPciLinksMoidReader is a Reader for the PatchPciLinksMoid structure.
type PatchPciLinksMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPciLinksMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchPciLinksMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchPciLinksMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchPciLinksMoidCreated creates a PatchPciLinksMoidCreated with default headers values
func NewPatchPciLinksMoidCreated() *PatchPciLinksMoidCreated {
	return &PatchPciLinksMoidCreated{}
}

/*PatchPciLinksMoidCreated handles this case with default header values.

Null response
*/
type PatchPciLinksMoidCreated struct {
}

func (o *PatchPciLinksMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /pci/Links/{moid}][%d] patchPciLinksMoidCreated ", 201)
}

func (o *PatchPciLinksMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPciLinksMoidDefault creates a PatchPciLinksMoidDefault with default headers values
func NewPatchPciLinksMoidDefault(code int) *PatchPciLinksMoidDefault {
	return &PatchPciLinksMoidDefault{
		_statusCode: code,
	}
}

/*PatchPciLinksMoidDefault handles this case with default header values.

unexpected error
*/
type PatchPciLinksMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch pci links moid default response
func (o *PatchPciLinksMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchPciLinksMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /pci/Links/{moid}][%d] PatchPciLinksMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchPciLinksMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchPciLinksMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
