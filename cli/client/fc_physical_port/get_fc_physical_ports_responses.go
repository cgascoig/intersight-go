// Code generated by go-swagger; DO NOT EDIT.

package fc_physical_port

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetFcPhysicalPortsReader is a Reader for the GetFcPhysicalPorts structure.
type GetFcPhysicalPortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFcPhysicalPortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFcPhysicalPortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFcPhysicalPortsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFcPhysicalPortsOK creates a GetFcPhysicalPortsOK with default headers values
func NewGetFcPhysicalPortsOK() *GetFcPhysicalPortsOK {
	return &GetFcPhysicalPortsOK{}
}

/*GetFcPhysicalPortsOK handles this case with default header values.

List of fcPhysicalPorts for the given filter criteria
*/
type GetFcPhysicalPortsOK struct {
	Payload *models.FcPhysicalPortList
}

func (o *GetFcPhysicalPortsOK) Error() string {
	return fmt.Sprintf("[GET /fc/PhysicalPorts][%d] getFcPhysicalPortsOK  %+v", 200, o.Payload)
}

func (o *GetFcPhysicalPortsOK) GetPayload() *models.FcPhysicalPortList {
	return o.Payload
}

func (o *GetFcPhysicalPortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcPhysicalPortList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFcPhysicalPortsDefault creates a GetFcPhysicalPortsDefault with default headers values
func NewGetFcPhysicalPortsDefault(code int) *GetFcPhysicalPortsDefault {
	return &GetFcPhysicalPortsDefault{
		_statusCode: code,
	}
}

/*GetFcPhysicalPortsDefault handles this case with default header values.

Unexpected error
*/
type GetFcPhysicalPortsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get fc physical ports default response
func (o *GetFcPhysicalPortsDefault) Code() int {
	return o._statusCode
}

func (o *GetFcPhysicalPortsDefault) Error() string {
	return fmt.Sprintf("[GET /fc/PhysicalPorts][%d] GetFcPhysicalPorts default  %+v", o._statusCode, o.Payload)
}

func (o *GetFcPhysicalPortsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFcPhysicalPortsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}