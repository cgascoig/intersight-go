// Code generated by go-swagger; DO NOT EDIT.

package compute_physical_summary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new compute physical summary API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute physical summary API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetComputePhysicalSummaries gets a list of compute physical summary instances
*/
func (a *Client) GetComputePhysicalSummaries(params *GetComputePhysicalSummariesParams) (*GetComputePhysicalSummariesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputePhysicalSummariesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetComputePhysicalSummaries",
		Method:             "GET",
		PathPattern:        "/compute/PhysicalSummaries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputePhysicalSummariesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComputePhysicalSummariesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetComputePhysicalSummariesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetComputePhysicalSummariesMoid gets a specific instance of compute physical summary
*/
func (a *Client) GetComputePhysicalSummariesMoid(params *GetComputePhysicalSummariesMoidParams) (*GetComputePhysicalSummariesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputePhysicalSummariesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetComputePhysicalSummariesMoid",
		Method:             "GET",
		PathPattern:        "/compute/PhysicalSummaries/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputePhysicalSummariesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComputePhysicalSummariesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetComputePhysicalSummariesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}