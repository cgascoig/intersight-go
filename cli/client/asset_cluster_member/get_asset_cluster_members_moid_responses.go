// Code generated by go-swagger; DO NOT EDIT.

package asset_cluster_member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetAssetClusterMembersMoidReader is a Reader for the GetAssetClusterMembersMoid structure.
type GetAssetClusterMembersMoidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetClusterMembersMoidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetClusterMembersMoidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAssetClusterMembersMoidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAssetClusterMembersMoidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssetClusterMembersMoidOK creates a GetAssetClusterMembersMoidOK with default headers values
func NewGetAssetClusterMembersMoidOK() *GetAssetClusterMembersMoidOK {
	return &GetAssetClusterMembersMoidOK{}
}

/*GetAssetClusterMembersMoidOK handles this case with default header values.

An instance of assetClusterMember
*/
type GetAssetClusterMembersMoidOK struct {
	Payload *models.AssetClusterMember
}

func (o *GetAssetClusterMembersMoidOK) Error() string {
	return fmt.Sprintf("[GET /asset/ClusterMembers/{moid}][%d] getAssetClusterMembersMoidOK  %+v", 200, o.Payload)
}

func (o *GetAssetClusterMembersMoidOK) GetPayload() *models.AssetClusterMember {
	return o.Payload
}

func (o *GetAssetClusterMembersMoidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetClusterMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetClusterMembersMoidNotFound creates a GetAssetClusterMembersMoidNotFound with default headers values
func NewGetAssetClusterMembersMoidNotFound() *GetAssetClusterMembersMoidNotFound {
	return &GetAssetClusterMembersMoidNotFound{}
}

/*GetAssetClusterMembersMoidNotFound handles this case with default header values.

Instance not found.
*/
type GetAssetClusterMembersMoidNotFound struct {
}

func (o *GetAssetClusterMembersMoidNotFound) Error() string {
	return fmt.Sprintf("[GET /asset/ClusterMembers/{moid}][%d] getAssetClusterMembersMoidNotFound ", 404)
}

func (o *GetAssetClusterMembersMoidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssetClusterMembersMoidDefault creates a GetAssetClusterMembersMoidDefault with default headers values
func NewGetAssetClusterMembersMoidDefault(code int) *GetAssetClusterMembersMoidDefault {
	return &GetAssetClusterMembersMoidDefault{
		_statusCode: code,
	}
}

/*GetAssetClusterMembersMoidDefault handles this case with default header values.

Unexpected error
*/
type GetAssetClusterMembersMoidDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get asset cluster members moid default response
func (o *GetAssetClusterMembersMoidDefault) Code() int {
	return o._statusCode
}

func (o *GetAssetClusterMembersMoidDefault) Error() string {
	return fmt.Sprintf("[GET /asset/ClusterMembers/{moid}][%d] GetAssetClusterMembersMoid default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssetClusterMembersMoidDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAssetClusterMembersMoidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
