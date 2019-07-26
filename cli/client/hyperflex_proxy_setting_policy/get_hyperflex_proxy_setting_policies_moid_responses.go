// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_proxy_setting_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexProxySettingPoliciesMoidReader is a Reader for the GetHyperflexProxySettingPoliciesMoid structure.
type GetHyperflexProxySettingPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexProxySettingPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexProxySettingPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHyperflexProxySettingPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHyperflexProxySettingPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexProxySettingPoliciesMoidOK creates a GetHyperflexProxySettingPoliciesMoidOK with default headers values
func NewGetHyperflexProxySettingPoliciesMoidOK() *GetHyperflexProxySettingPoliciesMoidOK {
	return &GetHyperflexProxySettingPoliciesMoidOK{}
}

/*GetHyperflexProxySettingPoliciesMoidOK handles this case with default header values.

An instance of hyperflexProxySettingPolicy
*/
type GetHyperflexProxySettingPoliciesMoidOK struct {
	Payload *models.HyperflexProxySettingPolicy
}

func (o *GetHyperflexProxySettingPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ProxySettingPolicies/{moid}][%d] getHyperflexProxySettingPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexProxySettingPoliciesMoidOK) GetPayload() *models.HyperflexProxySettingPolicy {
	return o.Payload
}

func (o *GetHyperflexProxySettingPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexProxySettingPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexProxySettingPoliciesMoidNotFound creates a GetHyperflexProxySettingPoliciesMoidNotFound with default headers values
func NewGetHyperflexProxySettingPoliciesMoidNotFound() *GetHyperflexProxySettingPoliciesMoidNotFound {
	return &GetHyperflexProxySettingPoliciesMoidNotFound{}
}

/*GetHyperflexProxySettingPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetHyperflexProxySettingPoliciesMoidNotFound struct {
}

func (o *GetHyperflexProxySettingPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ProxySettingPolicies/{moid}][%d] getHyperflexProxySettingPoliciesMoidNotFound ", 404)
}

func (o *GetHyperflexProxySettingPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHyperflexProxySettingPoliciesMoidDefault creates a GetHyperflexProxySettingPoliciesMoidDefault with default headers values
func NewGetHyperflexProxySettingPoliciesMoidDefault(code int) *GetHyperflexProxySettingPoliciesMoidDefault {
	return &GetHyperflexProxySettingPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetHyperflexProxySettingPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexProxySettingPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex proxy setting policies moid default response
func (o *GetHyperflexProxySettingPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexProxySettingPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/ProxySettingPolicies/{moid}][%d] GetHyperflexProxySettingPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexProxySettingPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexProxySettingPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
