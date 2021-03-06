// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_alarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hyperflex alarm API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hyperflex alarm API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHyperflexAlarms gets a list of hyperflex alarm instances
*/
func (a *Client) GetHyperflexAlarms(params *GetHyperflexAlarmsParams) (*GetHyperflexAlarmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHyperflexAlarmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHyperflexAlarms",
		Method:             "GET",
		PathPattern:        "/hyperflex/Alarms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHyperflexAlarmsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHyperflexAlarmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHyperflexAlarmsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHyperflexAlarmsMoid gets a specific instance of hyperflex alarm
*/
func (a *Client) GetHyperflexAlarmsMoid(params *GetHyperflexAlarmsMoidParams) (*GetHyperflexAlarmsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHyperflexAlarmsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHyperflexAlarmsMoid",
		Method:             "GET",
		PathPattern:        "/hyperflex/Alarms/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHyperflexAlarmsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHyperflexAlarmsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHyperflexAlarmsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
