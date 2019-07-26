// Code generated by go-swagger; DO NOT EDIT.

package niaapi_dcnm_release_recommend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiDcnmReleaseRecommendsReader is a Reader for the GetNiaapiDcnmReleaseRecommends structure.
type GetNiaapiDcnmReleaseRecommendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiDcnmReleaseRecommendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiDcnmReleaseRecommendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiDcnmReleaseRecommendsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiDcnmReleaseRecommendsOK creates a GetNiaapiDcnmReleaseRecommendsOK with default headers values
func NewGetNiaapiDcnmReleaseRecommendsOK() *GetNiaapiDcnmReleaseRecommendsOK {
	return &GetNiaapiDcnmReleaseRecommendsOK{}
}

/*GetNiaapiDcnmReleaseRecommendsOK handles this case with default header values.

List of niaapiDcnmReleaseRecommends for the given filter criteria
*/
type GetNiaapiDcnmReleaseRecommendsOK struct {
	Payload *models.NiaapiDcnmReleaseRecommendList
}

func (o *GetNiaapiDcnmReleaseRecommendsOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmReleaseRecommends][%d] getNiaapiDcnmReleaseRecommendsOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiDcnmReleaseRecommendsOK) GetPayload() *models.NiaapiDcnmReleaseRecommendList {
	return o.Payload
}

func (o *GetNiaapiDcnmReleaseRecommendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiDcnmReleaseRecommendList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiDcnmReleaseRecommendsDefault creates a GetNiaapiDcnmReleaseRecommendsDefault with default headers values
func NewGetNiaapiDcnmReleaseRecommendsDefault(code int) *GetNiaapiDcnmReleaseRecommendsDefault {
	return &GetNiaapiDcnmReleaseRecommendsDefault{
		_statusCode: code,
	}
}

/*GetNiaapiDcnmReleaseRecommendsDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiDcnmReleaseRecommendsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi dcnm release recommends default response
func (o *GetNiaapiDcnmReleaseRecommendsDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiDcnmReleaseRecommendsDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmReleaseRecommends][%d] GetNiaapiDcnmReleaseRecommends default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiDcnmReleaseRecommendsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiDcnmReleaseRecommendsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
