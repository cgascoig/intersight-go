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

// GetAssetDeviceConfigurationsReader is a Reader for the GetAssetDeviceConfigurations structure.
type GetAssetDeviceConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetDeviceConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetDeviceConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAssetDeviceConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssetDeviceConfigurationsOK creates a GetAssetDeviceConfigurationsOK with default headers values
func NewGetAssetDeviceConfigurationsOK() *GetAssetDeviceConfigurationsOK {
	return &GetAssetDeviceConfigurationsOK{}
}

/*GetAssetDeviceConfigurationsOK handles this case with default header values.

List of assetDeviceConfigurations for the given filter criteria
*/
type GetAssetDeviceConfigurationsOK struct {
	Payload *models.AssetDeviceConfigurationList
}

func (o *GetAssetDeviceConfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /asset/DeviceConfigurations][%d] getAssetDeviceConfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetAssetDeviceConfigurationsOK) GetPayload() *models.AssetDeviceConfigurationList {
	return o.Payload
}

func (o *GetAssetDeviceConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetDeviceConfigurationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetDeviceConfigurationsDefault creates a GetAssetDeviceConfigurationsDefault with default headers values
func NewGetAssetDeviceConfigurationsDefault(code int) *GetAssetDeviceConfigurationsDefault {
	return &GetAssetDeviceConfigurationsDefault{
		_statusCode: code,
	}
}

/*GetAssetDeviceConfigurationsDefault handles this case with default header values.

Unexpected error
*/
type GetAssetDeviceConfigurationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get asset device configurations default response
func (o *GetAssetDeviceConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *GetAssetDeviceConfigurationsDefault) Error() string {
	return fmt.Sprintf("[GET /asset/DeviceConfigurations][%d] GetAssetDeviceConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssetDeviceConfigurationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAssetDeviceConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
