// Code generated by go-swagger; DO NOT EDIT.

package appliance_diag_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostApplianceDiagSettingsReader is a Reader for the PostApplianceDiagSettings structure.
type PostApplianceDiagSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApplianceDiagSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostApplianceDiagSettingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostApplianceDiagSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostApplianceDiagSettingsCreated creates a PostApplianceDiagSettingsCreated with default headers values
func NewPostApplianceDiagSettingsCreated() *PostApplianceDiagSettingsCreated {
	return &PostApplianceDiagSettingsCreated{}
}

/*PostApplianceDiagSettingsCreated handles this case with default header values.

Null response
*/
type PostApplianceDiagSettingsCreated struct {
}

func (o *PostApplianceDiagSettingsCreated) Error() string {
	return fmt.Sprintf("[POST /appliance/DiagSettings][%d] postApplianceDiagSettingsCreated ", 201)
}

func (o *PostApplianceDiagSettingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApplianceDiagSettingsDefault creates a PostApplianceDiagSettingsDefault with default headers values
func NewPostApplianceDiagSettingsDefault(code int) *PostApplianceDiagSettingsDefault {
	return &PostApplianceDiagSettingsDefault{
		_statusCode: code,
	}
}

/*PostApplianceDiagSettingsDefault handles this case with default header values.

unexpected error
*/
type PostApplianceDiagSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post appliance diag settings default response
func (o *PostApplianceDiagSettingsDefault) Code() int {
	return o._statusCode
}

func (o *PostApplianceDiagSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /appliance/DiagSettings][%d] PostApplianceDiagSettings default  %+v", o._statusCode, o.Payload)
}

func (o *PostApplianceDiagSettingsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApplianceDiagSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}