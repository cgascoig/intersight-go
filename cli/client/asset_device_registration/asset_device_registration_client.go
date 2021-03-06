// Code generated by go-swagger; DO NOT EDIT.

package asset_device_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new asset device registration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for asset device registration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAssetDeviceRegistrationsMoid deletes the resource representing the device connector all associated r e s t resources will be deleted in particular inventory and operational data associated with this device will be deleted
*/
func (a *Client) DeleteAssetDeviceRegistrationsMoid(params *DeleteAssetDeviceRegistrationsMoidParams) (*DeleteAssetDeviceRegistrationsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAssetDeviceRegistrationsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAssetDeviceRegistrationsMoid",
		Method:             "DELETE",
		PathPattern:        "/asset/DeviceRegistrations/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAssetDeviceRegistrationsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAssetDeviceRegistrationsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAssetDeviceRegistrationsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAssetDeviceRegistrations gets a list of asset device registration instances
*/
func (a *Client) GetAssetDeviceRegistrations(params *GetAssetDeviceRegistrationsParams) (*GetAssetDeviceRegistrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssetDeviceRegistrationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAssetDeviceRegistrations",
		Method:             "GET",
		PathPattern:        "/asset/DeviceRegistrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAssetDeviceRegistrationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssetDeviceRegistrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAssetDeviceRegistrationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAssetDeviceRegistrationsMoid gets a specific instance of asset device registration
*/
func (a *Client) GetAssetDeviceRegistrationsMoid(params *GetAssetDeviceRegistrationsMoidParams) (*GetAssetDeviceRegistrationsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssetDeviceRegistrationsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAssetDeviceRegistrationsMoid",
		Method:             "GET",
		PathPattern:        "/asset/DeviceRegistrations/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAssetDeviceRegistrationsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssetDeviceRegistrationsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAssetDeviceRegistrationsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchAssetDeviceRegistrationsMoid updates the resource representing the device connector for example this can be used to annotate the device connector resource with user specified tags
*/
func (a *Client) PatchAssetDeviceRegistrationsMoid(params *PatchAssetDeviceRegistrationsMoidParams) (*PatchAssetDeviceRegistrationsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAssetDeviceRegistrationsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAssetDeviceRegistrationsMoid",
		Method:             "PATCH",
		PathPattern:        "/asset/DeviceRegistrations/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAssetDeviceRegistrationsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAssetDeviceRegistrationsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchAssetDeviceRegistrationsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostAssetDeviceRegistrationsMoid updates the resource representing the device connector for example this can be used to annotate the device connector resource with user specified tags
*/
func (a *Client) PostAssetDeviceRegistrationsMoid(params *PostAssetDeviceRegistrationsMoidParams) (*PostAssetDeviceRegistrationsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAssetDeviceRegistrationsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAssetDeviceRegistrationsMoid",
		Method:             "POST",
		PathPattern:        "/asset/DeviceRegistrations/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAssetDeviceRegistrationsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAssetDeviceRegistrationsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostAssetDeviceRegistrationsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
