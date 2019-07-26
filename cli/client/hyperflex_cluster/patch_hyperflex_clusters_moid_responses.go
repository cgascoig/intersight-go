// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// PatchHyperflexClustersMoidReader is a Reader for the PatchHyperflexClustersMoid structure.
type PatchHyperflexClustersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchHyperflexClustersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPatchHyperflexClustersMoidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchHyperflexClustersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchHyperflexClustersMoidCreated creates a PatchHyperflexClustersMoidCreated with default headers values
func NewPatchHyperflexClustersMoidCreated() *PatchHyperflexClustersMoidCreated {
	return &PatchHyperflexClustersMoidCreated{}
}

/*PatchHyperflexClustersMoidCreated handles this case with default header values.

Null response
*/
type PatchHyperflexClustersMoidCreated struct {
}

func (o *PatchHyperflexClustersMoidCreated) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/Clusters/{moid}][%d] patchHyperflexClustersMoidCreated ", 201)
}

func (o *PatchHyperflexClustersMoidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchHyperflexClustersMoidDefault creates a PatchHyperflexClustersMoidDefault with default headers values
func NewPatchHyperflexClustersMoidDefault(code int) *PatchHyperflexClustersMoidDefault {
	return &PatchHyperflexClustersMoidDefault{
		_statusCode: code,
	}
}

/*PatchHyperflexClustersMoidDefault handles this case with default header values.

unexpected error
*/
type PatchHyperflexClustersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch hyperflex clusters moid default response
func (o *PatchHyperflexClustersMoidDefault) Code() int {
	return o._statusCode
}

func (o *PatchHyperflexClustersMoidDefault) Error() string {
	return fmt.Sprintf("[PATCH /hyperflex/Clusters/{moid}][%d] PatchHyperflexClustersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *PatchHyperflexClustersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchHyperflexClustersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
