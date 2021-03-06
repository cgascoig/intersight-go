// Code generated by go-swagger; DO NOT EDIT.

package appliance_backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceBackupsMoidReader is a Reader for the GetApplianceBackupsMoid structure.
type GetApplianceBackupsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceBackupsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceBackupsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApplianceBackupsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplianceBackupsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceBackupsMoidOK creates a GetApplianceBackupsMoidOK with default headers values
func NewGetApplianceBackupsMoidOK() *GetApplianceBackupsMoidOK {
	return &GetApplianceBackupsMoidOK{}
}

/*GetApplianceBackupsMoidOK handles this case with default header values.

An instance of applianceBackup
*/
type GetApplianceBackupsMoidOK struct {
	Payload *models.ApplianceBackup
}

func (o *GetApplianceBackupsMoidOK) Error() string {
	return fmt.Sprintf("[GET /appliance/Backups/{moid}][%d] getApplianceBackupsMoidOK  %+v", 200, o.Payload)
}

func (o *GetApplianceBackupsMoidOK) GetPayload() *models.ApplianceBackup {
	return o.Payload
}

func (o *GetApplianceBackupsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceBackupsMoidNotFound creates a GetApplianceBackupsMoidNotFound with default headers values
func NewGetApplianceBackupsMoidNotFound() *GetApplianceBackupsMoidNotFound {
	return &GetApplianceBackupsMoidNotFound{}
}

/*GetApplianceBackupsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetApplianceBackupsMoidNotFound struct {
}

func (o *GetApplianceBackupsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /appliance/Backups/{moid}][%d] getApplianceBackupsMoidNotFound ", 404)
}

func (o *GetApplianceBackupsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplianceBackupsMoidDefault creates a GetApplianceBackupsMoidDefault with default headers values
func NewGetApplianceBackupsMoidDefault(code int) *GetApplianceBackupsMoidDefault {
	return &GetApplianceBackupsMoidDefault{
		_statusCode: code,
	}
}

/*GetApplianceBackupsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceBackupsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance backups moid default response
func (o *GetApplianceBackupsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceBackupsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/Backups/{moid}][%d] GetApplianceBackupsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceBackupsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceBackupsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
