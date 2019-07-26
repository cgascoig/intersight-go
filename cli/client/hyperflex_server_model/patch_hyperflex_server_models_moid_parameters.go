// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_server_model

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

// NewPatchHyperflexServerModelsMoidParams creates a new PatchHyperflexServerModelsMoidParams object
// with the default values initialized.
func NewPatchHyperflexServerModelsMoidParams() *PatchHyperflexServerModelsMoidParams {
	var ()
	return &PatchHyperflexServerModelsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexServerModelsMoidParamsWithTimeout creates a new PatchHyperflexServerModelsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexServerModelsMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexServerModelsMoidParams {
	var ()
	return &PatchHyperflexServerModelsMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexServerModelsMoidParamsWithContext creates a new PatchHyperflexServerModelsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexServerModelsMoidParamsWithContext(ctx context.Context) *PatchHyperflexServerModelsMoidParams {
	var ()
	return &PatchHyperflexServerModelsMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexServerModelsMoidParamsWithHTTPClient creates a new PatchHyperflexServerModelsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexServerModelsMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexServerModelsMoidParams {
	var ()
	return &PatchHyperflexServerModelsMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexServerModelsMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex server models moid operation typically these are written to a http.Request
*/
type PatchHyperflexServerModelsMoidParams struct {

	/*Body
	  hyperflexServerModel to update

	*/
	Body *models.HyperflexServerModel
	/*Moid
	  The moid of the hyperflexServerModel instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexServerModelsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) WithContext(ctx context.Context) *PatchHyperflexServerModelsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexServerModelsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) WithBody(body *models.HyperflexServerModel) *PatchHyperflexServerModelsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) SetBody(body *models.HyperflexServerModel) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) WithMoid(moid string) *PatchHyperflexServerModelsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex server models moid params
func (o *PatchHyperflexServerModelsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexServerModelsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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