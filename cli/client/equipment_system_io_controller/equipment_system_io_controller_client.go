// Code generated by go-swagger; DO NOT EDIT.

package equipment_system_io_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new equipment system io controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for equipment system io controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEquipmentSystemIoControllers gets a list of equipment system io controller instances
*/
func (a *Client) GetEquipmentSystemIoControllers(params *GetEquipmentSystemIoControllersParams) (*GetEquipmentSystemIoControllersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentSystemIoControllersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentSystemIoControllers",
		Method:             "GET",
		PathPattern:        "/equipment/SystemIoControllers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentSystemIoControllersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentSystemIoControllersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentSystemIoControllersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEquipmentSystemIoControllersMoid gets a specific instance of equipment system io controller
*/
func (a *Client) GetEquipmentSystemIoControllersMoid(params *GetEquipmentSystemIoControllersMoidParams) (*GetEquipmentSystemIoControllersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEquipmentSystemIoControllersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEquipmentSystemIoControllersMoid",
		Method:             "GET",
		PathPattern:        "/equipment/SystemIoControllers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEquipmentSystemIoControllersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEquipmentSystemIoControllersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEquipmentSystemIoControllersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchEquipmentSystemIoControllersMoid updates an instance of equipment system io controller
*/
func (a *Client) PatchEquipmentSystemIoControllersMoid(params *PatchEquipmentSystemIoControllersMoidParams) (*PatchEquipmentSystemIoControllersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEquipmentSystemIoControllersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEquipmentSystemIoControllersMoid",
		Method:             "PATCH",
		PathPattern:        "/equipment/SystemIoControllers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEquipmentSystemIoControllersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEquipmentSystemIoControllersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchEquipmentSystemIoControllersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostEquipmentSystemIoControllersMoid updates an instance of equipment system io controller
*/
func (a *Client) PostEquipmentSystemIoControllersMoid(params *PostEquipmentSystemIoControllersMoidParams) (*PostEquipmentSystemIoControllersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEquipmentSystemIoControllersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEquipmentSystemIoControllersMoid",
		Method:             "POST",
		PathPattern:        "/equipment/SystemIoControllers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEquipmentSystemIoControllersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEquipmentSystemIoControllersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEquipmentSystemIoControllersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
