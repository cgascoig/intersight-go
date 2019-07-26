// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_network_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexClusterNetworkPoliciesReader is a Reader for the GetHyperflexClusterNetworkPolicies structure.
type GetHyperflexClusterNetworkPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClusterNetworkPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClusterNetworkPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexClusterNetworkPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClusterNetworkPoliciesOK creates a GetHyperflexClusterNetworkPoliciesOK with default headers values
func NewGetHyperflexClusterNetworkPoliciesOK() *GetHyperflexClusterNetworkPoliciesOK {
	return &GetHyperflexClusterNetworkPoliciesOK{}
}

/*GetHyperflexClusterNetworkPoliciesOK handles this case with default header values.

List of hyperflexClusterNetworkPolicies for the given filter criteria
*/
type GetHyperflexClusterNetworkPoliciesOK struct {
	Payload *models.HyperflexClusterNetworkPolicyList
}

func (o *GetHyperflexClusterNetworkPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterNetworkPolicies][%d] getHyperflexClusterNetworkPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClusterNetworkPoliciesOK) GetPayload() *models.HyperflexClusterNetworkPolicyList {
	return o.Payload
}

func (o *GetHyperflexClusterNetworkPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexClusterNetworkPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClusterNetworkPoliciesDefault creates a GetHyperflexClusterNetworkPoliciesDefault with default headers values
func NewGetHyperflexClusterNetworkPoliciesDefault(code int) *GetHyperflexClusterNetworkPoliciesDefault {
	return &GetHyperflexClusterNetworkPoliciesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClusterNetworkPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClusterNetworkPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex cluster network policies default response
func (o *GetHyperflexClusterNetworkPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClusterNetworkPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterNetworkPolicies][%d] GetHyperflexClusterNetworkPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClusterNetworkPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClusterNetworkPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}