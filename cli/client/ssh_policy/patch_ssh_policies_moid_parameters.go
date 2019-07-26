// Code generated by go-swagger; DO NOT EDIT.

package ssh_policy

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

// NewPatchSSHPoliciesMoidParams creates a new PatchSSHPoliciesMoidParams object
// with the default values initialized.
func NewPatchSSHPoliciesMoidParams() *PatchSSHPoliciesMoidParams {
	var ()
	return &PatchSSHPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSSHPoliciesMoidParamsWithTimeout creates a new PatchSSHPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSSHPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchSSHPoliciesMoidParams {
	var ()
	return &PatchSSHPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchSSHPoliciesMoidParamsWithContext creates a new PatchSSHPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSSHPoliciesMoidParamsWithContext(ctx context.Context) *PatchSSHPoliciesMoidParams {
	var ()
	return &PatchSSHPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchSSHPoliciesMoidParamsWithHTTPClient creates a new PatchSSHPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSSHPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchSSHPoliciesMoidParams {
	var ()
	return &PatchSSHPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchSSHPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch SSH policies moid operation typically these are written to a http.Request
*/
type PatchSSHPoliciesMoidParams struct {

	/*Body
	  sshPolicy to update

	*/
	Body *models.SSHPolicy
	/*Moid
	  The moid of the sshPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchSSHPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) WithContext(ctx context.Context) *PatchSSHPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchSSHPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) WithBody(body *models.SSHPolicy) *PatchSSHPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) SetBody(body *models.SSHPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) WithMoid(moid string) *PatchSSHPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch SSH policies moid params
func (o *PatchSSHPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSSHPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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