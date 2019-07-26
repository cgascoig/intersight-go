// Code generated by go-swagger; DO NOT EDIT.

package niaapi_nia_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiNiaMetadataReader is a Reader for the GetNiaapiNiaMetadata structure.
type GetNiaapiNiaMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiNiaMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiNiaMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiNiaMetadataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiNiaMetadataOK creates a GetNiaapiNiaMetadataOK with default headers values
func NewGetNiaapiNiaMetadataOK() *GetNiaapiNiaMetadataOK {
	return &GetNiaapiNiaMetadataOK{}
}

/*GetNiaapiNiaMetadataOK handles this case with default header values.

List of niaapiNiaMetadata for the given filter criteria
*/
type GetNiaapiNiaMetadataOK struct {
	Payload *models.NiaapiNiaMetadataList
}

func (o *GetNiaapiNiaMetadataOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/NiaMetadata][%d] getNiaapiNiaMetadataOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiNiaMetadataOK) GetPayload() *models.NiaapiNiaMetadataList {
	return o.Payload
}

func (o *GetNiaapiNiaMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiNiaMetadataList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiNiaMetadataDefault creates a GetNiaapiNiaMetadataDefault with default headers values
func NewGetNiaapiNiaMetadataDefault(code int) *GetNiaapiNiaMetadataDefault {
	return &GetNiaapiNiaMetadataDefault{
		_statusCode: code,
	}
}

/*GetNiaapiNiaMetadataDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiNiaMetadataDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi nia metadata default response
func (o *GetNiaapiNiaMetadataDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiNiaMetadataDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/NiaMetadata][%d] GetNiaapiNiaMetadata default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiNiaMetadataDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiNiaMetadataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
