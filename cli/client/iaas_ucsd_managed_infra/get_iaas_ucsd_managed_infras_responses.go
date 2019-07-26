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

// GetIaasUcsdManagedInfrasReader is a Reader for the GetIaasUcsdManagedInfras structure.
type GetIaasUcsdManagedInfrasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIaasUcsdManagedInfrasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIaasUcsdManagedInfrasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIaasUcsdManagedInfrasDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIaasUcsdManagedInfrasOK creates a GetIaasUcsdManagedInfrasOK with default headers values
func NewGetIaasUcsdManagedInfrasOK() *GetIaasUcsdManagedInfrasOK {
	return &GetIaasUcsdManagedInfrasOK{}
}

/*GetIaasUcsdManagedInfrasOK handles this case with default header values.

List of iaasUcsdManagedInfras for the given filter criteria
*/
type GetIaasUcsdManagedInfrasOK struct {
	Payload *models.IaasUcsdManagedInfraList
}

func (o *GetIaasUcsdManagedInfrasOK) Error() string {
	return fmt.Sprintf("[GET /iaas/UcsdManagedInfras][%d] getIaasUcsdManagedInfrasOK  %+v", 200, o.Payload)
}

func (o *GetIaasUcsdManagedInfrasOK) GetPayload() *models.IaasUcsdManagedInfraList {
	return o.Payload
}

func (o *GetIaasUcsdManagedInfrasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaasUcsdManagedInfraList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIaasUcsdManagedInfrasDefault creates a GetIaasUcsdManagedInfrasDefault with default headers values
func NewGetIaasUcsdManagedInfrasDefault(code int) *GetIaasUcsdManagedInfrasDefault {
	return &GetIaasUcsdManagedInfrasDefault{
		_statusCode: code,
	}
}

/*GetIaasUcsdManagedInfrasDefault handles this case with default header values.

Unexpected error
*/
type GetIaasUcsdManagedInfrasDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iaas ucsd managed infras default response
func (o *GetIaasUcsdManagedInfrasDefault) Code() int {
	return o._statusCode
}

func (o *GetIaasUcsdManagedInfrasDefault) Error() string {
	return fmt.Sprintf("[GET /iaas/UcsdManagedInfras][%d] GetIaasUcsdManagedInfras default  %+v", o._statusCode, o.Payload)
}

func (o *GetIaasUcsdManagedInfrasDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIaasUcsdManagedInfrasDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}