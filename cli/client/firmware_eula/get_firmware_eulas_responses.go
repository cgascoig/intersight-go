// Code generated by go-swagger; DO NOT EDIT.

package firmware_eula

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetFirmwareEulasReader is a Reader for the GetFirmwareEulas structure.
type GetFirmwareEulasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirmwareEulasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirmwareEulasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFirmwareEulasDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirmwareEulasOK creates a GetFirmwareEulasOK with default headers values
func NewGetFirmwareEulasOK() *GetFirmwareEulasOK {
	return &GetFirmwareEulasOK{}
}

/*GetFirmwareEulasOK handles this case with default header values.

List of firmwareEulas for the given filter criteria
*/
type GetFirmwareEulasOK struct {
	Payload *models.FirmwareEulaList
}

func (o *GetFirmwareEulasOK) Error() string {
	return fmt.Sprintf("[GET /firmware/Eulas][%d] getFirmwareEulasOK  %+v", 200, o.Payload)
}

func (o *GetFirmwareEulasOK) GetPayload() *models.FirmwareEulaList {
	return o.Payload
}

func (o *GetFirmwareEulasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirmwareEulaList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwareEulasDefault creates a GetFirmwareEulasDefault with default headers values
func NewGetFirmwareEulasDefault(code int) *GetFirmwareEulasDefault {
	return &GetFirmwareEulasDefault{
		_statusCode: code,
	}
}

/*GetFirmwareEulasDefault handles this case with default header values.

Unexpected error
*/
type GetFirmwareEulasDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get firmware eulas default response
func (o *GetFirmwareEulasDefault) Code() int {
	return o._statusCode
}

func (o *GetFirmwareEulasDefault) Error() string {
	return fmt.Sprintf("[GET /firmware/Eulas][%d] GetFirmwareEulas default  %+v", o._statusCode, o.Payload)
}

func (o *GetFirmwareEulasDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwareEulasDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}