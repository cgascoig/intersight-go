// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_adapter_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchVnicEthAdapterPoliciesMoidReader is a Reader for the PatchVnicEthAdapterPoliciesMoid structure.
type PatchVnicEthAdapterPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVnicEthAdapterPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchVnicEthAdapterPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchVnicEthAdapterPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchVnicEthAdapterPoliciesMoidCreated creates a PatchVnicEthAdapterPoliciesMoidCreated with default headers values
func NewPatchVnicEthAdapterPoliciesMoidCreated() *PatchVnicEthAdapterPoliciesMoidCreated {
	return &PatchVnicEthAdapterPoliciesMoidCreated{}
}

/*PatchVnicEthAdapterPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchVnicEthAdapterPoliciesMoidCreated struct {
}

func (o *PatchVnicEthAdapterPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /vnic/EthAdapterPolicies/{moid}][%d] patchVnicEthAdapterPoliciesMoidCreated ", 201)
}

func (o *PatchVnicEthAdapterPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchVnicEthAdapterPoliciesMoidDefault creates a PatchVnicEthAdapterPoliciesMoidDefault with default headers values
func NewPatchVnicEthAdapterPoliciesMoidDefault(code int) *PatchVnicEthAdapterPoliciesMoidDefault {
	return &PatchVnicEthAdapterPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchVnicEthAdapterPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchVnicEthAdapterPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch vnic eth adapter policies moid default response
func (o *PatchVnicEthAdapterPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchVnicEthAdapterPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /vnic/EthAdapterPolicies/{moid}][%d] PatchVnicEthAdapterPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchVnicEthAdapterPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchVnicEthAdapterPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
