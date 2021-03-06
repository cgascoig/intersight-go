// Code generated by go-swagger; DO NOT EDIT.

package hcl_hyperflex_software_compatibility_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchHclHyperflexSoftwareCompatibilityInfosMoidReader is a Reader for the PatchHclHyperflexSoftwareCompatibilityInfosMoid structure.
type PatchHclHyperflexSoftwareCompatibilityInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchHclHyperflexSoftwareCompatibilityInfosMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchHclHyperflexSoftwareCompatibilityInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchHclHyperflexSoftwareCompatibilityInfosMoidCreated creates a PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated with default headers values
func NewPatchHclHyperflexSoftwareCompatibilityInfosMoidCreated() *PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated {
	return &PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated{}
}

/*PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated handles this case with default header values.

Null response
*/
type PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated struct {
}

func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /hcl/HyperflexSoftwareCompatibilityInfos/{moid}][%d] patchHclHyperflexSoftwareCompatibilityInfosMoidCreated ", 201)
}

func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchHclHyperflexSoftwareCompatibilityInfosMoidDefault creates a PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault with default headers values
func NewPatchHclHyperflexSoftwareCompatibilityInfosMoidDefault(code int) *PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault {
	return &PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault{
		_statusCode: code,
	}
}

/*PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault handles this case with default header values.

unexpected error
*/
type PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch hcl hyperflex software compatibility infos moid default response
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /hcl/HyperflexSoftwareCompatibilityInfos/{moid}][%d] PatchHclHyperflexSoftwareCompatibilityInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
