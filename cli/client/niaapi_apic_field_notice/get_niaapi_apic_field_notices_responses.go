// Code generated by go-swagger; DO NOT EDIT.

package niaapi_apic_field_notice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiApicFieldNoticesReader is a Reader for the GetNiaapiApicFieldNotices structure.
type GetNiaapiApicFieldNoticesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiApicFieldNoticesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiApicFieldNoticesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiApicFieldNoticesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiApicFieldNoticesOK creates a GetNiaapiApicFieldNoticesOK with default headers values
func NewGetNiaapiApicFieldNoticesOK() *GetNiaapiApicFieldNoticesOK {
	return &GetNiaapiApicFieldNoticesOK{}
}

/*GetNiaapiApicFieldNoticesOK handles this case with default header values.

List of niaapiApicFieldNotices for the given filter criteria
*/
type GetNiaapiApicFieldNoticesOK struct {
	Payload *models.NiaapiApicFieldNoticeList
}

func (o *GetNiaapiApicFieldNoticesOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/ApicFieldNotices][%d] getNiaapiApicFieldNoticesOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiApicFieldNoticesOK) GetPayload() *models.NiaapiApicFieldNoticeList {
	return o.Payload
}

func (o *GetNiaapiApicFieldNoticesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiApicFieldNoticeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiApicFieldNoticesDefault creates a GetNiaapiApicFieldNoticesDefault with default headers values
func NewGetNiaapiApicFieldNoticesDefault(code int) *GetNiaapiApicFieldNoticesDefault {
	return &GetNiaapiApicFieldNoticesDefault{
		_statusCode: code,
	}
}

/*GetNiaapiApicFieldNoticesDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiApicFieldNoticesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi apic field notices default response
func (o *GetNiaapiApicFieldNoticesDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiApicFieldNoticesDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/ApicFieldNotices][%d] GetNiaapiApicFieldNotices default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiApicFieldNoticesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiApicFieldNoticesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
