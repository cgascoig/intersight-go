// Code generated by go-swagger; DO NOT EDIT.

package equipment_chassis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostEquipmentChassesMoidReader is a Reader for the PostEquipmentChassesMoid structure.
type PostEquipmentChassesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEquipmentChassesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostEquipmentChassesMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEquipmentChassesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEquipmentChassesMoidCreated creates a PostEquipmentChassesMoidCreated with default headers values
func NewPostEquipmentChassesMoidCreated() *PostEquipmentChassesMoidCreated {
	return &PostEquipmentChassesMoidCreated{}
}

/*PostEquipmentChassesMoidCreated handles this case with default header values.

Null response
*/
type PostEquipmentChassesMoidCreated struct {
}

func (o *PostEquipmentChassesMoidCreated) Error() string {
	return fmt.Sprintf("[POST /equipment/Chasses/{moid}][%d] postEquipmentChassesMoidCreated ", 201)
}

func (o *PostEquipmentChassesMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEquipmentChassesMoidDefault creates a PostEquipmentChassesMoidDefault with default headers values
func NewPostEquipmentChassesMoidDefault(code int) *PostEquipmentChassesMoidDefault {
	return &PostEquipmentChassesMoidDefault{
		_statusCode: code,
	}
}

/*PostEquipmentChassesMoidDefault handles this case with default header values.

unexpected error
*/
type PostEquipmentChassesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post equipment chasses moid default response
func (o *PostEquipmentChassesMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostEquipmentChassesMoidDefault) Error() string {
	return fmt.Sprintf("[POST /equipment/Chasses/{moid}][%d] PostEquipmentChassesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostEquipmentChassesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostEquipmentChassesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
