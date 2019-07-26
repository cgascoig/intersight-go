// Code generated by go-swagger; DO NOT EDIT.

package cvd_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetCvdTemplatesMoidReader is a Reader for the GetCvdTemplatesMoid structure.
type GetCvdTemplatesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCvdTemplatesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCvdTemplatesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCvdTemplatesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCvdTemplatesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCvdTemplatesMoidOK creates a GetCvdTemplatesMoidOK with default headers values
func NewGetCvdTemplatesMoidOK() *GetCvdTemplatesMoidOK {
	return &GetCvdTemplatesMoidOK{}
}

/*GetCvdTemplatesMoidOK handles this case with default header values.

An instance of cvdTemplate
*/
type GetCvdTemplatesMoidOK struct {
	Payload *models.CvdTemplate
}

func (o *GetCvdTemplatesMoidOK) Error() string {
	return fmt.Sprintf("[GET /cvd/Templates/{moid}][%d] getCvdTemplatesMoidOK  %+v", 200, o.Payload)
}

func (o *GetCvdTemplatesMoidOK) GetPayload() *models.CvdTemplate {
	return o.Payload
}

func (o *GetCvdTemplatesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CvdTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCvdTemplatesMoidNotFound creates a GetCvdTemplatesMoidNotFound with default headers values
func NewGetCvdTemplatesMoidNotFound() *GetCvdTemplatesMoidNotFound {
	return &GetCvdTemplatesMoidNotFound{}
}

/*GetCvdTemplatesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetCvdTemplatesMoidNotFound struct {
}

func (o *GetCvdTemplatesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /cvd/Templates/{moid}][%d] getCvdTemplatesMoidNotFound ", 404)
}

func (o *GetCvdTemplatesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCvdTemplatesMoidDefault creates a GetCvdTemplatesMoidDefault with default headers values
func NewGetCvdTemplatesMoidDefault(code int) *GetCvdTemplatesMoidDefault {
	return &GetCvdTemplatesMoidDefault{
		_statusCode: code,
	}
}

/*GetCvdTemplatesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetCvdTemplatesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cvd templates moid default response
func (o *GetCvdTemplatesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetCvdTemplatesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /cvd/Templates/{moid}][%d] GetCvdTemplatesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCvdTemplatesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCvdTemplatesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}