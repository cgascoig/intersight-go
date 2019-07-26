// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_user_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iam end point user policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam end point user policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteIamEndPointUserPoliciesMoid deletes an instance of iam end point user policy
*/
func (a *Client) DeleteIamEndPointUserPoliciesMoid(params *DeleteIamEndPointUserPoliciesMoidParams) (*DeleteIamEndPointUserPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIamEndPointUserPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIamEndPointUserPoliciesMoid",
		Method:             "DELETE",
		PathPattern:        "/iam/EndPointUserPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIamEndPointUserPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIamEndPointUserPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIamEndPointUserPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamEndPointUserPolicies gets a list of iam end point user policy instances
*/
func (a *Client) GetIamEndPointUserPolicies(params *GetIamEndPointUserPoliciesParams) (*GetIamEndPointUserPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamEndPointUserPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamEndPointUserPolicies",
		Method:             "GET",
		PathPattern:        "/iam/EndPointUserPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamEndPointUserPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamEndPointUserPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamEndPointUserPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamEndPointUserPoliciesMoid gets a specific instance of iam end point user policy
*/
func (a *Client) GetIamEndPointUserPoliciesMoid(params *GetIamEndPointUserPoliciesMoidParams) (*GetIamEndPointUserPoliciesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamEndPointUserPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamEndPointUserPoliciesMoid",
		Method:             "GET",
		PathPattern:        "/iam/EndPointUserPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamEndPointUserPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamEndPointUserPoliciesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamEndPointUserPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchIamEndPointUserPoliciesMoid updates an instance of iam end point user policy
*/
func (a *Client) PatchIamEndPointUserPoliciesMoid(params *PatchIamEndPointUserPoliciesMoidParams) (*PatchIamEndPointUserPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchIamEndPointUserPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchIamEndPointUserPoliciesMoid",
		Method:             "PATCH",
		PathPattern:        "/iam/EndPointUserPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchIamEndPointUserPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchIamEndPointUserPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchIamEndPointUserPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamEndPointUserPolicies creates an instance of iam end point user policy
*/
func (a *Client) PostIamEndPointUserPolicies(params *PostIamEndPointUserPoliciesParams) (*PostIamEndPointUserPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamEndPointUserPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamEndPointUserPolicies",
		Method:             "POST",
		PathPattern:        "/iam/EndPointUserPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamEndPointUserPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamEndPointUserPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamEndPointUserPoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamEndPointUserPoliciesMoid updates an instance of iam end point user policy
*/
func (a *Client) PostIamEndPointUserPoliciesMoid(params *PostIamEndPointUserPoliciesMoidParams) (*PostIamEndPointUserPoliciesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamEndPointUserPoliciesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamEndPointUserPoliciesMoid",
		Method:             "POST",
		PathPattern:        "/iam/EndPointUserPolicies/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamEndPointUserPoliciesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamEndPointUserPoliciesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamEndPointUserPoliciesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}