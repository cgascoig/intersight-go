// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_capability_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexCapabilityInfosMoidReader is a Reader for the GetHyperflexCapabilityInfosMoid structure.
type GetHyperflexCapabilityInfosMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexCapabilityInfosMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexCapabilityInfosMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexCapabilityInfosMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexCapabilityInfosMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexCapabilityInfosMoidOK creates a GetHyperflexCapabilityInfosMoidOK with default headers values
func NewGetHyperflexCapabilityInfosMoidOK() *GetHyperflexCapabilityInfosMoidOK {
	return &GetHyperflexCapabilityInfosMoidOK{}
}

/*GetHyperflexCapabilityInfosMoidOK handles this case with default header values.

An instance of hyperflexCapabilityInfo
*/
type GetHyperflexCapabilityInfosMoidOK struct {
	Payload *models.HyperflexCapabilityInfo
}

func (o *GetHyperflexCapabilityInfosMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/CapabilityInfos/{moid}][%d] getHyperflexCapabilityInfosMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexCapabilityInfosMoidOK) GetPayload() *models.HyperflexCapabilityInfo {
	return o.Payload
}

func (o *GetHyperflexCapabilityInfosMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexCapabilityInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexCapabilityInfosMoidNotFound creates a GetHyperflexCapabilityInfosMoidNotFound with default headers values
func NewGetHyperflexCapabilityInfosMoidNotFound() *GetHyperflexCapabilityInfosMoidNotFound {
	return &GetHyperflexCapabilityInfosMoidNotFound{}
}

/*GetHyperflexCapabilityInfosMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexCapabilityInfosMoidNotFound struct {
}

func (o *GetHyperflexCapabilityInfosMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/CapabilityInfos/{moid}][%d] getHyperflexCapabilityInfosMoidNotFound ", 404)
}

func (o *GetHyperflexCapabilityInfosMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexCapabilityInfosMoidDefault creates a GetHyperflexCapabilityInfosMoidDefault with default headers values
func NewGetHyperflexCapabilityInfosMoidDefault(code int) *GetHyperflexCapabilityInfosMoidDefault {
	return &GetHyperflexCapabilityInfosMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexCapabilityInfosMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexCapabilityInfosMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex capability infos moid default response
func (o *GetHyperflexCapabilityInfosMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexCapabilityInfosMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/CapabilityInfos/{moid}][%d] GetHyperflexCapabilityInfosMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexCapabilityInfosMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexCapabilityInfosMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
