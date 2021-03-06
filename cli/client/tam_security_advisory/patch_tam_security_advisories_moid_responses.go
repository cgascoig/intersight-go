// Code generated by go-swagger; DO NOT EDIT.

package tam_security_advisory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchTamSecurityAdvisoriesMoidReader is a Reader for the PatchTamSecurityAdvisoriesMoid structure.
type PatchTamSecurityAdvisoriesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTamSecurityAdvisoriesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchTamSecurityAdvisoriesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchTamSecurityAdvisoriesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchTamSecurityAdvisoriesMoidCreated creates a PatchTamSecurityAdvisoriesMoidCreated with default headers values
func NewPatchTamSecurityAdvisoriesMoidCreated() *PatchTamSecurityAdvisoriesMoidCreated {
	return &PatchTamSecurityAdvisoriesMoidCreated{}
}

/*PatchTamSecurityAdvisoriesMoidCreated handles this case with default header values.

Null response
*/
type PatchTamSecurityAdvisoriesMoidCreated struct {
}

func (o *PatchTamSecurityAdvisoriesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /tam/SecurityAdvisories/{moid}][%d] patchTamSecurityAdvisoriesMoidCreated ", 201)
}

func (o *PatchTamSecurityAdvisoriesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchTamSecurityAdvisoriesMoidDefault creates a PatchTamSecurityAdvisoriesMoidDefault with default headers values
func NewPatchTamSecurityAdvisoriesMoidDefault(code int) *PatchTamSecurityAdvisoriesMoidDefault {
	return &PatchTamSecurityAdvisoriesMoidDefault{
		_statusCode: code,
	}
}

/*PatchTamSecurityAdvisoriesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchTamSecurityAdvisoriesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch tam security advisories moid default response
func (o *PatchTamSecurityAdvisoriesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchTamSecurityAdvisoriesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /tam/SecurityAdvisories/{moid}][%d] PatchTamSecurityAdvisoriesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchTamSecurityAdvisoriesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchTamSecurityAdvisoriesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
