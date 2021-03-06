// Code generated by go-swagger; DO NOT EDIT.

package appliance_upgrade_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceUpgradePoliciesReader is a Reader for the GetApplianceUpgradePolicies structure.
type GetApplianceUpgradePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceUpgradePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceUpgradePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetApplianceUpgradePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceUpgradePoliciesOK creates a GetApplianceUpgradePoliciesOK with default headers values
func NewGetApplianceUpgradePoliciesOK() *GetApplianceUpgradePoliciesOK {
	return &GetApplianceUpgradePoliciesOK{}
}

/*GetApplianceUpgradePoliciesOK handles this case with default header values.

List of applianceUpgradePolicies for the given filter criteria
*/
type GetApplianceUpgradePoliciesOK struct {
	Payload *models.ApplianceUpgradePolicyList
}

func (o *GetApplianceUpgradePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /appliance/UpgradePolicies][%d] getApplianceUpgradePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetApplianceUpgradePoliciesOK) GetPayload() *models.ApplianceUpgradePolicyList {
	return o.Payload
}

func (o *GetApplianceUpgradePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceUpgradePolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceUpgradePoliciesDefault creates a GetApplianceUpgradePoliciesDefault with default headers values
func NewGetApplianceUpgradePoliciesDefault(code int) *GetApplianceUpgradePoliciesDefault {
	return &GetApplianceUpgradePoliciesDefault{
		_statusCode: code,
	}
}

/*GetApplianceUpgradePoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceUpgradePoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance upgrade policies default response
func (o *GetApplianceUpgradePoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceUpgradePoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/UpgradePolicies][%d] GetApplianceUpgradePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceUpgradePoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceUpgradePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
