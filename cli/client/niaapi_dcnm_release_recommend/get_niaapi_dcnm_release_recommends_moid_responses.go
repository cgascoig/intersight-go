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

// GetNiaapiDcnmReleaseRecommendsMoidReader is a Reader for the GetNiaapiDcnmReleaseRecommendsMoid structure.
type GetNiaapiDcnmReleaseRecommendsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiDcnmReleaseRecommendsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiDcnmReleaseRecommendsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNiaapiDcnmReleaseRecommendsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNiaapiDcnmReleaseRecommendsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiDcnmReleaseRecommendsMoidOK creates a GetNiaapiDcnmReleaseRecommendsMoidOK with default headers values
func NewGetNiaapiDcnmReleaseRecommendsMoidOK() *GetNiaapiDcnmReleaseRecommendsMoidOK {
	return &GetNiaapiDcnmReleaseRecommendsMoidOK{}
}

/*GetNiaapiDcnmReleaseRecommendsMoidOK handles this case with default header values.

An instance of niaapiDcnmReleaseRecommend
*/
type GetNiaapiDcnmReleaseRecommendsMoidOK struct {
	Payload *models.NiaapiDcnmReleaseRecommend
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmReleaseRecommends/{moid}][%d] getNiaapiDcnmReleaseRecommendsMoidOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidOK) GetPayload() *models.NiaapiDcnmReleaseRecommend {
	return o.Payload
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiDcnmReleaseRecommend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiDcnmReleaseRecommendsMoidNotFound creates a GetNiaapiDcnmReleaseRecommendsMoidNotFound with default headers values
func NewGetNiaapiDcnmReleaseRecommendsMoidNotFound() *GetNiaapiDcnmReleaseRecommendsMoidNotFound {
	return &GetNiaapiDcnmReleaseRecommendsMoidNotFound{}
}

/*GetNiaapiDcnmReleaseRecommendsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetNiaapiDcnmReleaseRecommendsMoidNotFound struct {
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmReleaseRecommends/{moid}][%d] getNiaapiDcnmReleaseRecommendsMoidNotFound ", 404)
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNiaapiDcnmReleaseRecommendsMoidDefault creates a GetNiaapiDcnmReleaseRecommendsMoidDefault with default headers values
func NewGetNiaapiDcnmReleaseRecommendsMoidDefault(code int) *GetNiaapiDcnmReleaseRecommendsMoidDefault {
	return &GetNiaapiDcnmReleaseRecommendsMoidDefault{
		_statusCode: code,
	}
}

/*GetNiaapiDcnmReleaseRecommendsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiDcnmReleaseRecommendsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi dcnm release recommends moid default response
func (o *GetNiaapiDcnmReleaseRecommendsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmReleaseRecommends/{moid}][%d] GetNiaapiDcnmReleaseRecommendsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiDcnmReleaseRecommendsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}