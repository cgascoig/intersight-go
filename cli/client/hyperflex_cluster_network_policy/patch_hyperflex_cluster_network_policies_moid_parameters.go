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

	models "github.com/cgascoig/intersight-go/cli/models"
)

// NewPatchHyperflexClusterNetworkPoliciesMoidParams creates a new PatchHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized.
func NewPatchHyperflexClusterNetworkPoliciesMoidParams() *PatchHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &PatchHyperflexClusterNetworkPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexClusterNetworkPoliciesMoidParamsWithTimeout creates a new PatchHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexClusterNetworkPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &PatchHyperflexClusterNetworkPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexClusterNetworkPoliciesMoidParamsWithContext creates a new PatchHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexClusterNetworkPoliciesMoidParamsWithContext(ctx context.Context) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &PatchHyperflexClusterNetworkPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexClusterNetworkPoliciesMoidParamsWithHTTPClient creates a new PatchHyperflexClusterNetworkPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexClusterNetworkPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	var ()
	return &PatchHyperflexClusterNetworkPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexClusterNetworkPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex cluster network policies moid operation typically these are written to a http.Request
*/
type PatchHyperflexClusterNetworkPoliciesMoidParams struct {

	/*Body
	  hyperflexClusterNetworkPolicy to update

	*/
	Body *models.HyperflexClusterNetworkPolicy
	/*Moid
	  The moid of the hyperflexClusterNetworkPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) WithContext(ctx context.Context) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) WithBody(body *models.HyperflexClusterNetworkPolicy) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) SetBody(body *models.HyperflexClusterNetworkPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) WithMoid(moid string) *PatchHyperflexClusterNetworkPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex cluster network policies moid params
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexClusterNetworkPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
