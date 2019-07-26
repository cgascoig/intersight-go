// Code generated by go-swagger; DO NOT EDIT.

package search_search_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetSearchSearchItemsMoidReader is a Reader for the GetSearchSearchItemsMoid structure.
type GetSearchSearchItemsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchSearchItemsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchSearchItemsMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSearchSearchItemsMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSearchSearchItemsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSearchSearchItemsMoidOK creates a GetSearchSearchItemsMoidOK with default headers values
func NewGetSearchSearchItemsMoidOK() *GetSearchSearchItemsMoidOK {
	return &GetSearchSearchItemsMoidOK{}
}

/*GetSearchSearchItemsMoidOK handles this case with default header values.

An instance of searchSearchItem
*/
type GetSearchSearchItemsMoidOK struct {
	Payload *models.SearchSearchItem
}

func (o *GetSearchSearchItemsMoidOK) Error() string {
	return fmt.Sprintf("[GET /search/SearchItems/{moid}][%d] getSearchSearchItemsMoidOK  %+v", 200, o.Payload)
}

func (o *GetSearchSearchItemsMoidOK) GetPayload() *models.SearchSearchItem {
	return o.Payload
}

func (o *GetSearchSearchItemsMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchSearchItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchSearchItemsMoidNotFound creates a GetSearchSearchItemsMoidNotFound with default headers values
func NewGetSearchSearchItemsMoidNotFound() *GetSearchSearchItemsMoidNotFound {
	return &GetSearchSearchItemsMoidNotFound{}
}

/*GetSearchSearchItemsMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetSearchSearchItemsMoidNotFound struct {
}

func (o *GetSearchSearchItemsMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /search/SearchItems/{moid}][%d] getSearchSearchItemsMoidNotFound ", 404)
}

func (o *GetSearchSearchItemsMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchSearchItemsMoidDefault creates a GetSearchSearchItemsMoidDefault with default headers values
func NewGetSearchSearchItemsMoidDefault(code int) *GetSearchSearchItemsMoidDefault {
	return &GetSearchSearchItemsMoidDefault{
		_statusCode: code,
	}
}

/*GetSearchSearchItemsMoidDefault handles this case with default header values.

Unexpected error
*/
type GetSearchSearchItemsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get search search items moid default response
func (o *GetSearchSearchItemsMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetSearchSearchItemsMoidDefault) Error() string {
	return fmt.Sprintf("[GET /search/SearchItems/{moid}][%d] GetSearchSearchItemsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetSearchSearchItemsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSearchSearchItemsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
