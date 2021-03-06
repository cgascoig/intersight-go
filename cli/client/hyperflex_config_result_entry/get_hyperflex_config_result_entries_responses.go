// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_config_result_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexConfigResultEntriesReader is a Reader for the GetHyperflexConfigResultEntries structure.
type GetHyperflexConfigResultEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexConfigResultEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexConfigResultEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexConfigResultEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexConfigResultEntriesOK creates a GetHyperflexConfigResultEntriesOK with default headers values
func NewGetHyperflexConfigResultEntriesOK() *GetHyperflexConfigResultEntriesOK {
	return &GetHyperflexConfigResultEntriesOK{}
}

/*GetHyperflexConfigResultEntriesOK handles this case with default header values.

List of hyperflexConfigResultEntries for the given filter criteria
*/
type GetHyperflexConfigResultEntriesOK struct {
	Payload *models.HyperflexConfigResultEntryList
}

func (o *GetHyperflexConfigResultEntriesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ConfigResultEntries][%d] getHyperflexConfigResultEntriesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexConfigResultEntriesOK) GetPayload() *models.HyperflexConfigResultEntryList {
	return o.Payload
}

func (o *GetHyperflexConfigResultEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexConfigResultEntryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexConfigResultEntriesDefault creates a GetHyperflexConfigResultEntriesDefault with default headers values
func NewGetHyperflexConfigResultEntriesDefault(code int) *GetHyperflexConfigResultEntriesDefault {
	return &GetHyperflexConfigResultEntriesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexConfigResultEntriesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexConfigResultEntriesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex config result entries default response
func (o *GetHyperflexConfigResultEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexConfigResultEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ConfigResultEntries][%d] GetHyperflexConfigResultEntries default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexConfigResultEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexConfigResultEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
