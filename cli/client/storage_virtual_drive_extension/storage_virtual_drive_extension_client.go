// Code generated by go-swagger; DO NOT EDIT.

package storage_virtual_drive_extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage virtual drive extension API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage virtual drive extension API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStorageVirtualDriveExtensions gets a list of storage virtual drive extension instances
*/
func (a *Client) GetStorageVirtualDriveExtensions(params *GetStorageVirtualDriveExtensionsParams) (*GetStorageVirtualDriveExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageVirtualDriveExtensionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStorageVirtualDriveExtensions",
		Method:             "GET",
		PathPattern:        "/storage/VirtualDriveExtensions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageVirtualDriveExtensionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStorageVirtualDriveExtensionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStorageVirtualDriveExtensionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStorageVirtualDriveExtensionsMoid gets a specific instance of storage virtual drive extension
*/
func (a *Client) GetStorageVirtualDriveExtensionsMoid(params *GetStorageVirtualDriveExtensionsMoidParams) (*GetStorageVirtualDriveExtensionsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageVirtualDriveExtensionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStorageVirtualDriveExtensionsMoid",
		Method:             "GET",
		PathPattern:        "/storage/VirtualDriveExtensions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageVirtualDriveExtensionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStorageVirtualDriveExtensionsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStorageVirtualDriveExtensionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchStorageVirtualDriveExtensionsMoid updates an instance of storage virtual drive extension
*/
func (a *Client) PatchStorageVirtualDriveExtensionsMoid(params *PatchStorageVirtualDriveExtensionsMoidParams) (*PatchStorageVirtualDriveExtensionsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchStorageVirtualDriveExtensionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchStorageVirtualDriveExtensionsMoid",
		Method:             "PATCH",
		PathPattern:        "/storage/VirtualDriveExtensions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchStorageVirtualDriveExtensionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchStorageVirtualDriveExtensionsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchStorageVirtualDriveExtensionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostStorageVirtualDriveExtensionsMoid updates an instance of storage virtual drive extension
*/
func (a *Client) PostStorageVirtualDriveExtensionsMoid(params *PostStorageVirtualDriveExtensionsMoidParams) (*PostStorageVirtualDriveExtensionsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStorageVirtualDriveExtensionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStorageVirtualDriveExtensionsMoid",
		Method:             "POST",
		PathPattern:        "/storage/VirtualDriveExtensions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStorageVirtualDriveExtensionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostStorageVirtualDriveExtensionsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostStorageVirtualDriveExtensionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
