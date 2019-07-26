// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_node_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hyperflex node profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hyperflex node profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteHyperflexNodeProfilesMoid deletes an instance of hyperflex node profile
*/
func (a *Client) DeleteHyperflexNodeProfilesMoid(params *DeleteHyperflexNodeProfilesMoidParams) (*DeleteHyperflexNodeProfilesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHyperflexNodeProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteHyperflexNodeProfilesMoid",
		Method:             "DELETE",
		PathPattern:        "/hyperflex/NodeProfiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHyperflexNodeProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHyperflexNodeProfilesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteHyperflexNodeProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHyperflexNodeProfiles gets a list of hyperflex node profile instances
*/
func (a *Client) GetHyperflexNodeProfiles(params *GetHyperflexNodeProfilesParams) (*GetHyperflexNodeProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHyperflexNodeProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHyperflexNodeProfiles",
		Method:             "GET",
		PathPattern:        "/hyperflex/NodeProfiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHyperflexNodeProfilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHyperflexNodeProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHyperflexNodeProfilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHyperflexNodeProfilesMoid gets a specific instance of hyperflex node profile
*/
func (a *Client) GetHyperflexNodeProfilesMoid(params *GetHyperflexNodeProfilesMoidParams) (*GetHyperflexNodeProfilesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHyperflexNodeProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHyperflexNodeProfilesMoid",
		Method:             "GET",
		PathPattern:        "/hyperflex/NodeProfiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHyperflexNodeProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHyperflexNodeProfilesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHyperflexNodeProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchHyperflexNodeProfilesMoid updates an instance of hyperflex node profile
*/
func (a *Client) PatchHyperflexNodeProfilesMoid(params *PatchHyperflexNodeProfilesMoidParams) (*PatchHyperflexNodeProfilesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchHyperflexNodeProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchHyperflexNodeProfilesMoid",
		Method:             "PATCH",
		PathPattern:        "/hyperflex/NodeProfiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchHyperflexNodeProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchHyperflexNodeProfilesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchHyperflexNodeProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHyperflexNodeProfiles creates an instance of hyperflex node profile
*/
func (a *Client) PostHyperflexNodeProfiles(params *PostHyperflexNodeProfilesParams) (*PostHyperflexNodeProfilesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHyperflexNodeProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHyperflexNodeProfiles",
		Method:             "POST",
		PathPattern:        "/hyperflex/NodeProfiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHyperflexNodeProfilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHyperflexNodeProfilesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHyperflexNodeProfilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHyperflexNodeProfilesMoid updates an instance of hyperflex node profile
*/
func (a *Client) PostHyperflexNodeProfilesMoid(params *PostHyperflexNodeProfilesMoidParams) (*PostHyperflexNodeProfilesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHyperflexNodeProfilesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHyperflexNodeProfilesMoid",
		Method:             "POST",
		PathPattern:        "/hyperflex/NodeProfiles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHyperflexNodeProfilesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHyperflexNodeProfilesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHyperflexNodeProfilesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}