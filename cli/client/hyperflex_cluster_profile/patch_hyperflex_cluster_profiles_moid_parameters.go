// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_cluster_profile

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

// NewPatchHyperflexClusterProfilesMoidParams creates a new PatchHyperflexClusterProfilesMoidParams object
// with the default values initialized.
func NewPatchHyperflexClusterProfilesMoidParams() *PatchHyperflexClusterProfilesMoidParams {
	var ()
	return &PatchHyperflexClusterProfilesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexClusterProfilesMoidParamsWithTimeout creates a new PatchHyperflexClusterProfilesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexClusterProfilesMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexClusterProfilesMoidParams {
	var ()
	return &PatchHyperflexClusterProfilesMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexClusterProfilesMoidParamsWithContext creates a new PatchHyperflexClusterProfilesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexClusterProfilesMoidParamsWithContext(ctx context.Context) *PatchHyperflexClusterProfilesMoidParams {
	var ()
	return &PatchHyperflexClusterProfilesMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexClusterProfilesMoidParamsWithHTTPClient creates a new PatchHyperflexClusterProfilesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexClusterProfilesMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexClusterProfilesMoidParams {
	var ()
	return &PatchHyperflexClusterProfilesMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexClusterProfilesMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex cluster profiles moid operation typically these are written to a http.Request
*/
type PatchHyperflexClusterProfilesMoidParams struct {

	/*Body
	  hyperflexClusterProfile to update

	*/
	Body *models.HyperflexClusterProfile
	/*Moid
	  The moid of the hyperflexClusterProfile instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexClusterProfilesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) WithContext(ctx context.Context) *PatchHyperflexClusterProfilesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexClusterProfilesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) WithBody(body *models.HyperflexClusterProfile) *PatchHyperflexClusterProfilesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) SetBody(body *models.HyperflexClusterProfile) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) WithMoid(moid string) *PatchHyperflexClusterProfilesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex cluster profiles moid params
func (o *PatchHyperflexClusterProfilesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexClusterProfilesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
