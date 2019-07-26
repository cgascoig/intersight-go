// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_adapter_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetVnicFcAdapterPoliciesReader is a Reader for the GetVnicFcAdapterPolicies structure.
type GetVnicFcAdapterPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVnicFcAdapterPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVnicFcAdapterPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVnicFcAdapterPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVnicFcAdapterPoliciesOK creates a GetVnicFcAdapterPoliciesOK with default headers values
func NewGetVnicFcAdapterPoliciesOK() *GetVnicFcAdapterPoliciesOK {
	return &GetVnicFcAdapterPoliciesOK{}
}

/*GetVnicFcAdapterPoliciesOK handles this case with default header values.

List of vnicFcAdapterPolicies for the given filter criteria
*/
type GetVnicFcAdapterPoliciesOK struct {
	Payload *models.VnicFcAdapterPolicyList
}

func (o *GetVnicFcAdapterPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /vnic/FcAdapterPolicies][%d] getVnicFcAdapterPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetVnicFcAdapterPoliciesOK) GetPayload() *models.VnicFcAdapterPolicyList {
	return o.Payload
}

func (o *GetVnicFcAdapterPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VnicFcAdapterPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVnicFcAdapterPoliciesDefault creates a GetVnicFcAdapterPoliciesDefault with default headers values
func NewGetVnicFcAdapterPoliciesDefault(code int) *GetVnicFcAdapterPoliciesDefault {
	return &GetVnicFcAdapterPoliciesDefault{
		_statusCode: code,
	}
}

/*GetVnicFcAdapterPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetVnicFcAdapterPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get vnic fc adapter policies default response
func (o *GetVnicFcAdapterPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetVnicFcAdapterPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /vnic/FcAdapterPolicies][%d] GetVnicFcAdapterPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetVnicFcAdapterPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVnicFcAdapterPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
