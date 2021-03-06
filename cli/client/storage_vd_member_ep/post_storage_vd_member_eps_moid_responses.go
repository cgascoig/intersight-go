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

// PostStorageVdMemberEpsMoidReader is a Reader for the PostStorageVdMemberEpsMoid structure.
type PostStorageVdMemberEpsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageVdMemberEpsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStorageVdMemberEpsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostStorageVdMemberEpsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostStorageVdMemberEpsMoidCreated creates a PostStorageVdMemberEpsMoidCreated with default headers values
func NewPostStorageVdMemberEpsMoidCreated() *PostStorageVdMemberEpsMoidCreated {
	return &PostStorageVdMemberEpsMoidCreated{}
}

/*PostStorageVdMemberEpsMoidCreated handles this case with default header values.

Null response
*/
type PostStorageVdMemberEpsMoidCreated struct {
}

func (o *PostStorageVdMemberEpsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /storage/VdMemberEps/{moid}][%d] postStorageVdMemberEpsMoidCreated ", 201)
}

func (o *PostStorageVdMemberEpsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStorageVdMemberEpsMoidDefault creates a PostStorageVdMemberEpsMoidDefault with default headers values
func NewPostStorageVdMemberEpsMoidDefault(code int) *PostStorageVdMemberEpsMoidDefault {
	return &PostStorageVdMemberEpsMoidDefault{
		_statusCode: code,
	}
}

/*PostStorageVdMemberEpsMoidDefault handles this case with default header values.

unexpected error
*/
type PostStorageVdMemberEpsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post storage vd member eps moid default response
func (o *PostStorageVdMemberEpsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostStorageVdMemberEpsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /storage/VdMemberEps/{moid}][%d] PostStorageVdMemberEpsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostStorageVdMemberEpsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStorageVdMemberEpsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
