// Code generated by go-swagger; DO NOT EDIT.

package iaas_most_run_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iaas most run tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iaas most run tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetIaasMostRunTasks gets a list of iaas most run tasks instances
*/
func (a *Client) GetIaasMostRunTasks(params *GetIaasMostRunTasksParams) (*GetIaasMostRunTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIaasMostRunTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIaasMostRunTasks",
		Method:             "GET",
		PathPattern:        "/iaas/MostRunTasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIaasMostRunTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIaasMostRunTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIaasMostRunTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIaasMostRunTasksMoid gets a specific instance of iaas most run tasks
*/
func (a *Client) GetIaasMostRunTasksMoid(params *GetIaasMostRunTasksMoidParams) (*GetIaasMostRunTasksMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIaasMostRunTasksMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIaasMostRunTasksMoid",
		Method:             "GET",
		PathPattern:        "/iaas/MostRunTasks/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIaasMostRunTasksMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIaasMostRunTasksMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIaasMostRunTasksMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
