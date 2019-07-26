// Code generated by go-swagger; DO NOT EDIT.

package hcl_driver_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hcl driver image API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hcl driver image API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHclDriverImages gets a list of hcl driver image instances
*/
func (a *Client) GetHclDriverImages(params *GetHclDriverImagesParams) (*GetHclDriverImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHclDriverImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHclDriverImages",
		Method:             "GET",
		PathPattern:        "/hcl/DriverImages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHclDriverImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHclDriverImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHclDriverImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHclDriverImagesMoid gets a specific instance of hcl driver image
*/
func (a *Client) GetHclDriverImagesMoid(params *GetHclDriverImagesMoidParams) (*GetHclDriverImagesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHclDriverImagesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHclDriverImagesMoid",
		Method:             "GET",
		PathPattern:        "/hcl/DriverImages/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHclDriverImagesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHclDriverImagesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHclDriverImagesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}