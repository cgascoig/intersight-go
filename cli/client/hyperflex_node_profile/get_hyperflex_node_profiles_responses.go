// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_node_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexNodeProfilesReader is a Reader for the GetHyperflexNodeProfiles structure.
type GetHyperflexNodeProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexNodeProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexNodeProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexNodeProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexNodeProfilesOK creates a GetHyperflexNodeProfilesOK with default headers values
func NewGetHyperflexNodeProfilesOK() *GetHyperflexNodeProfilesOK {
	return &GetHyperflexNodeProfilesOK{}
}

/*GetHyperflexNodeProfilesOK handles this case with default header values.

List of hyperflexNodeProfiles for the given filter criteria
*/
type GetHyperflexNodeProfilesOK struct {
	Payload *models.HyperflexNodeProfileList
}

func (o *GetHyperflexNodeProfilesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/NodeProfiles][%d] getHyperflexNodeProfilesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexNodeProfilesOK) GetPayload() *models.HyperflexNodeProfileList {
	return o.Payload
}

func (o *GetHyperflexNodeProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexNodeProfileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexNodeProfilesDefault creates a GetHyperflexNodeProfilesDefault with default headers values
func NewGetHyperflexNodeProfilesDefault(code int) *GetHyperflexNodeProfilesDefault {
	return &GetHyperflexNodeProfilesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexNodeProfilesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexNodeProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex node profiles default response
func (o *GetHyperflexNodeProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexNodeProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/NodeProfiles][%d] GetHyperflexNodeProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexNodeProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexNodeProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
