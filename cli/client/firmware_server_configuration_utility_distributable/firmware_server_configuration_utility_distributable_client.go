// Code generated by go-swagger; DO NOT EDIT.

package firmware_server_configuration_utility_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new firmware server configuration utility distributable API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for firmware server configuration utility distributable API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteFirmwareServerConfigurationUtilityDistributablesMoid deletes an instance of firmware server configuration utility distributable
*/
func (a *Client) DeleteFirmwareServerConfigurationUtilityDistributablesMoid(params *DeleteFirmwareServerConfigurationUtilityDistributablesMoidParams) (*DeleteFirmwareServerConfigurationUtilityDistributablesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFirmwareServerConfigurationUtilityDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFirmwareServerConfigurationUtilityDistributablesMoid",
		Method:             "DELETE",
		PathPattern:        "/firmware/ServerConfigurationUtilityDistributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFirmwareServerConfigurationUtilityDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteFirmwareServerConfigurationUtilityDistributablesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFirmwareServerConfigurationUtilityDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFirmwareServerConfigurationUtilityDistributables gets a list of firmware server configuration utility distributable instances
*/
func (a *Client) GetFirmwareServerConfigurationUtilityDistributables(params *GetFirmwareServerConfigurationUtilityDistributablesParams) (*GetFirmwareServerConfigurationUtilityDistributablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFirmwareServerConfigurationUtilityDistributablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFirmwareServerConfigurationUtilityDistributables",
		Method:             "GET",
		PathPattern:        "/firmware/ServerConfigurationUtilityDistributables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFirmwareServerConfigurationUtilityDistributablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFirmwareServerConfigurationUtilityDistributablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFirmwareServerConfigurationUtilityDistributablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFirmwareServerConfigurationUtilityDistributablesMoid gets a specific instance of firmware server configuration utility distributable
*/
func (a *Client) GetFirmwareServerConfigurationUtilityDistributablesMoid(params *GetFirmwareServerConfigurationUtilityDistributablesMoidParams) (*GetFirmwareServerConfigurationUtilityDistributablesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFirmwareServerConfigurationUtilityDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFirmwareServerConfigurationUtilityDistributablesMoid",
		Method:             "GET",
		PathPattern:        "/firmware/ServerConfigurationUtilityDistributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFirmwareServerConfigurationUtilityDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFirmwareServerConfigurationUtilityDistributablesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFirmwareServerConfigurationUtilityDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchFirmwareServerConfigurationUtilityDistributablesMoid updates an instance of firmware server configuration utility distributable
*/
func (a *Client) PatchFirmwareServerConfigurationUtilityDistributablesMoid(params *PatchFirmwareServerConfigurationUtilityDistributablesMoidParams) (*PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchFirmwareServerConfigurationUtilityDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchFirmwareServerConfigurationUtilityDistributablesMoid",
		Method:             "PATCH",
		PathPattern:        "/firmware/ServerConfigurationUtilityDistributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchFirmwareServerConfigurationUtilityDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchFirmwareServerConfigurationUtilityDistributablesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchFirmwareServerConfigurationUtilityDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostFirmwareServerConfigurationUtilityDistributables creates an instance of firmware server configuration utility distributable
*/
func (a *Client) PostFirmwareServerConfigurationUtilityDistributables(params *PostFirmwareServerConfigurationUtilityDistributablesParams) (*PostFirmwareServerConfigurationUtilityDistributablesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFirmwareServerConfigurationUtilityDistributablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFirmwareServerConfigurationUtilityDistributables",
		Method:             "POST",
		PathPattern:        "/firmware/ServerConfigurationUtilityDistributables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFirmwareServerConfigurationUtilityDistributablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFirmwareServerConfigurationUtilityDistributablesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFirmwareServerConfigurationUtilityDistributablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostFirmwareServerConfigurationUtilityDistributablesMoid updates an instance of firmware server configuration utility distributable
*/
func (a *Client) PostFirmwareServerConfigurationUtilityDistributablesMoid(params *PostFirmwareServerConfigurationUtilityDistributablesMoidParams) (*PostFirmwareServerConfigurationUtilityDistributablesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFirmwareServerConfigurationUtilityDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFirmwareServerConfigurationUtilityDistributablesMoid",
		Method:             "POST",
		PathPattern:        "/firmware/ServerConfigurationUtilityDistributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFirmwareServerConfigurationUtilityDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFirmwareServerConfigurationUtilityDistributablesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFirmwareServerConfigurationUtilityDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
