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

// GetLicenseLicenseInfosReader is a Reader for the GetLicenseLicenseInfos structure.
type GetLicenseLicenseInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseLicenseInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseLicenseInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLicenseLicenseInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLicenseLicenseInfosOK creates a GetLicenseLicenseInfosOK with default headers values
func NewGetLicenseLicenseInfosOK() *GetLicenseLicenseInfosOK {
	return &GetLicenseLicenseInfosOK{}
}

/*GetLicenseLicenseInfosOK handles this case with default header values.

List of licenseLicenseInfos for the given filter criteria
*/
type GetLicenseLicenseInfosOK struct {
	Payload *models.LicenseLicenseInfoList
}

func (o *GetLicenseLicenseInfosOK) Error() string {
	return fmt.Sprintf("[GET /license/LicenseInfos][%d] getLicenseLicenseInfosOK  %+v", 200, o.Payload)
}

func (o *GetLicenseLicenseInfosOK) GetPayload() *models.LicenseLicenseInfoList {
	return o.Payload
}

func (o *GetLicenseLicenseInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseLicenseInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseLicenseInfosDefault creates a GetLicenseLicenseInfosDefault with default headers values
func NewGetLicenseLicenseInfosDefault(code int) *GetLicenseLicenseInfosDefault {
	return &GetLicenseLicenseInfosDefault{
		_statusCode: code,
	}
}

/*GetLicenseLicenseInfosDefault handles this case with default header values.

Unexpected error
*/
type GetLicenseLicenseInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get license license infos default response
func (o *GetLicenseLicenseInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetLicenseLicenseInfosDefault) Error() string {
	return fmt.Sprintf("[GET /license/LicenseInfos][%d] GetLicenseLicenseInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetLicenseLicenseInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLicenseLicenseInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
