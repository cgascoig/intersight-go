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

// PostHyperflexClusterNetworkPoliciesReader is a Reader for the PostHyperflexClusterNetworkPolicies structure.
type PostHyperflexClusterNetworkPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexClusterNetworkPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexClusterNetworkPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexClusterNetworkPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexClusterNetworkPoliciesCreated creates a PostHyperflexClusterNetworkPoliciesCreated with default headers values
func NewPostHyperflexClusterNetworkPoliciesCreated() *PostHyperflexClusterNetworkPoliciesCreated {
	return &PostHyperflexClusterNetworkPoliciesCreated{}
}

/*PostHyperflexClusterNetworkPoliciesCreated handles this case with default header values.

Null response
*/
type PostHyperflexClusterNetworkPoliciesCreated struct {
}

func (o *PostHyperflexClusterNetworkPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterNetworkPolicies][%d] postHyperflexClusterNetworkPoliciesCreated ", 201)
}

func (o *PostHyperflexClusterNetworkPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexClusterNetworkPoliciesDefault creates a PostHyperflexClusterNetworkPoliciesDefault with default headers values
func NewPostHyperflexClusterNetworkPoliciesDefault(code int) *PostHyperflexClusterNetworkPoliciesDefault {
	return &PostHyperflexClusterNetworkPoliciesDefault{
		_statusCode: code,
	}
}

/*PostHyperflexClusterNetworkPoliciesDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexClusterNetworkPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex cluster network policies default response
func (o *PostHyperflexClusterNetworkPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexClusterNetworkPoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterNetworkPolicies][%d] PostHyperflexClusterNetworkPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexClusterNetworkPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexClusterNetworkPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}