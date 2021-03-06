// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_feature_limit_external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetHyperflexFeatureLimitExternalsReader is a Reader for the GetHyperflexFeatureLimitExternals structure.
type GetHyperflexFeatureLimitExternalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHyperflexFeatureLimitExternalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHyperflexFeatureLimitExternalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHyperflexFeatureLimitExternalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHyperflexFeatureLimitExternalsOK creates a GetHyperflexFeatureLimitExternalsOK with default headers values
func NewGetHyperflexFeatureLimitExternalsOK() *GetHyperflexFeatureLimitExternalsOK {
	return &GetHyperflexFeatureLimitExternalsOK{}
}

/*GetHyperflexFeatureLimitExternalsOK handles this case with default header values.

List of hyperflexFeatureLimitExternals for the given filter criteria
*/
type GetHyperflexFeatureLimitExternalsOK struct {
	Payload *models.HyperflexFeatureLimitExternalList
}

func (o *GetHyperflexFeatureLimitExternalsOK) Error() string {
	return fmt.Sprintf("[GET /hyperflex/FeatureLimitExternals][%d] getHyperflexFeatureLimitExternalsOK  %+v", 200, o.Payload)
}

func (o *GetHyperflexFeatureLimitExternalsOK) GetPayload() *models.HyperflexFeatureLimitExternalList {
	return o.Payload
}

func (o *GetHyperflexFeatureLimitExternalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HyperflexFeatureLimitExternalList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHyperflexFeatureLimitExternalsDefault creates a GetHyperflexFeatureLimitExternalsDefault with default headers values
func NewGetHyperflexFeatureLimitExternalsDefault(code int) *GetHyperflexFeatureLimitExternalsDefault {
	return &GetHyperflexFeatureLimitExternalsDefault{
		_statusCode: code,
	}
}

/*GetHyperflexFeatureLimitExternalsDefault handles this case with default header values.

Unexpected error
*/
type GetHyperflexFeatureLimitExternalsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hyperflex feature limit externals default response
func (o *GetHyperflexFeatureLimitExternalsDefault) Code() int {
	return o._statusCode
}

func (o *GetHyperflexFeatureLimitExternalsDefault) Error() string {
	return fmt.Sprintf("[GET /hyperflex/FeatureLimitExternals][%d] GetHyperflexFeatureLimitExternals default  %+v", o._statusCode, o.Payload)
}

func (o *GetHyperflexFeatureLimitExternalsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHyperflexFeatureLimitExternalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
