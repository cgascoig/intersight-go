// Code generated by go-swagger; DO NOT EDIT.

package appliance_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceRestoresMoidReader is a Reader for the GetApplianceRestoresMoid structure.
type GetApplianceRestoresMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceRestoresMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceRestoresMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApplianceRestoresMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplianceRestoresMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceRestoresMoidOK creates a GetApplianceRestoresMoidOK with default headers values
func NewGetApplianceRestoresMoidOK() *GetApplianceRestoresMoidOK {
	return &GetApplianceRestoresMoidOK{}
}

/*GetApplianceRestoresMoidOK handles this case with default header values.

An instance of applianceRestore
*/
type GetApplianceRestoresMoidOK struct {
	Payload *models.ApplianceRestore
}

func (o *GetApplianceRestoresMoidOK) Error() string {
	return fmt.Sprintf("[GET /appliance/Restores/{moid}][%d] getApplianceRestoresMoidOK  %+v", 200, o.Payload)
}

func (o *GetApplianceRestoresMoidOK) GetPayload() *models.ApplianceRestore {
	return o.Payload
}

func (o *GetApplianceRestoresMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceRestore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceRestoresMoidNotFound creates a GetApplianceRestoresMoidNotFound with default headers values
func NewGetApplianceRestoresMoidNotFound() *GetApplianceRestoresMoidNotFound {
	return &GetApplianceRestoresMoidNotFound{}
}

/*GetApplianceRestoresMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetApplianceRestoresMoidNotFound struct {
}

func (o *GetApplianceRestoresMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /appliance/Restores/{moid}][%d] getApplianceRestoresMoidNotFound ", 404)
}

func (o *GetApplianceRestoresMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplianceRestoresMoidDefault creates a GetApplianceRestoresMoidDefault with default headers values
func NewGetApplianceRestoresMoidDefault(code int) *GetApplianceRestoresMoidDefault {
	return &GetApplianceRestoresMoidDefault{
		_statusCode: code,
	}
}

/*GetApplianceRestoresMoidDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceRestoresMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance restores moid default response
func (o *GetApplianceRestoresMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceRestoresMoidDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/Restores/{moid}][%d] GetApplianceRestoresMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceRestoresMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceRestoresMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
