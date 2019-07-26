// Code generated by go-swagger; DO NOT EDIT.

package adapter_config_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new adapter config policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for adapter config policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAdapterConfigPoliciesMoid deletes an instance of adapter config policy
*/
func (a *Client) DeleteAdapterConfigPoliciesMoid(params *DeleteAdapterConfigPoliciesMoidParams) (*DeleteAdapterConfigPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdapterConfigPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAdapterConfigPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/adapter/ConfigPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdapterConfigPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAdapterConfigPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAdapterConfigPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAdapterConfigPolicies gets a list of adapter config policy instances
*/
func (a *Client) GetAdapterConfigPolicies(params *GetAdapterConfigPoliciesParams) (*GetAdapterConfigPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdapterConfigPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAdapterConfigPolicies",
		Method:             "GET",
		PathPattern:        "/adapter/ConfigPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdapterConfigPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAdapterConfigPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdapterConfigPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAdapterConfigPoliciesMoid gets a specific instance of adapter config policy
*/
func (a *Client) GetAdapterConfigPoliciesMoid(params *GetAdapterConfigPoliciesMoidParams) (*GetAdapterConfigPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdapterConfigPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAdapterConfigPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/adapter/ConfigPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdapterConfigPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAdapterConfigPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdapterConfigPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchAdapterConfigPoliciesMoid updates an instance of adapter config policy
*/
func (a *Client) PatchAdapterConfigPoliciesMoid(params *PatchAdapterConfigPoliciesMoidParams) (*PatchAdapterConfigPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAdapterConfigPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAdapterConfigPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/adapter/ConfigPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAdapterConfigPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAdapterConfigPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchAdapterConfigPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostAdapterConfigPolicies creates an instance of adapter config policy
*/
func (a *Client) PostAdapterConfigPolicies(params *PostAdapterConfigPoliciesParams) (*PostAdapterConfigPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAdapterConfigPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAdapterConfigPolicies",
		Method:             "POST",
		PathPattern:        "/adapter/ConfigPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAdapterConfigPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAdapterConfigPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostAdapterConfigPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostAdapterConfigPoliciesMoid updates an instance of adapter config policy
*/
func (a *Client) PostAdapterConfigPoliciesMoid(params *PostAdapterConfigPoliciesMoidParams) (*PostAdapterConfigPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAdapterConfigPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAdapterConfigPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/adapter/ConfigPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAdapterConfigPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAdapterConfigPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostAdapterConfigPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}