// Code generated by go-swagger; DO NOT EDIT.

package equipment_io_expander

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostEquipmentIoExpandersMoidReader is a Reader for the PostEquipmentIoExpandersMoid structure.
type PostEquipmentIoExpandersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEquipmentIoExpandersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostEquipmentIoExpandersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEquipmentIoExpandersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEquipmentIoExpandersMoidCreated creates a PostEquipmentIoExpandersMoidCreated with default headers values
func NewPostEquipmentIoExpandersMoidCreated() *PostEquipmentIoExpandersMoidCreated {
	return &PostEquipmentIoExpandersMoidCreated{}
}

/*PostEquipmentIoExpandersMoidCreated handles this case with default header values.

Null response
*/
type PostEquipmentIoExpandersMoidCreated struct {
}

func (o *PostEquipmentIoExpandersMoidCreated) Error() string {
	return fmt.Sprintf("[POST /equipment/IoExpanders/{moid}][%d] postEquipmentIoExpandersMoidCreated ", 201)
}

func (o *PostEquipmentIoExpandersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEquipmentIoExpandersMoidDefault creates a PostEquipmentIoExpandersMoidDefault with default headers values
func NewPostEquipmentIoExpandersMoidDefault(code int) *PostEquipmentIoExpandersMoidDefault {
	return &PostEquipmentIoExpandersMoidDefault{
		_statusCode: code,
	}
}

/*PostEquipmentIoExpandersMoidDefault handles this case with default header values.

unexpected error
*/
type PostEquipmentIoExpandersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post equipment io expanders moid default response
func (o *PostEquipmentIoExpandersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostEquipmentIoExpandersMoidDefault) Error() string {
	return fmt.Sprintf("[POST /equipment/IoExpanders/{moid}][%d] PostEquipmentIoExpandersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostEquipmentIoExpandersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostEquipmentIoExpandersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
