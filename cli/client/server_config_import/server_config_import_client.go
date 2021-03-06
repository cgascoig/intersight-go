// Code generated by go-swagger; DO NOT EDIT.

package server_config_import

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new server config import API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server config import API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetServerConfigImports gets a list of server config import instances
*/
func (a *Client) GetServerConfigImports(params *GetServerConfigImportsParams) (*GetServerConfigImportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerConfigImportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServerConfigImports",
		Method:             "GET",
		PathPattern:        "/server/ConfigImports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerConfigImportsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerConfigImportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerConfigImportsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServerConfigImportsMoid gets a specific instance of server config import
*/
func (a *Client) GetServerConfigImportsMoid(params *GetServerConfigImportsMoidParams) (*GetServerConfigImportsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerConfigImportsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServerConfigImportsMoid",
		Method:             "GET",
		PathPattern:        "/server/ConfigImports/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerConfigImportsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerConfigImportsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerConfigImportsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostServerConfigImports creates an instance of server config import
*/
func (a *Client) PostServerConfigImports(params *PostServerConfigImportsParams) (*PostServerConfigImportsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostServerConfigImportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostServerConfigImports",
		Method:             "POST",
		PathPattern:        "/server/ConfigImports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostServerConfigImportsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostServerConfigImportsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostServerConfigImportsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
