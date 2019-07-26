// Code generated by go-swagger; DO NOT EDIT.

package adapter_config_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetAdapterConfigPoliciesReader is a Reader for the GetAdapterConfigPolicies structure.
type GetAdapterConfigPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdapterConfigPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdapterConfigPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAdapterConfigPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdapterConfigPoliciesOK creates a GetAdapterConfigPoliciesOK with default headers values
func NewGetAdapterConfigPoliciesOK() *GetAdapterConfigPoliciesOK {
	return &GetAdapterConfigPoliciesOK{}
}

/*GetAdapterConfigPoliciesOK handles this case with default header values.

List of adapterConfigPolicies for the given filter criteria
*/
type GetAdapterConfigPoliciesOK struct {
	Payload *models.AdapterConfigPolicyList
}

func (o *GetAdapterConfigPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /adapter/ConfigPolicies][%d] getAdapterConfigPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetAdapterConfigPoliciesOK) GetPayload() *models.AdapterConfigPolicyList {
	return o.Payload
}

func (o *GetAdapterConfigPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdapterConfigPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdapterConfigPoliciesDefault creates a GetAdapterConfigPoliciesDefault with default headers values
func NewGetAdapterConfigPoliciesDefault(code int) *GetAdapterConfigPoliciesDefault {
	return &GetAdapterConfigPoliciesDefault{
		_statusCode: code,
	}
}

/*GetAdapterConfigPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetAdapterConfigPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get adapter config policies default response
func (o *GetAdapterConfigPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetAdapterConfigPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /adapter/ConfigPolicies][%d] GetAdapterConfigPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdapterConfigPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAdapterConfigPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}