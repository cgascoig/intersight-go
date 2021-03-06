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

// GetNiaapiDcnmLatestMaintainedReleasesReader is a Reader for the GetNiaapiDcnmLatestMaintainedReleases structure.
type GetNiaapiDcnmLatestMaintainedReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiDcnmLatestMaintainedReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiDcnmLatestMaintainedReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiDcnmLatestMaintainedReleasesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiDcnmLatestMaintainedReleasesOK creates a GetNiaapiDcnmLatestMaintainedReleasesOK with default headers values
func NewGetNiaapiDcnmLatestMaintainedReleasesOK() *GetNiaapiDcnmLatestMaintainedReleasesOK {
	return &GetNiaapiDcnmLatestMaintainedReleasesOK{}
}

/*GetNiaapiDcnmLatestMaintainedReleasesOK handles this case with default header values.

List of niaapiDcnmLatestMaintainedReleases for the given filter criteria
*/
type GetNiaapiDcnmLatestMaintainedReleasesOK struct {
	Payload *models.NiaapiDcnmLatestMaintainedReleaseList
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmLatestMaintainedReleases][%d] getNiaapiDcnmLatestMaintainedReleasesOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesOK) GetPayload() *models.NiaapiDcnmLatestMaintainedReleaseList {
	return o.Payload
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiDcnmLatestMaintainedReleaseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiDcnmLatestMaintainedReleasesDefault creates a GetNiaapiDcnmLatestMaintainedReleasesDefault with default headers values
func NewGetNiaapiDcnmLatestMaintainedReleasesDefault(code int) *GetNiaapiDcnmLatestMaintainedReleasesDefault {
	return &GetNiaapiDcnmLatestMaintainedReleasesDefault{
		_statusCode: code,
	}
}

/*GetNiaapiDcnmLatestMaintainedReleasesDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiDcnmLatestMaintainedReleasesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi dcnm latest maintained releases default response
func (o *GetNiaapiDcnmLatestMaintainedReleasesDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/DcnmLatestMaintainedReleases][%d] GetNiaapiDcnmLatestMaintainedReleases default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiDcnmLatestMaintainedReleasesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
