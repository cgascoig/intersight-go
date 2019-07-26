// Code generated by go-swagger; DO NOT EDIT.

package management_entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new management entity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management entity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetManagementEntities gets a list of management entity instances
*/
func (a *Client) GetManagementEntities(params *GetManagementEntitiesParams) (*GetManagementEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagementEntitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetManagementEntities",
		Method:             "GET",
		PathPattern:        "/management/Entities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetManagementEntitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetManagementEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetManagementEntitiesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetManagementEntitiesMoid gets a specific instance of management entity
*/
func (a *Client) GetManagementEntitiesMoid(params *GetManagementEntitiesMoidParams) (*GetManagementEntitiesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagementEntitiesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetManagementEntitiesMoid",
		Method:             "GET",
		PathPattern:        "/management/Entities/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetManagementEntitiesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetManagementEntitiesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetManagementEntitiesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchManagementEntitiesMoid updates an instance of management entity
*/
func (a *Client) PatchManagementEntitiesMoid(params *PatchManagementEntitiesMoidParams) (*PatchManagementEntitiesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchManagementEntitiesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchManagementEntitiesMoid",
		Method:             "PATCH",
		PathPattern:        "/management/Entities/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchManagementEntitiesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchManagementEntitiesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchManagementEntitiesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostManagementEntitiesMoid updates an instance of management entity
*/
func (a *Client) PostManagementEntitiesMoid(params *PostManagementEntitiesMoidParams) (*PostManagementEntitiesMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostManagementEntitiesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostManagementEntitiesMoid",
		Method:             "POST",
		PathPattern:        "/management/Entities/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostManagementEntitiesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostManagementEntitiesMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostManagementEntitiesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}