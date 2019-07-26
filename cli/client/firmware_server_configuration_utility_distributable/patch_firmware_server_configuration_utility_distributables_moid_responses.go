// Code generated by go-swagger; DO NOT EDIT.

package firmware_server_configuration_utility_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchFirmwareServerConfigurationUtilityDistributablesMoidReader is a Reader for the PatchFirmwareServerConfigurationUtilityDistributablesMoid structure.
type PatchFirmwareServerConfigurationUtilityDistributablesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchFirmwareServerConfigurationUtilityDistributablesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchFirmwareServerConfigurationUtilityDistributablesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchFirmwareServerConfigurationUtilityDistributablesMoidCreated creates a PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated with default headers values
func NewPatchFirmwareServerConfigurationUtilityDistributablesMoidCreated() *PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated {
	return &PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated{}
}

/*PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated handles this case with default header values.

Null response
*/
type PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated struct {
}

func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /firmware/ServerConfigurationUtilityDistributables/{moid}][%d] patchFirmwareServerConfigurationUtilityDistributablesMoidCreated ", 201)
}

func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchFirmwareServerConfigurationUtilityDistributablesMoidDefault creates a PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault with default headers values
func NewPatchFirmwareServerConfigurationUtilityDistributablesMoidDefault(code int) *PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault {
	return &PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault{
		_statusCode: code,
	}
}

/*PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault handles this case with default header values.

unexpected error
*/
type PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch firmware server configuration utility distributables moid default response
func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /firmware/ServerConfigurationUtilityDistributables/{moid}][%d] PatchFirmwareServerConfigurationUtilityDistributablesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
