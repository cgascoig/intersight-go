// Code generated by go-swagger; DO NOT EDIT.

package asset_device_connector_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAssetDeviceConnectorManagersMoidParams creates a new GetAssetDeviceConnectorManagersMoidParams object
// with the default values initialized.
func NewGetAssetDeviceConnectorManagersMoidParams() *GetAssetDeviceConnectorManagersMoidParams {
	var ()
	return &GetAssetDeviceConnectorManagersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssetDeviceConnectorManagersMoidParamsWithTimeout creates a new GetAssetDeviceConnectorManagersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssetDeviceConnectorManagersMoidParamsWithTimeout(timeout time.Duration) *GetAssetDeviceConnectorManagersMoidParams {
	var ()
	return &GetAssetDeviceConnectorManagersMoidParams{

		timeout: timeout,
	}
}

// NewGetAssetDeviceConnectorManagersMoidParamsWithContext creates a new GetAssetDeviceConnectorManagersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssetDeviceConnectorManagersMoidParamsWithContext(ctx context.Context) *GetAssetDeviceConnectorManagersMoidParams {
	var ()
	return &GetAssetDeviceConnectorManagersMoidParams{

		Context: ctx,
	}
}

// NewGetAssetDeviceConnectorManagersMoidParamsWithHTTPClient creates a new GetAssetDeviceConnectorManagersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssetDeviceConnectorManagersMoidParamsWithHTTPClient(client *http.Client) *GetAssetDeviceConnectorManagersMoidParams {
	var ()
	return &GetAssetDeviceConnectorManagersMoidParams{
		HTTPClient: client,
	}
}

/*GetAssetDeviceConnectorManagersMoidParams contains all the parameters to send to the API endpoint
for the get asset device connector managers moid operation typically these are written to a http.Request
*/
type GetAssetDeviceConnectorManagersMoidParams struct {

	/*Moid
	  The moid of the assetDeviceConnectorManager instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) WithTimeout(timeout time.Duration) *GetAssetDeviceConnectorManagersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) WithContext(ctx context.Context) *GetAssetDeviceConnectorManagersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) WithHTTPClient(client *http.Client) *GetAssetDeviceConnectorManagersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) WithMoid(moid string) *GetAssetDeviceConnectorManagersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get asset device connector managers moid params
func (o *GetAssetDeviceConnectorManagersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssetDeviceConnectorManagersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
