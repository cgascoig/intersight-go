// Code generated by go-swagger; DO NOT EDIT.

package asset_device_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetAssetDeviceConfigurationsMoidReader is a Reader for the GetAssetDeviceConfigurationsMoid structure.
type GetAssetDeviceConfigurationsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetDeviceConfigurationsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetDeviceConfigurationsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAssetDeviceConfigurationsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAssetDeviceConfigurationsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssetDeviceConfigurationsMoidOK creates a GetAssetDeviceConfigurationsMoidOK with default headers values
func NewGetAssetDeviceConfigurationsMoidOK() *GetAssetDeviceConfigurationsMoidOK {
	return &GetAssetDeviceConfigurationsMoidOK{}
}

/*GetAssetDeviceConfigurationsMoidOK handles this case with default header values.

An instance of assetDeviceConfiguration
*/
type GetAssetDeviceConfigurationsMoidOK struct {
	Payload *models.AssetDeviceConfiguration
}

func (o *GetAssetDeviceConfigurationsMoidOK) Error() string {
	return fmt.Sprintf("[GET /asset/DeviceConfigurations/{moid}][%d] getAssetDeviceConfigurationsMoidOK  %+v", 200, o.Payload)
}

func (o *GetAssetDeviceConfigurationsMoidOK) GetPayload() *models.AssetDeviceConfiguration {
	return o.Payload
}

func (o *GetAssetDeviceConfigurationsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetDeviceConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetDeviceConfigurationsMoidNotFound creates a GetAssetDeviceConfigurationsMoidNotFound with default headers values
func NewGetAssetDeviceConfigurationsMoidNotFound() *GetAssetDeviceConfigurationsMoidNotFound {
	return &GetAssetDeviceConfigurationsMoidNotFound{}
}

/*GetAssetDeviceConfigurationsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetAssetDeviceConfigurationsMoidNotFound struct {
}

func (o *GetAssetDeviceConfigurationsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /asset/DeviceConfigurations/{moid}][%d] getAssetDeviceConfigurationsMoidNotFound ", 404)
}

func (o *GetAssetDeviceConfigurationsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetDeviceConfigurationsMoidDefault creates a GetAssetDeviceConfigurationsMoidDefault with default headers values
func NewGetAssetDeviceConfigurationsMoidDefault(code int) *GetAssetDeviceConfigurationsMoidDefault {
	return &GetAssetDeviceConfigurationsMoidDefault{
		_statusCode: code,
	}
}

/*GetAssetDeviceConfigurationsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetAssetDeviceConfigurationsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get asset device configurations moid default response
func (o *GetAssetDeviceConfigurationsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetAssetDeviceConfigurationsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /asset/DeviceConfigurations/{moid}][%d] GetAssetDeviceConfigurationsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssetDeviceConfigurationsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAssetDeviceConfigurationsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}