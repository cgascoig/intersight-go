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

// GetHyperflexClusterNetworkPoliciesMoidReader is a Reader for the GetHyperflexClusterNetworkPoliciesMoid structure.
type GetHyperflexClusterNetworkPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClusterNetworkPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClusterNetworkPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexClusterNetworkPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexClusterNetworkPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClusterNetworkPoliciesMoidOK creates a GetHyperflexClusterNetworkPoliciesMoidOK with default headers values
func NewGetHyperflexClusterNetworkPoliciesMoidOK() *GetHyperflexClusterNetworkPoliciesMoidOK {
	return &GetHyperflexClusterNetworkPoliciesMoidOK{}
}

/*GetHyperflexClusterNetworkPoliciesMoidOK handles this case with default header values.

An instance of hyperflexClusterNetworkPolicy
*/
type GetHyperflexClusterNetworkPoliciesMoidOK struct {
	Payload *models.HyperflexClusterNetworkPolicy
}

func (o *GetHyperflexClusterNetworkPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterNetworkPolicies/{moid}][%d] getHyperflexClusterNetworkPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClusterNetworkPoliciesMoidOK) GetPayload() *models.HyperflexClusterNetworkPolicy {
	return o.Payload
}

func (o *GetHyperflexClusterNetworkPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexClusterNetworkPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClusterNetworkPoliciesMoidNotFound creates a GetHyperflexClusterNetworkPoliciesMoidNotFound with default headers values
func NewGetHyperflexClusterNetworkPoliciesMoidNotFound() *GetHyperflexClusterNetworkPoliciesMoidNotFound {
	return &GetHyperflexClusterNetworkPoliciesMoidNotFound{}
}

/*GetHyperflexClusterNetworkPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexClusterNetworkPoliciesMoidNotFound struct {
}

func (o *GetHyperflexClusterNetworkPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterNetworkPolicies/{moid}][%d] getHyperflexClusterNetworkPoliciesMoidNotFound ", 404)
}

func (o *GetHyperflexClusterNetworkPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexClusterNetworkPoliciesMoidDefault creates a GetHyperflexClusterNetworkPoliciesMoidDefault with default headers values
func NewGetHyperflexClusterNetworkPoliciesMoidDefault(code int) *GetHyperflexClusterNetworkPoliciesMoidDefault {
	return &GetHyperflexClusterNetworkPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClusterNetworkPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClusterNetworkPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex cluster network policies moid default response
func (o *GetHyperflexClusterNetworkPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClusterNetworkPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterNetworkPolicies/{moid}][%d] GetHyperflexClusterNetworkPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClusterNetworkPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClusterNetworkPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
