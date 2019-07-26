// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_flash_controller_props

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageFlexFlashControllerPropsReader is a Reader for the GetStorageFlexFlashControllerProps structure.
type GetStorageFlexFlashControllerPropsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageFlexFlashControllerPropsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageFlexFlashControllerPropsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStorageFlexFlashControllerPropsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageFlexFlashControllerPropsOK creates a GetStorageFlexFlashControllerPropsOK with default headers values
func NewGetStorageFlexFlashControllerPropsOK() *GetStorageFlexFlashControllerPropsOK {
	return &GetStorageFlexFlashControllerPropsOK{}
}

/*GetStorageFlexFlashControllerPropsOK handles this case with default header values.

List of storageFlexFlashControllerProps for the given filter criteria
*/
type GetStorageFlexFlashControllerPropsOK struct {
	Payload *models.StorageFlexFlashControllerPropsList
}

func (o *GetStorageFlexFlashControllerPropsOK) Error() string {
	return fmt.Sprintf("[GET /storage/FlexFlashControllerProps][%d] getStorageFlexFlashControllerPropsOK  %+v", 200, o.Payload)
}

func (o *GetStorageFlexFlashControllerPropsOK) GetPayload() *models.StorageFlexFlashControllerPropsList {
	return o.Payload
}

func (o *GetStorageFlexFlashControllerPropsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageFlexFlashControllerPropsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageFlexFlashControllerPropsDefault creates a GetStorageFlexFlashControllerPropsDefault with default headers values
func NewGetStorageFlexFlashControllerPropsDefault(code int) *GetStorageFlexFlashControllerPropsDefault {
	return &GetStorageFlexFlashControllerPropsDefault{
		_statusCode: code,
	}
}

/*GetStorageFlexFlashControllerPropsDefault handles this case with default header values.

Unexpected error
*/
type GetStorageFlexFlashControllerPropsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage flex flash controller props default response
func (o *GetStorageFlexFlashControllerPropsDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageFlexFlashControllerPropsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/FlexFlashControllerProps][%d] GetStorageFlexFlashControllerProps default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageFlexFlashControllerPropsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageFlexFlashControllerPropsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
