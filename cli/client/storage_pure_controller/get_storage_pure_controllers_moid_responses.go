// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureControllersMoidReader is a Reader for the GetStoragePureControllersMoid structure.
type GetStoragePureControllersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureControllersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureControllersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStoragePureControllersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStoragePureControllersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureControllersMoidOK creates a GetStoragePureControllersMoidOK with default headers values
func NewGetStoragePureControllersMoidOK() *GetStoragePureControllersMoidOK {
	return &GetStoragePureControllersMoidOK{}
}

/*GetStoragePureControllersMoidOK handles this case with default header values.

An instance of storagePureController
*/
type GetStoragePureControllersMoidOK struct {
	Payload *models.StoragePureController
}

func (o *GetStoragePureControllersMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureControllers/{moid}][%d] getStoragePureControllersMoidOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureControllersMoidOK) GetPayload() *models.StoragePureController {
	return o.Payload
}

func (o *GetStoragePureControllersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureController)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureControllersMoidNotFound creates a GetStoragePureControllersMoidNotFound with default headers values
func NewGetStoragePureControllersMoidNotFound() *GetStoragePureControllersMoidNotFound {
	return &GetStoragePureControllersMoidNotFound{}
}

/*GetStoragePureControllersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStoragePureControllersMoidNotFound struct {
}

func (o *GetStoragePureControllersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/PureControllers/{moid}][%d] getStoragePureControllersMoidNotFound ", 404)
}

func (o *GetStoragePureControllersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoragePureControllersMoidDefault creates a GetStoragePureControllersMoidDefault with default headers values
func NewGetStoragePureControllersMoidDefault(code int) *GetStoragePureControllersMoidDefault {
	return &GetStoragePureControllersMoidDefault{
		_statusCode: code,
	}
}

/*GetStoragePureControllersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureControllersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure controllers moid default response
func (o *GetStoragePureControllersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureControllersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureControllers/{moid}][%d] GetStoragePureControllersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureControllersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureControllersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
