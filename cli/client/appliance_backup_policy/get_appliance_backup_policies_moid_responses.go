// Code generated by go-swagger; DO NOT EDIT.

package appliance_backup_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceBackupPoliciesMoidReader is a Reader for the GetApplianceBackupPoliciesMoid structure.
type GetApplianceBackupPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceBackupPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceBackupPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApplianceBackupPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplianceBackupPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceBackupPoliciesMoidOK creates a GetApplianceBackupPoliciesMoidOK with default headers values
func NewGetApplianceBackupPoliciesMoidOK() *GetApplianceBackupPoliciesMoidOK {
	return &GetApplianceBackupPoliciesMoidOK{}
}

/*GetApplianceBackupPoliciesMoidOK handles this case with default header values.

An instance of applianceBackupPolicy
*/
type GetApplianceBackupPoliciesMoidOK struct {
	Payload *models.ApplianceBackupPolicy
}

func (o *GetApplianceBackupPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /appliance/BackupPolicies/{moid}][%d] getApplianceBackupPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetApplianceBackupPoliciesMoidOK) GetPayload() *models.ApplianceBackupPolicy {
	return o.Payload
}

func (o *GetApplianceBackupPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceBackupPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceBackupPoliciesMoidNotFound creates a GetApplianceBackupPoliciesMoidNotFound with default headers values
func NewGetApplianceBackupPoliciesMoidNotFound() *GetApplianceBackupPoliciesMoidNotFound {
	return &GetApplianceBackupPoliciesMoidNotFound{}
}

/*GetApplianceBackupPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetApplianceBackupPoliciesMoidNotFound struct {
}

func (o *GetApplianceBackupPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /appliance/BackupPolicies/{moid}][%d] getApplianceBackupPoliciesMoidNotFound ", 404)
}

func (o *GetApplianceBackupPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplianceBackupPoliciesMoidDefault creates a GetApplianceBackupPoliciesMoidDefault with default headers values
func NewGetApplianceBackupPoliciesMoidDefault(code int) *GetApplianceBackupPoliciesMoidDefault {
	return &GetApplianceBackupPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetApplianceBackupPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceBackupPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance backup policies moid default response
func (o *GetApplianceBackupPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceBackupPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/BackupPolicies/{moid}][%d] GetApplianceBackupPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceBackupPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceBackupPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
