// Code generated by go-swagger; DO NOT EDIT.

package iam_privilege

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iam privilege API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam privilege API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetIamPrivileges gets a list of iam privilege instances
*/
func (a *Client) GetIamPrivileges(params *GetIamPrivilegesParams) (*GetIamPrivilegesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamPrivilegesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamPrivileges",
		Method:             "GET",
		PathPattern:        "/iam/Privileges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamPrivilegesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamPrivilegesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamPrivilegesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamPrivilegesMoid gets a specific instance of iam privilege
*/
func (a *Client) GetIamPrivilegesMoid(params *GetIamPrivilegesMoidParams) (*GetIamPrivilegesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamPrivilegesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamPrivilegesMoid",
		Method:             "GET",
		PathPattern:        "/iam/Privileges/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamPrivilegesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamPrivilegesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamPrivilegesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}