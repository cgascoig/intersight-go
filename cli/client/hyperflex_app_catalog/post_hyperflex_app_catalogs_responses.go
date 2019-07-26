// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_app_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostHyperflexAppCatalogsReader is a Reader for the PostHyperflexAppCatalogs structure.
type PostHyperflexAppCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHyperflexAppCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHyperflexAppCatalogsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHyperflexAppCatalogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHyperflexAppCatalogsCreated creates a PostHyperflexAppCatalogsCreated with default headers values
func NewPostHyperflexAppCatalogsCreated() *PostHyperflexAppCatalogsCreated {
	return &PostHyperflexAppCatalogsCreated{}
}

/*PostHyperflexAppCatalogsCreated handles this case with default header values.

Null response
*/
type PostHyperflexAppCatalogsCreated struct {
}

func (o *PostHyperflexAppCatalogsCreated) Error() string {
	return fmt.Sprintf("[POST /hyperflex/AppCatalogs][%d] postHyperflexAppCatalogsCreated ", 201)
}

func (o *PostHyperflexAppCatalogsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHyperflexAppCatalogsDefault creates a PostHyperflexAppCatalogsDefault with default headers values
func NewPostHyperflexAppCatalogsDefault(code int) *PostHyperflexAppCatalogsDefault {
	return &PostHyperflexAppCatalogsDefault{
		_statusCode: code,
	}
}

/*PostHyperflexAppCatalogsDefault handles this case with default header values.

unexpected error
*/
type PostHyperflexAppCatalogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post hyperflex app catalogs default response
func (o *PostHyperflexAppCatalogsDefault) Code() int {
	return o._statusCode
}

func (o *PostHyperflexAppCatalogsDefault) Error() string {
	return fmt.Sprintf("[POST /hyperflex/AppCatalogs][%d] PostHyperflexAppCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *PostHyperflexAppCatalogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHyperflexAppCatalogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}