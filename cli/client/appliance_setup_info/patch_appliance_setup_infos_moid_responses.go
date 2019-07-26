// Code generated by go-swagger; DO NOT EDIT.

package appliance_setup_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchApplianceSetupInfosMoidReader is a Reader for the PatchApplianceSetupInfosMoid structure.
type PatchApplianceSetupInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchApplianceSetupInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchApplianceSetupInfosMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchApplianceSetupInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchApplianceSetupInfosMoidCreated creates a PatchApplianceSetupInfosMoidCreated with default headers values
func NewPatchApplianceSetupInfosMoidCreated() *PatchApplianceSetupInfosMoidCreated {
	return &PatchApplianceSetupInfosMoidCreated{}
}

/*PatchApplianceSetupInfosMoidCreated handles this case with default header values.

Null response
*/
type PatchApplianceSetupInfosMoidCreated struct {
}

func (o *PatchApplianceSetupInfosMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /appliance/SetupInfos/{moid}][%d] patchApplianceSetupInfosMoidCreated ", 201)
}

func (o *PatchApplianceSetupInfosMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchApplianceSetupInfosMoidDefault creates a PatchApplianceSetupInfosMoidDefault with default headers values
func NewPatchApplianceSetupInfosMoidDefault(code int) *PatchApplianceSetupInfosMoidDefault {
	return &PatchApplianceSetupInfosMoidDefault{
		_statusCode: code,
	}
}

/*PatchApplianceSetupInfosMoidDefault handles this case with default header values.

unexpected error
*/
type PatchApplianceSetupInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch appliance setup infos moid default response
func (o *PatchApplianceSetupInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchApplianceSetupInfosMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /appliance/SetupInfos/{moid}][%d] PatchApplianceSetupInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchApplianceSetupInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplianceSetupInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
