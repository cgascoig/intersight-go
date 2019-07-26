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

// PostSoftwareHyperflexDistributablesReader is a Reader for the PostSoftwareHyperflexDistributables structure.
type PostSoftwareHyperflexDistributablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSoftwareHyperflexDistributablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSoftwareHyperflexDistributablesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostSoftwareHyperflexDistributablesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSoftwareHyperflexDistributablesCreated creates a PostSoftwareHyperflexDistributablesCreated with default headers values
func NewPostSoftwareHyperflexDistributablesCreated() *PostSoftwareHyperflexDistributablesCreated {
	return &PostSoftwareHyperflexDistributablesCreated{}
}

/*PostSoftwareHyperflexDistributablesCreated handles this case with default header values.

Null response
*/
type PostSoftwareHyperflexDistributablesCreated struct {
}

func (o *PostSoftwareHyperflexDistributablesCreated) Error() string {
	return fmt.Sprintf("[POST /software/HyperflexDistributables][%d] postSoftwareHyperflexDistributablesCreated ", 201)
}

func (o *PostSoftwareHyperflexDistributablesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSoftwareHyperflexDistributablesDefault creates a PostSoftwareHyperflexDistributablesDefault with default headers values
func NewPostSoftwareHyperflexDistributablesDefault(code int) *PostSoftwareHyperflexDistributablesDefault {
	return &PostSoftwareHyperflexDistributablesDefault{
		_statusCode: code,
	}
}

/*PostSoftwareHyperflexDistributablesDefault handles this case with default header values.

unexpected error
*/
type PostSoftwareHyperflexDistributablesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post software hyperflex distributables default response
func (o *PostSoftwareHyperflexDistributablesDefault) Code() int {
	return o._statusCode
}

func (o *PostSoftwareHyperflexDistributablesDefault) Error() string {
	return fmt.Sprintf("[POST /software/HyperflexDistributables][%d] PostSoftwareHyperflexDistributables default  %+v", o._statusCode, o.Payload)
}

func (o *PostSoftwareHyperflexDistributablesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSoftwareHyperflexDistributablesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
