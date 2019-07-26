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

// GetHyperflexClustersReader is a Reader for the GetHyperflexClusters structure.
type GetHyperflexClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClustersOK creates a GetHyperflexClustersOK with default headers values
func NewGetHyperflexClustersOK() *GetHyperflexClustersOK {
	return &GetHyperflexClustersOK{}
}

/*GetHyperflexClustersOK handles this case with default header values.

List of hyperflexClusters for the given filter criteria
*/
type GetHyperflexClustersOK struct {
	Payload *models.HyperflexClusterList
}

func (o *GetHyperflexClustersOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/Clusters][%d] getHyperflexClustersOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClustersOK) GetPayload() *models.HyperflexClusterList {
	return o.Payload
}

func (o *GetHyperflexClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexClusterList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClustersDefault creates a GetHyperflexClustersDefault with default headers values
func NewGetHyperflexClustersDefault(code int) *GetHyperflexClustersDefault {
	return &GetHyperflexClustersDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClustersDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClustersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex clusters default response
func (o *GetHyperflexClustersDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClustersDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/Clusters][%d] GetHyperflexClusters default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClustersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
