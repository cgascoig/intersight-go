// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_auto_support_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchHyperflexAutoSupportPoliciesMoidReader is a Reader for the PatchHyperflexAutoSupportPoliciesMoid structure.
type PatchHyperflexAutoSupportPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchHyperflexAutoSupportPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchHyperflexAutoSupportPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchHyperflexAutoSupportPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchHyperflexAutoSupportPoliciesMoidCreated creates a PatchHyperflexAutoSupportPoliciesMoidCreated with default headers values
func NewPatchHyperflexAutoSupportPoliciesMoidCreated() *PatchHyperflexAutoSupportPoliciesMoidCreated {
	return &PatchHyperflexAutoSupportPoliciesMoidCreated{}
}

/*PatchHyperflexAutoSupportPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchHyperflexAutoSupportPoliciesMoidCreated struct {
}

func (o *PatchHyperflexAutoSupportPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/AutoSupportPolicies/{moid}][%d] patchHyperflexAutoSupportPoliciesMoidCreated ", 201)
}

func (o *PatchHyperflexAutoSupportPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchHyperflexAutoSupportPoliciesMoidDefault creates a PatchHyperflexAutoSupportPoliciesMoidDefault with default headers values
func NewPatchHyperflexAutoSupportPoliciesMoidDefault(code int) *PatchHyperflexAutoSupportPoliciesMoidDefault {
	return &PatchHyperflexAutoSupportPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchHyperflexAutoSupportPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchHyperflexAutoSupportPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch hyperflex auto support policies moid default response
func (o *PatchHyperflexAutoSupportPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchHyperflexAutoSupportPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/AutoSupportPolicies/{moid}][%d] PatchHyperflexAutoSupportPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchHyperflexAutoSupportPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchHyperflexAutoSupportPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
