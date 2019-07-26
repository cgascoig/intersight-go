// Code generated by go-swagger; DO NOT EDIT.

package oauth_oauth_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetOauthOauthUsersMoidReader is a Reader for the GetOauthOauthUsersMoid structure.
type GetOauthOauthUsersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOauthOauthUsersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOauthOauthUsersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOauthOauthUsersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOauthOauthUsersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOauthOauthUsersMoidOK creates a GetOauthOauthUsersMoidOK with default headers values
func NewGetOauthOauthUsersMoidOK() *GetOauthOauthUsersMoidOK {
	return &GetOauthOauthUsersMoidOK{}
}

/*GetOauthOauthUsersMoidOK handles this case with default header values.

An instance of oauthOauthUser
*/
type GetOauthOauthUsersMoidOK struct {
	Payload *models.OauthOauthUser
}

func (o *GetOauthOauthUsersMoidOK) Error() string {
	return fmt.Sprintf("[GET /oauth/OauthUsers/{moid}][%d] getOauthOauthUsersMoidOK  %+v", 200, o.Payload)
}

func (o *GetOauthOauthUsersMoidOK) GetPayload() *models.OauthOauthUser {
	return o.Payload
}

func (o *GetOauthOauthUsersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthOauthUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthOauthUsersMoidNotFound creates a GetOauthOauthUsersMoidNotFound with default headers values
func NewGetOauthOauthUsersMoidNotFound() *GetOauthOauthUsersMoidNotFound {
	return &GetOauthOauthUsersMoidNotFound{}
}

/*GetOauthOauthUsersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetOauthOauthUsersMoidNotFound struct {
}

func (o *GetOauthOauthUsersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /oauth/OauthUsers/{moid}][%d] getOauthOauthUsersMoidNotFound ", 404)
}

func (o *GetOauthOauthUsersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOauthOauthUsersMoidDefault creates a GetOauthOauthUsersMoidDefault with default headers values
func NewGetOauthOauthUsersMoidDefault(code int) *GetOauthOauthUsersMoidDefault {
	return &GetOauthOauthUsersMoidDefault{
		_statusCode: code,
	}
}

/*GetOauthOauthUsersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetOauthOauthUsersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get oauth oauth users moid default response
func (o *GetOauthOauthUsersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetOauthOauthUsersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /oauth/OauthUsers/{moid}][%d] GetOauthOauthUsersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetOauthOauthUsersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOauthOauthUsersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}