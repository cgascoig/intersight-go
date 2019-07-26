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

// PostHyperflexClusterStoragePoliciesMoidReader is a Reader for the PostHyperflexClusterStoragePoliciesMoid structure.
type PostHyperflexClusterStoragePoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexClusterStoragePoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexClusterStoragePoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexClusterStoragePoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexClusterStoragePoliciesMoidCreated creates a PostHyperflexClusterStoragePoliciesMoidCreated with default headers values
func NewPostHyperflexClusterStoragePoliciesMoidCreated() *PostHyperflexClusterStoragePoliciesMoidCreated {
	return &PostHyperflexClusterStoragePoliciesMoidCreated{}
}

/*PostHyperflexClusterStoragePoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexClusterStoragePoliciesMoidCreated struct {
}

func (o *PostHyperflexClusterStoragePoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterStoragePolicies/{moid}][%d] postHyperflexClusterStoragePoliciesMoidCreated ", 201)
}

func (o *PostHyperflexClusterStoragePoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexClusterStoragePoliciesMoidDefault creates a PostHyperflexClusterStoragePoliciesMoidDefault with default headers values
func NewPostHyperflexClusterStoragePoliciesMoidDefault(code int) *PostHyperflexClusterStoragePoliciesMoidDefault {
	return &PostHyperflexClusterStoragePoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexClusterStoragePoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexClusterStoragePoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex cluster storage policies moid default response
func (o *PostHyperflexClusterStoragePoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexClusterStoragePoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterStoragePolicies/{moid}][%d] PostHyperflexClusterStoragePoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexClusterStoragePoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexClusterStoragePoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
