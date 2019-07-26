// Code generated by go-swagger; DO NOT EDIT.

package security_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchSecurityUnitsMoidReader is a Reader for the PatchSecurityUnitsMoid structure.
type PatchSecurityUnitsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSecurityUnitsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchSecurityUnitsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchSecurityUnitsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSecurityUnitsMoidCreated creates a PatchSecurityUnitsMoidCreated with default headers values
func NewPatchSecurityUnitsMoidCreated() *PatchSecurityUnitsMoidCreated {
	return &PatchSecurityUnitsMoidCreated{}
}

/*PatchSecurityUnitsMoidCreated handles this case with default header values.

Null response
*/
type PatchSecurityUnitsMoidCreated struct {
}

func (o *PatchSecurityUnitsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /security/Units/{moid}][%d] patchSecurityUnitsMoidCreated ", 201)
}

func (o *PatchSecurityUnitsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSecurityUnitsMoidDefault creates a PatchSecurityUnitsMoidDefault with default headers values
func NewPatchSecurityUnitsMoidDefault(code int) *PatchSecurityUnitsMoidDefault {
	return &PatchSecurityUnitsMoidDefault{
		_statusCode: code,
	}
}

/*PatchSecurityUnitsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchSecurityUnitsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch security units moid default response
func (o *PatchSecurityUnitsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchSecurityUnitsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/Units/{moid}][%d] PatchSecurityUnitsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSecurityUnitsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchSecurityUnitsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
