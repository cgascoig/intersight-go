// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_vcenter_config_policy

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

// NewPatchHyperflexVcenterConfigPoliciesMoidParams creates a new PatchHyperflexVcenterConfigPoliciesMoidParams object
// with the default values initialized.
func NewPatchHyperflexVcenterConfigPoliciesMoidParams() *PatchHyperflexVcenterConfigPoliciesMoidParams {
	var ()
	return &PatchHyperflexVcenterConfigPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexVcenterConfigPoliciesMoidParamsWithTimeout creates a new PatchHyperflexVcenterConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexVcenterConfigPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	var ()
	return &PatchHyperflexVcenterConfigPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexVcenterConfigPoliciesMoidParamsWithContext creates a new PatchHyperflexVcenterConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexVcenterConfigPoliciesMoidParamsWithContext(ctx context.Context) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	var ()
	return &PatchHyperflexVcenterConfigPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexVcenterConfigPoliciesMoidParamsWithHTTPClient creates a new PatchHyperflexVcenterConfigPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexVcenterConfigPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	var ()
	return &PatchHyperflexVcenterConfigPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexVcenterConfigPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex vcenter config policies moid operation typically these are written to a http.Request
*/
type PatchHyperflexVcenterConfigPoliciesMoidParams struct {

	/*Body
	  hyperflexVcenterConfigPolicy to update

	*/
	Body *models.HyperflexVcenterConfigPolicy
	/*Moid
	  The moid of the hyperflexVcenterConfigPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) WithContext(ctx context.Context) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) WithBody(body *models.HyperflexVcenterConfigPolicy) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) SetBody(body *models.HyperflexVcenterConfigPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) WithMoid(moid string) *PatchHyperflexVcenterConfigPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex vcenter config policies moid params
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexVcenterConfigPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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