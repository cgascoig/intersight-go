// Code generated by go-swagger; DO NOT EDIT.

package asset_device_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchAssetDeviceRegistrationsMoidReader is a Reader for the PatchAssetDeviceRegistrationsMoid structure.
type PatchAssetDeviceRegistrationsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAssetDeviceRegistrationsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchAssetDeviceRegistrationsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchAssetDeviceRegistrationsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAssetDeviceRegistrationsMoidCreated creates a PatchAssetDeviceRegistrationsMoidCreated with default headers values
func NewPatchAssetDeviceRegistrationsMoidCreated() *PatchAssetDeviceRegistrationsMoidCreated {
	return &PatchAssetDeviceRegistrationsMoidCreated{}
}

/*PatchAssetDeviceRegistrationsMoidCreated handles this case with default header values.

Null response
*/
type PatchAssetDeviceRegistrationsMoidCreated struct {
}

func (o *PatchAssetDeviceRegistrationsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /asset/DeviceRegistrations/{moid}][%d] patchAssetDeviceRegistrationsMoidCreated ", 201)
}

func (o *PatchAssetDeviceRegistrationsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAssetDeviceRegistrationsMoidDefault creates a PatchAssetDeviceRegistrationsMoidDefault with default headers values
func NewPatchAssetDeviceRegistrationsMoidDefault(code int) *PatchAssetDeviceRegistrationsMoidDefault {
	return &PatchAssetDeviceRegistrationsMoidDefault{
		_statusCode: code,
	}
}

/*PatchAssetDeviceRegistrationsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchAssetDeviceRegistrationsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch asset device registrations moid default response
func (o *PatchAssetDeviceRegistrationsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchAssetDeviceRegistrationsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /asset/DeviceRegistrations/{moid}][%d] PatchAssetDeviceRegistrationsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAssetDeviceRegistrationsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchAssetDeviceRegistrationsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
