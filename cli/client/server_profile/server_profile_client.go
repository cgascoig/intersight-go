// Code generated by go-swagger; DO NOT EDIT.

package server_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new server profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteServerProfilesMoid deletes an instance of server profile
*/
func (a *Client) DeleteServerProfilesMoid(params *DeleteServerProfilesMoidParams) (*DeleteServerProfilesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServerProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteServerProfilesMoid",
		Method:             "DELETE",
		PathPattern:        "/server/Profiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServerProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServerProfilesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteServerProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServerProfiles gets a list of server profile instances
*/
func (a *Client) GetServerProfiles(params *GetServerProfilesParams) (*GetServerProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServerProfiles",
		Method:             "GET",
		PathPattern:        "/server/Profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerProfilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerProfilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServerProfilesMoid gets a specific instance of server profile
*/
func (a *Client) GetServerProfilesMoid(params *GetServerProfilesMoidParams) (*GetServerProfilesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServerProfilesMoid",
		Method:             "GET",
		PathPattern:        "/server/Profiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerProfilesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchServerProfilesMoid updates an instance of server profile
*/
func (a *Client) PatchServerProfilesMoid(params *PatchServerProfilesMoidParams) (*PatchServerProfilesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchServerProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchServerProfilesMoid",
		Method:             "PATCH",
		PathPattern:        "/server/Profiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchServerProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchServerProfilesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchServerProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostServerProfiles creates an instance of server profile
*/
func (a *Client) PostServerProfiles(params *PostServerProfilesParams) (*PostServerProfilesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostServerProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostServerProfiles",
		Method:             "POST",
		PathPattern:        "/server/Profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostServerProfilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostServerProfilesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostServerProfilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostServerProfilesMoid updates an instance of server profile
*/
func (a *Client) PostServerProfilesMoid(params *PostServerProfilesMoidParams) (*PostServerProfilesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostServerProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostServerProfilesMoid",
		Method:             "POST",
		PathPattern:        "/server/Profiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostServerProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostServerProfilesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostServerProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}