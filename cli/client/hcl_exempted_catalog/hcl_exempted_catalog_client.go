// Code generated by go-swagger; DO NOT EDIT.

package hcl_exempted_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hcl exempted catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hcl exempted catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHclExemptedCatalogs gets a list of hcl exempted catalog instances
*/
func (a *Client) GetHclExemptedCatalogs(params *GetHclExemptedCatalogsParams) (*GetHclExemptedCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHclExemptedCatalogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHclExemptedCatalogs",
		Method:             "GET",
		PathPattern:        "/hcl/ExemptedCatalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHclExemptedCatalogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHclExemptedCatalogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHclExemptedCatalogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHclExemptedCatalogsMoid gets a specific instance of hcl exempted catalog
*/
func (a *Client) GetHclExemptedCatalogsMoid(params *GetHclExemptedCatalogsMoidParams) (*GetHclExemptedCatalogsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHclExemptedCatalogsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHclExemptedCatalogsMoid",
		Method:             "GET",
		PathPattern:        "/hcl/ExemptedCatalogs/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHclExemptedCatalogsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHclExemptedCatalogsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHclExemptedCatalogsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
