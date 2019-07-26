// Code generated by go-swagger; DO NOT EDIT.

package firmware_running_firmware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new firmware running firmware API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for firmware running firmware API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFirmwareRunningFirmwares gets a list of firmware running firmware instances
*/
func (a *Client) GetFirmwareRunningFirmwares(params *GetFirmwareRunningFirmwaresParams) (*GetFirmwareRunningFirmwaresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFirmwareRunningFirmwaresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFirmwareRunningFirmwares",
		Method:             "GET",
		PathPattern:        "/firmware/RunningFirmwares",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFirmwareRunningFirmwaresReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFirmwareRunningFirmwaresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFirmwareRunningFirmwaresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFirmwareRunningFirmwaresMoid gets a specific instance of firmware running firmware
*/
func (a *Client) GetFirmwareRunningFirmwaresMoid(params *GetFirmwareRunningFirmwaresMoidParams) (*GetFirmwareRunningFirmwaresMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFirmwareRunningFirmwaresMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFirmwareRunningFirmwaresMoid",
		Method:             "GET",
		PathPattern:        "/firmware/RunningFirmwares/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFirmwareRunningFirmwaresMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFirmwareRunningFirmwaresMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFirmwareRunningFirmwaresMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchFirmwareRunningFirmwaresMoid updates an instance of firmware running firmware
*/
func (a *Client) PatchFirmwareRunningFirmwaresMoid(params *PatchFirmwareRunningFirmwaresMoidParams) (*PatchFirmwareRunningFirmwaresMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchFirmwareRunningFirmwaresMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchFirmwareRunningFirmwaresMoid",
		Method:             "PATCH",
		PathPattern:        "/firmware/RunningFirmwares/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchFirmwareRunningFirmwaresMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchFirmwareRunningFirmwaresMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchFirmwareRunningFirmwaresMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostFirmwareRunningFirmwaresMoid updates an instance of firmware running firmware
*/
func (a *Client) PostFirmwareRunningFirmwaresMoid(params *PostFirmwareRunningFirmwaresMoidParams) (*PostFirmwareRunningFirmwaresMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFirmwareRunningFirmwaresMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFirmwareRunningFirmwaresMoid",
		Method:             "POST",
		PathPattern:        "/firmware/RunningFirmwares/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFirmwareRunningFirmwaresMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFirmwareRunningFirmwaresMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFirmwareRunningFirmwaresMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}