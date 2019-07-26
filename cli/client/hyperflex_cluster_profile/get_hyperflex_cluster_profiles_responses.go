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

// GetHyperflexClusterProfilesReader is a Reader for the GetHyperflexClusterProfiles structure.
type GetHyperflexClusterProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexClusterProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexClusterProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexClusterProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexClusterProfilesOK creates a GetHyperflexClusterProfilesOK with default headers values
func NewGetHyperflexClusterProfilesOK() *GetHyperflexClusterProfilesOK {
	return &GetHyperflexClusterProfilesOK{}
}

/*GetHyperflexClusterProfilesOK handles this case with default header values.

List of hyperflexClusterProfiles for the given filter criteria
*/
type GetHyperflexClusterProfilesOK struct {
	Payload *models.HyperflexClusterProfileList
}

func (o *GetHyperflexClusterProfilesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterProfiles][%d] getHyperflexClusterProfilesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexClusterProfilesOK) GetPayload() *models.HyperflexClusterProfileList {
	return o.Payload
}

func (o *GetHyperflexClusterProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexClusterProfileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexClusterProfilesDefault creates a GetHyperflexClusterProfilesDefault with default headers values
func NewGetHyperflexClusterProfilesDefault(code int) *GetHyperflexClusterProfilesDefault {
	return &GetHyperflexClusterProfilesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexClusterProfilesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexClusterProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex cluster profiles default response
func (o *GetHyperflexClusterProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexClusterProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ClusterProfiles][%d] GetHyperflexClusterProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexClusterProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexClusterProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
