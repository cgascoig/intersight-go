// Code generated by go-swagger; DO NOT EDIT.

package os_install

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetOsInstallsReader is a Reader for the GetOsInstalls structure.
type GetOsInstallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOsInstallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOsInstallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOsInstallsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOsInstallsOK creates a GetOsInstallsOK with default headers values
func NewGetOsInstallsOK() *GetOsInstallsOK {
	return &GetOsInstallsOK{}
}

/*GetOsInstallsOK handles this case with default header values.

List of osInstalls for the given filter criteria
*/
type GetOsInstallsOK struct {
	Payload *models.OsInstallList
}

func (o *GetOsInstallsOK) Error() string {
	return fmt.Sprintf("[GET /os/Installs][%d] getOsInstallsOK  %+v", 200, o.Payload)
}

func (o *GetOsInstallsOK) GetPayload() *models.OsInstallList {
	return o.Payload
}

func (o *GetOsInstallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OsInstallList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOsInstallsDefault creates a GetOsInstallsDefault with default headers values
func NewGetOsInstallsDefault(code int) *GetOsInstallsDefault {
	return &GetOsInstallsDefault{
		_statusCode: code,
	}
}

/*GetOsInstallsDefault handles this case with default header values.

Unexpected error
*/
type GetOsInstallsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get os installs default response
func (o *GetOsInstallsDefault) Code() int {
	return o._statusCode
}

func (o *GetOsInstallsDefault) Error() string {
	return fmt.Sprintf("[GET /os/Installs][%d] GetOsInstalls default  %+v", o._statusCode, o.Payload)
}

func (o *GetOsInstallsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOsInstallsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}