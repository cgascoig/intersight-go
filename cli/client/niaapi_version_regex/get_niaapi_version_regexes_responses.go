// Code generated by go-swagger; DO NOT EDIT.

package niaapi_version_regex

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiVersionRegexesReader is a Reader for the GetNiaapiVersionRegexes structure.
type GetNiaapiVersionRegexesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiVersionRegexesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiVersionRegexesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiVersionRegexesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiVersionRegexesOK creates a GetNiaapiVersionRegexesOK with default headers values
func NewGetNiaapiVersionRegexesOK() *GetNiaapiVersionRegexesOK {
	return &GetNiaapiVersionRegexesOK{}
}

/*GetNiaapiVersionRegexesOK handles this case with default header values.

List of niaapiVersionRegexes for the given filter criteria
*/
type GetNiaapiVersionRegexesOK struct {
	Payload *models.NiaapiVersionRegexList
}

func (o *GetNiaapiVersionRegexesOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/VersionRegexes][%d] getNiaapiVersionRegexesOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiVersionRegexesOK) GetPayload() *models.NiaapiVersionRegexList {
	return o.Payload
}

func (o *GetNiaapiVersionRegexesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiVersionRegexList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiVersionRegexesDefault creates a GetNiaapiVersionRegexesDefault with default headers values
func NewGetNiaapiVersionRegexesDefault(code int) *GetNiaapiVersionRegexesDefault {
	return &GetNiaapiVersionRegexesDefault{
		_statusCode: code,
	}
}

/*GetNiaapiVersionRegexesDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiVersionRegexesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi version regexes default response
func (o *GetNiaapiVersionRegexesDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiVersionRegexesDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/VersionRegexes][%d] GetNiaapiVersionRegexes default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiVersionRegexesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiVersionRegexesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
