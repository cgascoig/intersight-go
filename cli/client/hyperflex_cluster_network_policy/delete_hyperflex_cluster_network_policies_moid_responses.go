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

// DeleteHyperflexClusterNetworkPoliciesMoidReader is a Reader for the DeleteHyperflexClusterNetworkPoliciesMoid structure.
type DeleteHyperflexClusterNetworkPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHyperflexClusterNetworkPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHyperflexClusterNetworkPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHyperflexClusterNetworkPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHyperflexClusterNetworkPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHyperflexClusterNetworkPoliciesMoidOK creates a DeleteHyperflexClusterNetworkPoliciesMoidOK with default headers values
func NewDeleteHyperflexClusterNetworkPoliciesMoidOK() *DeleteHyperflexClusterNetworkPoliciesMoidOK {
	return &DeleteHyperflexClusterNetworkPoliciesMoidOK{}
}

/*DeleteHyperflexClusterNetworkPoliciesMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteHyperflexClusterNetworkPoliciesMoidOK struct {
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ClusterNetworkPolicies/{moid}][%d] deleteHyperflexClusterNetworkPoliciesMoidOK ", 200)
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexClusterNetworkPoliciesMoidNotFound creates a DeleteHyperflexClusterNetworkPoliciesMoidNotFound with default headers values
func NewDeleteHyperflexClusterNetworkPoliciesMoidNotFound() *DeleteHyperflexClusterNetworkPoliciesMoidNotFound {
	return &DeleteHyperflexClusterNetworkPoliciesMoidNotFound{}
}

/*DeleteHyperflexClusterNetworkPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteHyperflexClusterNetworkPoliciesMoidNotFound struct {
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ClusterNetworkPolicies/{moid}][%d] deleteHyperflexClusterNetworkPoliciesMoidNotFound ", 404)
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHyperflexClusterNetworkPoliciesMoidDefault creates a DeleteHyperflexClusterNetworkPoliciesMoidDefault with default headers values
func NewDeleteHyperflexClusterNetworkPoliciesMoidDefault(code int) *DeleteHyperflexClusterNetworkPoliciesMoidDefault {
	return &DeleteHyperflexClusterNetworkPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*DeleteHyperflexClusterNetworkPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteHyperflexClusterNetworkPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hyperflex cluster network policies moid default response
func (o *DeleteHyperflexClusterNetworkPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /hyperflex/ClusterNetworkPolicies/{moid}][%d] DeleteHyperflexClusterNetworkPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHyperflexClusterNetworkPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
