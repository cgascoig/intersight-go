// Code generated by go-swagger; DO NOT EDIT.

package hcl_hyperflex_software_compatibility_info

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

// NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParams creates a new PatchHclHyperflexSoftwareCompatibilityInfosMoidParams object
// with the default values initialized.
func NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParams() *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	var ()
	return &PatchHclHyperflexSoftwareCompatibilityInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParamsWithTimeout creates a new PatchHclHyperflexSoftwareCompatibilityInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParamsWithTimeout(timeout time.Duration) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	var ()
	return &PatchHclHyperflexSoftwareCompatibilityInfosMoidParams{

		timeout: timeout,
	}
}

// NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParamsWithContext creates a new PatchHclHyperflexSoftwareCompatibilityInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParamsWithContext(ctx context.Context) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	var ()
	return &PatchHclHyperflexSoftwareCompatibilityInfosMoidParams{

		Context: ctx,
	}
}

// NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParamsWithHTTPClient creates a new PatchHclHyperflexSoftwareCompatibilityInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHclHyperflexSoftwareCompatibilityInfosMoidParamsWithHTTPClient(client *http.Client) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	var ()
	return &PatchHclHyperflexSoftwareCompatibilityInfosMoidParams{
		HTTPClient: client,
	}
}

/*PatchHclHyperflexSoftwareCompatibilityInfosMoidParams contains all the parameters to send to the API endpoint
for the patch hcl hyperflex software compatibility infos moid operation typically these are written to a http.Request
*/
type PatchHclHyperflexSoftwareCompatibilityInfosMoidParams struct {

	/*Body
	  hclHyperflexSoftwareCompatibilityInfo to update

	*/
	Body *models.HclHyperflexSoftwareCompatibilityInfo
	/*Moid
	  The moid of the hclHyperflexSoftwareCompatibilityInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) WithTimeout(timeout time.Duration) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) WithContext(ctx context.Context) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) WithHTTPClient(client *http.Client) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) WithBody(body *models.HclHyperflexSoftwareCompatibilityInfo) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) SetBody(body *models.HclHyperflexSoftwareCompatibilityInfo) {
	o.Body = body
}

// WithMoid adds the moid to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) WithMoid(moid string) *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hcl hyperflex software compatibility infos moid params
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHclHyperflexSoftwareCompatibilityInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
