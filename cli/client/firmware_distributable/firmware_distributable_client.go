// Code generated by go-swagger; DO NOT EDIT.

package firmware_distributable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new firmware distributable API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for firmware distributable API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteFirmwareDistributablesMoid deletes an instance of firmware distributable
*/
func (a *Client) DeleteFirmwareDistributablesMoid(params *DeleteFirmwareDistributablesMoidParams) (*DeleteFirmwareDistributablesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFirmwareDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFirmwareDistributablesMoid",
		Method:             "DELETE",
		PathPattern:        "/firmware/Distributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFirmwareDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteFirmwareDistributablesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFirmwareDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFirmwareDistributables gets a list of firmware distributable instances
*/
func (a *Client) GetFirmwareDistributables(params *GetFirmwareDistributablesParams) (*GetFirmwareDistributablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFirmwareDistributablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFirmwareDistributables",
		Method:             "GET",
		PathPattern:        "/firmware/Distributables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFirmwareDistributablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFirmwareDistributablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFirmwareDistributablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFirmwareDistributablesMoid gets a specific instance of firmware distributable
*/
func (a *Client) GetFirmwareDistributablesMoid(params *GetFirmwareDistributablesMoidParams) (*GetFirmwareDistributablesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFirmwareDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFirmwareDistributablesMoid",
		Method:             "GET",
		PathPattern:        "/firmware/Distributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFirmwareDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFirmwareDistributablesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFirmwareDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchFirmwareDistributablesMoid updates an instance of firmware distributable
*/
func (a *Client) PatchFirmwareDistributablesMoid(params *PatchFirmwareDistributablesMoidParams) (*PatchFirmwareDistributablesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchFirmwareDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchFirmwareDistributablesMoid",
		Method:             "PATCH",
		PathPattern:        "/firmware/Distributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchFirmwareDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchFirmwareDistributablesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchFirmwareDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostFirmwareDistributables creates an instance of firmware distributable
*/
func (a *Client) PostFirmwareDistributables(params *PostFirmwareDistributablesParams) (*PostFirmwareDistributablesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFirmwareDistributablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFirmwareDistributables",
		Method:             "POST",
		PathPattern:        "/firmware/Distributables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFirmwareDistributablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFirmwareDistributablesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFirmwareDistributablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostFirmwareDistributablesMoid updates an instance of firmware distributable
*/
func (a *Client) PostFirmwareDistributablesMoid(params *PostFirmwareDistributablesMoidParams) (*PostFirmwareDistributablesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFirmwareDistributablesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFirmwareDistributablesMoid",
		Method:             "POST",
		PathPattern:        "/firmware/Distributables/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFirmwareDistributablesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFirmwareDistributablesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostFirmwareDistributablesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
