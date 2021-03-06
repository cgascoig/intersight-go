// Code generated by go-swagger; DO NOT EDIT.

package iam_privilege_set

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIamPrivilegeSetsReader is a Reader for the GetIamPrivilegeSets structure.
type GetIamPrivilegeSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIamPrivilegeSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIamPrivilegeSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIamPrivilegeSetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIamPrivilegeSetsOK creates a GetIamPrivilegeSetsOK with default headers values
func NewGetIamPrivilegeSetsOK() *GetIamPrivilegeSetsOK {
	return &GetIamPrivilegeSetsOK{}
}

/*GetIamPrivilegeSetsOK handles this case with default header values.

List of iamPrivilegeSets for the given filter criteria
*/
type GetIamPrivilegeSetsOK struct {
	Payload *models.IamPrivilegeSetList
}

func (o *GetIamPrivilegeSetsOK) Error() string {
	return fmt.Sprintf("[GET /iam/PrivilegeSets][%d] getIamPrivilegeSetsOK  %+v", 200, o.Payload)
}

func (o *GetIamPrivilegeSetsOK) GetPayload() *models.IamPrivilegeSetList {
	return o.Payload
}

func (o *GetIamPrivilegeSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamPrivilegeSetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIamPrivilegeSetsDefault creates a GetIamPrivilegeSetsDefault with default headers values
func NewGetIamPrivilegeSetsDefault(code int) *GetIamPrivilegeSetsDefault {
	return &GetIamPrivilegeSetsDefault{
		_statusCode: code,
	}
}

/*GetIamPrivilegeSetsDefault handles this case with default header values.

Unexpected error
*/
type GetIamPrivilegeSetsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iam privilege sets default response
func (o *GetIamPrivilegeSetsDefault) Code() int {
	return o._statusCode
}

func (o *GetIamPrivilegeSetsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/PrivilegeSets][%d] GetIamPrivilegeSets default  %+v", o._statusCode, o.Payload)
}

func (o *GetIamPrivilegeSetsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIamPrivilegeSetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
