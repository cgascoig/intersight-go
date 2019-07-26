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

// GetServerProfilesReader is a Reader for the GetServerProfiles structure.
type GetServerProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServerProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerProfilesOK creates a GetServerProfilesOK with default headers values
func NewGetServerProfilesOK() *GetServerProfilesOK {
	return &GetServerProfilesOK{}
}

/*GetServerProfilesOK handles this case with default header values.

List of serverProfiles for the given filter criteria
*/
type GetServerProfilesOK struct {
	Payload *models.ServerProfileList
}

func (o *GetServerProfilesOK) Error() string {
	return fmt.Sprintf("[GET /server/Profiles][%d] getServerProfilesOK  %+v", 200, o.Payload)
}

func (o *GetServerProfilesOK) GetPayload() *models.ServerProfileList {
	return o.Payload
}

func (o *GetServerProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerProfileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerProfilesDefault creates a GetServerProfilesDefault with default headers values
func NewGetServerProfilesDefault(code int) *GetServerProfilesDefault {
	return &GetServerProfilesDefault{
		_statusCode: code,
	}
}

/*GetServerProfilesDefault handles this case with default header values.

Unexpected error
*/
type GetServerProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get server profiles default response
func (o *GetServerProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetServerProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /server/Profiles][%d] GetServerProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}