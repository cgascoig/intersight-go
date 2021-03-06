// Code generated by go-swagger; DO NOT EDIT.

package iaas_ucsd_managed_infra

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIaasUcsdManagedInfrasMoidReader is a Reader for the GetIaasUcsdManagedInfrasMoid structure.
type GetIaasUcsdManagedInfrasMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIaasUcsdManagedInfrasMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIaasUcsdManagedInfrasMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIaasUcsdManagedInfrasMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIaasUcsdManagedInfrasMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIaasUcsdManagedInfrasMoidOK creates a GetIaasUcsdManagedInfrasMoidOK with default headers values
func NewGetIaasUcsdManagedInfrasMoidOK() *GetIaasUcsdManagedInfrasMoidOK {
	return &GetIaasUcsdManagedInfrasMoidOK{}
}

/*GetIaasUcsdManagedInfrasMoidOK handles this case with default header values.

An instance of iaasUcsdManagedInfra
*/
type GetIaasUcsdManagedInfrasMoidOK struct {
	Payload *models.IaasUcsdManagedInfra
}

func (o *GetIaasUcsdManagedInfrasMoidOK) Error() string {
	return fmt.Sprintf("[GET /iaas/UcsdManagedInfras/{moid}][%d] getIaasUcsdManagedInfrasMoidOK  %+v", 200, o.Payload)
}

func (o *GetIaasUcsdManagedInfrasMoidOK) GetPayload() *models.IaasUcsdManagedInfra {
	return o.Payload
}

func (o *GetIaasUcsdManagedInfrasMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaasUcsdManagedInfra)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIaasUcsdManagedInfrasMoidNotFound creates a GetIaasUcsdManagedInfrasMoidNotFound with default headers values
func NewGetIaasUcsdManagedInfrasMoidNotFound() *GetIaasUcsdManagedInfrasMoidNotFound {
	return &GetIaasUcsdManagedInfrasMoidNotFound{}
}

/*GetIaasUcsdManagedInfrasMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetIaasUcsdManagedInfrasMoidNotFound struct {
}

func (o *GetIaasUcsdManagedInfrasMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/UcsdManagedInfras/{moid}][%d] getIaasUcsdManagedInfrasMoidNotFound ", 404)
}

func (o *GetIaasUcsdManagedInfrasMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIaasUcsdManagedInfrasMoidDefault creates a GetIaasUcsdManagedInfrasMoidDefault with default headers values
func NewGetIaasUcsdManagedInfrasMoidDefault(code int) *GetIaasUcsdManagedInfrasMoidDefault {
	return &GetIaasUcsdManagedInfrasMoidDefault{
		_statusCode: code,
	}
}

/*GetIaasUcsdManagedInfrasMoidDefault handles this case with default header values.

Unexpected error
*/
type GetIaasUcsdManagedInfrasMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iaas ucsd managed infras moid default response
func (o *GetIaasUcsdManagedInfrasMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetIaasUcsdManagedInfrasMoidDefault) Error() string {
	return fmt.Sprintf("[GET /iaas/UcsdManagedInfras/{moid}][%d] GetIaasUcsdManagedInfrasMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetIaasUcsdManagedInfrasMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIaasUcsdManagedInfrasMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
