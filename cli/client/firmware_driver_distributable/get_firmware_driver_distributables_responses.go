// Code generated by go-swagger; DO NOT EDIT.

package firmware_driver_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetFirmwareDriverDistributablesReader is a Reader for the GetFirmwareDriverDistributables structure.
type GetFirmwareDriverDistributablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirmwareDriverDistributablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirmwareDriverDistributablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFirmwareDriverDistributablesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirmwareDriverDistributablesOK creates a GetFirmwareDriverDistributablesOK with default headers values
func NewGetFirmwareDriverDistributablesOK() *GetFirmwareDriverDistributablesOK {
	return &GetFirmwareDriverDistributablesOK{}
}

/*GetFirmwareDriverDistributablesOK handles this case with default header values.

List of firmwareDriverDistributables for the given filter criteria
*/
type GetFirmwareDriverDistributablesOK struct {
	Payload *models.FirmwareDriverDistributableList
}

func (o *GetFirmwareDriverDistributablesOK) Error() string {
	return fmt.Sprintf("[GET /firmware/DriverDistributables][%d] getFirmwareDriverDistributablesOK  %+v", 200, o.Payload)
}

func (o *GetFirmwareDriverDistributablesOK) GetPayload() *models.FirmwareDriverDistributableList {
	return o.Payload
}

func (o *GetFirmwareDriverDistributablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirmwareDriverDistributableList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwareDriverDistributablesDefault creates a GetFirmwareDriverDistributablesDefault with default headers values
func NewGetFirmwareDriverDistributablesDefault(code int) *GetFirmwareDriverDistributablesDefault {
	return &GetFirmwareDriverDistributablesDefault{
		_statusCode: code,
	}
}

/*GetFirmwareDriverDistributablesDefault handles this case with default header values.

Unexpected error
*/
type GetFirmwareDriverDistributablesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get firmware driver distributables default response
func (o *GetFirmwareDriverDistributablesDefault) Code() int {
	return o._statusCode
}

func (o *GetFirmwareDriverDistributablesDefault) Error() string {
	return fmt.Sprintf("[GET /firmware/DriverDistributables][%d] GetFirmwareDriverDistributables default  %+v", o._statusCode, o.Payload)
}

func (o *GetFirmwareDriverDistributablesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwareDriverDistributablesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
