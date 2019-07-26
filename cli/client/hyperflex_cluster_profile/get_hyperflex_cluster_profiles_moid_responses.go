// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexClusterProfilesMoidReader is a Reader for the GetHyperflexClusterProfilesMoid structure.
type GetHyperflexClusterProfilesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClusterProfilesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClusterProfilesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexClusterProfilesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexClusterProfilesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClusterProfilesMoidOK creates a GetHyperflexClusterProfilesMoidOK with default headers values
func NewGetHyperflexClusterProfilesMoidOK() *GetHyperflexClusterProfilesMoidOK {
	return &GetHyperflexClusterProfilesMoidOK{}
}

/*GetHyperflexClusterProfilesMoidOK handles this case with default header values.

An instance of hyperflexClusterProfile
*/
type GetHyperflexClusterProfilesMoidOK struct {
	Payload *models.HyperflexClusterProfile
}

func (o *GetHyperflexClusterProfilesMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterProfiles/{moid}][%d] getHyperflexClusterProfilesMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClusterProfilesMoidOK) GetPayload() *models.HyperflexClusterProfile {
	return o.Payload
}

func (o *GetHyperflexClusterProfilesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexClusterProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClusterProfilesMoidNotFound creates a GetHyperflexClusterProfilesMoidNotFound with default headers values
func NewGetHyperflexClusterProfilesMoidNotFound() *GetHyperflexClusterProfilesMoidNotFound {
	return &GetHyperflexClusterProfilesMoidNotFound{}
}

/*GetHyperflexClusterProfilesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexClusterProfilesMoidNotFound struct {
}

func (o *GetHyperflexClusterProfilesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterProfiles/{moid}][%d] getHyperflexClusterProfilesMoidNotFound ", 404)
}

func (o *GetHyperflexClusterProfilesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexClusterProfilesMoidDefault creates a GetHyperflexClusterProfilesMoidDefault with default headers values
func NewGetHyperflexClusterProfilesMoidDefault(code int) *GetHyperflexClusterProfilesMoidDefault {
	return &GetHyperflexClusterProfilesMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClusterProfilesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClusterProfilesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex cluster profiles moid default response
func (o *GetHyperflexClusterProfilesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClusterProfilesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterProfiles/{moid}][%d] GetHyperflexClusterProfilesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClusterProfilesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClusterProfilesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
