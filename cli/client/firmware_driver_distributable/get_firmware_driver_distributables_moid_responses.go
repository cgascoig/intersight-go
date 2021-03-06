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

// GetFirmwareDriverDistributablesMoidReader is a Reader for the GetFirmwareDriverDistributablesMoid structure.
type GetFirmwareDriverDistributablesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirmwareDriverDistributablesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirmwareDriverDistributablesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetFirmwareDriverDistributablesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetFirmwareDriverDistributablesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirmwareDriverDistributablesMoidOK creates a GetFirmwareDriverDistributablesMoidOK with default headers values
func NewGetFirmwareDriverDistributablesMoidOK() *GetFirmwareDriverDistributablesMoidOK {
	return &GetFirmwareDriverDistributablesMoidOK{}
}

/*GetFirmwareDriverDistributablesMoidOK handles this case with default header values.

An instance of firmwareDriverDistributable
*/
type GetFirmwareDriverDistributablesMoidOK struct {
	Payload *models.FirmwareDriverDistributable
}

func (o *GetFirmwareDriverDistributablesMoidOK) Error() string {
	return fmt.Sprintf("[GET /firmware/DriverDistributables/{moid}][%d] getFirmwareDriverDistributablesMoidOK  %+v", 200, o.Payload)
}

func (o *GetFirmwareDriverDistributablesMoidOK) GetPayload() *models.FirmwareDriverDistributable {
	return o.Payload
}

func (o *GetFirmwareDriverDistributablesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirmwareDriverDistributable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwareDriverDistributablesMoidNotFound creates a GetFirmwareDriverDistributablesMoidNotFound with default headers values
func NewGetFirmwareDriverDistributablesMoidNotFound() *GetFirmwareDriverDistributablesMoidNotFound {
	return &GetFirmwareDriverDistributablesMoidNotFound{}
}

/*GetFirmwareDriverDistributablesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetFirmwareDriverDistributablesMoidNotFound struct {
}

func (o *GetFirmwareDriverDistributablesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /firmware/DriverDistributables/{moid}][%d] getFirmwareDriverDistributablesMoidNotFound ", 404)
}

func (o *GetFirmwareDriverDistributablesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFirmwareDriverDistributablesMoidDefault creates a GetFirmwareDriverDistributablesMoidDefault with default headers values
func NewGetFirmwareDriverDistributablesMoidDefault(code int) *GetFirmwareDriverDistributablesMoidDefault {
	return &GetFirmwareDriverDistributablesMoidDefault{
		_statusCode: code,
	}
}

/*GetFirmwareDriverDistributablesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetFirmwareDriverDistributablesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get firmware driver distributables moid default response
func (o *GetFirmwareDriverDistributablesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetFirmwareDriverDistributablesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /firmware/DriverDistributables/{moid}][%d] GetFirmwareDriverDistributablesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetFirmwareDriverDistributablesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwareDriverDistributablesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
