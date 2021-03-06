// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexClustersMoidReader is a Reader for the GetHyperflexClustersMoid structure.
type GetHyperflexClustersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClustersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClustersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexClustersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexClustersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClustersMoidOK creates a GetHyperflexClustersMoidOK with default headers values
func NewGetHyperflexClustersMoidOK() *GetHyperflexClustersMoidOK {
	return &GetHyperflexClustersMoidOK{}
}

/*GetHyperflexClustersMoidOK handles this case with default header values.

An instance of hyperflexCluster
*/
type GetHyperflexClustersMoidOK struct {
	Payload *models.HyperflexCluster
}

func (o *GetHyperflexClustersMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/Clusters/{moid}][%d] getHyperflexClustersMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClustersMoidOK) GetPayload() *models.HyperflexCluster {
	return o.Payload
}

func (o *GetHyperflexClustersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClustersMoidNotFound creates a GetHyperflexClustersMoidNotFound with default headers values
func NewGetHyperflexClustersMoidNotFound() *GetHyperflexClustersMoidNotFound {
	return &GetHyperflexClustersMoidNotFound{}
}

/*GetHyperflexClustersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexClustersMoidNotFound struct {
}

func (o *GetHyperflexClustersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/Clusters/{moid}][%d] getHyperflexClustersMoidNotFound ", 404)
}

func (o *GetHyperflexClustersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexClustersMoidDefault creates a GetHyperflexClustersMoidDefault with default headers values
func NewGetHyperflexClustersMoidDefault(code int) *GetHyperflexClustersMoidDefault {
	return &GetHyperflexClustersMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClustersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClustersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex clusters moid default response
func (o *GetHyperflexClustersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClustersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/Clusters/{moid}][%d] GetHyperflexClustersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClustersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClustersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
