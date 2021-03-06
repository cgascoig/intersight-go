// Code generated by go-swagger; DO NOT EDIT.

package appliance_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// DeleteApplianceRestoresMoidReader is a Reader for the DeleteApplianceRestoresMoid structure.
type DeleteApplianceRestoresMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplianceRestoresMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApplianceRestoresMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteApplianceRestoresMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteApplianceRestoresMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteApplianceRestoresMoidOK creates a DeleteApplianceRestoresMoidOK with default headers values
func NewDeleteApplianceRestoresMoidOK() *DeleteApplianceRestoresMoidOK {
	return &DeleteApplianceRestoresMoidOK{}
}

/*DeleteApplianceRestoresMoidOK handles this case with default header values.

Delete successful.
*/
type DeleteApplianceRestoresMoidOK struct {
}

func (o *DeleteApplianceRestoresMoidOK) Error() string {
	return fmt.Sprintf("[DELETE /appliance/Restores/{moid}][%d] deleteApplianceRestoresMoidOK ", 200)
}

func (o *DeleteApplianceRestoresMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplianceRestoresMoidNotFound creates a DeleteApplianceRestoresMoidNotFound with default headers values
func NewDeleteApplianceRestoresMoidNotFound() *DeleteApplianceRestoresMoidNotFound {
	return &DeleteApplianceRestoresMoidNotFound{}
}

/*DeleteApplianceRestoresMoidNotFound handles this case with default header values.

Instance not found.
*/
type DeleteApplianceRestoresMoidNotFound struct {
}

func (o *DeleteApplianceRestoresMoidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /appliance/Restores/{moid}][%d] deleteApplianceRestoresMoidNotFound ", 404)
}

func (o *DeleteApplianceRestoresMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplianceRestoresMoidDefault creates a DeleteApplianceRestoresMoidDefault with default headers values
func NewDeleteApplianceRestoresMoidDefault(code int) *DeleteApplianceRestoresMoidDefault {
	return &DeleteApplianceRestoresMoidDefault{
		_statusCode: code,
	}
}

/*DeleteApplianceRestoresMoidDefault handles this case with default header values.

Unexpected error
*/
type DeleteApplianceRestoresMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete appliance restores moid default response
func (o *DeleteApplianceRestoresMoidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteApplianceRestoresMoidDefault) Error() string {
	return fmt.Sprintf("[DELETE /appliance/Restores/{moid}][%d] DeleteApplianceRestoresMoid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteApplianceRestoresMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApplianceRestoresMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
