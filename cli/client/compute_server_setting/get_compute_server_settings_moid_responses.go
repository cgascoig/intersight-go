// Code generated by go-swagger; DO NOT EDIT.

package compute_server_setting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetComputeServerSettingsMoidReader is a Reader for the GetComputeServerSettingsMoid structure.
type GetComputeServerSettingsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeServerSettingsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeServerSettingsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetComputeServerSettingsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetComputeServerSettingsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetComputeServerSettingsMoidOK creates a GetComputeServerSettingsMoidOK with default headers values
func NewGetComputeServerSettingsMoidOK() *GetComputeServerSettingsMoidOK {
	return &GetComputeServerSettingsMoidOK{}
}

/*GetComputeServerSettingsMoidOK handles this case with default header values.

An instance of computeServerSetting
*/
type GetComputeServerSettingsMoidOK struct {
	Payload *models.ComputeServerSetting
}

func (o *GetComputeServerSettingsMoidOK) Error() string {
	return fmt.Sprintf("[GET /compute/ServerSettings/{moid}][%d] getComputeServerSettingsMoidOK  %+v", 200, o.Payload)
}

func (o *GetComputeServerSettingsMoidOK) GetPayload() *models.ComputeServerSetting {
	return o.Payload
}

func (o *GetComputeServerSettingsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeServerSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeServerSettingsMoidNotFound creates a GetComputeServerSettingsMoidNotFound with default headers values
func NewGetComputeServerSettingsMoidNotFound() *GetComputeServerSettingsMoidNotFound {
	return &GetComputeServerSettingsMoidNotFound{}
}

/*GetComputeServerSettingsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetComputeServerSettingsMoidNotFound struct {
}

func (o *GetComputeServerSettingsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /compute/ServerSettings/{moid}][%d] getComputeServerSettingsMoidNotFound ", 404)
}

func (o *GetComputeServerSettingsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComputeServerSettingsMoidDefault creates a GetComputeServerSettingsMoidDefault with default headers values
func NewGetComputeServerSettingsMoidDefault(code int) *GetComputeServerSettingsMoidDefault {
	return &GetComputeServerSettingsMoidDefault{
		_statusCode: code,
	}
}

/*GetComputeServerSettingsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetComputeServerSettingsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get compute server settings moid default response
func (o *GetComputeServerSettingsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetComputeServerSettingsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /compute/ServerSettings/{moid}][%d] GetComputeServerSettingsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetComputeServerSettingsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetComputeServerSettingsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
