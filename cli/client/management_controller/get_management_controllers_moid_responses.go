// Code generated by go-swagger; DO NOT EDIT.

package management_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetManagementControllersMoidReader is a Reader for the GetManagementControllersMoid structure.
type GetManagementControllersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetManagementControllersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetManagementControllersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetManagementControllersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetManagementControllersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetManagementControllersMoidOK creates a GetManagementControllersMoidOK with default headers values
func NewGetManagementControllersMoidOK() *GetManagementControllersMoidOK {
	return &GetManagementControllersMoidOK{}
}

/*GetManagementControllersMoidOK handles this case with default header values.

An instance of managementController
*/
type GetManagementControllersMoidOK struct {
	Payload *models.ManagementController
}

func (o *GetManagementControllersMoidOK) Error() string {
	return fmt.Sprintf("[GET /management/Controllers/{moid}][%d] getManagementControllersMoidOK  %+v", 200, o.Payload)
}

func (o *GetManagementControllersMoidOK) GetPayload() *models.ManagementController {
	return o.Payload
}

func (o *GetManagementControllersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementController)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetManagementControllersMoidNotFound creates a GetManagementControllersMoidNotFound with default headers values
func NewGetManagementControllersMoidNotFound() *GetManagementControllersMoidNotFound {
	return &GetManagementControllersMoidNotFound{}
}

/*GetManagementControllersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetManagementControllersMoidNotFound struct {
}

func (o *GetManagementControllersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /management/Controllers/{moid}][%d] getManagementControllersMoidNotFound ", 404)
}

func (o *GetManagementControllersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagementControllersMoidDefault creates a GetManagementControllersMoidDefault with default headers values
func NewGetManagementControllersMoidDefault(code int) *GetManagementControllersMoidDefault {
	return &GetManagementControllersMoidDefault{
		_statusCode: code,
	}
}

/*GetManagementControllersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetManagementControllersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get management controllers moid default response
func (o *GetManagementControllersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetManagementControllersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /management/Controllers/{moid}][%d] GetManagementControllersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetManagementControllersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetManagementControllersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}