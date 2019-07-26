// Code generated by go-swagger; DO NOT EDIT.

package ntp_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ntp policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ntp policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteNtpPoliciesMoid deletes an instance of ntp policy
*/
func (a *Client) DeleteNtpPoliciesMoid(params *DeleteNtpPoliciesMoidParams) (*DeleteNtpPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNtpPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNtpPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/ntp/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNtpPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNtpPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNtpPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNtpPolicies gets a list of ntp policy instances
*/
func (a *Client) GetNtpPolicies(params *GetNtpPoliciesParams) (*GetNtpPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNtpPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNtpPolicies",
		Method:             "GET",
		PathPattern:        "/ntp/Policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNtpPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNtpPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNtpPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNtpPoliciesMoid gets a specific instance of ntp policy
*/
func (a *Client) GetNtpPoliciesMoid(params *GetNtpPoliciesMoidParams) (*GetNtpPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNtpPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNtpPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/ntp/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNtpPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNtpPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNtpPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchNtpPoliciesMoid updates an instance of ntp policy
*/
func (a *Client) PatchNtpPoliciesMoid(params *PatchNtpPoliciesMoidParams) (*PatchNtpPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNtpPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchNtpPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/ntp/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchNtpPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchNtpPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchNtpPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostNtpPolicies creates an instance of ntp policy
*/
func (a *Client) PostNtpPolicies(params *PostNtpPoliciesParams) (*PostNtpPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNtpPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNtpPolicies",
		Method:             "POST",
		PathPattern:        "/ntp/Policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostNtpPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNtpPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNtpPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostNtpPoliciesMoid updates an instance of ntp policy
*/
func (a *Client) PostNtpPoliciesMoid(params *PostNtpPoliciesMoidParams) (*PostNtpPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNtpPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNtpPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/ntp/Policies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostNtpPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNtpPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNtpPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}