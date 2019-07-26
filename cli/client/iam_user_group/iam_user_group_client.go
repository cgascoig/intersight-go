// Code generated by go-swagger; DO NOT EDIT.

package iam_user_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iam user group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam user group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteIamUserGroupsMoid deletes an instance of iam user group
*/
func (a *Client) DeleteIamUserGroupsMoid(params *DeleteIamUserGroupsMoidParams) (*DeleteIamUserGroupsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIamUserGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIamUserGroupsMoid",
		Method:             "DELETE",
		PathPattern:        "/iam/UserGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIamUserGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIamUserGroupsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIamUserGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamUserGroups gets a list of iam user group instances
*/
func (a *Client) GetIamUserGroups(params *GetIamUserGroupsParams) (*GetIamUserGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamUserGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamUserGroups",
		Method:             "GET",
		PathPattern:        "/iam/UserGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamUserGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamUserGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamUserGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamUserGroupsMoid gets a specific instance of iam user group
*/
func (a *Client) GetIamUserGroupsMoid(params *GetIamUserGroupsMoidParams) (*GetIamUserGroupsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamUserGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamUserGroupsMoid",
		Method:             "GET",
		PathPattern:        "/iam/UserGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamUserGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamUserGroupsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamUserGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchIamUserGroupsMoid updates an instance of iam user group
*/
func (a *Client) PatchIamUserGroupsMoid(params *PatchIamUserGroupsMoidParams) (*PatchIamUserGroupsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchIamUserGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchIamUserGroupsMoid",
		Method:             "PATCH",
		PathPattern:        "/iam/UserGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchIamUserGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchIamUserGroupsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchIamUserGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamUserGroups creates an instance of iam user group
*/
func (a *Client) PostIamUserGroups(params *PostIamUserGroupsParams) (*PostIamUserGroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamUserGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamUserGroups",
		Method:             "POST",
		PathPattern:        "/iam/UserGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamUserGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamUserGroupsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamUserGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamUserGroupsMoid updates an instance of iam user group
*/
func (a *Client) PostIamUserGroupsMoid(params *PostIamUserGroupsMoidParams) (*PostIamUserGroupsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamUserGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamUserGroupsMoid",
		Method:             "POST",
		PathPattern:        "/iam/UserGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamUserGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamUserGroupsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamUserGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
