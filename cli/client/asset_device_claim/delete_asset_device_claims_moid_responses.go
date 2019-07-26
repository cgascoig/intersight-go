// Code generated by go-swagger; DO NOT EDIT.

package asset_device_claim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteAssetDeviceClaimsMoidReader is a Reader for the DeleteAssetDeviceClaimsMoid structure.
type DeleteAssetDeviceClaimsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAssetDeviceClaimsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAssetDeviceClaimsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAssetDeviceClaimsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAssetDeviceClaimsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAssetDeviceClaimsMoidOK creates a DeleteAssetDeviceClaimsMoidOK with default headers values
func NewDeleteAssetDeviceClaimsMoidOK() *DeleteAssetDeviceClaimsMoidOK {
	return &DeleteAssetDeviceClaimsMoidOK{}
}

/*DeleteAssetDeviceClaimsMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteAssetDeviceClaimsMoidOK struct {
}

func (o *DeleteAssetDeviceClaimsMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /asset/DeviceClaims/{moid}][%d] deleteAssetDeviceClaimsMoidOK ", 200)
}

func (o *DeleteAssetDeviceClaimsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAssetDeviceClaimsMoidNotFound creates a DeleteAssetDeviceClaimsMoidNotFound with default headers values
func NewDeleteAssetDeviceClaimsMoidNotFound() *DeleteAssetDeviceClaimsMoidNotFound {
	return &DeleteAssetDeviceClaimsMoidNotFound{}
}

/*DeleteAssetDeviceClaimsMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteAssetDeviceClaimsMoidNotFound struct {
}

func (o *DeleteAssetDeviceClaimsMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /asset/DeviceClaims/{moid}][%d] deleteAssetDeviceClaimsMoidNotFound ", 404)
}

func (o *DeleteAssetDeviceClaimsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAssetDeviceClaimsMoidDefault creates a DeleteAssetDeviceClaimsMoidDefault with default headers values
func NewDeleteAssetDeviceClaimsMoidDefault(code int) *DeleteAssetDeviceClaimsMoidDefault {
	return &DeleteAssetDeviceClaimsMoidDefault{
		_statusCode: code,
	}
}

/*DeleteAssetDeviceClaimsMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteAssetDeviceClaimsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete asset device claims moid default response
func (o *DeleteAssetDeviceClaimsMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAssetDeviceClaimsMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /asset/DeviceClaims/{moid}][%d] DeleteAssetDeviceClaimsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAssetDeviceClaimsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAssetDeviceClaimsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}