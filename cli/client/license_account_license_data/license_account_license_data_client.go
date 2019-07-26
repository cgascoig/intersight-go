// Code generated by go-swagger; DO NOT EDIT.

package license_account_license_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new license account license data API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for license account license data API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLicenseAccountLicenseData gets a list of license account license data instances
*/
func (a *Client) GetLicenseAccountLicenseData(params *GetLicenseAccountLicenseDataParams) (*GetLicenseAccountLicenseDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseAccountLicenseDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLicenseAccountLicenseData",
		Method:             "GET",
		PathPattern:        "/license/AccountLicenseData",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseAccountLicenseDataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLicenseAccountLicenseDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLicenseAccountLicenseDataDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLicenseAccountLicenseDataMoid gets a specific instance of license account license data
*/
func (a *Client) GetLicenseAccountLicenseDataMoid(params *GetLicenseAccountLicenseDataMoidParams) (*GetLicenseAccountLicenseDataMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseAccountLicenseDataMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLicenseAccountLicenseDataMoid",
		Method:             "GET",
		PathPattern:        "/license/AccountLicenseData/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseAccountLicenseDataMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLicenseAccountLicenseDataMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLicenseAccountLicenseDataMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
