// Code generated by go-swagger; DO NOT EDIT.

package iam_user_preference

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamUserPreferencesReader is a Reader for the GetIamUserPreferences structure.
type GetIamUserPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamUserPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamUserPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamUserPreferencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamUserPreferencesOK creates a GetIamUserPreferencesOK with default headers values
func NewGetIamUserPreferencesOK() *GetIamUserPreferencesOK {
	return &GetIamUserPreferencesOK{}
}

/*GetIamUserPreferencesOK handles this case with default header values.

List of iamUserPreferences for the given filter criteria
*/
type GetIamUserPreferencesOK struct {
	Payload *models.IamUserPreferenceList
}

func (o *GetIamUserPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /iam/UserPreferences][%d] getIamUserPreferencesOK  %+v", 200, o.Payload)
}

func (o *GetIamUserPreferencesOK) GetPayload() *models.IamUserPreferenceList {
	return o.Payload
}

func (o *GetIamUserPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamUserPreferenceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamUserPreferencesDefault creates a GetIamUserPreferencesDefault with default headers values
func NewGetIamUserPreferencesDefault(code int) *GetIamUserPreferencesDefault {
	return &GetIamUserPreferencesDefault{
		_statusCode: code,
	}
}

/*GetIamUserPreferencesDefault handles this case with default header values.

Unexpected error
*/
type GetIamUserPreferencesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam user preferences default response
func (o *GetIamUserPreferencesDefault) Code() int {
	return o._statusCode
}

func (o *GetIamUserPreferencesDefault) Error() string {
	return fmt.Sprintf("[GET /iam/UserPreferences][%d] GetIamUserPreferences default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamUserPreferencesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamUserPreferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
