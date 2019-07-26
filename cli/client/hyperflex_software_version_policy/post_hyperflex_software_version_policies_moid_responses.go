// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_software_version_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexSoftwareVersionPoliciesMoidReader is a Reader for the PostHyperflexSoftwareVersionPoliciesMoid structure.
type PostHyperflexSoftwareVersionPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexSoftwareVersionPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexSoftwareVersionPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexSoftwareVersionPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexSoftwareVersionPoliciesMoidCreated creates a PostHyperflexSoftwareVersionPoliciesMoidCreated with default headers values
func NewPostHyperflexSoftwareVersionPoliciesMoidCreated() *PostHyperflexSoftwareVersionPoliciesMoidCreated {
	return &PostHyperflexSoftwareVersionPoliciesMoidCreated{}
}

/*PostHyperflexSoftwareVersionPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexSoftwareVersionPoliciesMoidCreated struct {
}

func (o *PostHyperflexSoftwareVersionPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/SoftwareVersionPolicies/{moid}][%d] postHyperflexSoftwareVersionPoliciesMoidCreated ", 201)
}

func (o *PostHyperflexSoftwareVersionPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexSoftwareVersionPoliciesMoidDefault creates a PostHyperflexSoftwareVersionPoliciesMoidDefault with default headers values
func NewPostHyperflexSoftwareVersionPoliciesMoidDefault(code int) *PostHyperflexSoftwareVersionPoliciesMoidDefault {
	return &PostHyperflexSoftwareVersionPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexSoftwareVersionPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexSoftwareVersionPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex software version policies moid default response
func (o *PostHyperflexSoftwareVersionPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexSoftwareVersionPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/SoftwareVersionPolicies/{moid}][%d] PostHyperflexSoftwareVersionPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexSoftwareVersionPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexSoftwareVersionPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}