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

// GetHyperflexCapabilityInfosReader is a Reader for the GetHyperflexCapabilityInfos structure.
type GetHyperflexCapabilityInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexCapabilityInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexCapabilityInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexCapabilityInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexCapabilityInfosOK creates a GetHyperflexCapabilityInfosOK with default headers values
func NewGetHyperflexCapabilityInfosOK() *GetHyperflexCapabilityInfosOK {
	return &GetHyperflexCapabilityInfosOK{}
}

/*GetHyperflexCapabilityInfosOK handles this case with default header values.

List of hyperflexCapabilityInfos for the given filter criteria
*/
type GetHyperflexCapabilityInfosOK struct {
	Payload *models.HyperflexCapabilityInfoList
}

func (o *GetHyperflexCapabilityInfosOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/CapabilityInfos][%d] getHyperflexCapabilityInfosOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexCapabilityInfosOK) GetPayload() *models.HyperflexCapabilityInfoList {
	return o.Payload
}

func (o *GetHyperflexCapabilityInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexCapabilityInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexCapabilityInfosDefault creates a GetHyperflexCapabilityInfosDefault with default headers values
func NewGetHyperflexCapabilityInfosDefault(code int) *GetHyperflexCapabilityInfosDefault {
	return &GetHyperflexCapabilityInfosDefault{
		_statusCode: code,
	}
}

/*GetHyperflexCapabilityInfosDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexCapabilityInfosDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex capability infos default response
func (o *GetHyperflexCapabilityInfosDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexCapabilityInfosDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/CapabilityInfos][%d] GetHyperflexCapabilityInfos default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexCapabilityInfosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexCapabilityInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}