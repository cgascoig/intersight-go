// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_if

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vnic eth if API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vnic eth if API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVnicEthIfsMoid deletes an instance of vnic eth if
*/
func (a *Client) DeleteVnicEthIfsMoid(params *DeleteVnicEthIfsMoidParams) (*DeleteVnicEthIfsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVnicEthIfsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVnicEthIfsMoid",
		Method:             "DELETE",
		PathPattern:        "/vnic/EthIfs/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVnicEthIfsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVnicEthIfsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVnicEthIfsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicEthIfs gets a list of vnic eth if instances
*/
func (a *Client) GetVnicEthIfs(params *GetVnicEthIfsParams) (*GetVnicEthIfsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicEthIfsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicEthIfs",
		Method:             "GET",
		PathPattern:        "/vnic/EthIfs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicEthIfsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicEthIfsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicEthIfsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicEthIfsMoid gets a specific instance of vnic eth if
*/
func (a *Client) GetVnicEthIfsMoid(params *GetVnicEthIfsMoidParams) (*GetVnicEthIfsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicEthIfsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicEthIfsMoid",
		Method:             "GET",
		PathPattern:        "/vnic/EthIfs/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicEthIfsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicEthIfsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicEthIfsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchVnicEthIfsMoid updates an instance of vnic eth if
*/
func (a *Client) PatchVnicEthIfsMoid(params *PatchVnicEthIfsMoidParams) (*PatchVnicEthIfsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchVnicEthIfsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchVnicEthIfsMoid",
		Method:             "PATCH",
		PathPattern:        "/vnic/EthIfs/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVnicEthIfsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchVnicEthIfsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchVnicEthIfsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicEthIfs creates an instance of vnic eth if
*/
func (a *Client) PostVnicEthIfs(params *PostVnicEthIfsParams) (*PostVnicEthIfsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicEthIfsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicEthIfs",
		Method:             "POST",
		PathPattern:        "/vnic/EthIfs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicEthIfsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicEthIfsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicEthIfsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicEthIfsMoid updates an instance of vnic eth if
*/
func (a *Client) PostVnicEthIfsMoid(params *PostVnicEthIfsMoidParams) (*PostVnicEthIfsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicEthIfsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicEthIfsMoid",
		Method:             "POST",
		PathPattern:        "/vnic/EthIfs/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicEthIfsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicEthIfsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicEthIfsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
