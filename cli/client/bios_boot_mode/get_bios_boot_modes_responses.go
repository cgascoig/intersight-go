// Code generated by go-swagger; DO NOT EDIT.

package bios_boot_mode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetBiosBootModesReader is a Reader for the GetBiosBootModes structure.
type GetBiosBootModesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBiosBootModesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBiosBootModesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBiosBootModesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBiosBootModesOK creates a GetBiosBootModesOK with default headers values
func NewGetBiosBootModesOK() *GetBiosBootModesOK {
	return &GetBiosBootModesOK{}
}

/*GetBiosBootModesOK handles this case with default header values.

List of biosBootModes for the given filter criteria
*/
type GetBiosBootModesOK struct {
	Payload *models.BiosBootModeList
}

func (o *GetBiosBootModesOK) Error() string {
	return fmt.Sprintf("[GET /bios/BootModes][%d] getBiosBootModesOK  %+v", 200, o.Payload)
}

func (o *GetBiosBootModesOK) GetPayload() *models.BiosBootModeList {
	return o.Payload
}

func (o *GetBiosBootModesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BiosBootModeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBiosBootModesDefault creates a GetBiosBootModesDefault with default headers values
func NewGetBiosBootModesDefault(code int) *GetBiosBootModesDefault {
	return &GetBiosBootModesDefault{
		_statusCode: code,
	}
}

/*GetBiosBootModesDefault handles this case with default header values.

Unexpected error
*/
type GetBiosBootModesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get bios boot modes default response
func (o *GetBiosBootModesDefault) Code() int {
	return o._statusCode
}

func (o *GetBiosBootModesDefault) Error() string {
	return fmt.Sprintf("[GET /bios/BootModes][%d] GetBiosBootModes default  %+v", o._statusCode, o.Payload)
}

func (o *GetBiosBootModesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBiosBootModesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}