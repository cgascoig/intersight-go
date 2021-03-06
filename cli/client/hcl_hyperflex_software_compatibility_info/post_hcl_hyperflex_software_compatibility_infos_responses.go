// Code generated by go-swagger; DO NOT EDIT.

package hcl_hyperflex_software_compatibility_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHclHyperflexSoftwareCompatibilityInfosReader is a Reader for the PostHclHyperflexSoftwareCompatibilityInfos structure.
type PostHclHyperflexSoftwareCompatibilityInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHclHyperflexSoftwareCompatibilityInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHclHyperflexSoftwareCompatibilityInfosCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHclHyperflexSoftwareCompatibilityInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHclHyperflexSoftwareCompatibilityInfosCreated creates a PostHclHyperflexSoftwareCompatibilityInfosCreated with default headers values
func NewPostHclHyperflexSoftwareCompatibilityInfosCreated() *PostHclHyperflexSoftwareCompatibilityInfosCreated {
	return &PostHclHyperflexSoftwareCompatibilityInfosCreated{}
}

/*PostHclHyperflexSoftwareCompatibilityInfosCreated handles this case with default header values.

Null response
*/
type PostHclHyperflexSoftwareCompatibilityInfosCreated struct {
}

func (o *PostHclHyperflexSoftwareCompatibilityInfosCreated) Error() string {
	return fmt.Sprintf("[POST /hcl/HyperflexSoftwareCompatibilityInfos][%d] postHclHyperflexSoftwareCompatibilityInfosCreated ", 201)
}

func (o *PostHclHyperflexSoftwareCompatibilityInfosCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHclHyperflexSoftwareCompatibilityInfosDefault creates a PostHclHyperflexSoftwareCompatibilityInfosDefault with default headers values
func NewPostHclHyperflexSoftwareCompatibilityInfosDefault(code int) *PostHclHyperflexSoftwareCompatibilityInfosDefault {
	return &PostHclHyperflexSoftwareCompatibilityInfosDefault{
		_statusCode: code,
	}
}

/*PostHclHyperflexSoftwareCompatibilityInfosDefault handles this case with default header values.

unexpected error
*/
type PostHclHyperflexSoftwareCompatibilityInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hcl hyperflex software compatibility infos default response
func (o *PostHclHyperflexSoftwareCompatibilityInfosDefault) Code() int {
	return o._statusCode
}

func (o *PostHclHyperflexSoftwareCompatibilityInfosDefault) Error() string {
	return fmt.Sprintf("[POST /hcl/HyperflexSoftwareCompatibilityInfos][%d] PostHclHyperflexSoftwareCompatibilityInfos default  %+v", o._statusCode, o.Payload)
}

func (o *PostHclHyperflexSoftwareCompatibilityInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHclHyperflexSoftwareCompatibilityInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
