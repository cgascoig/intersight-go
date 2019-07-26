// Code generated by go-swagger; DO NOT EDIT.

package niaapi_dcnm_latest_maintained_release

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new niaapi dcnm latest maintained release API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for niaapi dcnm latest maintained release API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNiaapiDcnmLatestMaintainedReleases gets a list of niaapi dcnm latest maintained release instances
*/
func (a *Client) GetNiaapiDcnmLatestMaintainedReleases(params *GetNiaapiDcnmLatestMaintainedReleasesParams) (*GetNiaapiDcnmLatestMaintainedReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNiaapiDcnmLatestMaintainedReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNiaapiDcnmLatestMaintainedReleases",
		Method:             "GET",
		PathPattern:        "/niaapi/DcnmLatestMaintainedReleases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNiaapiDcnmLatestMaintainedReleasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNiaapiDcnmLatestMaintainedReleasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNiaapiDcnmLatestMaintainedReleasesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNiaapiDcnmLatestMaintainedReleasesMoid gets a specific instance of niaapi dcnm latest maintained release
*/
func (a *Client) GetNiaapiDcnmLatestMaintainedReleasesMoid(params *GetNiaapiDcnmLatestMaintainedReleasesMoidParams) (*GetNiaapiDcnmLatestMaintainedReleasesMoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNiaapiDcnmLatestMaintainedReleasesMoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNiaapiDcnmLatestMaintainedReleasesMoid",
		Method:             "GET",
		PathPattern:        "/niaapi/DcnmLatestMaintainedReleases/{moid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNiaapiDcnmLatestMaintainedReleasesMoidReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNiaapiDcnmLatestMaintainedReleasesMoidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNiaapiDcnmLatestMaintainedReleasesMoidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
