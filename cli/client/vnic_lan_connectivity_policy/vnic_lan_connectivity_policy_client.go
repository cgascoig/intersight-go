// Code generated by go-swagger; DO NOT EDIT.

package vnic_lan_connectivity_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vnic lan connectivity policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vnic lan connectivity policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVnicLanConnectivityPoliciesMoid deletes an instance of vnic lan connectivity policy
*/
func (a *Client) DeleteVnicLanConnectivityPoliciesMoid(params *DeleteVnicLanConnectivityPoliciesMoidParams) (*DeleteVnicLanConnectivityPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVnicLanConnectivityPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVnicLanConnectivityPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/vnic/LanConnectivityPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVnicLanConnectivityPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVnicLanConnectivityPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVnicLanConnectivityPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicLanConnectivityPolicies gets a list of vnic lan connectivity policy instances
*/
func (a *Client) GetVnicLanConnectivityPolicies(params *GetVnicLanConnectivityPoliciesParams) (*GetVnicLanConnectivityPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicLanConnectivityPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicLanConnectivityPolicies",
		Method:             "GET",
		PathPattern:        "/vnic/LanConnectivityPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicLanConnectivityPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicLanConnectivityPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicLanConnectivityPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicLanConnectivityPoliciesMoid gets a specific instance of vnic lan connectivity policy
*/
func (a *Client) GetVnicLanConnectivityPoliciesMoid(params *GetVnicLanConnectivityPoliciesMoidParams) (*GetVnicLanConnectivityPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicLanConnectivityPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicLanConnectivityPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/vnic/LanConnectivityPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicLanConnectivityPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicLanConnectivityPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicLanConnectivityPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchVnicLanConnectivityPoliciesMoid updates an instance of vnic lan connectivity policy
*/
func (a *Client) PatchVnicLanConnectivityPoliciesMoid(params *PatchVnicLanConnectivityPoliciesMoidParams) (*PatchVnicLanConnectivityPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchVnicLanConnectivityPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchVnicLanConnectivityPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/vnic/LanConnectivityPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVnicLanConnectivityPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchVnicLanConnectivityPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchVnicLanConnectivityPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicLanConnectivityPolicies creates an instance of vnic lan connectivity policy
*/
func (a *Client) PostVnicLanConnectivityPolicies(params *PostVnicLanConnectivityPoliciesParams) (*PostVnicLanConnectivityPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicLanConnectivityPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicLanConnectivityPolicies",
		Method:             "POST",
		PathPattern:        "/vnic/LanConnectivityPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicLanConnectivityPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicLanConnectivityPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicLanConnectivityPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicLanConnectivityPoliciesMoid updates an instance of vnic lan connectivity policy
*/
func (a *Client) PostVnicLanConnectivityPoliciesMoid(params *PostVnicLanConnectivityPoliciesMoidParams) (*PostVnicLanConnectivityPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicLanConnectivityPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicLanConnectivityPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/vnic/LanConnectivityPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicLanConnectivityPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicLanConnectivityPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicLanConnectivityPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
