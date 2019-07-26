// Code generated by go-swagger; DO NOT EDIT.

package boot_device_boot_mode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostBootDeviceBootModesMoidReader is a Reader for the PostBootDeviceBootModesMoid structure.
type PostBootDeviceBootModesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBootDeviceBootModesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostBootDeviceBootModesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostBootDeviceBootModesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostBootDeviceBootModesMoidCreated creates a PostBootDeviceBootModesMoidCreated with default headers values
func NewPostBootDeviceBootModesMoidCreated() *PostBootDeviceBootModesMoidCreated {
	return &PostBootDeviceBootModesMoidCreated{}
}

/*PostBootDeviceBootModesMoidCreated handles this case with default header values.

Null response
*/
type PostBootDeviceBootModesMoidCreated struct {
}

func (o *PostBootDeviceBootModesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /boot/DeviceBootModes/{moid}][%d] postBootDeviceBootModesMoidCreated ", 201)
}

func (o *PostBootDeviceBootModesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBootDeviceBootModesMoidDefault creates a PostBootDeviceBootModesMoidDefault with default headers values
func NewPostBootDeviceBootModesMoidDefault(code int) *PostBootDeviceBootModesMoidDefault {
	return &PostBootDeviceBootModesMoidDefault{
		_statusCode: code,
	}
}

/*PostBootDeviceBootModesMoidDefault handles this case with default header values.

unexpected error
*/
type PostBootDeviceBootModesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post boot device boot modes moid default response
func (o *PostBootDeviceBootModesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostBootDeviceBootModesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /boot/DeviceBootModes/{moid}][%d] PostBootDeviceBootModesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostBootDeviceBootModesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostBootDeviceBootModesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
