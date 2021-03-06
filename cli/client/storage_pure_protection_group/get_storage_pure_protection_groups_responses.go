// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_protection_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetStoragePureProtectionGroupsReader is a Reader for the GetStoragePureProtectionGroups structure.
type GetStoragePureProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragePureProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoragePureProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStoragePureProtectionGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragePureProtectionGroupsOK creates a GetStoragePureProtectionGroupsOK with default headers values
func NewGetStoragePureProtectionGroupsOK() *GetStoragePureProtectionGroupsOK {
	return &GetStoragePureProtectionGroupsOK{}
}

/*GetStoragePureProtectionGroupsOK handles this case with default header values.

List of storagePureProtectionGroups for the given filter criteria
*/
type GetStoragePureProtectionGroupsOK struct {
	Payload *models.StoragePureProtectionGroupList
}

func (o *GetStoragePureProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroups][%d] getStoragePureProtectionGroupsOK  %+v", 200, o.Payload)
}

func (o *GetStoragePureProtectionGroupsOK) GetPayload() *models.StoragePureProtectionGroupList {
	return o.Payload
}

func (o *GetStoragePureProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePureProtectionGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragePureProtectionGroupsDefault creates a GetStoragePureProtectionGroupsDefault with default headers values
func NewGetStoragePureProtectionGroupsDefault(code int) *GetStoragePureProtectionGroupsDefault {
	return &GetStoragePureProtectionGroupsDefault{
		_statusCode: code,
	}
}

/*GetStoragePureProtectionGroupsDefault handles this case with default header values.

Unexpected error
*/
type GetStoragePureProtectionGroupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storage pure protection groups default response
func (o *GetStoragePureProtectionGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragePureProtectionGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/PureProtectionGroups][%d] GetStoragePureProtectionGroups default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragePureProtectionGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStoragePureProtectionGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
