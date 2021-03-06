// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_user_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchIamEndPointUserPoliciesMoidReader is a Reader for the PatchIamEndPointUserPoliciesMoid structure.
type PatchIamEndPointUserPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIamEndPointUserPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchIamEndPointUserPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchIamEndPointUserPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchIamEndPointUserPoliciesMoidCreated creates a PatchIamEndPointUserPoliciesMoidCreated with default headers values
func NewPatchIamEndPointUserPoliciesMoidCreated() *PatchIamEndPointUserPoliciesMoidCreated {
	return &PatchIamEndPointUserPoliciesMoidCreated{}
}

/*PatchIamEndPointUserPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchIamEndPointUserPoliciesMoidCreated struct {
}

func (o *PatchIamEndPointUserPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /iam/EndPointUserPolicies/{moid}][%d] patchIamEndPointUserPoliciesMoidCreated ", 201)
}

func (o *PatchIamEndPointUserPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchIamEndPointUserPoliciesMoidDefault creates a PatchIamEndPointUserPoliciesMoidDefault with default headers values
func NewPatchIamEndPointUserPoliciesMoidDefault(code int) *PatchIamEndPointUserPoliciesMoidDefault {
	return &PatchIamEndPointUserPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchIamEndPointUserPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchIamEndPointUserPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch iam end point user policies moid default response
func (o *PatchIamEndPointUserPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchIamEndPointUserPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /iam/EndPointUserPolicies/{moid}][%d] PatchIamEndPointUserPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchIamEndPointUserPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchIamEndPointUserPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
