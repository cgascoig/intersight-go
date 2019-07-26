// Code generated by go-swagger; DO NOT EDIT.

package iam_ldap_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iam ldap group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam ldap group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteIamLdapGroupsMoid deletes an instance of iam ldap group
*/
func (a *Client) DeleteIamLdapGroupsMoid(params *DeleteIamLdapGroupsMoidParams) (*DeleteIamLdapGroupsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIamLdapGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIamLdapGroupsMoid",
		Method:             "DELETE",
		PathPattern:        "/iam/LdapGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIamLdapGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIamLdapGroupsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIamLdapGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamLdapGroups gets a list of iam ldap group instances
*/
func (a *Client) GetIamLdapGroups(params *GetIamLdapGroupsParams) (*GetIamLdapGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamLdapGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamLdapGroups",
		Method:             "GET",
		PathPattern:        "/iam/LdapGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamLdapGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamLdapGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamLdapGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIamLdapGroupsMoid gets a specific instance of iam ldap group
*/
func (a *Client) GetIamLdapGroupsMoid(params *GetIamLdapGroupsMoidParams) (*GetIamLdapGroupsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIamLdapGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIamLdapGroupsMoid",
		Method:             "GET",
		PathPattern:        "/iam/LdapGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIamLdapGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIamLdapGroupsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIamLdapGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchIamLdapGroupsMoid updates an instance of iam ldap group
*/
func (a *Client) PatchIamLdapGroupsMoid(params *PatchIamLdapGroupsMoidParams) (*PatchIamLdapGroupsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchIamLdapGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchIamLdapGroupsMoid",
		Method:             "PATCH",
		PathPattern:        "/iam/LdapGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchIamLdapGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchIamLdapGroupsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchIamLdapGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamLdapGroups creates an instance of iam ldap group
*/
func (a *Client) PostIamLdapGroups(params *PostIamLdapGroupsParams) (*PostIamLdapGroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamLdapGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamLdapGroups",
		Method:             "POST",
		PathPattern:        "/iam/LdapGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamLdapGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamLdapGroupsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamLdapGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostIamLdapGroupsMoid updates an instance of iam ldap group
*/
func (a *Client) PostIamLdapGroupsMoid(params *PostIamLdapGroupsMoidParams) (*PostIamLdapGroupsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIamLdapGroupsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIamLdapGroupsMoid",
		Method:             "POST",
		PathPattern:        "/iam/LdapGroups/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIamLdapGroupsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIamLdapGroupsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostIamLdapGroupsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}