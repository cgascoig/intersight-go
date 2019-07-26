// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_privilege

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamEndPointPrivilegesReader is a Reader for the GetIamEndPointPrivileges structure.
type GetIamEndPointPrivilegesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamEndPointPrivilegesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamEndPointPrivilegesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamEndPointPrivilegesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamEndPointPrivilegesOK creates a GetIamEndPointPrivilegesOK with default headers values
func NewGetIamEndPointPrivilegesOK() *GetIamEndPointPrivilegesOK {
	return &GetIamEndPointPrivilegesOK{}
}

/*GetIamEndPointPrivilegesOK handles this case with default header values.

List of iamEndPointPrivileges for the given filter criteria
*/
type GetIamEndPointPrivilegesOK struct {
	Payload *models.IamEndPointPrivilegeList
}

func (o *GetIamEndPointPrivilegesOK) Error() string {
	return fmt.Sprintf("[GET /iam/EndPointPrivileges][%d] getIamEndPointPrivilegesOK  %+v", 200, o.Payload)
}

func (o *GetIamEndPointPrivilegesOK) GetPayload() *models.IamEndPointPrivilegeList {
	return o.Payload
}

func (o *GetIamEndPointPrivilegesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamEndPointPrivilegeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamEndPointPrivilegesDefault creates a GetIamEndPointPrivilegesDefault with default headers values
func NewGetIamEndPointPrivilegesDefault(code int) *GetIamEndPointPrivilegesDefault {
	return &GetIamEndPointPrivilegesDefault{
		_statusCode: code,
	}
}

/*GetIamEndPointPrivilegesDefault handles this case with default header values.

Unexpected error
*/
type GetIamEndPointPrivilegesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam end point privileges default response
func (o *GetIamEndPointPrivilegesDefault) Code() int {
	return o._statusCode
}

func (o *GetIamEndPointPrivilegesDefault) Error() string {
	return fmt.Sprintf("[GET /iam/EndPointPrivileges][%d] GetIamEndPointPrivileges default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamEndPointPrivilegesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamEndPointPrivilegesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
