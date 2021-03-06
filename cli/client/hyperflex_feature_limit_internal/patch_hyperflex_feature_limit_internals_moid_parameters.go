// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_feature_limit_internal

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

// NewPatchHyperflexFeatureLimitInternalsMoidParams creates a new PatchHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized.
func NewPatchHyperflexFeatureLimitInternalsMoidParams() *PatchHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &PatchHyperflexFeatureLimitInternalsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexFeatureLimitInternalsMoidParamsWithTimeout creates a new PatchHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexFeatureLimitInternalsMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &PatchHyperflexFeatureLimitInternalsMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexFeatureLimitInternalsMoidParamsWithContext creates a new PatchHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexFeatureLimitInternalsMoidParamsWithContext(ctx context.Context) *PatchHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &PatchHyperflexFeatureLimitInternalsMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexFeatureLimitInternalsMoidParamsWithHTTPClient creates a new PatchHyperflexFeatureLimitInternalsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexFeatureLimitInternalsMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexFeatureLimitInternalsMoidParams {
	var ()
	return &PatchHyperflexFeatureLimitInternalsMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexFeatureLimitInternalsMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex feature limit internals moid operation typically these are written to a http.Request
*/
type PatchHyperflexFeatureLimitInternalsMoidParams struct {

	/*Body
	  hyperflexFeatureLimitInternal to update

	*/
	Body *models.HyperflexFeatureLimitInternal
	/*Moid
	  The moid of the hyperflexFeatureLimitInternal instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexFeatureLimitInternalsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) WithContext(ctx context.Context) *PatchHyperflexFeatureLimitInternalsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexFeatureLimitInternalsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) WithBody(body *models.HyperflexFeatureLimitInternal) *PatchHyperflexFeatureLimitInternalsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) SetBody(body *models.HyperflexFeatureLimitInternal) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) WithMoid(moid string) *PatchHyperflexFeatureLimitInternalsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex feature limit internals moid params
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexFeatureLimitInternalsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
