// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_ext_fc_storage_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexExtFcStoragePoliciesMoidReader is a Reader for the PostHyperflexExtFcStoragePoliciesMoid structure.
type PostHyperflexExtFcStoragePoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexExtFcStoragePoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexExtFcStoragePoliciesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexExtFcStoragePoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexExtFcStoragePoliciesMoidCreated creates a PostHyperflexExtFcStoragePoliciesMoidCreated with default headers values
func NewPostHyperflexExtFcStoragePoliciesMoidCreated() *PostHyperflexExtFcStoragePoliciesMoidCreated {
	return &PostHyperflexExtFcStoragePoliciesMoidCreated{}
}

/*PostHyperflexExtFcStoragePoliciesMoidCreated handles this case with default header values.

Null response
*/
type PostHyperflexExtFcStoragePoliciesMoidCreated struct {
}

func (o *PostHyperflexExtFcStoragePoliciesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ExtFcStoragePolicies/{moid}][%d] postHyperflexExtFcStoragePoliciesMoidCreated ", 201)
}

func (o *PostHyperflexExtFcStoragePoliciesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexExtFcStoragePoliciesMoidDefault creates a PostHyperflexExtFcStoragePoliciesMoidDefault with default headers values
func NewPostHyperflexExtFcStoragePoliciesMoidDefault(code int) *PostHyperflexExtFcStoragePoliciesMoidDefault {
	return &PostHyperflexExtFcStoragePoliciesMoidDefault{
		_statusCode: code,
	}
}

/*PostHyperflexExtFcStoragePoliciesMoidDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexExtFcStoragePoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex ext fc storage policies moid default response
func (o *PostHyperflexExtFcStoragePoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexExtFcStoragePoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/ExtFcStoragePolicies/{moid}][%d] PostHyperflexExtFcStoragePoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexExtFcStoragePoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexExtFcStoragePoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
