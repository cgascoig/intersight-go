// Code generated by go-swagger; DO NOT EDIT.

package storage_storage_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostStorageStoragePoliciesReader is a Reader for the PostStorageStoragePolicies structure.
type PostStorageStoragePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageStoragePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStorageStoragePoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStorageStoragePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStorageStoragePoliciesCreated creates a PostStorageStoragePoliciesCreated with default headers values
func NewPostStorageStoragePoliciesCreated() *PostStorageStoragePoliciesCreated {
	return &PostStorageStoragePoliciesCreated{}
}

/*PostStorageStoragePoliciesCreated handles this case with default header values.

Null response
*/
type PostStorageStoragePoliciesCreated struct {
}

func (o *PostStorageStoragePoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /storage/StoragePolicies][%d] postStorageStoragePoliciesCreated ", 201)
}

func (o *PostStorageStoragePoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStorageStoragePoliciesDefault creates a PostStorageStoragePoliciesDefault with default headers values
func NewPostStorageStoragePoliciesDefault(code int) *PostStorageStoragePoliciesDefault {
	return &PostStorageStoragePoliciesDefault{
		_statusCode: code,
	}
}

/*PostStorageStoragePoliciesDefault handles this case with default header values.

unexpected error
*/
type PostStorageStoragePoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage storage policies default response
func (o *PostStorageStoragePoliciesDefault) Code() int {
	return o._statusCode
}

func (o *PostStorageStoragePoliciesDefault) Error() string {
	return fmt.Sprintf("[POST /storage/StoragePolicies][%d] PostStorageStoragePolicies default  %+v", o._statusCode, o.Payload)
}

func (o *PostStorageStoragePoliciesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStorageStoragePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}