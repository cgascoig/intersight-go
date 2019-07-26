// Code generated by go-swagger; DO NOT EDIT.

package iam_end_point_role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iam end point role API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam end point role API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetIamEndPointRoles gets a list of iam end point role instances
*/
func (a *Client) GetIamEndPointRoles(params *GetIamEndPointRolesParams) (*GetIamEndPointRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamEndPointRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamEndPointRoles",
		Method:             "GET",
		PathPattern:        "/iam/EndPointRoles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamEndPointRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamEndPointRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamEndPointRolesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamEndPointRolesMoid gets a specific instance of iam end point role
*/
func (a *Client) GetIamEndPointRolesMoid(params *GetIamEndPointRolesMoidParams) (*GetIamEndPointRolesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamEndPointRolesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamEndPointRolesMoid",
		Method:             "GET",
		PathPattern:        "/iam/EndPointRoles/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamEndPointRolesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamEndPointRolesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamEndPointRolesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}