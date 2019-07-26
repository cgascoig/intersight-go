// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexClustersMoidReader is a Reader for the PostHyperflexClustersMoid structure.
type PostHyperflexClustersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexClustersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexClustersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexClustersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexClustersMoidCreated creates a PostHyperflexClustersMoidCreated with default headers values
func NewPostHyperflexClustersMoidCreated() *PostHyperflexClustersMoidCreated {
	return &PostHyperflexClustersMoidCreated{}
}

/*PostHyperflexClustersMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexClustersMoidCreated struct {
}

func (o *PostHyperflexClustersMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/Clusters/{moid}][%d] postHyperflexClustersMoidCreated ", 201)
}

func (o *PostHyperflexClustersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexClustersMoidDefault creates a PostHyperflexClustersMoidDefault with default headers values
func NewPostHyperflexClustersMoidDefault(code int) *PostHyperflexClustersMoidDefault {
	return &PostHyperflexClustersMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexClustersMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexClustersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex clusters moid default response
func (o *PostHyperflexClustersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexClustersMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/Clusters/{moid}][%d] PostHyperflexClustersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexClustersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexClustersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
