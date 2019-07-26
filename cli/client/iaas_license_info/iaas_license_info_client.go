// Code generated by go-swagger; DO NOT EDIT.

package iaas_license_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iaas license info API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iaas license info API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetIaasLicenseInfos gets a list of iaas license info instances
*/
func (a *Client) GetIaasLicenseInfos(params *GetIaasLicenseInfosParams) (*GetIaasLicenseInfosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIaasLicenseInfosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIaasLicenseInfos",
		Method:             "GET",
		PathPattern:        "/iaas/LicenseInfos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIaasLicenseInfosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIaasLicenseInfosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIaasLicenseInfosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIaasLicenseInfosMoid gets a specific instance of iaas license info
*/
func (a *Client) GetIaasLicenseInfosMoid(params *GetIaasLicenseInfosMoidParams) (*GetIaasLicenseInfosMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIaasLicenseInfosMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIaasLicenseInfosMoid",
		Method:             "GET",
		PathPattern:        "/iaas/LicenseInfos/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIaasLicenseInfosMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIaasLicenseInfosMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIaasLicenseInfosMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
