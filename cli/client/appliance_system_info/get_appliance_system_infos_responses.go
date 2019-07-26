// Code generated by go-swagger; DO NOT EDIT.

package appliance_system_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceSystemInfosReader is a Reader for the GetApplianceSystemInfos structure.
type GetApplianceSystemInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceSystemInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceSystemInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetApplianceSystemInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceSystemInfosOK creates a GetApplianceSystemInfosOK with default headers values
func NewGetApplianceSystemInfosOK() *GetApplianceSystemInfosOK {
	return &GetApplianceSystemInfosOK{}
}

/*GetApplianceSystemInfosOK handles this case with default header values.

List of applianceSystemInfos for the given filter criteria
*/
type GetApplianceSystemInfosOK struct {
	Payload *models.ApplianceSystemInfoList
}

func (o *GetApplianceSystemInfosOK) Error() string {
	return fmt.Sprintf("[GET /appliance/SystemInfos][%d] getApplianceSystemInfosOK  %+v", 200, o.Payload)
}

func (o *GetApplianceSystemInfosOK) GetPayload() *models.ApplianceSystemInfoList {
	return o.Payload
}

func (o *GetApplianceSystemInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceSystemInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceSystemInfosDefault creates a GetApplianceSystemInfosDefault with default headers values
func NewGetApplianceSystemInfosDefault(code int) *GetApplianceSystemInfosDefault {
	return &GetApplianceSystemInfosDefault{
		_statusCode: code,
	}
}

/*GetApplianceSystemInfosDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceSystemInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance system infos default response
func (o *GetApplianceSystemInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceSystemInfosDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/SystemInfos][%d] GetApplianceSystemInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceSystemInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceSystemInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}