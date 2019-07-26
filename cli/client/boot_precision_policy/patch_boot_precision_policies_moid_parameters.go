// Code generated by go-swagger; DO NOT EDIT.

package boot_precision_policy

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

// NewPatchBootPrecisionPoliciesMoidParams creates a new PatchBootPrecisionPoliciesMoidParams object
// with the default values initialized.
func NewPatchBootPrecisionPoliciesMoidParams() *PatchBootPrecisionPoliciesMoidParams {
	var ()
	return &PatchBootPrecisionPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchBootPrecisionPoliciesMoidParamsWithTimeout creates a new PatchBootPrecisionPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchBootPrecisionPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchBootPrecisionPoliciesMoidParams {
	var ()
	return &PatchBootPrecisionPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchBootPrecisionPoliciesMoidParamsWithContext creates a new PatchBootPrecisionPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchBootPrecisionPoliciesMoidParamsWithContext(ctx context.Context) *PatchBootPrecisionPoliciesMoidParams {
	var ()
	return &PatchBootPrecisionPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchBootPrecisionPoliciesMoidParamsWithHTTPClient creates a new PatchBootPrecisionPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchBootPrecisionPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchBootPrecisionPoliciesMoidParams {
	var ()
	return &PatchBootPrecisionPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchBootPrecisionPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch boot precision policies moid operation typically these are written to a http.Request
*/
type PatchBootPrecisionPoliciesMoidParams struct {

	/*Body
	  bootPrecisionPolicy to update

	*/
	Body *models.BootPrecisionPolicy
	/*Moid
	  The moid of the bootPrecisionPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchBootPrecisionPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) WithContext(ctx context.Context) *PatchBootPrecisionPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchBootPrecisionPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) WithBody(body *models.BootPrecisionPolicy) *PatchBootPrecisionPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) SetBody(body *models.BootPrecisionPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) WithMoid(moid string) *PatchBootPrecisionPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch boot precision policies moid params
func (o *PatchBootPrecisionPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchBootPrecisionPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
