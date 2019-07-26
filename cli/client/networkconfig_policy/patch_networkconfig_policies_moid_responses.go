// Code generated by go-swagger; DO NOT EDIT.

package networkconfig_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchNetworkconfigPoliciesMoidReader is a Reader for the PatchNetworkconfigPoliciesMoid structure.
type PatchNetworkconfigPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchNetworkconfigPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchNetworkconfigPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchNetworkconfigPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchNetworkconfigPoliciesMoidCreated creates a PatchNetworkconfigPoliciesMoidCreated with default headers values
func NewPatchNetworkconfigPoliciesMoidCreated() *PatchNetworkconfigPoliciesMoidCreated {
	return &PatchNetworkconfigPoliciesMoidCreated{}
}

/*PatchNetworkconfigPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchNetworkconfigPoliciesMoidCreated struct {
}

func (o *PatchNetworkconfigPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /networkconfig/Policies/{moid}][%d] patchNetworkconfigPoliciesMoidCreated ", 201)
}

func (o *PatchNetworkconfigPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchNetworkconfigPoliciesMoidDefault creates a PatchNetworkconfigPoliciesMoidDefault with default headers values
func NewPatchNetworkconfigPoliciesMoidDefault(code int) *PatchNetworkconfigPoliciesMoidDefault {
	return &PatchNetworkconfigPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchNetworkconfigPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchNetworkconfigPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch networkconfig policies moid default response
func (o *PatchNetworkconfigPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchNetworkconfigPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /networkconfig/Policies/{moid}][%d] PatchNetworkconfigPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchNetworkconfigPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchNetworkconfigPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
