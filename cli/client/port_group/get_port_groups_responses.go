// Code generated by go-swagger; DO NOT EDIT.

package port_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetPortGroupsReader is a Reader for the GetPortGroups structure.
type GetPortGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPortGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPortGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPortGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPortGroupsOK creates a GetPortGroupsOK with default headers values
func NewGetPortGroupsOK() *GetPortGroupsOK {
	return &GetPortGroupsOK{}
}

/*GetPortGroupsOK handles this case with default header values.

List of portGroups for the given filter criteria
*/
type GetPortGroupsOK struct {
	Payload *models.PortGroupList
}

func (o *GetPortGroupsOK) Error() string {
	return fmt.Sprintf("[GET /port/Groups][%d] getPortGroupsOK  %+v", 200, o.Payload)
}

func (o *GetPortGroupsOK) GetPayload() *models.PortGroupList {
	return o.Payload
}

func (o *GetPortGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPortGroupsDefault creates a GetPortGroupsDefault with default headers values
func NewGetPortGroupsDefault(code int) *GetPortGroupsDefault {
	return &GetPortGroupsDefault{
		_statusCode: code,
	}
}

/*GetPortGroupsDefault handles this case with default header values.

Unexpected error
*/
type GetPortGroupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get port groups default response
func (o *GetPortGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GetPortGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /port/Groups][%d] GetPortGroups default  %+v", o._statusCode, o.Payload)
}

func (o *GetPortGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPortGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}