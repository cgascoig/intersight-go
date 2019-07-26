// Code generated by go-swagger; DO NOT EDIT.

package server_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostServerProfilesReader is a Reader for the PostServerProfiles structure.
type PostServerProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServerProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostServerProfilesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostServerProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostServerProfilesCreated creates a PostServerProfilesCreated with default headers values
func NewPostServerProfilesCreated() *PostServerProfilesCreated {
	return &PostServerProfilesCreated{}
}

/*PostServerProfilesCreated handles this case with default header values.

Null response
*/
type PostServerProfilesCreated struct {
}

func (o *PostServerProfilesCreated) Error() string {
	return fmt.Sprintf("[POST /server/Profiles][%d] postServerProfilesCreated ", 201)
}

func (o *PostServerProfilesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServerProfilesDefault creates a PostServerProfilesDefault with default headers values
func NewPostServerProfilesDefault(code int) *PostServerProfilesDefault {
	return &PostServerProfilesDefault{
		_statusCode: code,
	}
}

/*PostServerProfilesDefault handles this case with default header values.

unexpected error
*/
type PostServerProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post server profiles default response
func (o *PostServerProfilesDefault) Code() int {
	return o._statusCode
}

func (o *PostServerProfilesDefault) Error() string {
	return fmt.Sprintf("[POST /server/Profiles][%d] PostServerProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *PostServerProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostServerProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
