// Code generated by go-swagger; DO NOT EDIT.

package equipment_io_card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostEquipmentIoCardsMoidReader is a Reader for the PostEquipmentIoCardsMoid structure.
type PostEquipmentIoCardsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEquipmentIoCardsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostEquipmentIoCardsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEquipmentIoCardsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEquipmentIoCardsMoidCreated creates a PostEquipmentIoCardsMoidCreated with default headers values
func NewPostEquipmentIoCardsMoidCreated() *PostEquipmentIoCardsMoidCreated {
	return &PostEquipmentIoCardsMoidCreated{}
}

/*PostEquipmentIoCardsMoidCreated handles this case with default header values.

Null response
*/
type PostEquipmentIoCardsMoidCreated struct {
}

func (o *PostEquipmentIoCardsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /equipment/IoCards/{moid}][%d] postEquipmentIoCardsMoidCreated ", 201)
}

func (o *PostEquipmentIoCardsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEquipmentIoCardsMoidDefault creates a PostEquipmentIoCardsMoidDefault with default headers values
func NewPostEquipmentIoCardsMoidDefault(code int) *PostEquipmentIoCardsMoidDefault {
	return &PostEquipmentIoCardsMoidDefault{
		_statusCode: code,
	}
}

/*PostEquipmentIoCardsMoidDefault handles this case with default header values.

unexpected error
*/
type PostEquipmentIoCardsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post equipment io cards moid default response
func (o *PostEquipmentIoCardsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostEquipmentIoCardsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /equipment/IoCards/{moid}][%d] PostEquipmentIoCardsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostEquipmentIoCardsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostEquipmentIoCardsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}