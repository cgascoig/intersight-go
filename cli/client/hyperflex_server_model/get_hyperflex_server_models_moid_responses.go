// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexServerModelsMoidReader is a Reader for the GetHyperflexServerModelsMoid structure.
type GetHyperflexServerModelsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexServerModelsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexServerModelsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexServerModelsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexServerModelsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexServerModelsMoidOK creates a GetHyperflexServerModelsMoidOK with default headers values
func NewGetHyperflexServerModelsMoidOK() *GetHyperflexServerModelsMoidOK {
	return &GetHyperflexServerModelsMoidOK{}
}

/*GetHyperflexServerModelsMoidOK handles this case with default header values.

An instance of hyperflexServerModel
*/
type GetHyperflexServerModelsMoidOK struct {
	Payload *models.HyperflexServerModel
}

func (o *GetHyperflexServerModelsMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ServerModels/{moid}][%d] getHyperflexServerModelsMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexServerModelsMoidOK) GetPayload() *models.HyperflexServerModel {
	return o.Payload
}

func (o *GetHyperflexServerModelsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexServerModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexServerModelsMoidNotFound creates a GetHyperflexServerModelsMoidNotFound with default headers values
func NewGetHyperflexServerModelsMoidNotFound() *GetHyperflexServerModelsMoidNotFound {
	return &GetHyperflexServerModelsMoidNotFound{}
}

/*GetHyperflexServerModelsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexServerModelsMoidNotFound struct {
}

func (o *GetHyperflexServerModelsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ServerModels/{moid}][%d] getHyperflexServerModelsMoidNotFound ", 404)
}

func (o *GetHyperflexServerModelsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexServerModelsMoidDefault creates a GetHyperflexServerModelsMoidDefault with default headers values
func NewGetHyperflexServerModelsMoidDefault(code int) *GetHyperflexServerModelsMoidDefault {
	return &GetHyperflexServerModelsMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexServerModelsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexServerModelsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex server models moid default response
func (o *GetHyperflexServerModelsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexServerModelsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ServerModels/{moid}][%d] GetHyperflexServerModelsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexServerModelsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexServerModelsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
