// Code generated by go-swagger; DO NOT EDIT.

package vmedia_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vmedia policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vmedia policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVmediaPoliciesMoid deletes an instance of vmedia policy
*/
func (a *Client) DeleteVmediaPoliciesMoid(params *DeleteVmediaPoliciesMoidParams) (*DeleteVmediaPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVmediaPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVmediaPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/vmedia/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVmediaPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVmediaPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVmediaPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVmediaPolicies gets a list of vmedia policy instances
*/
func (a *Client) GetVmediaPolicies(params *GetVmediaPoliciesParams) (*GetVmediaPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVmediaPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVmediaPolicies",
		Method:             "GET",
		PathPattern:        "/vmedia/Policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVmediaPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVmediaPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVmediaPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVmediaPoliciesMoid gets a specific instance of vmedia policy
*/
func (a *Client) GetVmediaPoliciesMoid(params *GetVmediaPoliciesMoidParams) (*GetVmediaPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVmediaPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVmediaPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/vmedia/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVmediaPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVmediaPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVmediaPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchVmediaPoliciesMoid updates an instance of vmedia policy
*/
func (a *Client) PatchVmediaPoliciesMoid(params *PatchVmediaPoliciesMoidParams) (*PatchVmediaPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchVmediaPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchVmediaPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/vmedia/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVmediaPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchVmediaPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchVmediaPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVmediaPolicies creates an instance of vmedia policy
*/
func (a *Client) PostVmediaPolicies(params *PostVmediaPoliciesParams) (*PostVmediaPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVmediaPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVmediaPolicies",
		Method:             "POST",
		PathPattern:        "/vmedia/Policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVmediaPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVmediaPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVmediaPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostVmediaPoliciesMoid updates an instance of vmedia policy
*/
func (a *Client) PostVmediaPoliciesMoid(params *PostVmediaPoliciesMoidParams) (*PostVmediaPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVmediaPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVmediaPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/vmedia/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVmediaPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVmediaPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVmediaPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
