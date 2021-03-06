// Code generated by go-swagger; DO NOT EDIT.

package storage_vd_member_ep

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStorageVdMemberEpsReader is a Reader for the GetStorageVdMemberEps structure.
type GetStorageVdMemberEpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageVdMemberEpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageVdMemberEpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStorageVdMemberEpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageVdMemberEpsOK creates a GetStorageVdMemberEpsOK with default headers values
func NewGetStorageVdMemberEpsOK() *GetStorageVdMemberEpsOK {
	return &GetStorageVdMemberEpsOK{}
}

/*GetStorageVdMemberEpsOK handles this case with default header values.

List of storageVdMemberEps for the given filter criteria
*/
type GetStorageVdMemberEpsOK struct {
	Payload *models.StorageVdMemberEpList
}

func (o *GetStorageVdMemberEpsOK) Error() string {
	return fmt.Sprintf("[GET /storage/VdMemberEps][%d] getStorageVdMemberEpsOK  %+v", 200, o.Payload)
}

func (o *GetStorageVdMemberEpsOK) GetPayload() *models.StorageVdMemberEpList {
	return o.Payload
}

func (o *GetStorageVdMemberEpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageVdMemberEpList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageVdMemberEpsDefault creates a GetStorageVdMemberEpsDefault with default headers values
func NewGetStorageVdMemberEpsDefault(code int) *GetStorageVdMemberEpsDefault {
	return &GetStorageVdMemberEpsDefault{
		_statusCode: code,
	}
}

/*GetStorageVdMemberEpsDefault handles this case with default header values.

Unexpected error
*/
type GetStorageVdMemberEpsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage vd member eps default response
func (o *GetStorageVdMemberEpsDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageVdMemberEpsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/VdMemberEps][%d] GetStorageVdMemberEps default  %+v", o._statusCode, o.Payload)
}

func (o *GetStorageVdMemberEpsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageVdMemberEpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
