// Code generated by go-swagger; DO NOT EDIT.

package port_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchPortGroupsMoidReader is a Reader for the PatchPortGroupsMoid structure.
type PatchPortGroupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPortGroupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchPortGroupsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchPortGroupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchPortGroupsMoidCreated creates a PatchPortGroupsMoidCreated with default headers values
func NewPatchPortGroupsMoidCreated() *PatchPortGroupsMoidCreated {
	return &PatchPortGroupsMoidCreated{}
}

/*PatchPortGroupsMoidCreated handles this case with default header values.

Null response
*/
type PatchPortGroupsMoidCreated struct {
}

func (o *PatchPortGroupsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /port/Groups/{moid}][%d] patchPortGroupsMoidCreated ", 201)
}

func (o *PatchPortGroupsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPortGroupsMoidDefault creates a PatchPortGroupsMoidDefault with default headers values
func NewPatchPortGroupsMoidDefault(code int) *PatchPortGroupsMoidDefault {
	return &PatchPortGroupsMoidDefault{
		_statusCode: code,
	}
}

/*PatchPortGroupsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchPortGroupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch port groups moid default response
func (o *PatchPortGroupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchPortGroupsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /port/Groups/{moid}][%d] PatchPortGroupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchPortGroupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchPortGroupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
