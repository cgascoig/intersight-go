// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_feature_limit_internal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexFeatureLimitInternalsReader is a Reader for the PostHyperflexFeatureLimitInternals structure.
type PostHyperflexFeatureLimitInternalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexFeatureLimitInternalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexFeatureLimitInternalsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexFeatureLimitInternalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexFeatureLimitInternalsCreated creates a PostHyperflexFeatureLimitInternalsCreated with default headers values
func NewPostHyperflexFeatureLimitInternalsCreated() *PostHyperflexFeatureLimitInternalsCreated {
	return &PostHyperflexFeatureLimitInternalsCreated{}
}

/*PostHyperflexFeatureLimitInternalsCreated handles this case with default header values.

Null response
*/
type PostHyperflexFeatureLimitInternalsCreated struct {
}

func (o *PostHyperflexFeatureLimitInternalsCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/FeatureLimitInternals][%d] postHyperflexFeatureLimitInternalsCreated ", 201)
}

func (o *PostHyperflexFeatureLimitInternalsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexFeatureLimitInternalsDefault creates a PostHyperflexFeatureLimitInternalsDefault with default headers values
func NewPostHyperflexFeatureLimitInternalsDefault(code int) *PostHyperflexFeatureLimitInternalsDefault {
	return &PostHyperflexFeatureLimitInternalsDefault{
		_statusCode: code,
	}
}

/*PostHyperflexFeatureLimitInternalsDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexFeatureLimitInternalsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex feature limit internals default response
func (o *PostHyperflexFeatureLimitInternalsDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexFeatureLimitInternalsDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/FeatureLimitInternals][%d] PostHyperflexFeatureLimitInternals default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexFeatureLimitInternalsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexFeatureLimitInternalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
