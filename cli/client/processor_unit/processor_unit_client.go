// Code generated by go-swagger; DO NOT EDIT.

package processor_unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new processor unit API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for processor unit API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetProcessorUnits gets a list of processor unit instances
*/
func (a *Client) GetProcessorUnits(params *GetProcessorUnitsParams) (*GetProcessorUnitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessorUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProcessorUnits",
		Method:             "GET",
		PathPattern:        "/processor/Units",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProcessorUnitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProcessorUnitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProcessorUnitsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProcessorUnitsMoid gets a specific instance of processor unit
*/
func (a *Client) GetProcessorUnitsMoid(params *GetProcessorUnitsMoidParams) (*GetProcessorUnitsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessorUnitsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProcessorUnitsMoid",
		Method:             "GET",
		PathPattern:        "/processor/Units/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProcessorUnitsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProcessorUnitsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProcessorUnitsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchProcessorUnitsMoid updates an instance of processor unit
*/
func (a *Client) PatchProcessorUnitsMoid(params *PatchProcessorUnitsMoidParams) (*PatchProcessorUnitsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchProcessorUnitsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchProcessorUnitsMoid",
		Method:             "PATCH",
		PathPattern:        "/processor/Units/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchProcessorUnitsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchProcessorUnitsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchProcessorUnitsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostProcessorUnitsMoid updates an instance of processor unit
*/
func (a *Client) PostProcessorUnitsMoid(params *PostProcessorUnitsMoidParams) (*PostProcessorUnitsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProcessorUnitsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostProcessorUnitsMoid",
		Method:             "POST",
		PathPattern:        "/processor/Units/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostProcessorUnitsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostProcessorUnitsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostProcessorUnitsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
