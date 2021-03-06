// Code generated by go-swagger; DO NOT EDIT.

package storage_physical_disk_extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage physical disk extension API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage physical disk extension API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStoragePhysicalDiskExtensions gets a list of storage physical disk extension instances
*/
func (a *Client) GetStoragePhysicalDiskExtensions(params *GetStoragePhysicalDiskExtensionsParams) (*GetStoragePhysicalDiskExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePhysicalDiskExtensionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePhysicalDiskExtensions",
		Method:             "GET",
		PathPattern:        "/storage/PhysicalDiskExtensions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePhysicalDiskExtensionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePhysicalDiskExtensionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePhysicalDiskExtensionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStoragePhysicalDiskExtensionsMoid gets a specific instance of storage physical disk extension
*/
func (a *Client) GetStoragePhysicalDiskExtensionsMoid(params *GetStoragePhysicalDiskExtensionsMoidParams) (*GetStoragePhysicalDiskExtensionsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoragePhysicalDiskExtensionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoragePhysicalDiskExtensionsMoid",
		Method:             "GET",
		PathPattern:        "/storage/PhysicalDiskExtensions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoragePhysicalDiskExtensionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoragePhysicalDiskExtensionsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoragePhysicalDiskExtensionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchStoragePhysicalDiskExtensionsMoid updates an instance of storage physical disk extension
*/
func (a *Client) PatchStoragePhysicalDiskExtensionsMoid(params *PatchStoragePhysicalDiskExtensionsMoidParams) (*PatchStoragePhysicalDiskExtensionsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchStoragePhysicalDiskExtensionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchStoragePhysicalDiskExtensionsMoid",
		Method:             "PATCH",
		PathPattern:        "/storage/PhysicalDiskExtensions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchStoragePhysicalDiskExtensionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchStoragePhysicalDiskExtensionsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchStoragePhysicalDiskExtensionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostStoragePhysicalDiskExtensionsMoid updates an instance of storage physical disk extension
*/
func (a *Client) PostStoragePhysicalDiskExtensionsMoid(params *PostStoragePhysicalDiskExtensionsMoidParams) (*PostStoragePhysicalDiskExtensionsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStoragePhysicalDiskExtensionsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStoragePhysicalDiskExtensionsMoid",
		Method:             "POST",
		PathPattern:        "/storage/PhysicalDiskExtensions/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStoragePhysicalDiskExtensionsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostStoragePhysicalDiskExtensionsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostStoragePhysicalDiskExtensionsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
