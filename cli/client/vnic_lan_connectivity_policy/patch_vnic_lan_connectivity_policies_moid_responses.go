// Code generated by go-swagger; DO NOT EDIT.

package vnic_lan_connectivity_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchVnicLanConnectivityPoliciesMoidReader is a Reader for the PatchVnicLanConnectivityPoliciesMoid structure.
type PatchVnicLanConnectivityPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVnicLanConnectivityPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchVnicLanConnectivityPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchVnicLanConnectivityPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchVnicLanConnectivityPoliciesMoidCreated creates a PatchVnicLanConnectivityPoliciesMoidCreated with default headers values
func NewPatchVnicLanConnectivityPoliciesMoidCreated() *PatchVnicLanConnectivityPoliciesMoidCreated {
	return &PatchVnicLanConnectivityPoliciesMoidCreated{}
}

/*PatchVnicLanConnectivityPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchVnicLanConnectivityPoliciesMoidCreated struct {
}

func (o *PatchVnicLanConnectivityPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /vnic/LanConnectivityPolicies/{moid}][%d] patchVnicLanConnectivityPoliciesMoidCreated ", 201)
}

func (o *PatchVnicLanConnectivityPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchVnicLanConnectivityPoliciesMoidDefault creates a PatchVnicLanConnectivityPoliciesMoidDefault with default headers values
func NewPatchVnicLanConnectivityPoliciesMoidDefault(code int) *PatchVnicLanConnectivityPoliciesMoidDefault {
	return &PatchVnicLanConnectivityPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchVnicLanConnectivityPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchVnicLanConnectivityPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch vnic lan connectivity policies moid default response
func (o *PatchVnicLanConnectivityPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchVnicLanConnectivityPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /vnic/LanConnectivityPolicies/{moid}][%d] PatchVnicLanConnectivityPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchVnicLanConnectivityPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchVnicLanConnectivityPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
