// Code generated by go-swagger; DO NOT EDIT.

package software_hyperflex_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostSoftwareHyperflexDistributablesMoidReader is a Reader for the PostSoftwareHyperflexDistributablesMoid structure.
type PostSoftwareHyperflexDistributablesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSoftwareHyperflexDistributablesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSoftwareHyperflexDistributablesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostSoftwareHyperflexDistributablesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSoftwareHyperflexDistributablesMoidCreated creates a PostSoftwareHyperflexDistributablesMoidCreated with default headers values
func NewPostSoftwareHyperflexDistributablesMoidCreated() *PostSoftwareHyperflexDistributablesMoidCreated {
	return &PostSoftwareHyperflexDistributablesMoidCreated{}
}

/*PostSoftwareHyperflexDistributablesMoidCreated handles this case with default header values.

Null response
*/
type PostSoftwareHyperflexDistributablesMoidCreated struct {
}

func (o *PostSoftwareHyperflexDistributablesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /software/HyperflexDistributables/{moid}][%d] postSoftwareHyperflexDistributablesMoidCreated ", 201)
}

func (o *PostSoftwareHyperflexDistributablesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSoftwareHyperflexDistributablesMoidDefault creates a PostSoftwareHyperflexDistributablesMoidDefault with default headers values
func NewPostSoftwareHyperflexDistributablesMoidDefault(code int) *PostSoftwareHyperflexDistributablesMoidDefault {
	return &PostSoftwareHyperflexDistributablesMoidDefault{
		_statusCode: code,
	}
}

/*PostSoftwareHyperflexDistributablesMoidDefault handles this case with default header values.

unexpected error
*/
type PostSoftwareHyperflexDistributablesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post software hyperflex distributables moid default response
func (o *PostSoftwareHyperflexDistributablesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostSoftwareHyperflexDistributablesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /software/HyperflexDistributables/{moid}][%d] PostSoftwareHyperflexDistributablesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostSoftwareHyperflexDistributablesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSoftwareHyperflexDistributablesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
