// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hyperflex cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hyperflex cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHyperflexClusters gets a list of hyperflex cluster instances
*/
func (a *Client) GetHyperflexClusters(params *GetHyperflexClustersParams) (*GetHyperflexClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHyperflexClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHyperflexClusters",
		Method:             "GET",
		PathPattern:        "/hyperflex/Clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHyperflexClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHyperflexClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHyperflexClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHyperflexClustersMoid gets a specific instance of hyperflex cluster
*/
func (a *Client) GetHyperflexClustersMoid(params *GetHyperflexClustersMoidParams) (*GetHyperflexClustersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHyperflexClustersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHyperflexClustersMoid",
		Method:             "GET",
		PathPattern:        "/hyperflex/Clusters/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHyperflexClustersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHyperflexClustersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHyperflexClustersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchHyperflexClustersMoid updates an instance of hyperflex cluster
*/
func (a *Client) PatchHyperflexClustersMoid(params *PatchHyperflexClustersMoidParams) (*PatchHyperflexClustersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchHyperflexClustersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchHyperflexClustersMoid",
		Method:             "PATCH",
		PathPattern:        "/hyperflex/Clusters/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchHyperflexClustersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchHyperflexClustersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchHyperflexClustersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHyperflexClustersMoid updates an instance of hyperflex cluster
*/
func (a *Client) PostHyperflexClustersMoid(params *PostHyperflexClustersMoidParams) (*PostHyperflexClustersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHyperflexClustersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHyperflexClustersMoid",
		Method:             "POST",
		PathPattern:        "/hyperflex/Clusters/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHyperflexClustersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHyperflexClustersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHyperflexClustersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}