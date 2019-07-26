// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ucsm_config_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexUcsmConfigPoliciesReader is a Reader for the GetHyperflexUcsmConfigPolicies structure.
type GetHyperflexUcsmConfigPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexUcsmConfigPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexUcsmConfigPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexUcsmConfigPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexUcsmConfigPoliciesOK creates a GetHyperflexUcsmConfigPoliciesOK with default headers values
func NewGetHyperflexUcsmConfigPoliciesOK() *GetHyperflexUcsmConfigPoliciesOK {
	return &GetHyperflexUcsmConfigPoliciesOK{}
}

/*GetHyperflexUcsmConfigPoliciesOK handles this case with default header values.

List of hyperflexUcsmConfigPolicies for the given filter criteria
*/
type GetHyperflexUcsmConfigPoliciesOK struct {
	Payload *models.HyperflexUcsmConfigPolicyList
}

func (o *GetHyperflexUcsmConfigPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/UcsmConfigPolicies][%d] getHyperflexUcsmConfigPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexUcsmConfigPoliciesOK) GetPayload() *models.HyperflexUcsmConfigPolicyList {
	return o.Payload
}

func (o *GetHyperflexUcsmConfigPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexUcsmConfigPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexUcsmConfigPoliciesDefault creates a GetHyperflexUcsmConfigPoliciesDefault with default headers values
func NewGetHyperflexUcsmConfigPoliciesDefault(code int) *GetHyperflexUcsmConfigPoliciesDefault {
	return &GetHyperflexUcsmConfigPoliciesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexUcsmConfigPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexUcsmConfigPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex ucsm config policies default response
func (o *GetHyperflexUcsmConfigPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexUcsmConfigPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/UcsmConfigPolicies][%d] GetHyperflexUcsmConfigPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexUcsmConfigPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexUcsmConfigPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
