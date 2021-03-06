// Code generated by go-swagger; DO NOT EDIT.

package oauth_oauth_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new oauth oauth user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for oauth oauth user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteOauthOauthUsersMoid deletes an instance of oauth oauth user
*/
func (a *Client) DeleteOauthOauthUsersMoid(params *DeleteOauthOauthUsersMoidParams) (*DeleteOauthOauthUsersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOauthOauthUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOauthOauthUsersMoid",
		Method:             "DELETE",
		PathPattern:        "/oauth/OauthUsers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOauthOauthUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOauthOauthUsersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOauthOauthUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOauthOauthUsers gets a list of oauth oauth user instances
*/
func (a *Client) GetOauthOauthUsers(params *GetOauthOauthUsersParams) (*GetOauthOauthUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOauthOauthUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOauthOauthUsers",
		Method:             "GET",
		PathPattern:        "/oauth/OauthUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOauthOauthUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOauthOauthUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOauthOauthUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOauthOauthUsersMoid gets a specific instance of oauth oauth user
*/
func (a *Client) GetOauthOauthUsersMoid(params *GetOauthOauthUsersMoidParams) (*GetOauthOauthUsersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOauthOauthUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOauthOauthUsersMoid",
		Method:             "GET",
		PathPattern:        "/oauth/OauthUsers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOauthOauthUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOauthOauthUsersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOauthOauthUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchOauthOauthUsersMoid updates an instance of oauth oauth user
*/
func (a *Client) PatchOauthOauthUsersMoid(params *PatchOauthOauthUsersMoidParams) (*PatchOauthOauthUsersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOauthOauthUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchOauthOauthUsersMoid",
		Method:             "PATCH",
		PathPattern:        "/oauth/OauthUsers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchOauthOauthUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchOauthOauthUsersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchOauthOauthUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostOauthOauthUsers creates an instance of oauth oauth user
*/
func (a *Client) PostOauthOauthUsers(params *PostOauthOauthUsersParams) (*PostOauthOauthUsersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthOauthUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthOauthUsers",
		Method:             "POST",
		PathPattern:        "/oauth/OauthUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthOauthUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOauthOauthUsersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostOauthOauthUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostOauthOauthUsersMoid updates an instance of oauth oauth user
*/
func (a *Client) PostOauthOauthUsersMoid(params *PostOauthOauthUsersMoidParams) (*PostOauthOauthUsersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthOauthUsersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthOauthUsersMoid",
		Method:             "POST",
		PathPattern:        "/oauth/OauthUsers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthOauthUsersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOauthOauthUsersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostOauthOauthUsersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
