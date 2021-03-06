// Code generated by go-swagger; DO NOT EDIT.

package iam_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iam user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteIamUsersMoid deletes an instance of iam user
*/
func (a *Client) DeleteIamUsersMoid(params *DeleteIamUsersMoidParams) (*DeleteIamUsersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIamUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIamUsersMoid",
		Method:             "DELETE",
		PathPattern:        "/iam/Users/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIamUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIamUsersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIamUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamUsers gets a list of iam user instances
*/
func (a *Client) GetIamUsers(params *GetIamUsersParams) (*GetIamUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamUsers",
		Method:             "GET",
		PathPattern:        "/iam/Users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamUsersMoid gets a specific instance of iam user
*/
func (a *Client) GetIamUsersMoid(params *GetIamUsersMoidParams) (*GetIamUsersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamUsersMoid",
		Method:             "GET",
		PathPattern:        "/iam/Users/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamUsersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchIamUsersMoid updates an instance of iam user
*/
func (a *Client) PatchIamUsersMoid(params *PatchIamUsersMoidParams) (*PatchIamUsersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchIamUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchIamUsersMoid",
		Method:             "PATCH",
		PathPattern:        "/iam/Users/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchIamUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchIamUsersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchIamUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamUsers creates an instance of iam user
*/
func (a *Client) PostIamUsers(params *PostIamUsersParams) (*PostIamUsersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamUsers",
		Method:             "POST",
		PathPattern:        "/iam/Users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamUsersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamUsersMoid updates an instance of iam user
*/
func (a *Client) PostIamUsersMoid(params *PostIamUsersMoidParams) (*PostIamUsersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamUsersMoid",
		Method:             "POST",
		PathPattern:        "/iam/Users/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamUsersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
