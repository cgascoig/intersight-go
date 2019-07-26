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

// GetHyperflexLocalCredentialPoliciesReader is a Reader for the GetHyperflexLocalCredentialPolicies structure.
type GetHyperflexLocalCredentialPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexLocalCredentialPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexLocalCredentialPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexLocalCredentialPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexLocalCredentialPoliciesOK creates a GetHyperflexLocalCredentialPoliciesOK with default headers values
func NewGetHyperflexLocalCredentialPoliciesOK() *GetHyperflexLocalCredentialPoliciesOK {
	return &GetHyperflexLocalCredentialPoliciesOK{}
}

/*GetHyperflexLocalCredentialPoliciesOK handles this case with default header values.

List of hyperflexLocalCredentialPolicies for the given filter criteria
*/
type GetHyperflexLocalCredentialPoliciesOK struct {
	Payload *models.HyperflexLocalCredentialPolicyList
}

func (o *GetHyperflexLocalCredentialPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/LocalCredentialPolicies][%d] getHyperflexLocalCredentialPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexLocalCredentialPoliciesOK) GetPayload() *models.HyperflexLocalCredentialPolicyList {
	return o.Payload
}

func (o *GetHyperflexLocalCredentialPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexLocalCredentialPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexLocalCredentialPoliciesDefault creates a GetHyperflexLocalCredentialPoliciesDefault with default headers values
func NewGetHyperflexLocalCredentialPoliciesDefault(code int) *GetHyperflexLocalCredentialPoliciesDefault {
	return &GetHyperflexLocalCredentialPoliciesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexLocalCredentialPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexLocalCredentialPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex local credential policies default response
func (o *GetHyperflexLocalCredentialPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexLocalCredentialPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/LocalCredentialPolicies][%d] GetHyperflexLocalCredentialPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexLocalCredentialPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexLocalCredentialPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}