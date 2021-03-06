// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_software_version_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexSoftwareVersionPoliciesReader is a Reader for the GetHyperflexSoftwareVersionPolicies structure.
type GetHyperflexSoftwareVersionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexSoftwareVersionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexSoftwareVersionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexSoftwareVersionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexSoftwareVersionPoliciesOK creates a GetHyperflexSoftwareVersionPoliciesOK with default headers values
func NewGetHyperflexSoftwareVersionPoliciesOK() *GetHyperflexSoftwareVersionPoliciesOK {
	return &GetHyperflexSoftwareVersionPoliciesOK{}
}

/*GetHyperflexSoftwareVersionPoliciesOK handles this case with default header values.

List of hyperflexSoftwareVersionPolicies for the given filter criteria
*/
type GetHyperflexSoftwareVersionPoliciesOK struct {
	Payload *models.HyperflexSoftwareVersionPolicyList
}

func (o *GetHyperflexSoftwareVersionPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/SoftwareVersionPolicies][%d] getHyperflexSoftwareVersionPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexSoftwareVersionPoliciesOK) GetPayload() *models.HyperflexSoftwareVersionPolicyList {
	return o.Payload
}

func (o *GetHyperflexSoftwareVersionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexSoftwareVersionPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexSoftwareVersionPoliciesDefault creates a GetHyperflexSoftwareVersionPoliciesDefault with default headers values
func NewGetHyperflexSoftwareVersionPoliciesDefault(code int) *GetHyperflexSoftwareVersionPoliciesDefault {
	return &GetHyperflexSoftwareVersionPoliciesDefault{
		_statusCode: code,
	}
}

/*GetHyperflexSoftwareVersionPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexSoftwareVersionPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex software version policies default response
func (o *GetHyperflexSoftwareVersionPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexSoftwareVersionPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/SoftwareVersionPolicies][%d] GetHyperflexSoftwareVersionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexSoftwareVersionPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexSoftwareVersionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
