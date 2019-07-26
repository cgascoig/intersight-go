// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_storage_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexClusterStoragePoliciesReader is a Reader for the GetHyperflexClusterStoragePolicies structure.
type GetHyperflexClusterStoragePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClusterStoragePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClusterStoragePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexClusterStoragePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClusterStoragePoliciesOK creates a GetHyperflexClusterStoragePoliciesOK with default headers values
func NewGetHyperflexClusterStoragePoliciesOK() *GetHyperflexClusterStoragePoliciesOK {
	return &GetHyperflexClusterStoragePoliciesOK{}
}

/*GetHyperflexClusterStoragePoliciesOK handles this case with default header values.

List of hyperflexClusterStoragePolicies for the given filter criteria
*/
type GetHyperflexClusterStoragePoliciesOK struct {
	Payload *models.HyperflexClusterStoragePolicyList
}

func (o *GetHyperflexClusterStoragePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterStoragePolicies][%d] getHyperflexClusterStoragePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClusterStoragePoliciesOK) GetPayload() *models.HyperflexClusterStoragePolicyList {
	return o.Payload
}

func (o *GetHyperflexClusterStoragePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexClusterStoragePolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClusterStoragePoliciesDefault creates a GetHyperflexClusterStoragePoliciesDefault with default headers values
func NewGetHyperflexClusterStoragePoliciesDefault(code int) *GetHyperflexClusterStoragePoliciesDefault {
	return &GetHyperflexClusterStoragePoliciesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClusterStoragePoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClusterStoragePoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex cluster storage policies default response
func (o *GetHyperflexClusterStoragePoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClusterStoragePoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterStoragePolicies][%d] GetHyperflexClusterStoragePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClusterStoragePoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClusterStoragePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
