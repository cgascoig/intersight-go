// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_array

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage pure array API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage pure array API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStoragePureArrays gets a list of storage pure array instances
*/
func (a *Client) GetStoragePureArrays(params *GetStoragePureArraysParams) (*GetStoragePureArraysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePureArraysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePureArrays",
		Method:             "GET",
		PathPattern:        "/storage/PureArrays",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePureArraysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePureArraysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePureArraysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStoragePureArraysMoid gets a specific instance of storage pure array
*/
func (a *Client) GetStoragePureArraysMoid(params *GetStoragePureArraysMoidParams) (*GetStoragePureArraysMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePureArraysMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePureArraysMoid",
		Method:             "GET",
		PathPattern:        "/storage/PureArrays/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePureArraysMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePureArraysMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePureArraysMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchStoragePureArraysMoid updates an instance of storage pure array
*/
func (a *Client) PatchStoragePureArraysMoid(params *PatchStoragePureArraysMoidParams) (*PatchStoragePureArraysMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchStoragePureArraysMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchStoragePureArraysMoid",
		Method:             "PATCH",
		PathPattern:        "/storage/PureArrays/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchStoragePureArraysMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchStoragePureArraysMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchStoragePureArraysMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostStoragePureArraysMoid updates an instance of storage pure array
*/
func (a *Client) PostStoragePureArraysMoid(params *PostStoragePureArraysMoidParams) (*PostStoragePureArraysMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStoragePureArraysMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStoragePureArraysMoid",
		Method:             "POST",
		PathPattern:        "/storage/PureArrays/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStoragePureArraysMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostStoragePureArraysMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostStoragePureArraysMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}