// Code generated by go-swagger; DO NOT EDIT.

package workflow_batch_api_executor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workflow batch api executor API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow batch api executor API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteWorkflowBatchAPIExecutorsMoid deletes an instance of workflow batch Api executor
*/
func (a *Client) DeleteWorkflowBatchAPIExecutorsMoid(params *DeleteWorkflowBatchAPIExecutorsMoidParams) (*DeleteWorkflowBatchAPIExecutorsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkflowBatchAPIExecutorsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteWorkflowBatchAPIExecutorsMoid",
		Method:             "DELETE",
		PathPattern:        "/workflow/BatchApiExecutors/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWorkflowBatchAPIExecutorsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWorkflowBatchAPIExecutorsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWorkflowBatchAPIExecutorsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetWorkflowBatchAPIExecutors gets a list of workflow batch Api executor instances
*/
func (a *Client) GetWorkflowBatchAPIExecutors(params *GetWorkflowBatchAPIExecutorsParams) (*GetWorkflowBatchAPIExecutorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowBatchAPIExecutorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowBatchAPIExecutors",
		Method:             "GET",
		PathPattern:        "/workflow/BatchApiExecutors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkflowBatchAPIExecutorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkflowBatchAPIExecutorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowBatchAPIExecutorsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetWorkflowBatchAPIExecutorsMoid gets a specific instance of workflow batch Api executor
*/
func (a *Client) GetWorkflowBatchAPIExecutorsMoid(params *GetWorkflowBatchAPIExecutorsMoidParams) (*GetWorkflowBatchAPIExecutorsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowBatchAPIExecutorsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWorkflowBatchAPIExecutorsMoid",
		Method:             "GET",
		PathPattern:        "/workflow/BatchApiExecutors/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkflowBatchAPIExecutorsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkflowBatchAPIExecutorsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowBatchAPIExecutorsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchWorkflowBatchAPIExecutorsMoid updates an instance of workflow batch Api executor
*/
func (a *Client) PatchWorkflowBatchAPIExecutorsMoid(params *PatchWorkflowBatchAPIExecutorsMoidParams) (*PatchWorkflowBatchAPIExecutorsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchWorkflowBatchAPIExecutorsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchWorkflowBatchAPIExecutorsMoid",
		Method:             "PATCH",
		PathPattern:        "/workflow/BatchApiExecutors/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchWorkflowBatchAPIExecutorsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchWorkflowBatchAPIExecutorsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchWorkflowBatchAPIExecutorsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostWorkflowBatchAPIExecutors creates an instance of workflow batch Api executor
*/
func (a *Client) PostWorkflowBatchAPIExecutors(params *PostWorkflowBatchAPIExecutorsParams) (*PostWorkflowBatchAPIExecutorsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWorkflowBatchAPIExecutorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWorkflowBatchAPIExecutors",
		Method:             "POST",
		PathPattern:        "/workflow/BatchApiExecutors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWorkflowBatchAPIExecutorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWorkflowBatchAPIExecutorsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostWorkflowBatchAPIExecutorsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostWorkflowBatchAPIExecutorsMoid updates an instance of workflow batch Api executor
*/
func (a *Client) PostWorkflowBatchAPIExecutorsMoid(params *PostWorkflowBatchAPIExecutorsMoidParams) (*PostWorkflowBatchAPIExecutorsMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWorkflowBatchAPIExecutorsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWorkflowBatchAPIExecutorsMoid",
		Method:             "POST",
		PathPattern:        "/workflow/BatchApiExecutors/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWorkflowBatchAPIExecutorsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWorkflowBatchAPIExecutorsMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostWorkflowBatchAPIExecutorsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
