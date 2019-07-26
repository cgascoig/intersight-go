// Code generated by go-swagger; DO NOT EDIT.

package testcrypt_read_only_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteTestcryptReadOnlyUsersMoidReader is a Reader for the DeleteTestcryptReadOnlyUsersMoid structure.
type DeleteTestcryptReadOnlyUsersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTestcryptReadOnlyUsersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTestcryptReadOnlyUsersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTestcryptReadOnlyUsersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTestcryptReadOnlyUsersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTestcryptReadOnlyUsersMoidOK creates a DeleteTestcryptReadOnlyUsersMoidOK with default headers values
func NewDeleteTestcryptReadOnlyUsersMoidOK() *DeleteTestcryptReadOnlyUsersMoidOK {
	return &DeleteTestcryptReadOnlyUsersMoidOK{}
}

/*DeleteTestcryptReadOnlyUsersMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteTestcryptReadOnlyUsersMoidOK struct {
}

func (o *DeleteTestcryptReadOnlyUsersMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /testcrypt/ReadOnlyUsers/{moid}][%d] deleteTestcryptReadOnlyUsersMoidOK ", 200)
}

func (o *DeleteTestcryptReadOnlyUsersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTestcryptReadOnlyUsersMoidNotFound creates a DeleteTestcryptReadOnlyUsersMoidNotFound with default headers values
func NewDeleteTestcryptReadOnlyUsersMoidNotFound() *DeleteTestcryptReadOnlyUsersMoidNotFound {
	return &DeleteTestcryptReadOnlyUsersMoidNotFound{}
}

/*DeleteTestcryptReadOnlyUsersMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteTestcryptReadOnlyUsersMoidNotFound struct {
}

func (o *DeleteTestcryptReadOnlyUsersMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /testcrypt/ReadOnlyUsers/{moid}][%d] deleteTestcryptReadOnlyUsersMoidNotFound ", 404)
}

func (o *DeleteTestcryptReadOnlyUsersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTestcryptReadOnlyUsersMoidDefault creates a DeleteTestcryptReadOnlyUsersMoidDefault with default headers values
func NewDeleteTestcryptReadOnlyUsersMoidDefault(code int) *DeleteTestcryptReadOnlyUsersMoidDefault {
	return &DeleteTestcryptReadOnlyUsersMoidDefault{
		_statusCode: code,
	}
}

/*DeleteTestcryptReadOnlyUsersMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteTestcryptReadOnlyUsersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete testcrypt read only users moid default response
func (o *DeleteTestcryptReadOnlyUsersMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTestcryptReadOnlyUsersMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /testcrypt/ReadOnlyUsers/{moid}][%d] DeleteTestcryptReadOnlyUsersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTestcryptReadOnlyUsersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTestcryptReadOnlyUsersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}