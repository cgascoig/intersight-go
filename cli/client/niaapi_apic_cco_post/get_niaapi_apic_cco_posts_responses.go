// Code generated by go-swagger; DO NOT EDIT.

package niaapi_apic_cco_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNiaapiApicCcoPostsReader is a Reader for the GetNiaapiApicCcoPosts structure.
type GetNiaapiApicCcoPostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNiaapiApicCcoPostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNiaapiApicCcoPostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNiaapiApicCcoPostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNiaapiApicCcoPostsOK creates a GetNiaapiApicCcoPostsOK with default headers values
func NewGetNiaapiApicCcoPostsOK() *GetNiaapiApicCcoPostsOK {
	return &GetNiaapiApicCcoPostsOK{}
}

/*GetNiaapiApicCcoPostsOK handles this case with default header values.

List of niaapiApicCcoPosts for the given filter criteria
*/
type GetNiaapiApicCcoPostsOK struct {
	Payload *models.NiaapiApicCcoPostList
}

func (o *GetNiaapiApicCcoPostsOK) Error() string {
	return fmt.Sprintf("[GET /niaapi/ApicCcoPosts][%d] getNiaapiApicCcoPostsOK  %+v", 200, o.Payload)
}

func (o *GetNiaapiApicCcoPostsOK) GetPayload() *models.NiaapiApicCcoPostList {
	return o.Payload
}

func (o *GetNiaapiApicCcoPostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NiaapiApicCcoPostList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNiaapiApicCcoPostsDefault creates a GetNiaapiApicCcoPostsDefault with default headers values
func NewGetNiaapiApicCcoPostsDefault(code int) *GetNiaapiApicCcoPostsDefault {
	return &GetNiaapiApicCcoPostsDefault{
		_statusCode: code,
	}
}

/*GetNiaapiApicCcoPostsDefault handles this case with default header values.

Unexpected error
*/
type GetNiaapiApicCcoPostsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get niaapi apic cco posts default response
func (o *GetNiaapiApicCcoPostsDefault) Code() int {
	return o._statusCode
}

func (o *GetNiaapiApicCcoPostsDefault) Error() string {
	return fmt.Sprintf("[GET /niaapi/ApicCcoPosts][%d] GetNiaapiApicCcoPosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetNiaapiApicCcoPostsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNiaapiApicCcoPostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}