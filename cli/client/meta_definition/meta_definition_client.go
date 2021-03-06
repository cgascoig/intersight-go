// Code generated by go-swagger; DO NOT EDIT.

package meta_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new meta definition API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for meta definition API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteMetaDefinitionsMoid deletes an instance of meta definition
*/
func (a *Client) DeleteMetaDefinitionsMoid(params *DeleteMetaDefinitionsMoidParams) (*DeleteMetaDefinitionsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMetaDefinitionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMetaDefinitionsMoid",
		Method:             "DELETE",
		PathPattern:        "/meta/Definitions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMetaDefinitionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteMetaDefinitionsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteMetaDefinitionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetMetaDefinitions gets a list of meta definition instances
*/
func (a *Client) GetMetaDefinitions(params *GetMetaDefinitionsParams) (*GetMetaDefinitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetaDefinitionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetaDefinitions",
		Method:             "GET",
		PathPattern:        "/meta/Definitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetaDefinitionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMetaDefinitionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMetaDefinitionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetMetaDefinitionsMoid gets a specific instance of meta definition
*/
func (a *Client) GetMetaDefinitionsMoid(params *GetMetaDefinitionsMoidParams) (*GetMetaDefinitionsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetaDefinitionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetaDefinitionsMoid",
		Method:             "GET",
		PathPattern:        "/meta/Definitions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetaDefinitionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMetaDefinitionsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMetaDefinitionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
