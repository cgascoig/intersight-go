// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_protection_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage pure protection group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage pure protection group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStoragePureProtectionGroups gets a list of storage pure protection group instances
*/
func (a *Client) GetStoragePureProtectionGroups(params *GetStoragePureProtectionGroupsParams) (*GetStoragePureProtectionGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePureProtectionGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePureProtectionGroups",
		Method:             "GET",
		PathPattern:        "/storage/PureProtectionGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePureProtectionGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePureProtectionGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePureProtectionGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStoragePureProtectionGroupsMoid gets a specific instance of storage pure protection group
*/
func (a *Client) GetStoragePureProtectionGroupsMoid(params *GetStoragePureProtectionGroupsMoidParams) (*GetStoragePureProtectionGroupsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePureProtectionGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePureProtectionGroupsMoid",
		Method:             "GET",
		PathPattern:        "/storage/PureProtectionGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePureProtectionGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePureProtectionGroupsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePureProtectionGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
