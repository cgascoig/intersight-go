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

// GetAssetClusterMembersReader is a Reader for the GetAssetClusterMembers structure.
type GetAssetClusterMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetClusterMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetClusterMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAssetClusterMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssetClusterMembersOK creates a GetAssetClusterMembersOK with default headers values
func NewGetAssetClusterMembersOK() *GetAssetClusterMembersOK {
	return &GetAssetClusterMembersOK{}
}

/*GetAssetClusterMembersOK handles this case with default header values.

List of assetClusterMembers for the given filter criteria
*/
type GetAssetClusterMembersOK struct {
	Payload *models.AssetClusterMemberList
}

func (o *GetAssetClusterMembersOK) Error() string {
	return fmt.Sprintf("[GET /asset/ClusterMembers][%d] getAssetClusterMembersOK  %+v", 200, o.Payload)
}

func (o *GetAssetClusterMembersOK) GetPayload() *models.AssetClusterMemberList {
	return o.Payload
}

func (o *GetAssetClusterMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetClusterMemberList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetClusterMembersDefault creates a GetAssetClusterMembersDefault with default headers values
func NewGetAssetClusterMembersDefault(code int) *GetAssetClusterMembersDefault {
	return &GetAssetClusterMembersDefault{
		_statusCode: code,
	}
}

/*GetAssetClusterMembersDefault handles this case with default header values.

Unexpected error
*/
type GetAssetClusterMembersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get asset cluster members default response
func (o *GetAssetClusterMembersDefault) Code() int {
	return o._statusCode
}

func (o *GetAssetClusterMembersDefault) Error() string {
	return fmt.Sprintf("[GET /asset/ClusterMembers][%d] GetAssetClusterMembers default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssetClusterMembersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAssetClusterMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
