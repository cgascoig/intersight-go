// Code generated by go-swagger; DO NOT EDIT.

package vnic_eth_adapter_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vnic eth adapter policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vnic eth adapter policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVnicEthAdapterPoliciesMoid deletes an instance of vnic eth adapter policy
*/
func (a *Client) DeleteVnicEthAdapterPoliciesMoid(params *DeleteVnicEthAdapterPoliciesMoidParams) (*DeleteVnicEthAdapterPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVnicEthAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVnicEthAdapterPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/vnic/EthAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVnicEthAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVnicEthAdapterPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVnicEthAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicEthAdapterPolicies gets a list of vnic eth adapter policy instances
*/
func (a *Client) GetVnicEthAdapterPolicies(params *GetVnicEthAdapterPoliciesParams) (*GetVnicEthAdapterPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicEthAdapterPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicEthAdapterPolicies",
		Method:             "GET",
		PathPattern:        "/vnic/EthAdapterPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicEthAdapterPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicEthAdapterPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicEthAdapterPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicEthAdapterPoliciesMoid gets a specific instance of vnic eth adapter policy
*/
func (a *Client) GetVnicEthAdapterPoliciesMoid(params *GetVnicEthAdapterPoliciesMoidParams) (*GetVnicEthAdapterPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicEthAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicEthAdapterPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/vnic/EthAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicEthAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicEthAdapterPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicEthAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchVnicEthAdapterPoliciesMoid updates an instance of vnic eth adapter policy
*/
func (a *Client) PatchVnicEthAdapterPoliciesMoid(params *PatchVnicEthAdapterPoliciesMoidParams) (*PatchVnicEthAdapterPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchVnicEthAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchVnicEthAdapterPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/vnic/EthAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVnicEthAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchVnicEthAdapterPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchVnicEthAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicEthAdapterPolicies creates an instance of vnic eth adapter policy
*/
func (a *Client) PostVnicEthAdapterPolicies(params *PostVnicEthAdapterPoliciesParams) (*PostVnicEthAdapterPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicEthAdapterPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicEthAdapterPolicies",
		Method:             "POST",
		PathPattern:        "/vnic/EthAdapterPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicEthAdapterPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicEthAdapterPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicEthAdapterPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicEthAdapterPoliciesMoid updates an instance of vnic eth adapter policy
*/
func (a *Client) PostVnicEthAdapterPoliciesMoid(params *PostVnicEthAdapterPoliciesMoidParams) (*PostVnicEthAdapterPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicEthAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicEthAdapterPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/vnic/EthAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicEthAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicEthAdapterPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicEthAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
