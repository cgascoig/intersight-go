// Code generated by go-swagger; DO NOT EDIT.

package firmware_running_firmware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetFirmwareRunningFirmwaresReader is a Reader for the GetFirmwareRunningFirmwares structure.
type GetFirmwareRunningFirmwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirmwareRunningFirmwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirmwareRunningFirmwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFirmwareRunningFirmwaresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirmwareRunningFirmwaresOK creates a GetFirmwareRunningFirmwaresOK with default headers values
func NewGetFirmwareRunningFirmwaresOK() *GetFirmwareRunningFirmwaresOK {
	return &GetFirmwareRunningFirmwaresOK{}
}

/*GetFirmwareRunningFirmwaresOK handles this case with default header values.

List of firmwareRunningFirmwares for the given filter criteria
*/
type GetFirmwareRunningFirmwaresOK struct {
	Payload *models.FirmwareRunningFirmwareList
}

func (o *GetFirmwareRunningFirmwaresOK) Error() string {
	return fmt.Sprintf("[GET /firmware/RunningFirmwares][%d] getFirmwareRunningFirmwaresOK  %+v", 200, o.Payload)
}

func (o *GetFirmwareRunningFirmwaresOK) GetPayload() *models.FirmwareRunningFirmwareList {
	return o.Payload
}

func (o *GetFirmwareRunningFirmwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirmwareRunningFirmwareList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwareRunningFirmwaresDefault creates a GetFirmwareRunningFirmwaresDefault with default headers values
func NewGetFirmwareRunningFirmwaresDefault(code int) *GetFirmwareRunningFirmwaresDefault {
	return &GetFirmwareRunningFirmwaresDefault{
		_statusCode: code,
	}
}

/*GetFirmwareRunningFirmwaresDefault handles this case with default header values.

Unexpected error
*/
type GetFirmwareRunningFirmwaresDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get firmware running firmwares default response
func (o *GetFirmwareRunningFirmwaresDefault) Code() int {
	return o._statusCode
}

func (o *GetFirmwareRunningFirmwaresDefault) Error() string {
	return fmt.Sprintf("[GET /firmware/RunningFirmwares][%d] GetFirmwareRunningFirmwares default  %+v", o._statusCode, o.Payload)
}

func (o *GetFirmwareRunningFirmwaresDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwareRunningFirmwaresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}