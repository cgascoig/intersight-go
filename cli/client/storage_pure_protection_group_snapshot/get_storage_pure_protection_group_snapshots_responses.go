// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_protection_group_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureProtectionGroupSnapshotsReader is a Reader for the GetStoragePureProtectionGroupSnapshots structure.
type GetStoragePureProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStoragePureProtectionGroupSnapshotsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureProtectionGroupSnapshotsOK creates a GetStoragePureProtectionGroupSnapshotsOK with default headers values
func NewGetStoragePureProtectionGroupSnapshotsOK() *GetStoragePureProtectionGroupSnapshotsOK {
	return &GetStoragePureProtectionGroupSnapshotsOK{}
}

/*GetStoragePureProtectionGroupSnapshotsOK handles this case with default header values.

List of storagePureProtectionGroupSnapshots for the given filter criteria
*/
type GetStoragePureProtectionGroupSnapshotsOK struct {
	Payload *models.StoragePureProtectionGroupSnapshotList
}

func (o *GetStoragePureProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroupSnapshots][%d] getStoragePureProtectionGroupSnapshotsOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureProtectionGroupSnapshotsOK) GetPayload() *models.StoragePureProtectionGroupSnapshotList {
	return o.Payload
}

func (o *GetStoragePureProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureProtectionGroupSnapshotList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureProtectionGroupSnapshotsDefault creates a GetStoragePureProtectionGroupSnapshotsDefault with default headers values
func NewGetStoragePureProtectionGroupSnapshotsDefault(code int) *GetStoragePureProtectionGroupSnapshotsDefault {
	return &GetStoragePureProtectionGroupSnapshotsDefault{
		_statusCode: code,
	}
}

/*GetStoragePureProtectionGroupSnapshotsDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureProtectionGroupSnapshotsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure protection group snapshots default response
func (o *GetStoragePureProtectionGroupSnapshotsDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureProtectionGroupSnapshotsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroupSnapshots][%d] GetStoragePureProtectionGroupSnapshots default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureProtectionGroupSnapshotsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureProtectionGroupSnapshotsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}