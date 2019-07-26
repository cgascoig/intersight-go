// Code generated by go-swagger; DO NOT EDIT.

package vmedia_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchVmediaPoliciesMoidReader is a Reader for the PatchVmediaPoliciesMoid structure.
type PatchVmediaPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVmediaPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchVmediaPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchVmediaPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchVmediaPoliciesMoidCreated creates a PatchVmediaPoliciesMoidCreated with default headers values
func NewPatchVmediaPoliciesMoidCreated() *PatchVmediaPoliciesMoidCreated {
	return &PatchVmediaPoliciesMoidCreated{}
}

/*PatchVmediaPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PatchVmediaPoliciesMoidCreated struct {
}

func (o *PatchVmediaPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /vmedia/Policies/{moid}][%d] patchVmediaPoliciesMoidCreated ", 201)
}

func (o *PatchVmediaPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchVmediaPoliciesMoidDefault creates a PatchVmediaPoliciesMoidDefault with default headers values
func NewPatchVmediaPoliciesMoidDefault(code int) *PatchVmediaPoliciesMoidDefault {
	return &PatchVmediaPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PatchVmediaPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchVmediaPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch vmedia policies moid default response
func (o *PatchVmediaPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchVmediaPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /vmedia/Policies/{moid}][%d] PatchVmediaPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchVmediaPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchVmediaPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}