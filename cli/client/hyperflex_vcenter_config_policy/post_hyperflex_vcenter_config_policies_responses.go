// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_vcenter_config_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexVcenterConfigPoliciesReader is a Reader for the PostHyperflexVcenterConfigPolicies structure.
type PostHyperflexVcenterConfigPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexVcenterConfigPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexVcenterConfigPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexVcenterConfigPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexVcenterConfigPoliciesCreated creates a PostHyperflexVcenterConfigPoliciesCreated with default headers values
func NewPostHyperflexVcenterConfigPoliciesCreated() *PostHyperflexVcenterConfigPoliciesCreated {
	return &PostHyperflexVcenterConfigPoliciesCreated{}
}

/*PostHyperflexVcenterConfigPoliciesCreated handles this case with default header values.

Null response
*/
type PostHyperflexVcenterConfigPoliciesCreated struct {
}

func (o *PostHyperflexVcenterConfigPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/VcenterConfigPolicies][%d] postHyperflexVcenterConfigPoliciesCreated ", 201)
}

func (o *PostHyperflexVcenterConfigPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexVcenterConfigPoliciesDefault creates a PostHyperflexVcenterConfigPoliciesDefault with default headers values
func NewPostHyperflexVcenterConfigPoliciesDefault(code int) *PostHyperflexVcenterConfigPoliciesDefault {
	return &PostHyperflexVcenterConfigPoliciesDefault{
		_statusCode: code,
	}
}

/*PostHyperflexVcenterConfigPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexVcenterConfigPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex vcenter config policies default response
func (o *PostHyperflexVcenterConfigPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexVcenterConfigPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/VcenterConfigPolicies][%d] PostHyperflexVcenterConfigPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexVcenterConfigPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexVcenterConfigPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
