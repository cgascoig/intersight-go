// Code generated by go-swagger; DO NOT EDIT.

package ls_service_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetLsServiceProfilesReader is a Reader for the GetLsServiceProfiles structure.
type GetLsServiceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLsServiceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLsServiceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLsServiceProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLsServiceProfilesOK creates a GetLsServiceProfilesOK with default headers values
func NewGetLsServiceProfilesOK() *GetLsServiceProfilesOK {
	return &GetLsServiceProfilesOK{}
}

/*GetLsServiceProfilesOK handles this case with default header values.

List of lsServiceProfiles for the given filter criteria
*/
type GetLsServiceProfilesOK struct {
	Payload *models.LsServiceProfileList
}

func (o *GetLsServiceProfilesOK) Error() string {
	return fmt.Sprintf("[GET /ls/ServiceProfiles][%d] getLsServiceProfilesOK  %+v", 200, o.Payload)
}

func (o *GetLsServiceProfilesOK) GetPayload() *models.LsServiceProfileList {
	return o.Payload
}

func (o *GetLsServiceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LsServiceProfileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLsServiceProfilesDefault creates a GetLsServiceProfilesDefault with default headers values
func NewGetLsServiceProfilesDefault(code int) *GetLsServiceProfilesDefault {
	return &GetLsServiceProfilesDefault{
		_statusCode: code,
	}
}

/*GetLsServiceProfilesDefault handles this case with default header values.

Unexpected error
*/
type GetLsServiceProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get ls service profiles default response
func (o *GetLsServiceProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetLsServiceProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /ls/ServiceProfiles][%d] GetLsServiceProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetLsServiceProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLsServiceProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
