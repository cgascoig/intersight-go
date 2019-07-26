// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostResourceGroupsReader is a Reader for the PostResourceGroups structure.
type PostResourceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostResourceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostResourceGroupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostResourceGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostResourceGroupsCreated creates a PostResourceGroupsCreated with default headers values
func NewPostResourceGroupsCreated() *PostResourceGroupsCreated {
	return &PostResourceGroupsCreated{}
}

/*PostResourceGroupsCreated handles this case with default header values.

Null response
*/
type PostResourceGroupsCreated struct {
}

func (o *PostResourceGroupsCreated) Error() string {
	return fmt.Sprintf("[POST /resource/Groups][%d] postResourceGroupsCreated ", 201)
}

func (o *PostResourceGroupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostResourceGroupsDefault creates a PostResourceGroupsDefault with default headers values
func NewPostResourceGroupsDefault(code int) *PostResourceGroupsDefault {
	return &PostResourceGroupsDefault{
		_statusCode: code,
	}
}

/*PostResourceGroupsDefault handles this case with default header values.

unexpected error
*/
type PostResourceGroupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post resource groups default response
func (o *PostResourceGroupsDefault) Code() int {
	return o._statusCode
}

func (o *PostResourceGroupsDefault) Error() string {
	return fmt.Sprintf("[POST /resource/Groups][%d] PostResourceGroups default  %+v", o._statusCode, o.Payload)
}

func (o *PostResourceGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostResourceGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
