// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_initiate_hx_cluster_upgrade

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

// NewGetHyperflexInitiateHxClusterUpgradesMoidParams creates a new GetHyperflexInitiateHxClusterUpgradesMoidParams object
// with the default values initialized.
func NewGetHyperflexInitiateHxClusterUpgradesMoidParams() *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	var ()
	return &GetHyperflexInitiateHxClusterUpgradesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexInitiateHxClusterUpgradesMoidParamsWithTimeout creates a new GetHyperflexInitiateHxClusterUpgradesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexInitiateHxClusterUpgradesMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	var ()
	return &GetHyperflexInitiateHxClusterUpgradesMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexInitiateHxClusterUpgradesMoidParamsWithContext creates a new GetHyperflexInitiateHxClusterUpgradesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexInitiateHxClusterUpgradesMoidParamsWithContext(ctx context.Context) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	var ()
	return &GetHyperflexInitiateHxClusterUpgradesMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexInitiateHxClusterUpgradesMoidParamsWithHTTPClient creates a new GetHyperflexInitiateHxClusterUpgradesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexInitiateHxClusterUpgradesMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	var ()
	return &GetHyperflexInitiateHxClusterUpgradesMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexInitiateHxClusterUpgradesMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex initiate hx cluster upgrades moid operation typically these are written to a http.Request
*/
type GetHyperflexInitiateHxClusterUpgradesMoidParams struct {

	/*Moid
	  The moid of the hyperflexInitiateHxClusterUpgrade instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) WithContext(ctx context.Context) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) WithMoid(moid string) *GetHyperflexInitiateHxClusterUpgradesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex initiate hx cluster upgrades moid params
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexInitiateHxClusterUpgradesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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