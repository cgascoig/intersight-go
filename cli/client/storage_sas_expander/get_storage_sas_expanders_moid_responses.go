// Code generated by go-swagger; DO NOT EDIT.

package storage_sas_expander

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageSasExpandersMoidReader is a Reader for the GetStorageSasExpandersMoid structure.
type GetStorageSasExpandersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageSasExpandersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageSasExpandersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStorageSasExpandersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStorageSasExpandersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageSasExpandersMoidOK creates a GetStorageSasExpandersMoidOK with default headers values
func NewGetStorageSasExpandersMoidOK() *GetStorageSasExpandersMoidOK {
	return &GetStorageSasExpandersMoidOK{}
}

/*GetStorageSasExpandersMoidOK handles this case with default header values.

An instance of storageSasExpander
*/
type GetStorageSasExpandersMoidOK struct {
	Payload *models.StorageSasExpander
}

func (o *GetStorageSasExpandersMoidOK) Error() string {
	return fmt.Sprintf("[GET /storage/SasExpanders/{moid}][%d] getStorageSasExpandersMoidOK  %+v", 200, o.Payload)
}

func (o *GetStorageSasExpandersMoidOK) GetPayload() *models.StorageSasExpander {
	return o.Payload
}

func (o *GetStorageSasExpandersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageSasExpander)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageSasExpandersMoidNotFound creates a GetStorageSasExpandersMoidNotFound with default headers values
func NewGetStorageSasExpandersMoidNotFound() *GetStorageSasExpandersMoidNotFound {
	return &GetStorageSasExpandersMoidNotFound{}
}

/*GetStorageSasExpandersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetStorageSasExpandersMoidNotFound struct {
}

func (o *GetStorageSasExpandersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/SasExpanders/{moid}][%d] getStorageSasExpandersMoidNotFound ", 404)
}

func (o *GetStorageSasExpandersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageSasExpandersMoidDefault creates a GetStorageSasExpandersMoidDefault with default headers values
func NewGetStorageSasExpandersMoidDefault(code int) *GetStorageSasExpandersMoidDefault {
	return &GetStorageSasExpandersMoidDefault{
		_statusCode: code,
	}
}

/*GetStorageSasExpandersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetStorageSasExpandersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage sas expanders moid default response
func (o *GetStorageSasExpandersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageSasExpandersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /storage/SasExpanders/{moid}][%d] GetStorageSasExpandersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageSasExpandersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageSasExpandersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}