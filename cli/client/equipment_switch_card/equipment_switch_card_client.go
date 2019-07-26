// Code generated by go-swagger; DO NOT EDIT.

package equipment_switch_card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new equipment switch card API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for equipment switch card API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEquipmentSwitchCards gets a list of equipment switch card instances
*/
func (a *Client) GetEquipmentSwitchCards(params *GetEquipmentSwitchCardsParams) (*GetEquipmentSwitchCardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentSwitchCardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentSwitchCards",
		Method:             "GET",
		PathPattern:        "/equipment/SwitchCards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentSwitchCardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentSwitchCardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentSwitchCardsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEquipmentSwitchCardsMoid gets a specific instance of equipment switch card
*/
func (a *Client) GetEquipmentSwitchCardsMoid(params *GetEquipmentSwitchCardsMoidParams) (*GetEquipmentSwitchCardsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentSwitchCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentSwitchCardsMoid",
		Method:             "GET",
		PathPattern:        "/equipment/SwitchCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentSwitchCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentSwitchCardsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentSwitchCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchEquipmentSwitchCardsMoid updates an instance of equipment switch card
*/
func (a *Client) PatchEquipmentSwitchCardsMoid(params *PatchEquipmentSwitchCardsMoidParams) (*PatchEquipmentSwitchCardsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEquipmentSwitchCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEquipmentSwitchCardsMoid",
		Method:             "PATCH",
		PathPattern:        "/equipment/SwitchCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEquipmentSwitchCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEquipmentSwitchCardsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchEquipmentSwitchCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostEquipmentSwitchCardsMoid updates an instance of equipment switch card
*/
func (a *Client) PostEquipmentSwitchCardsMoid(params *PostEquipmentSwitchCardsMoidParams) (*PostEquipmentSwitchCardsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEquipmentSwitchCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEquipmentSwitchCardsMoid",
		Method:             "POST",
		PathPattern:        "/equipment/SwitchCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEquipmentSwitchCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEquipmentSwitchCardsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEquipmentSwitchCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}