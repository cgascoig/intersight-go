// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureVolumesReader is a Reader for the GetStoragePureVolumes structure.
type GetStoragePureVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStoragePureVolumesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureVolumesOK creates a GetStoragePureVolumesOK with default headers values
func NewGetStoragePureVolumesOK() *GetStoragePureVolumesOK {
	return &GetStoragePureVolumesOK{}
}

/*GetStoragePureVolumesOK handles this case with default header values.

List of storagePureVolumes for the given filter criteria
*/
type GetStoragePureVolumesOK struct {
	Payload *models.StoragePureVolumeList
}

func (o *GetStoragePureVolumesOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureVolumes][%d] getStoragePureVolumesOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureVolumesOK) GetPayload() *models.StoragePureVolumeList {
	return o.Payload
}

func (o *GetStoragePureVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureVolumeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureVolumesDefault creates a GetStoragePureVolumesDefault with default headers values
func NewGetStoragePureVolumesDefault(code int) *GetStoragePureVolumesDefault {
	return &GetStoragePureVolumesDefault{
		_statusCode: code,
	}
}

/*GetStoragePureVolumesDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureVolumesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure volumes default response
func (o *GetStoragePureVolumesDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureVolumesDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureVolumes][%d] GetStoragePureVolumes default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureVolumesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureVolumesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}