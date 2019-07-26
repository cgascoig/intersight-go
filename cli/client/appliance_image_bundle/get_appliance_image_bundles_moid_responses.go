// Code generated by go-swagger; DO NOT EDIT.

package appliance_image_bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetApplianceImageBundlesMoidReader is a Reader for the GetApplianceImageBundlesMoid structure.
type GetApplianceImageBundlesMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplianceImageBundlesMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplianceImageBundlesMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApplianceImageBundlesMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplianceImageBundlesMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplianceImageBundlesMoidOK creates a GetApplianceImageBundlesMoidOK with default headers values
func NewGetApplianceImageBundlesMoidOK() *GetApplianceImageBundlesMoidOK {
	return &GetApplianceImageBundlesMoidOK{}
}

/*GetApplianceImageBundlesMoidOK handles this case with default header values.

An instance of applianceImageBundle
*/
type GetApplianceImageBundlesMoidOK struct {
	Payload *models.ApplianceImageBundle
}

func (o *GetApplianceImageBundlesMoidOK) Error() string {
	return fmt.Sprintf("[GET /appliance/ImageBundles/{moid}][%d] getApplianceImageBundlesMoidOK  %+v", 200, o.Payload)
}

func (o *GetApplianceImageBundlesMoidOK) GetPayload() *models.ApplianceImageBundle {
	return o.Payload
}

func (o *GetApplianceImageBundlesMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplianceImageBundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplianceImageBundlesMoidNotFound creates a GetApplianceImageBundlesMoidNotFound with default headers values
func NewGetApplianceImageBundlesMoidNotFound() *GetApplianceImageBundlesMoidNotFound {
	return &GetApplianceImageBundlesMoidNotFound{}
}

/*GetApplianceImageBundlesMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetApplianceImageBundlesMoidNotFound struct {
}

func (o *GetApplianceImageBundlesMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /appliance/ImageBundles/{moid}][%d] getApplianceImageBundlesMoidNotFound ", 404)
}

func (o *GetApplianceImageBundlesMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplianceImageBundlesMoidDefault creates a GetApplianceImageBundlesMoidDefault with default headers values
func NewGetApplianceImageBundlesMoidDefault(code int) *GetApplianceImageBundlesMoidDefault {
	return &GetApplianceImageBundlesMoidDefault{
		_statusCode: code,
	}
}

/*GetApplianceImageBundlesMoidDefault handles this case with default header values.

Unexpected error
*/
type GetApplianceImageBundlesMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get appliance image bundles moid default response
func (o *GetApplianceImageBundlesMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetApplianceImageBundlesMoidDefault) Error() string {
	return fmt.Sprintf("[GET /appliance/ImageBundles/{moid}][%d] GetApplianceImageBundlesMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplianceImageBundlesMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplianceImageBundlesMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}