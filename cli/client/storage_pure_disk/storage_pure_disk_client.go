// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage pure disk API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage pure disk API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStoragePureDisks gets a list of storage pure disk instances
*/
func (a *Client) GetStoragePureDisks(params *GetStoragePureDisksParams) (*GetStoragePureDisksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePureDisksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePureDisks",
		Method:             "GET",
		PathPattern:        "/storage/PureDisks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePureDisksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePureDisksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePureDisksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStoragePureDisksMoid gets a specific instance of storage pure disk
*/
func (a *Client) GetStoragePureDisksMoid(params *GetStoragePureDisksMoidParams) (*GetStoragePureDisksMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePureDisksMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePureDisksMoid",
		Method:             "GET",
		PathPattern:        "/storage/PureDisks/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePureDisksMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePureDisksMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePureDisksMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
