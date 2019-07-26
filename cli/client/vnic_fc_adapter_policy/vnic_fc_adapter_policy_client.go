// Code generated by go-swagger; DO NOT EDIT.

package vnic_fc_adapter_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vnic fc adapter policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vnic fc adapter policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVnicFcAdapterPoliciesMoid deletes an instance of vnic fc adapter policy
*/
func (a *Client) DeleteVnicFcAdapterPoliciesMoid(params *DeleteVnicFcAdapterPoliciesMoidParams) (*DeleteVnicFcAdapterPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVnicFcAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVnicFcAdapterPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/vnic/FcAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVnicFcAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVnicFcAdapterPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVnicFcAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicFcAdapterPolicies gets a list of vnic fc adapter policy instances
*/
func (a *Client) GetVnicFcAdapterPolicies(params *GetVnicFcAdapterPoliciesParams) (*GetVnicFcAdapterPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicFcAdapterPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicFcAdapterPolicies",
		Method:             "GET",
		PathPattern:        "/vnic/FcAdapterPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicFcAdapterPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicFcAdapterPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicFcAdapterPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVnicFcAdapterPoliciesMoid gets a specific instance of vnic fc adapter policy
*/
func (a *Client) GetVnicFcAdapterPoliciesMoid(params *GetVnicFcAdapterPoliciesMoidParams) (*GetVnicFcAdapterPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicFcAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicFcAdapterPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/vnic/FcAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVnicFcAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVnicFcAdapterPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVnicFcAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchVnicFcAdapterPoliciesMoid updates an instance of vnic fc adapter policy
*/
func (a *Client) PatchVnicFcAdapterPoliciesMoid(params *PatchVnicFcAdapterPoliciesMoidParams) (*PatchVnicFcAdapterPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchVnicFcAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchVnicFcAdapterPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/vnic/FcAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVnicFcAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchVnicFcAdapterPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchVnicFcAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicFcAdapterPolicies creates an instance of vnic fc adapter policy
*/
func (a *Client) PostVnicFcAdapterPolicies(params *PostVnicFcAdapterPoliciesParams) (*PostVnicFcAdapterPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicFcAdapterPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicFcAdapterPolicies",
		Method:             "POST",
		PathPattern:        "/vnic/FcAdapterPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicFcAdapterPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicFcAdapterPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicFcAdapterPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVnicFcAdapterPoliciesMoid updates an instance of vnic fc adapter policy
*/
func (a *Client) PostVnicFcAdapterPoliciesMoid(params *PostVnicFcAdapterPoliciesMoidParams) (*PostVnicFcAdapterPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVnicFcAdapterPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVnicFcAdapterPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/vnic/FcAdapterPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVnicFcAdapterPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVnicFcAdapterPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVnicFcAdapterPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
