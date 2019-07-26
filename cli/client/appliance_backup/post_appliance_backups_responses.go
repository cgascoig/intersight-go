// Code generated by go-swagger; DO NOT EDIT.

package appliance_backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostApplianceBackupsReader is a Reader for the PostApplianceBackups structure.
type PostApplianceBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApplianceBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostApplianceBackupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostApplianceBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostApplianceBackupsCreated creates a PostApplianceBackupsCreated with default headers values
func NewPostApplianceBackupsCreated() *PostApplianceBackupsCreated {
	return &PostApplianceBackupsCreated{}
}

/*PostApplianceBackupsCreated handles this case with default header values.

Null response
*/
type PostApplianceBackupsCreated struct {
}

func (o *PostApplianceBackupsCreated) Error() string {
	return fmt.Sprintf("[POST /appliance/Backups][%d] postApplianceBackupsCreated ", 201)
}

func (o *PostApplianceBackupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApplianceBackupsDefault creates a PostApplianceBackupsDefault with default headers values
func NewPostApplianceBackupsDefault(code int) *PostApplianceBackupsDefault {
	return &PostApplianceBackupsDefault{
		_statusCode: code,
	}
}

/*PostApplianceBackupsDefault handles this case with default header values.

unexpected error
*/
type PostApplianceBackupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post appliance backups default response
func (o *PostApplianceBackupsDefault) Code() int {
	return o._statusCode
}

func (o *PostApplianceBackupsDefault) Error() string {
	return fmt.Sprintf("[POST /appliance/Backups][%d] PostApplianceBackups default  %+v", o._statusCode, o.Payload)
}

func (o *PostApplianceBackupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApplianceBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
