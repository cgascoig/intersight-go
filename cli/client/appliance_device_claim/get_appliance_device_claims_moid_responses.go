// Code generated by go-swagger; DO NOT EDIT.

package appliance_device_claim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceDeviceClaimsMoidReader is a Reader for the GetApplianceDeviceClaimsMoid structure.
type GetApplianceDeviceClaimsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceDeviceClaimsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceDeviceClaimsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApplianceDeviceClaimsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplianceDeviceClaimsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceDeviceClaimsMoidOK creates a GetApplianceDeviceClaimsMoidOK with default headers values
func NewGetApplianceDeviceClaimsMoidOK() *GetApplianceDeviceClaimsMoidOK {
	return &GetApplianceDeviceClaimsMoidOK{}
}

/*GetApplianceDeviceClaimsMoidOK handles this case with default header values.

An instance of applianceDeviceClaim
*/
type GetApplianceDeviceClaimsMoidOK struct {
	Payload *models.ApplianceDeviceClaim
}

func (o *GetApplianceDeviceClaimsMoidOK) Error() string {
	return fmt.Sprintf("[GET /appliance/DeviceClaims/{moid}][%d] getApplianceDeviceClaimsMoidOK  %+v", 200, o.Payload)
}

func (o *GetApplianceDeviceClaimsMoidOK) GetPayload() *models.ApplianceDeviceClaim {
	return o.Payload
}

func (o *GetApplianceDeviceClaimsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceDeviceClaim)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceDeviceClaimsMoidNotFound creates a GetApplianceDeviceClaimsMoidNotFound with default headers values
func NewGetApplianceDeviceClaimsMoidNotFound() *GetApplianceDeviceClaimsMoidNotFound {
	return &GetApplianceDeviceClaimsMoidNotFound{}
}

/*GetApplianceDeviceClaimsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetApplianceDeviceClaimsMoidNotFound struct {
}

func (o *GetApplianceDeviceClaimsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /appliance/DeviceClaims/{moid}][%d] getApplianceDeviceClaimsMoidNotFound ", 404)
}

func (o *GetApplianceDeviceClaimsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplianceDeviceClaimsMoidDefault creates a GetApplianceDeviceClaimsMoidDefault with default headers values
func NewGetApplianceDeviceClaimsMoidDefault(code int) *GetApplianceDeviceClaimsMoidDefault {
	return &GetApplianceDeviceClaimsMoidDefault{
		_statusCode: code,
	}
}

/*GetApplianceDeviceClaimsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceDeviceClaimsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance device claims moid default response
func (o *GetApplianceDeviceClaimsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceDeviceClaimsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/DeviceClaims/{moid}][%d] GetApplianceDeviceClaimsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceDeviceClaimsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceDeviceClaimsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}