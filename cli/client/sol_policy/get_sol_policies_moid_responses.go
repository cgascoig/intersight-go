// Code generated by go-swagger; DO NOT EDIT.

package sol_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetSolPoliciesMoidReader is a Reader for the GetSolPoliciesMoid structure.
type GetSolPoliciesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSolPoliciesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSolPoliciesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSolPoliciesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSolPoliciesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSolPoliciesMoidOK creates a GetSolPoliciesMoidOK with default headers values
func NewGetSolPoliciesMoidOK() *GetSolPoliciesMoidOK {
	return &GetSolPoliciesMoidOK{}
}

/*GetSolPoliciesMoidOK handles this case with default header values.

An instance of solPolicy
*/
type GetSolPoliciesMoidOK struct {
	Payload *models.SolPolicy
}

func (o *GetSolPoliciesMoidOK) Error() string {
	return fmt.Sprintf("[GET /sol/Policies/{moid}][%d] getSolPoliciesMoidOK  %+v", 200, o.Payload)
}

func (o *GetSolPoliciesMoidOK) GetPayload() *models.SolPolicy {
	return o.Payload
}

func (o *GetSolPoliciesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SolPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSolPoliciesMoidNotFound creates a GetSolPoliciesMoidNotFound with default headers values
func NewGetSolPoliciesMoidNotFound() *GetSolPoliciesMoidNotFound {
	return &GetSolPoliciesMoidNotFound{}
}

/*GetSolPoliciesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetSolPoliciesMoidNotFound struct {
}

func (o *GetSolPoliciesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /sol/Policies/{moid}][%d] getSolPoliciesMoidNotFound ", 404)
}

func (o *GetSolPoliciesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSolPoliciesMoidDefault creates a GetSolPoliciesMoidDefault with default headers values
func NewGetSolPoliciesMoidDefault(code int) *GetSolPoliciesMoidDefault {
	return &GetSolPoliciesMoidDefault{
		_statusCode: code,
	}
}

/*GetSolPoliciesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetSolPoliciesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get sol policies moid default response
func (o *GetSolPoliciesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetSolPoliciesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /sol/Policies/{moid}][%d] GetSolPoliciesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetSolPoliciesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSolPoliciesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
