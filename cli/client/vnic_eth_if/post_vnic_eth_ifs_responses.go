// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_if

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PostVnicEthIfsReader is a Reader for the PostVnicEthIfs structure.
type PostVnicEthIfsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVnicEthIfsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostVnicEthIfsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostVnicEthIfsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVnicEthIfsCreated creates a PostVnicEthIfsCreated with default headers values
func NewPostVnicEthIfsCreated() *PostVnicEthIfsCreated {
	return &PostVnicEthIfsCreated{}
}

/*PostVnicEthIfsCreated handles this case with default header values.

Null response
*/
type PostVnicEthIfsCreated struct {
}

func (o *PostVnicEthIfsCreated) Error() string {
	return fmt.Sprintf("[POST /vnic/EthIfs][%d] postVnicEthIfsCreated ", 201)
}

func (o *PostVnicEthIfsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVnicEthIfsDefault creates a PostVnicEthIfsDefault with default headers values
func NewPostVnicEthIfsDefault(code int) *PostVnicEthIfsDefault {
	return &PostVnicEthIfsDefault{
		_statusCode: code,
	}
}

/*PostVnicEthIfsDefault handles this case with default header values.

unexpected error
*/
type PostVnicEthIfsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post vnic eth ifs default response
func (o *PostVnicEthIfsDefault) Code() int {
	return o._statusCode
}

func (o *PostVnicEthIfsDefault) Error() string {
	return fmt.Sprintf("[POST /vnic/EthIfs][%d] PostVnicEthIfs default  %+v", o._statusCode, o.Payload)
}

func (o *PostVnicEthIfsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVnicEthIfsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
