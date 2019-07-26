// Code generated by go-swagger; DO NOT EDIT.

package hyperflex_local_credential_policy

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

// NewPatchHyperflexLocalCredentialPoliciesMoidParams creates a new PatchHyperflexLocalCredentialPoliciesMoidParams object
// with the default values initialized.
func NewPatchHyperflexLocalCredentialPoliciesMoidParams() *PatchHyperflexLocalCredentialPoliciesMoidParams {
	var ()
	return &PatchHyperflexLocalCredentialPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHyperflexLocalCredentialPoliciesMoidParamsWithTimeout creates a new PatchHyperflexLocalCredentialPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchHyperflexLocalCredentialPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	var ()
	return &PatchHyperflexLocalCredentialPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchHyperflexLocalCredentialPoliciesMoidParamsWithContext creates a new PatchHyperflexLocalCredentialPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchHyperflexLocalCredentialPoliciesMoidParamsWithContext(ctx context.Context) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	var ()
	return &PatchHyperflexLocalCredentialPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchHyperflexLocalCredentialPoliciesMoidParamsWithHTTPClient creates a new PatchHyperflexLocalCredentialPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchHyperflexLocalCredentialPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	var ()
	return &PatchHyperflexLocalCredentialPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchHyperflexLocalCredentialPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch hyperflex local credential policies moid operation typically these are written to a http.Request
*/
type PatchHyperflexLocalCredentialPoliciesMoidParams struct {

	/*Body
	  hyperflexLocalCredentialPolicy to update

	*/
	Body *models.HyperflexLocalCredentialPolicy
	/*Moid
	  The moid of the hyperflexLocalCredentialPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) WithContext(ctx context.Context) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) WithBody(body *models.HyperflexLocalCredentialPolicy) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) SetBody(body *models.HyperflexLocalCredentialPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) WithMoid(moid string) *PatchHyperflexLocalCredentialPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch hyperflex local credential policies moid params
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHyperflexLocalCredentialPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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