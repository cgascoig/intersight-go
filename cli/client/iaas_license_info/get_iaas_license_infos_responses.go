// Code generated by go-swagger; DO NOT EDIT.

package iaas_license_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIaasLicenseInfosReader is a Reader for the GetIaasLicenseInfos structure.
type GetIaasLicenseInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIaasLicenseInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIaasLicenseInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIaasLicenseInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIaasLicenseInfosOK creates a GetIaasLicenseInfosOK with default headers values
func NewGetIaasLicenseInfosOK() *GetIaasLicenseInfosOK {
	return &GetIaasLicenseInfosOK{}
}

/*GetIaasLicenseInfosOK handles this case with default header values.

List of iaasLicenseInfos for the given filter criteria
*/
type GetIaasLicenseInfosOK struct {
	Payload *models.IaasLicenseInfoList
}

func (o *GetIaasLicenseInfosOK) Error() string {
	return fmt.Sprintf("[GET /iaas/LicenseInfos][%d] getIaasLicenseInfosOK  %+v", 200, o.Payload)
}

func (o *GetIaasLicenseInfosOK) GetPayload() *models.IaasLicenseInfoList {
	return o.Payload
}

func (o *GetIaasLicenseInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaasLicenseInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIaasLicenseInfosDefault creates a GetIaasLicenseInfosDefault with default headers values
func NewGetIaasLicenseInfosDefault(code int) *GetIaasLicenseInfosDefault {
	return &GetIaasLicenseInfosDefault{
		_statusCode: code,
	}
}

/*GetIaasLicenseInfosDefault handles this case with default header values.

Unexpected error
*/
type GetIaasLicenseInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iaas license infos default response
func (o *GetIaasLicenseInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetIaasLicenseInfosDefault) Error() string {
	return fmt.Sprintf("[GET /iaas/LicenseInfos][%d] GetIaasLicenseInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetIaasLicenseInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIaasLicenseInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
