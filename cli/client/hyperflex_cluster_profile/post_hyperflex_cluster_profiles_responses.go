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

// PostHyperflexClusterProfilesReader is a Reader for the PostHyperflexClusterProfiles structure.
type PostHyperflexClusterProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexClusterProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexClusterProfilesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexClusterProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexClusterProfilesCreated creates a PostHyperflexClusterProfilesCreated with default headers values
func NewPostHyperflexClusterProfilesCreated() *PostHyperflexClusterProfilesCreated {
	return &PostHyperflexClusterProfilesCreated{}
}

/*PostHyperflexClusterProfilesCreated handles this case with default header values.

Null response
*/
type PostHyperflexClusterProfilesCreated struct {
}

func (o *PostHyperflexClusterProfilesCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterProfiles][%d] postHyperflexClusterProfilesCreated ", 201)
}

func (o *PostHyperflexClusterProfilesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexClusterProfilesDefault creates a PostHyperflexClusterProfilesDefault with default headers values
func NewPostHyperflexClusterProfilesDefault(code int) *PostHyperflexClusterProfilesDefault {
	return &PostHyperflexClusterProfilesDefault{
		_statusCode: code,
	}
}

/*PostHyperflexClusterProfilesDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexClusterProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex cluster profiles default response
func (o *PostHyperflexClusterProfilesDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexClusterProfilesDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ClusterProfiles][%d] PostHyperflexClusterProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexClusterProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexClusterProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
