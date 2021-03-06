// Code generated by go-swagger; DO NOT EDIT.

package cond_hcl_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cond hcl status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cond hcl status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCondHclStatuses gets a list of cond hcl status instances
*/
func (a *Client) GetCondHclStatuses(params *GetCondHclStatusesParams) (*GetCondHclStatusesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCondHclStatusesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCondHclStatuses",
		Method:             "GET",
		PathPattern:        "/cond/HclStatuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCondHclStatusesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCondHclStatusesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCondHclStatusesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCondHclStatusesMoid gets a specific instance of cond hcl status
*/
func (a *Client) GetCondHclStatusesMoid(params *GetCondHclStatusesMoidParams) (*GetCondHclStatusesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCondHclStatusesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCondHclStatusesMoid",
		Method:             "GET",
		PathPattern:        "/cond/HclStatuses/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCondHclStatusesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCondHclStatusesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCondHclStatusesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
