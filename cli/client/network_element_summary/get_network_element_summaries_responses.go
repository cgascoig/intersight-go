// Code generated by go-swagger; DO NOT EDIT.

package network_element_summary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetNetworkElementSummariesReader is a Reader for the GetNetworkElementSummaries structure.
type GetNetworkElementSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkElementSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkElementSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworkElementSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkElementSummariesOK creates a GetNetworkElementSummariesOK with default headers values
func NewGetNetworkElementSummariesOK() *GetNetworkElementSummariesOK {
	return &GetNetworkElementSummariesOK{}
}

/*GetNetworkElementSummariesOK handles this case with default header values.

List of networkElementSummaries for the given filter criteria
*/
type GetNetworkElementSummariesOK struct {
	Payload *models.NetworkElementSummaryList
}

func (o *GetNetworkElementSummariesOK) Error() string {
	return fmt.Sprintf("[GET /network/ElementSummaries][%d] getNetworkElementSummariesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkElementSummariesOK) GetPayload() *models.NetworkElementSummaryList {
	return o.Payload
}

func (o *GetNetworkElementSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkElementSummaryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkElementSummariesDefault creates a GetNetworkElementSummariesDefault with default headers values
func NewGetNetworkElementSummariesDefault(code int) *GetNetworkElementSummariesDefault {
	return &GetNetworkElementSummariesDefault{
		_statusCode: code,
	}
}

/*GetNetworkElementSummariesDefault handles this case with default header values.

Unexpected error
*/
type GetNetworkElementSummariesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get network element summaries default response
func (o *GetNetworkElementSummariesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkElementSummariesDefault) Error() string {
	return fmt.Sprintf("[GET /network/ElementSummaries][%d] GetNetworkElementSummaries default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworkElementSummariesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworkElementSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}