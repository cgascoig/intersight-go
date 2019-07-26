// Code generated by go-swagger; DO NOT EDIT.

package iaas_connector_pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetIaasConnectorPacksReader is a Reader for the GetIaasConnectorPacks structure.
type GetIaasConnectorPacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIaasConnectorPacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIaasConnectorPacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIaasConnectorPacksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIaasConnectorPacksOK creates a GetIaasConnectorPacksOK with default headers values
func NewGetIaasConnectorPacksOK() *GetIaasConnectorPacksOK {
	return &GetIaasConnectorPacksOK{}
}

/*GetIaasConnectorPacksOK handles this case with default header values.

List of iaasConnectorPacks for the given filter criteria
*/
type GetIaasConnectorPacksOK struct {
	Payload *models.IaasConnectorPackList
}

func (o *GetIaasConnectorPacksOK) Error() string {
	return fmt.Sprintf("[GET /iaas/ConnectorPacks][%d] getIaasConnectorPacksOK  %+v", 200, o.Payload)
}

func (o *GetIaasConnectorPacksOK) GetPayload() *models.IaasConnectorPackList {
	return o.Payload
}

func (o *GetIaasConnectorPacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaasConnectorPackList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIaasConnectorPacksDefault creates a GetIaasConnectorPacksDefault with default headers values
func NewGetIaasConnectorPacksDefault(code int) *GetIaasConnectorPacksDefault {
	return &GetIaasConnectorPacksDefault{
		_statusCode: code,
	}
}

/*GetIaasConnectorPacksDefault handles this case with default header values.

Unexpected error
*/
type GetIaasConnectorPacksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get iaas connector packs default response
func (o *GetIaasConnectorPacksDefault) Code() int {
	return o._statusCode
}

func (o *GetIaasConnectorPacksDefault) Error() string {
	return fmt.Sprintf("[GET /iaas/ConnectorPacks][%d] GetIaasConnectorPacks default  %+v", o._statusCode, o.Payload)
}

func (o *GetIaasConnectorPacksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIaasConnectorPacksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
