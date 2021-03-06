// Code generated by go-swagger; DO NOT EDIT.

package network_element

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostNetworkElementsMoidReader is a Reader for the PostNetworkElementsMoid structure.
type PostNetworkElementsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworkElementsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworkElementsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworkElementsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworkElementsMoidCreated creates a PostNetworkElementsMoidCreated with default headers values
func NewPostNetworkElementsMoidCreated() *PostNetworkElementsMoidCreated {
	return &PostNetworkElementsMoidCreated{}
}

/*PostNetworkElementsMoidCreated handles this case with default header values.

Null response
*/
type PostNetworkElementsMoidCreated struct {
}

func (o *PostNetworkElementsMoidCreated) Error() string {
	return fmt.Sprintf("[POST /network/Elements/{moid}][%d] postNetworkElementsMoidCreated ", 201)
}

func (o *PostNetworkElementsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworkElementsMoidDefault creates a PostNetworkElementsMoidDefault with default headers values
func NewPostNetworkElementsMoidDefault(code int) *PostNetworkElementsMoidDefault {
	return &PostNetworkElementsMoidDefault{
		_statusCode: code,
	}
}

/*PostNetworkElementsMoidDefault handles this case with default header values.

unexpected error
*/
type PostNetworkElementsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post network elements moid default response
func (o *PostNetworkElementsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworkElementsMoidDefault) Error() string {
	return fmt.Sprintf("[POST /network/Elements/{moid}][%d] PostNetworkElementsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PostNetworkElementsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworkElementsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
