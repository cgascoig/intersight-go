// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_app_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchHyperflexAppCatalogsMoidReader is a Reader for the PatchHyperflexAppCatalogsMoid structure.
type PatchHyperflexAppCatalogsMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchHyperflexAppCatalogsMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchHyperflexAppCatalogsMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchHyperflexAppCatalogsMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchHyperflexAppCatalogsMoidCreated creates a PatchHyperflexAppCatalogsMoidCreated with default headers values
func NewPatchHyperflexAppCatalogsMoidCreated() *PatchHyperflexAppCatalogsMoidCreated {
	return &PatchHyperflexAppCatalogsMoidCreated{}
}

/*PatchHyperflexAppCatalogsMoidCreated handles this case with default header values.

Null response
*/
type PatchHyperflexAppCatalogsMoidCreated struct {
}

func (o *PatchHyperflexAppCatalogsMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/AppCatalogs/{moid}][%d] patchHyperflexAppCatalogsMoidCreated ", 201)
}

func (o *PatchHyperflexAppCatalogsMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchHyperflexAppCatalogsMoidDefault creates a PatchHyperflexAppCatalogsMoidDefault with default headers values
func NewPatchHyperflexAppCatalogsMoidDefault(code int) *PatchHyperflexAppCatalogsMoidDefault {
	return &PatchHyperflexAppCatalogsMoidDefault{
		_statusCode: code,
	}
}

/*PatchHyperflexAppCatalogsMoidDefault handles this case with default header values.

unexpected error
*/
type PatchHyperflexAppCatalogsMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch hyperflex app catalogs moid default response
func (o *PatchHyperflexAppCatalogsMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchHyperflexAppCatalogsMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/AppCatalogs/{moid}][%d] PatchHyperflexAppCatalogsMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchHyperflexAppCatalogsMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchHyperflexAppCatalogsMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
