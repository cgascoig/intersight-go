// Code generated by go-swagger; DO NOT EDIT.

package ipmioverlan_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ipmioverlan policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipmioverlan policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteIpmioverlanPoliciesMoid deletes an instance of ipmioverlan policy
*/
func (a *Client) DeleteIpmioverlanPoliciesMoid(params *DeleteIpmioverlanPoliciesMoidParams) (*DeleteIpmioverlanPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIpmioverlanPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIpmioverlanPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/ipmioverlan/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIpmioverlanPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIpmioverlanPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIpmioverlanPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIpmioverlanPolicies gets a list of ipmioverlan policy instances
*/
func (a *Client) GetIpmioverlanPolicies(params *GetIpmioverlanPoliciesParams) (*GetIpmioverlanPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIpmioverlanPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIpmioverlanPolicies",
		Method:             "GET",
		PathPattern:        "/ipmioverlan/Policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIpmioverlanPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIpmioverlanPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIpmioverlanPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIpmioverlanPoliciesMoid gets a specific instance of ipmioverlan policy
*/
func (a *Client) GetIpmioverlanPoliciesMoid(params *GetIpmioverlanPoliciesMoidParams) (*GetIpmioverlanPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIpmioverlanPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIpmioverlanPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/ipmioverlan/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIpmioverlanPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIpmioverlanPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIpmioverlanPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchIpmioverlanPoliciesMoid updates an instance of ipmioverlan policy
*/
func (a *Client) PatchIpmioverlanPoliciesMoid(params *PatchIpmioverlanPoliciesMoidParams) (*PatchIpmioverlanPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchIpmioverlanPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchIpmioverlanPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/ipmioverlan/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchIpmioverlanPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchIpmioverlanPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchIpmioverlanPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIpmioverlanPolicies creates an instance of ipmioverlan policy
*/
func (a *Client) PostIpmioverlanPolicies(params *PostIpmioverlanPoliciesParams) (*PostIpmioverlanPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIpmioverlanPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIpmioverlanPolicies",
		Method:             "POST",
		PathPattern:        "/ipmioverlan/Policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIpmioverlanPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIpmioverlanPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIpmioverlanPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIpmioverlanPoliciesMoid updates an instance of ipmioverlan policy
*/
func (a *Client) PostIpmioverlanPoliciesMoid(params *PostIpmioverlanPoliciesMoidParams) (*PostIpmioverlanPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIpmioverlanPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIpmioverlanPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/ipmioverlan/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIpmioverlanPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIpmioverlanPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIpmioverlanPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
