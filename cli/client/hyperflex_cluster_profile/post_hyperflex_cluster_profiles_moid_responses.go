// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexClusterProfilesMoidReader is a Reader for the PostHyperflexClusterProfilesMoid structure.
type PostHyperflexClusterProfilesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexClusterProfilesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexClusterProfilesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexClusterProfilesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexClusterProfilesMoidCreated creates a PostHyperflexClusterProfilesMoidCreated with default headers values
func NewPostHyperflexClusterProfilesMoidCreated() *PostHyperflexClusterProfilesMoidCreated {
	return &PostHyperflexClusterProfilesMoidCreated{}
}

/*PostHyperflexClusterProfilesMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexClusterProfilesMoidCreated struct {
}

func (o *PostHyperflexClusterProfilesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterProfiles/{moid}][%d] postHyperflexClusterProfilesMoidCreated ", 201)
}

func (o *PostHyperflexClusterProfilesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexClusterProfilesMoidDefault creates a PostHyperflexClusterProfilesMoidDefault with default headers values
func NewPostHyperflexClusterProfilesMoidDefault(code int) *PostHyperflexClusterProfilesMoidDefault {
	return &PostHyperflexClusterProfilesMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexClusterProfilesMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexClusterProfilesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex cluster profiles moid default response
func (o *PostHyperflexClusterProfilesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexClusterProfilesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterProfiles/{moid}][%d] PostHyperflexClusterProfilesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexClusterProfilesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexClusterProfilesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
