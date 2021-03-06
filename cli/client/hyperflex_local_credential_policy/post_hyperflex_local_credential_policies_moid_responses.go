// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_local_credential_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexLocalCredentialPoliciesMoidReader is a Reader for the PostHyperflexLocalCredentialPoliciesMoid structure.
type PostHyperflexLocalCredentialPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexLocalCredentialPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexLocalCredentialPoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexLocalCredentialPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexLocalCredentialPoliciesMoidCreated creates a PostHyperflexLocalCredentialPoliciesMoidCreated with default headers values
func NewPostHyperflexLocalCredentialPoliciesMoidCreated() *PostHyperflexLocalCredentialPoliciesMoidCreated {
	return &PostHyperflexLocalCredentialPoliciesMoidCreated{}
}

/*PostHyperflexLocalCredentialPoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexLocalCredentialPoliciesMoidCreated struct {
}

func (o *PostHyperflexLocalCredentialPoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/LocalCredentialPolicies/{moid}][%d] postHyperflexLocalCredentialPoliciesMoidCreated ", 201)
}

func (o *PostHyperflexLocalCredentialPoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexLocalCredentialPoliciesMoidDefault creates a PostHyperflexLocalCredentialPoliciesMoidDefault with default headers values
func NewPostHyperflexLocalCredentialPoliciesMoidDefault(code int) *PostHyperflexLocalCredentialPoliciesMoidDefault {
	return &PostHyperflexLocalCredentialPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexLocalCredentialPoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexLocalCredentialPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex local credential policies moid default response
func (o *PostHyperflexLocalCredentialPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexLocalCredentialPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/LocalCredentialPolicies/{moid}][%d] PostHyperflexLocalCredentialPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexLocalCredentialPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexLocalCredentialPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
