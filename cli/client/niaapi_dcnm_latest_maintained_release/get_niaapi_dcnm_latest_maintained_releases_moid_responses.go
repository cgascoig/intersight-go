// Code generated by go-swagger; DO NOT EDIT.

package niaapi_dcnm_latest_maintained_release

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiDcnmLatestMaintainedReleasesMoidReader is a Reader for the GetNiaapiDcnmLatestMaintainedReleasesMoid structure.
type GetNiaapiDcnmLatestMaintainedReleasesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiDcnmLatestMaintainedReleasesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNiaapiDcnmLatestMaintainedReleasesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNiaapiDcnmLatestMaintainedReleasesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidOK creates a GetNiaapiDcnmLatestMaintainedReleasesMoidOK with default headers values
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidOK() *GetNiaapiDcnmLatestMaintainedReleasesMoidOK {
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidOK{}
}

/*GetNiaapiDcnmLatestMaintainedReleasesMoidOK handles this case with default header values.

An instance of niaapiDcnmLatestMaintainedRelease
*/
type GetNiaapiDcnmLatestMaintainedReleasesMoidOK struct {
	Payload *models.NiaapiDcnmLatestMaintainedRelease
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmLatestMaintainedReleases/{moid}][%d] getNiaapiDcnmLatestMaintainedReleasesMoidOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidOK) GetPayload() *models.NiaapiDcnmLatestMaintainedRelease {
	return o.Payload
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiDcnmLatestMaintainedRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidNotFound creates a GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound with default headers values
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidNotFound() *GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound {
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound{}
}

/*GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound struct {
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmLatestMaintainedReleases/{moid}][%d] getNiaapiDcnmLatestMaintainedReleasesMoidNotFound ", 404)
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNiaapiDcnmLatestMaintainedReleasesMoidDefault creates a GetNiaapiDcnmLatestMaintainedReleasesMoidDefault with default headers values
func NewGetNiaapiDcnmLatestMaintainedReleasesMoidDefault(code int) *GetNiaapiDcnmLatestMaintainedReleasesMoidDefault {
	return &GetNiaapiDcnmLatestMaintainedReleasesMoidDefault{
		_statusCode: code,
	}
}

/*GetNiaapiDcnmLatestMaintainedReleasesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiDcnmLatestMaintainedReleasesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi dcnm latest maintained releases moid default response
func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmLatestMaintainedReleases/{moid}][%d] GetNiaapiDcnmLatestMaintainedReleasesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
