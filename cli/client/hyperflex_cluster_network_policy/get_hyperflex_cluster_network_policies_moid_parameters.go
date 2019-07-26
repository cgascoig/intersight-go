// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_network_policy

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

// NewGetHyperflexClusterNetworkPoliciesMoidParams creates a new GetHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized.
func NewGetHyperflexClusterNetworkPoliciesMoidParams() *GetHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &GetHyperflexClusterNetworkPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHyperflexClusterNetworkPoliciesMoidParamsWithTimeout creates a new GetHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHyperflexClusterNetworkPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &GetHyperflexClusterNetworkPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetHyperflexClusterNetworkPoliciesMoidParamsWithContext creates a new GetHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHyperflexClusterNetworkPoliciesMoidParamsWithContext(ctx context.Context) *GetHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &GetHyperflexClusterNetworkPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetHyperflexClusterNetworkPoliciesMoidParamsWithHTTPClient creates a new GetHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHyperflexClusterNetworkPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &GetHyperflexClusterNetworkPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetHyperflexClusterNetworkPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get hyperflex cluster network policies moid operation typically these are written to a http.Request
*/
type GetHyperflexClusterNetworkPoliciesMoidParams struct {

	/*Moid
	  The moid of the hyperflexClusterNetworkPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetHyperflexClusterNetworkPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) WithContext(ctx context.Context) *GetHyperflexClusterNetworkPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetHyperflexClusterNetworkPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) WithMoid(moid string) *GetHyperflexClusterNetworkPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hyperflex cluster network policies moid params
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHyperflexClusterNetworkPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
