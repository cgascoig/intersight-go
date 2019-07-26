// Code generated by go-swagger; DO NOT EDIT.

package pci_coprocessor_card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pci coprocessor card API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pci coprocessor card API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPciCoprocessorCards gets a list of pci coprocessor card instances
*/
func (a *Client) GetPciCoprocessorCards(params *GetPciCoprocessorCardsParams) (*GetPciCoprocessorCardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPciCoprocessorCardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPciCoprocessorCards",
		Method:             "GET",
		PathPattern:        "/pci/CoprocessorCards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPciCoprocessorCardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPciCoprocessorCardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPciCoprocessorCardsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPciCoprocessorCardsMoid gets a specific instance of pci coprocessor card
*/
func (a *Client) GetPciCoprocessorCardsMoid(params *GetPciCoprocessorCardsMoidParams) (*GetPciCoprocessorCardsMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPciCoprocessorCardsMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPciCoprocessorCardsMoid",
		Method:             "GET",
		PathPattern:        "/pci/CoprocessorCards/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPciCoprocessorCardsMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPciCoprocessorCardsMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPciCoprocessorCardsMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
