// Code generated by go-swagger; DO NOT EDIT.

package iam_api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchIamAPIKeysMoidReader is a Reader for the PatchIamAPIKeysMoid structure.
type PatchIamAPIKeysMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIamAPIKeysMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchIamAPIKeysMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchIamAPIKeysMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchIamAPIKeysMoidCreated creates a PatchIamAPIKeysMoidCreated with default headers values
func NewPatchIamAPIKeysMoidCreated() *PatchIamAPIKeysMoidCreated {
	return &PatchIamAPIKeysMoidCreated{}
}

/*PatchIamAPIKeysMoidCreated handles this case with default header values.

Null response
*/
type PatchIamAPIKeysMoidCreated struct {
}

func (o *PatchIamAPIKeysMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /iam/ApiKeys/{moid}][%d] patchIamApiKeysMoidCreated ", 201)
}

func (o *PatchIamAPIKeysMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchIamAPIKeysMoidDefault creates a PatchIamAPIKeysMoidDefault with default headers values
func NewPatchIamAPIKeysMoidDefault(code int) *PatchIamAPIKeysMoidDefault {
	return &PatchIamAPIKeysMoidDefault{
		_statusCode: code,
	}
}

/*PatchIamAPIKeysMoidDefault handles this case with default header values.

unexpected error
*/
type PatchIamAPIKeysMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch iam API keys moid default response
func (o *PatchIamAPIKeysMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchIamAPIKeysMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /iam/ApiKeys/{moid}][%d] PatchIamAPIKeysMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchIamAPIKeysMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchIamAPIKeysMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}