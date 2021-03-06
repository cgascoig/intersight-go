// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_network_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetVnicEthNetworkPoliciesReader is a Reader for the GetVnicEthNetworkPolicies structure.
type GetVnicEthNetworkPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVnicEthNetworkPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVnicEthNetworkPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVnicEthNetworkPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVnicEthNetworkPoliciesOK creates a GetVnicEthNetworkPoliciesOK with default headers values
func NewGetVnicEthNetworkPoliciesOK() *GetVnicEthNetworkPoliciesOK {
	return &GetVnicEthNetworkPoliciesOK{}
}

/*GetVnicEthNetworkPoliciesOK handles this case with default header values.

List of vnicEthNetworkPolicies for the given filter criteria
*/
type GetVnicEthNetworkPoliciesOK struct {
	Payload *models.VnicEthNetworkPolicyList
}

func (o *GetVnicEthNetworkPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /vnic/EthNetworkPolicies][%d] getVnicEthNetworkPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetVnicEthNetworkPoliciesOK) GetPayload() *models.VnicEthNetworkPolicyList {
	return o.Payload
}

func (o *GetVnicEthNetworkPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VnicEthNetworkPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVnicEthNetworkPoliciesDefault creates a GetVnicEthNetworkPoliciesDefault with default headers values
func NewGetVnicEthNetworkPoliciesDefault(code int) *GetVnicEthNetworkPoliciesDefault {
	return &GetVnicEthNetworkPoliciesDefault{
		_statusCode: code,
	}
}

/*GetVnicEthNetworkPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetVnicEthNetworkPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get vnic eth network policies default response
func (o *GetVnicEthNetworkPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetVnicEthNetworkPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /vnic/EthNetworkPolicies][%d] GetVnicEthNetworkPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetVnicEthNetworkPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVnicEthNetworkPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
