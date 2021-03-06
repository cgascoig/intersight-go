// Code generated by go-swagger; DO NOT EDIT.

package search_search_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search search item API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search search item API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSearchSearchItems gets a list of search search item instances
*/
func (a *Client) GetSearchSearchItems(params *GetSearchSearchItemsParams) (*GetSearchSearchItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchSearchItemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchSearchItems",
		Method:             "GET",
		PathPattern:        "/search/SearchItems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchSearchItemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSearchSearchItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSearchSearchItemsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSearchSearchItemsMoid gets a specific instance of search search item
*/
func (a *Client) GetSearchSearchItemsMoid(params *GetSearchSearchItemsMoidParams) (*GetSearchSearchItemsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchSearchItemsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchSearchItemsMoid",
		Method:             "GET",
		PathPattern:        "/search/SearchItems/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchSearchItemsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSearchSearchItemsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSearchSearchItemsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
