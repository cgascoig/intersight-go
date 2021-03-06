// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_feature_limit_external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchHyperflexFeatureLimitExternalsMoidReader is a Reader for the PatchHyperflexFeatureLimitExternalsMoid structure.
type PatchHyperflexFeatureLimitExternalsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchHyperflexFeatureLimitExternalsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchHyperflexFeatureLimitExternalsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchHyperflexFeatureLimitExternalsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchHyperflexFeatureLimitExternalsMoidCreated creates a PatchHyperflexFeatureLimitExternalsMoidCreated with default headers values
func NewPatchHyperflexFeatureLimitExternalsMoidCreated() *PatchHyperflexFeatureLimitExternalsMoidCreated {
	return &PatchHyperflexFeatureLimitExternalsMoidCreated{}
}

/*PatchHyperflexFeatureLimitExternalsMoidCreated handles this case with default header values.

Null response
*/
type PatchHyperflexFeatureLimitExternalsMoidCreated struct {
}

func (o *PatchHyperflexFeatureLimitExternalsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/FeatureLimitExternals/{moid}][%d] patchHyperflexFeatureLimitExternalsMoidCreated ", 201)
}

func (o *PatchHyperflexFeatureLimitExternalsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchHyperflexFeatureLimitExternalsMoidDefault creates a PatchHyperflexFeatureLimitExternalsMoidDefault with default headers values
func NewPatchHyperflexFeatureLimitExternalsMoidDefault(code int) *PatchHyperflexFeatureLimitExternalsMoidDefault {
	return &PatchHyperflexFeatureLimitExternalsMoidDefault{
		_statusCode: code,
	}
}

/*PatchHyperflexFeatureLimitExternalsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchHyperflexFeatureLimitExternalsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch hyperflex feature limit externals moid default response
func (o *PatchHyperflexFeatureLimitExternalsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchHyperflexFeatureLimitExternalsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/FeatureLimitExternals/{moid}][%d] PatchHyperflexFeatureLimitExternalsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchHyperflexFeatureLimitExternalsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchHyperflexFeatureLimitExternalsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
