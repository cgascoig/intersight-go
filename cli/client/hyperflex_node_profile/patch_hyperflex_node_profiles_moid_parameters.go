// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_node_profile

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

// NewPatchHyperflexNodeProfilesMoidParams creates a new PatchHyperflexNodeProfilesMoidParams object
// with the default values initialized.
func NewPatchHyperflexNodeProfilesMoidParams() *PatchHyperflexNodeProfilesMoidParams {
	var ()
	return &PatchHyperflexNodeProfilesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexNodeProfilesMoidParamsWithTimeout creates a new PatchHyperflexNodeProfilesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexNodeProfilesMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexNodeProfilesMoidParams {
	var ()
	return &PatchHyperflexNodeProfilesMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexNodeProfilesMoidParamsWithContext creates a new PatchHyperflexNodeProfilesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexNodeProfilesMoidParamsWithContext(ctx context.Context) *PatchHyperflexNodeProfilesMoidParams {
	var ()
	return &PatchHyperflexNodeProfilesMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexNodeProfilesMoidParamsWithHTTPClient creates a new PatchHyperflexNodeProfilesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexNodeProfilesMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexNodeProfilesMoidParams {
	var ()
	return &PatchHyperflexNodeProfilesMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexNodeProfilesMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex node profiles moid operation typically these are written to a http.Request
*/
type PatchHyperflexNodeProfilesMoidParams struct {

	/*Body
	  hyperflexNodeProfile to update

	*/
	Body *models.HyperflexNodeProfile
	/*Moid
	  The moid of the hyperflexNodeProfile instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexNodeProfilesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) WithContext(ctx context.Context) *PatchHyperflexNodeProfilesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexNodeProfilesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) WithBody(body *models.HyperflexNodeProfile) *PatchHyperflexNodeProfilesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) SetBody(body *models.HyperflexNodeProfile) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) WithMoid(moid string) *PatchHyperflexNodeProfilesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex node profiles moid params
func (o *PatchHyperflexNodeProfilesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexNodeProfilesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
