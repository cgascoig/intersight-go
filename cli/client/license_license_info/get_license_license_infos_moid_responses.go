// Code generated by go-swagger; DO NOT EDIT.

package license_license_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetLicenseLicenseInfosMoidReader is a Reader for the GetLicenseLicenseInfosMoid structure.
type GetLicenseLicenseInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseLicenseInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseLicenseInfosMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLicenseLicenseInfosMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLicenseLicenseInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLicenseLicenseInfosMoidOK creates a GetLicenseLicenseInfosMoidOK with default headers values
func NewGetLicenseLicenseInfosMoidOK() *GetLicenseLicenseInfosMoidOK {
	return &GetLicenseLicenseInfosMoidOK{}
}

/*GetLicenseLicenseInfosMoidOK handles this case with default header values.

An instance of licenseLicenseInfo
*/
type GetLicenseLicenseInfosMoidOK struct {
	Payload *models.LicenseLicenseInfo
}

func (o *GetLicenseLicenseInfosMoidOK) Error() string {
	return fmt.Sprintf("[GET /license/LicenseInfos/{moid}][%d] getLicenseLicenseInfosMoidOK  %+v", 200, o.Payload)
}

func (o *GetLicenseLicenseInfosMoidOK) GetPayload() *models.LicenseLicenseInfo {
	return o.Payload
}

func (o *GetLicenseLicenseInfosMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseLicenseInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseLicenseInfosMoidNotFound creates a GetLicenseLicenseInfosMoidNotFound with default headers values
func NewGetLicenseLicenseInfosMoidNotFound() *GetLicenseLicenseInfosMoidNotFound {
	return &GetLicenseLicenseInfosMoidNotFound{}
}

/*GetLicenseLicenseInfosMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetLicenseLicenseInfosMoidNotFound struct {
}

func (o *GetLicenseLicenseInfosMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /license/LicenseInfos/{moid}][%d] getLicenseLicenseInfosMoidNotFound ", 404)
}

func (o *GetLicenseLicenseInfosMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLicenseLicenseInfosMoidDefault creates a GetLicenseLicenseInfosMoidDefault with default headers values
func NewGetLicenseLicenseInfosMoidDefault(code int) *GetLicenseLicenseInfosMoidDefault {
	return &GetLicenseLicenseInfosMoidDefault{
		_statusCode: code,
	}
}

/*GetLicenseLicenseInfosMoidDefault handles this case with default header values.

Unexpected error
*/
type GetLicenseLicenseInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get license license infos moid default response
func (o *GetLicenseLicenseInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetLicenseLicenseInfosMoidDefault) Error() string {
	return fmt.Sprintf("[GET /license/LicenseInfos/{moid}][%d] GetLicenseLicenseInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetLicenseLicenseInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLicenseLicenseInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}