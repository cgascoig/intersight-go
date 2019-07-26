// Code generated by go-swagger; DO NOT EDIT.

package equipment_rack_enclosure_slot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new equipment rack enclosure slot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for equipment rack enclosure slot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEquipmentRackEnclosureSlots gets a list of equipment rack enclosure slot instances
*/
func (a *Client) GetEquipmentRackEnclosureSlots(params *GetEquipmentRackEnclosureSlotsParams) (*GetEquipmentRackEnclosureSlotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentRackEnclosureSlotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentRackEnclosureSlots",
		Method:             "GET",
		PathPattern:        "/equipment/RackEnclosureSlots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentRackEnclosureSlotsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentRackEnclosureSlotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentRackEnclosureSlotsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEquipmentRackEnclosureSlotsMoid gets a specific instance of equipment rack enclosure slot
*/
func (a *Client) GetEquipmentRackEnclosureSlotsMoid(params *GetEquipmentRackEnclosureSlotsMoidParams) (*GetEquipmentRackEnclosureSlotsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentRackEnclosureSlotsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentRackEnclosureSlotsMoid",
		Method:             "GET",
		PathPattern:        "/equipment/RackEnclosureSlots/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentRackEnclosureSlotsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentRackEnclosureSlotsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentRackEnclosureSlotsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchEquipmentRackEnclosureSlotsMoid updates an instance of equipment rack enclosure slot
*/
func (a *Client) PatchEquipmentRackEnclosureSlotsMoid(params *PatchEquipmentRackEnclosureSlotsMoidParams) (*PatchEquipmentRackEnclosureSlotsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEquipmentRackEnclosureSlotsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEquipmentRackEnclosureSlotsMoid",
		Method:             "PATCH",
		PathPattern:        "/equipment/RackEnclosureSlots/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEquipmentRackEnclosureSlotsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEquipmentRackEnclosureSlotsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchEquipmentRackEnclosureSlotsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostEquipmentRackEnclosureSlotsMoid updates an instance of equipment rack enclosure slot
*/
func (a *Client) PostEquipmentRackEnclosureSlotsMoid(params *PostEquipmentRackEnclosureSlotsMoidParams) (*PostEquipmentRackEnclosureSlotsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEquipmentRackEnclosureSlotsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEquipmentRackEnclosureSlotsMoid",
		Method:             "POST",
		PathPattern:        "/equipment/RackEnclosureSlots/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEquipmentRackEnclosureSlotsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEquipmentRackEnclosureSlotsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEquipmentRackEnclosureSlotsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
