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

// GetStoragePureControllersReader is a Reader for the GetStoragePureControllers structure.
type GetStoragePureControllersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureControllersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureControllersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStoragePureControllersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureControllersOK creates a GetStoragePureControllersOK with default headers values
func NewGetStoragePureControllersOK() *GetStoragePureControllersOK {
	return &GetStoragePureControllersOK{}
}

/*GetStoragePureControllersOK handles this case with default header values.

List of storagePureControllers for the given filter criteria
*/
type GetStoragePureControllersOK struct {
	Payload *models.StoragePureControllerList
}

func (o *GetStoragePureControllersOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureControllers][%d] getStoragePureControllersOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureControllersOK) GetPayload() *models.StoragePureControllerList {
	return o.Payload
}

func (o *GetStoragePureControllersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureControllerList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureControllersDefault creates a GetStoragePureControllersDefault with default headers values
func NewGetStoragePureControllersDefault(code int) *GetStoragePureControllersDefault {
	return &GetStoragePureControllersDefault{
		_statusCode: code,
	}
}

/*GetStoragePureControllersDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureControllersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure controllers default response
func (o *GetStoragePureControllersDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureControllersDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureControllers][%d] GetStoragePureControllers default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureControllersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureControllersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}