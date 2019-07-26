// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_adapter_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchVnicFcAdapterPoliciesMoidReader is a Reader for the PatchVnicFcAdapterPoliciesMoid structure.
type PatchVnicFcAdapterPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVnicFcAdapterPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchVnicFcAdapterPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchVnicFcAdapterPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchVnicFcAdapterPoliciesMoidCreated creates a PatchVnicFcAdapterPoliciesMoidCreated with default headers values
func NewPatchVnicFcAdapterPoliciesMoidCreated() *PatchVnicFcAdapterPoliciesMoidCreated {
	return &PatchVnicFcAdapterPoliciesMoidCreated{}
}

/*PatchVnicFcAdapterPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchVnicFcAdapterPoliciesMoidCreated struct {
}

func (o *PatchVnicFcAdapterPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /vnic/FcAdapterPolicies/{moid}][%d] patchVnicFcAdapterPoliciesMoidCreated ", 201)
}

func (o *PatchVnicFcAdapterPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchVnicFcAdapterPoliciesMoidDefault creates a PatchVnicFcAdapterPoliciesMoidDefault with default headers values
func NewPatchVnicFcAdapterPoliciesMoidDefault(code int) *PatchVnicFcAdapterPoliciesMoidDefault {
	return &PatchVnicFcAdapterPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchVnicFcAdapterPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchVnicFcAdapterPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch vnic fc adapter policies moid default response
func (o *PatchVnicFcAdapterPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchVnicFcAdapterPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /vnic/FcAdapterPolicies/{moid}][%d] PatchVnicFcAdapterPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchVnicFcAdapterPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchVnicFcAdapterPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}