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

// GetBootDeviceBootModesReader is a Reader for the GetBootDeviceBootModes structure.
type GetBootDeviceBootModesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBootDeviceBootModesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBootDeviceBootModesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBootDeviceBootModesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBootDeviceBootModesOK creates a GetBootDeviceBootModesOK with default headers values
func NewGetBootDeviceBootModesOK() *GetBootDeviceBootModesOK {
	return &GetBootDeviceBootModesOK{}
}

/*GetBootDeviceBootModesOK handles this case with default header values.

List of bootDeviceBootModes for the given filter criteria
*/
type GetBootDeviceBootModesOK struct {
	Payload *models.BootDeviceBootModeList
}

func (o *GetBootDeviceBootModesOK) Error() string {
	return fmt.Sprintf("[GET /boot/DeviceBootModes][%d] getBootDeviceBootModesOK  %+v", 200, o.Payload)
}

func (o *GetBootDeviceBootModesOK) GetPayload() *models.BootDeviceBootModeList {
	return o.Payload
}

func (o *GetBootDeviceBootModesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BootDeviceBootModeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBootDeviceBootModesDefault creates a GetBootDeviceBootModesDefault with default headers values
func NewGetBootDeviceBootModesDefault(code int) *GetBootDeviceBootModesDefault {
	return &GetBootDeviceBootModesDefault{
		_statusCode: code,
	}
}

/*GetBootDeviceBootModesDefault handles this case with default header values.

Unexpected error
*/
type GetBootDeviceBootModesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get boot device boot modes default response
func (o *GetBootDeviceBootModesDefault) Code() int {
	return o._statusCode
}

func (o *GetBootDeviceBootModesDefault) Error() string {
	return fmt.Sprintf("[GET /boot/DeviceBootModes][%d] GetBootDeviceBootModes default  %+v", o._statusCode, o.Payload)
}

func (o *GetBootDeviceBootModesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBootDeviceBootModesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
