// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_network_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vnic fc network policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vnic fc network policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVnicFcNetworkPoliciesMoid deletes an instance of vnic fc network policy
*/
func (a *Client) DeleteVnicFcNetworkPoliciesMoid(params *DeleteVnicFcNetworkPoliciesMoidParams) (*DeleteVnicFcNetworkPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVnicFcNetworkPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVnicFcNetworkPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/vnic/FcNetworkPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVnicFcNetworkPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVnicFcNetworkPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVnicFcNetworkPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicFcNetworkPolicies gets a list of vnic fc network policy instances
*/
func (a *Client) GetVnicFcNetworkPolicies(params *GetVnicFcNetworkPoliciesParams) (*GetVnicFcNetworkPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicFcNetworkPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicFcNetworkPolicies",
		Method:             "GET",
		PathPattern:        "/vnic/FcNetworkPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicFcNetworkPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicFcNetworkPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicFcNetworkPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicFcNetworkPoliciesMoid gets a specific instance of vnic fc network policy
*/
func (a *Client) GetVnicFcNetworkPoliciesMoid(params *GetVnicFcNetworkPoliciesMoidParams) (*GetVnicFcNetworkPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicFcNetworkPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicFcNetworkPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/vnic/FcNetworkPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicFcNetworkPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicFcNetworkPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicFcNetworkPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchVnicFcNetworkPoliciesMoid updates an instance of vnic fc network policy
*/
func (a *Client) PatchVnicFcNetworkPoliciesMoid(params *PatchVnicFcNetworkPoliciesMoidParams) (*PatchVnicFcNetworkPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchVnicFcNetworkPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchVnicFcNetworkPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/vnic/FcNetworkPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVnicFcNetworkPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchVnicFcNetworkPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchVnicFcNetworkPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicFcNetworkPolicies creates an instance of vnic fc network policy
*/
func (a *Client) PostVnicFcNetworkPolicies(params *PostVnicFcNetworkPoliciesParams) (*PostVnicFcNetworkPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicFcNetworkPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicFcNetworkPolicies",
		Method:             "POST",
		PathPattern:        "/vnic/FcNetworkPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicFcNetworkPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicFcNetworkPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicFcNetworkPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicFcNetworkPoliciesMoid updates an instance of vnic fc network policy
*/
func (a *Client) PostVnicFcNetworkPoliciesMoid(params *PostVnicFcNetworkPoliciesMoidParams) (*PostVnicFcNetworkPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicFcNetworkPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicFcNetworkPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/vnic/FcNetworkPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicFcNetworkPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicFcNetworkPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicFcNetworkPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
