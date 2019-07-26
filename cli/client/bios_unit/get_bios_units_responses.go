// Code generated by go-swagger; DO NOT EDIT.

package bios_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetBiosUnitsReader is a Reader for the GetBiosUnits structure.
type GetBiosUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBiosUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBiosUnitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBiosUnitsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBiosUnitsOK creates a GetBiosUnitsOK with default headers values
func NewGetBiosUnitsOK() *GetBiosUnitsOK {
	return &GetBiosUnitsOK{}
}

/*GetBiosUnitsOK handles this case with default header values.

List of biosUnits for the given filter criteria
*/
type GetBiosUnitsOK struct {
	Payload *models.BiosUnitList
}

func (o *GetBiosUnitsOK) Error() string {
	return fmt.Sprintf("[GET /bios/Units][%d] getBiosUnitsOK  %+v", 200, o.Payload)
}

func (o *GetBiosUnitsOK) GetPayload() *models.BiosUnitList {
	return o.Payload
}

func (o *GetBiosUnitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BiosUnitList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBiosUnitsDefault creates a GetBiosUnitsDefault with default headers values
func NewGetBiosUnitsDefault(code int) *GetBiosUnitsDefault {
	return &GetBiosUnitsDefault{
		_statusCode: code,
	}
}

/*GetBiosUnitsDefault handles this case with default header values.

Unexpected error
*/
type GetBiosUnitsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get bios units default response
func (o *GetBiosUnitsDefault) Code() int {
	return o._statusCode
}

func (o *GetBiosUnitsDefault) Error() string {
	return fmt.Sprintf("[GET /bios/Units][%d] GetBiosUnits default  %+v", o._statusCode, o.Payload)
}

func (o *GetBiosUnitsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBiosUnitsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
