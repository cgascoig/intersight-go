// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_hxdp_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexHxdpVersionsReader is a Reader for the PostHyperflexHxdpVersions structure.
type PostHyperflexHxdpVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexHxdpVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexHxdpVersionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexHxdpVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexHxdpVersionsCreated creates a PostHyperflexHxdpVersionsCreated with default headers values
func NewPostHyperflexHxdpVersionsCreated() *PostHyperflexHxdpVersionsCreated {
	return &PostHyperflexHxdpVersionsCreated{}
}

/*PostHyperflexHxdpVersionsCreated handles this case with default header values.

Null response
*/
type PostHyperflexHxdpVersionsCreated struct {
}

func (o *PostHyperflexHxdpVersionsCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/HxdpVersions][%d] postHyperflexHxdpVersionsCreated ", 201)
}

func (o *PostHyperflexHxdpVersionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexHxdpVersionsDefault creates a PostHyperflexHxdpVersionsDefault with default headers values
func NewPostHyperflexHxdpVersionsDefault(code int) *PostHyperflexHxdpVersionsDefault {
	return &PostHyperflexHxdpVersionsDefault{
		_statusCode: code,
	}
}

/*PostHyperflexHxdpVersionsDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexHxdpVersionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex hxdp versions default response
func (o *PostHyperflexHxdpVersionsDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexHxdpVersionsDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/HxdpVersions][%d] PostHyperflexHxdpVersions default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexHxdpVersionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexHxdpVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
