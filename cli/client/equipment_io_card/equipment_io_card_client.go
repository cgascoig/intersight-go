// Code generated by go-swagger; DO NOT EDIT.

package equipment_io_card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new equipment io card API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for equipment io card API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEquipmentIoCards gets a list of equipment io card instances
*/
func (a *Client) GetEquipmentIoCards(params *GetEquipmentIoCardsParams) (*GetEquipmentIoCardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentIoCardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentIoCards",
		Method:             "GET",
		PathPattern:        "/equipment/IoCards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentIoCardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentIoCardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentIoCardsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEquipmentIoCardsMoid gets a specific instance of equipment io card
*/
func (a *Client) GetEquipmentIoCardsMoid(params *GetEquipmentIoCardsMoidParams) (*GetEquipmentIoCardsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentIoCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentIoCardsMoid",
		Method:             "GET",
		PathPattern:        "/equipment/IoCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentIoCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentIoCardsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentIoCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchEquipmentIoCardsMoid updates an instance of equipment io card
*/
func (a *Client) PatchEquipmentIoCardsMoid(params *PatchEquipmentIoCardsMoidParams) (*PatchEquipmentIoCardsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEquipmentIoCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEquipmentIoCardsMoid",
		Method:             "PATCH",
		PathPattern:        "/equipment/IoCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEquipmentIoCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEquipmentIoCardsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchEquipmentIoCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostEquipmentIoCardsMoid updates an instance of equipment io card
*/
func (a *Client) PostEquipmentIoCardsMoid(params *PostEquipmentIoCardsMoidParams) (*PostEquipmentIoCardsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEquipmentIoCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEquipmentIoCardsMoid",
		Method:             "POST",
		PathPattern:        "/equipment/IoCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEquipmentIoCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEquipmentIoCardsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEquipmentIoCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
