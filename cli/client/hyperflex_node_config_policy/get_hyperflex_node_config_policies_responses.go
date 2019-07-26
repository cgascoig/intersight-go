// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_node_config_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexNodeConfigPoliciesReader is a Reader for the GetHyperflexNodeConfigPolicies structure.
type GetHyperflexNodeConfigPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexNodeConfigPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexNodeConfigPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexNodeConfigPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexNodeConfigPoliciesOK creates a GetHyperflexNodeConfigPoliciesOK with default headers values
func NewGetHyperflexNodeConfigPoliciesOK() *GetHyperflexNodeConfigPoliciesOK {
	return &GetHyperflexNodeConfigPoliciesOK{}
}

/*GetHyperflexNodeConfigPoliciesOK handles this case with default header values.

List of hyperflexNodeConfigPolicies for the given filter criteria
*/
type GetHyperflexNodeConfigPoliciesOK struct {
	Payload *models.HyperflexNodeConfigPolicyList
}

func (o *GetHyperflexNodeConfigPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/NodeConfigPolicies][%d] getHyperflexNodeConfigPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexNodeConfigPoliciesOK) GetPayload() *models.HyperflexNodeConfigPolicyList {
	return o.Payload
}

func (o *GetHyperflexNodeConfigPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexNodeConfigPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexNodeConfigPoliciesDefault creates a GetHyperflexNodeConfigPoliciesDefault with default headers values
func NewGetHyperflexNodeConfigPoliciesDefault(code int) *GetHyperflexNodeConfigPoliciesDefault {
	return &GetHyperflexNodeConfigPoliciesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexNodeConfigPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexNodeConfigPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex node config policies default response
func (o *GetHyperflexNodeConfigPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexNodeConfigPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/NodeConfigPolicies][%d] GetHyperflexNodeConfigPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexNodeConfigPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexNodeConfigPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
