// Code generated by go-swagger; DO NOT EDIT.

package cond_hcl_status_detail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cond hcl status detail API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cond hcl status detail API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCondHclStatusDetails gets a list of cond hcl status detail instances
*/
func (a *Client) GetCondHclStatusDetails(params *GetCondHclStatusDetailsParams) (*GetCondHclStatusDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCondHclStatusDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCondHclStatusDetails",
		Method:             "GET",
		PathPattern:        "/cond/HclStatusDetails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCondHclStatusDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCondHclStatusDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCondHclStatusDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCondHclStatusDetailsMoid gets a specific instance of cond hcl status detail
*/
func (a *Client) GetCondHclStatusDetailsMoid(params *GetCondHclStatusDetailsMoidParams) (*GetCondHclStatusDetailsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCondHclStatusDetailsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCondHclStatusDetailsMoid",
		Method:             "GET",
		PathPattern:        "/cond/HclStatusDetails/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCondHclStatusDetailsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCondHclStatusDetailsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCondHclStatusDetailsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}