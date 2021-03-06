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

// GetApplianceBackupPoliciesReader is a Reader for the GetApplianceBackupPolicies structure.
type GetApplianceBackupPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceBackupPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceBackupPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetApplianceBackupPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceBackupPoliciesOK creates a GetApplianceBackupPoliciesOK with default headers values
func NewGetApplianceBackupPoliciesOK() *GetApplianceBackupPoliciesOK {
	return &GetApplianceBackupPoliciesOK{}
}

/*GetApplianceBackupPoliciesOK handles this case with default header values.

List of applianceBackupPolicies for the given filter criteria
*/
type GetApplianceBackupPoliciesOK struct {
	Payload *models.ApplianceBackupPolicyList
}

func (o *GetApplianceBackupPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /appliance/BackupPolicies][%d] getApplianceBackupPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetApplianceBackupPoliciesOK) GetPayload() *models.ApplianceBackupPolicyList {
	return o.Payload
}

func (o *GetApplianceBackupPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceBackupPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceBackupPoliciesDefault creates a GetApplianceBackupPoliciesDefault with default headers values
func NewGetApplianceBackupPoliciesDefault(code int) *GetApplianceBackupPoliciesDefault {
	return &GetApplianceBackupPoliciesDefault{
		_statusCode: code,
	}
}

/*GetApplianceBackupPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceBackupPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance backup policies default response
func (o *GetApplianceBackupPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceBackupPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/BackupPolicies][%d] GetApplianceBackupPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceBackupPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceBackupPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
