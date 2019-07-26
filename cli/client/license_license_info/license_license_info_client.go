// Code generated by go-swagger; DO NOT EDIT.

package license_license_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new license license info API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for license license info API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLicenseLicenseInfos gets a list of license license info instances
*/
func (a *Client) GetLicenseLicenseInfos(params *GetLicenseLicenseInfosParams) (*GetLicenseLicenseInfosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseLicenseInfosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLicenseLicenseInfos",
		Method:             "GET",
		PathPattern:        "/license/LicenseInfos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseLicenseInfosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLicenseLicenseInfosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLicenseLicenseInfosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLicenseLicenseInfosMoid gets a specific instance of license license info
*/
func (a *Client) GetLicenseLicenseInfosMoid(params *GetLicenseLicenseInfosMoidParams) (*GetLicenseLicenseInfosMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseLicenseInfosMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLicenseLicenseInfosMoid",
		Method:             "GET",
		PathPattern:        "/license/LicenseInfos/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseLicenseInfosMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLicenseLicenseInfosMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLicenseLicenseInfosMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchLicenseLicenseInfosMoid updates an instance of license license info
*/
func (a *Client) PatchLicenseLicenseInfosMoid(params *PatchLicenseLicenseInfosMoidParams) (*PatchLicenseLicenseInfosMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLicenseLicenseInfosMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLicenseLicenseInfosMoid",
		Method:             "PATCH",
		PathPattern:        "/license/LicenseInfos/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchLicenseLicenseInfosMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchLicenseLicenseInfosMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchLicenseLicenseInfosMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostLicenseLicenseInfos creates an instance of license license info
*/
func (a *Client) PostLicenseLicenseInfos(params *PostLicenseLicenseInfosParams) (*PostLicenseLicenseInfosCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLicenseLicenseInfosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLicenseLicenseInfos",
		Method:             "POST",
		PathPattern:        "/license/LicenseInfos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLicenseLicenseInfosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLicenseLicenseInfosCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostLicenseLicenseInfosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostLicenseLicenseInfosMoid updates an instance of license license info
*/
func (a *Client) PostLicenseLicenseInfosMoid(params *PostLicenseLicenseInfosMoidParams) (*PostLicenseLicenseInfosMoidCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLicenseLicenseInfosMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLicenseLicenseInfosMoid",
		Method:             "POST",
		PathPattern:        "/license/LicenseInfos/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLicenseLicenseInfosMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLicenseLicenseInfosMoidCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostLicenseLicenseInfosMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
