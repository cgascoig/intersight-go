// Code generated by go-swagger; DO NOT EDIT.

package niaapi_apic_release_recommend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiApicReleaseRecommendsReader is a Reader for the GetNiaapiApicReleaseRecommends structure.
type GetNiaapiApicReleaseRecommendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiApicReleaseRecommendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiApicReleaseRecommendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiApicReleaseRecommendsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiApicReleaseRecommendsOK creates a GetNiaapiApicReleaseRecommendsOK with default headers values
func NewGetNiaapiApicReleaseRecommendsOK() *GetNiaapiApicReleaseRecommendsOK {
	return &GetNiaapiApicReleaseRecommendsOK{}
}

/*GetNiaapiApicReleaseRecommendsOK handles this case with default header values.

List of niaapiApicReleaseRecommends for the given filter criteria
*/
type GetNiaapiApicReleaseRecommendsOK struct {
	Payload *models.NiaapiApicReleaseRecommendList
}

func (o *GetNiaapiApicReleaseRecommendsOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/ApicReleaseRecommends][%d] getNiaapiApicReleaseRecommendsOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiApicReleaseRecommendsOK) GetPayload() *models.NiaapiApicReleaseRecommendList {
	return o.Payload
}

func (o *GetNiaapiApicReleaseRecommendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiApicReleaseRecommendList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiApicReleaseRecommendsDefault creates a GetNiaapiApicReleaseRecommendsDefault with default headers values
func NewGetNiaapiApicReleaseRecommendsDefault(code int) *GetNiaapiApicReleaseRecommendsDefault {
	return &GetNiaapiApicReleaseRecommendsDefault{
		_statusCode: code,
	}
}

/*GetNiaapiApicReleaseRecommendsDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiApicReleaseRecommendsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi apic release recommends default response
func (o *GetNiaapiApicReleaseRecommendsDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiApicReleaseRecommendsDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/ApicReleaseRecommends][%d] GetNiaapiApicReleaseRecommends default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiApicReleaseRecommendsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiApicReleaseRecommendsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
