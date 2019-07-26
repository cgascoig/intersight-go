// Code generated by go-swagger; DO NOT EDIT.

package niaapi_apic_latest_maintained_release

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new niaapi apic latest maintained release API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for niaapi apic latest maintained release API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNiaapiApicLatestMaintainedReleases gets a list of niaapi apic latest maintained release instances
*/
func (a *Client) GetNiaapiApicLatestMaintainedReleases(params *GetNiaapiApicLatestMaintainedReleasesParams) (*GetNiaapiApicLatestMaintainedReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNiaapiApicLatestMaintainedReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNiaapiApicLatestMaintainedReleases",
		Method:             "GET",
		PathPattern:        "/niaapi/ApicLatestMaintainedReleases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNiaapiApicLatestMaintainedReleasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNiaapiApicLatestMaintainedReleasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNiaapiApicLatestMaintainedReleasesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNiaapiApicLatestMaintainedReleasesMoid gets a specific instance of niaapi apic latest maintained release
*/
func (a *Client) GetNiaapiApicLatestMaintainedReleasesMoid(params *GetNiaapiApicLatestMaintainedReleasesMoidParams) (*GetNiaapiApicLatestMaintainedReleasesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNiaapiApicLatestMaintainedReleasesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNiaapiApicLatestMaintainedReleasesMoid",
		Method:             "GET",
		PathPattern:        "/niaapi/ApicLatestMaintainedReleases/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNiaapiApicLatestMaintainedReleasesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNiaapiApicLatestMaintainedReleasesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNiaapiApicLatestMaintainedReleasesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}