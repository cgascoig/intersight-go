// Code generated by go-swagger; DO NOT EDIT.

package management_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new management controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetManagementControllers gets a list of management controller instances
*/
func (a *Client) GetManagementControllers(params *GetManagementControllersParams) (*GetManagementControllersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagementControllersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetManagementControllers",
		Method:             "GET",
		PathPattern:        "/management/Controllers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetManagementControllersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetManagementControllersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetManagementControllersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetManagementControllersMoid gets a specific instance of management controller
*/
func (a *Client) GetManagementControllersMoid(params *GetManagementControllersMoidParams) (*GetManagementControllersMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagementControllersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetManagementControllersMoid",
		Method:             "GET",
		PathPattern:        "/management/Controllers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetManagementControllersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetManagementControllersMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetManagementControllersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchManagementControllersMoid updates an instance of management controller
*/
func (a *Client) PatchManagementControllersMoid(params *PatchManagementControllersMoidParams) (*PatchManagementControllersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchManagementControllersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchManagementControllersMoid",
		Method:             "PATCH",
		PathPattern:        "/management/Controllers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchManagementControllersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchManagementControllersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchManagementControllersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostManagementControllersMoid updates an instance of management controller
*/
func (a *Client) PostManagementControllersMoid(params *PostManagementControllersMoidParams) (*PostManagementControllersMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostManagementControllersMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostManagementControllersMoid",
		Method:             "POST",
		PathPattern:        "/management/Controllers/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostManagementControllersMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostManagementControllersMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostManagementControllersMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
