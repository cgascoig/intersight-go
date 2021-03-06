// Code generated by go-swagger; DO NOT EDIT.

package management_entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetManagementEntitiesReader is a Reader for the GetManagementEntities structure.
type GetManagementEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetManagementEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetManagementEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetManagementEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetManagementEntitiesOK creates a GetManagementEntitiesOK with default headers values
func NewGetManagementEntitiesOK() *GetManagementEntitiesOK {
	return &GetManagementEntitiesOK{}
}

/*GetManagementEntitiesOK handles this case with default header values.

List of managementEntities for the given filter criteria
*/
type GetManagementEntitiesOK struct {
	Payload *models.ManagementEntityList
}

func (o *GetManagementEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /management/Entities][%d] getManagementEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetManagementEntitiesOK) GetPayload() *models.ManagementEntityList {
	return o.Payload
}

func (o *GetManagementEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementEntityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetManagementEntitiesDefault creates a GetManagementEntitiesDefault with default headers values
func NewGetManagementEntitiesDefault(code int) *GetManagementEntitiesDefault {
	return &GetManagementEntitiesDefault{
		_statusCode: code,
	}
}

/*GetManagementEntitiesDefault handles this case with default header values.

Unexpected error
*/
type GetManagementEntitiesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get management entities default response
func (o *GetManagementEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *GetManagementEntitiesDefault) Error() string {
	return fmt.Sprintf("[GET /management/Entities][%d] GetManagementEntities default  %+v", o._statusCode, o.Payload)
}

func (o *GetManagementEntitiesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetManagementEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
