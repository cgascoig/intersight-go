// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamLdapPoliciesReader is a Reader for the GetIamLdapPolicies structure.
type GetIamLdapPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamLdapPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamLdapPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamLdapPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamLdapPoliciesOK creates a GetIamLdapPoliciesOK with default headers values
func NewGetIamLdapPoliciesOK() *GetIamLdapPoliciesOK {
	return &GetIamLdapPoliciesOK{}
}

/*GetIamLdapPoliciesOK handles this case with default header values.

List of iamLdapPolicies for the given filter criteria
*/
type GetIamLdapPoliciesOK struct {
	Payload *models.IamLdapPolicyList
}

func (o *GetIamLdapPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /iam/LdapPolicies][%d] getIamLdapPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetIamLdapPoliciesOK) GetPayload() *models.IamLdapPolicyList {
	return o.Payload
}

func (o *GetIamLdapPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamLdapPolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamLdapPoliciesDefault creates a GetIamLdapPoliciesDefault with default headers values
func NewGetIamLdapPoliciesDefault(code int) *GetIamLdapPoliciesDefault {
	return &GetIamLdapPoliciesDefault{
		_statusCode: code,
	}
}

/*GetIamLdapPoliciesDefault handles this case with default header values.

Unexpected error
*/
type GetIamLdapPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam ldap policies default response
func (o *GetIamLdapPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetIamLdapPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /iam/LdapPolicies][%d] GetIamLdapPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamLdapPoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamLdapPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}