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

// PostAssetDeviceRegistrationsMoidReader is a Reader for the PostAssetDeviceRegistrationsMoid structure.
type PostAssetDeviceRegistrationsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAssetDeviceRegistrationsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAssetDeviceRegistrationsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAssetDeviceRegistrationsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAssetDeviceRegistrationsMoidCreated creates a PostAssetDeviceRegistrationsMoidCreated with default headers values
func NewPostAssetDeviceRegistrationsMoidCreated() *PostAssetDeviceRegistrationsMoidCreated {
	return &PostAssetDeviceRegistrationsMoidCreated{}
}

/*PostAssetDeviceRegistrationsMoidCreated handles this case with default header values.

Null response
*/
type PostAssetDeviceRegistrationsMoidCreated struct {
}

func (o *PostAssetDeviceRegistrationsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /asset/DeviceRegistrations/{moid}][%d] postAssetDeviceRegistrationsMoidCreated ", 201)
}

func (o *PostAssetDeviceRegistrationsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAssetDeviceRegistrationsMoidDefault creates a PostAssetDeviceRegistrationsMoidDefault with default headers values
func NewPostAssetDeviceRegistrationsMoidDefault(code int) *PostAssetDeviceRegistrationsMoidDefault {
	return &PostAssetDeviceRegistrationsMoidDefault{
		_statusCode: code,
	}
}

/*PostAssetDeviceRegistrationsMoidDefault handles this case with default header values.

unexpected error
*/
type PostAssetDeviceRegistrationsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post asset device registrations moid default response
func (o *PostAssetDeviceRegistrationsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostAssetDeviceRegistrationsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /asset/DeviceRegistrations/{moid}][%d] PostAssetDeviceRegistrationsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostAssetDeviceRegistrationsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostAssetDeviceRegistrationsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
