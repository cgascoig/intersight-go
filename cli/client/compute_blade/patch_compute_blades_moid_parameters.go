// Code generated by go-swagger; DO NOT EDIT.

package compute_blade

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

// NewPatchComputeBladesMoidParams creates a new PatchComputeBladesMoidParams object
// with the default values initialized.
func NewPatchComputeBladesMoidParams() *PatchComputeBladesMoidParams {
	var ()
	return &PatchComputeBladesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchComputeBladesMoidParamsWithTimeout creates a new PatchComputeBladesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchComputeBladesMoidParamsWithTimeout(timeout time.Duration) *PatchComputeBladesMoidParams {
	var ()
	return &PatchComputeBladesMoidParams{

		timeout: timeout,
	}
}

// NewPatchComputeBladesMoidParamsWithContext creates a new PatchComputeBladesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchComputeBladesMoidParamsWithContext(ctx context.Context) *PatchComputeBladesMoidParams {
	var ()
	return &PatchComputeBladesMoidParams{

		Context: ctx,
	}
}

// NewPatchComputeBladesMoidParamsWithHTTPClient creates a new PatchComputeBladesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchComputeBladesMoidParamsWithHTTPClient(client *http.Client) *PatchComputeBladesMoidParams {
	var ()
	return &PatchComputeBladesMoidParams{
		HTTPClient: client,
	}
}

/*PatchComputeBladesMoidParams contains all the parameters to send to the API endpoint
for the patch compute blades moid operation typically these are written to a http.Request
*/
type PatchComputeBladesMoidParams struct {

	/*Body
	  computeBlade to update

	*/
	Body *models.ComputeBlade
	/*Moid
	  The moid of the computeBlade instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) WithTimeout(timeout time.Duration) *PatchComputeBladesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) WithContext(ctx context.Context) *PatchComputeBladesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) WithHTTPClient(client *http.Client) *PatchComputeBladesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) WithBody(body *models.ComputeBlade) *PatchComputeBladesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) SetBody(body *models.ComputeBlade) {
	o.Body = body
}

// WithMoid adds the moid to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) WithMoid(moid string) *PatchComputeBladesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch compute blades moid params
func (o *PatchComputeBladesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchComputeBladesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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