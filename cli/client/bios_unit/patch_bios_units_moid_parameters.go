// Code generated by go-swagger; DO NOT EDIT.

package bios_unit

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

// NewPatchBiosUnitsMoidParams creates a new PatchBiosUnitsMoidParams object
// with the default values initialized.
func NewPatchBiosUnitsMoidParams() *PatchBiosUnitsMoidParams {
	var ()
	return &PatchBiosUnitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchBiosUnitsMoidParamsWithTimeout creates a new PatchBiosUnitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchBiosUnitsMoidParamsWithTimeout(timeout time.Duration) *PatchBiosUnitsMoidParams {
	var ()
	return &PatchBiosUnitsMoidParams{

		timeout: timeout,
	}
}

// NewPatchBiosUnitsMoidParamsWithContext creates a new PatchBiosUnitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchBiosUnitsMoidParamsWithContext(ctx context.Context) *PatchBiosUnitsMoidParams {
	var ()
	return &PatchBiosUnitsMoidParams{

		Context: ctx,
	}
}

// NewPatchBiosUnitsMoidParamsWithHTTPClient creates a new PatchBiosUnitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchBiosUnitsMoidParamsWithHTTPClient(client *http.Client) *PatchBiosUnitsMoidParams {
	var ()
	return &PatchBiosUnitsMoidParams{
		HTTPClient: client,
	}
}

/*PatchBiosUnitsMoidParams contains all the parameters to send to the API endpoint
for the patch bios units moid operation typically these are written to a http.Request
*/
type PatchBiosUnitsMoidParams struct {

	/*Body
	  biosUnit to update

	*/
	Body *models.BiosUnit
	/*Moid
	  The moid of the biosUnit instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) WithTimeout(timeout time.Duration) *PatchBiosUnitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) WithContext(ctx context.Context) *PatchBiosUnitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) WithHTTPClient(client *http.Client) *PatchBiosUnitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) WithBody(body *models.BiosUnit) *PatchBiosUnitsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) SetBody(body *models.BiosUnit) {
	o.Body = body
}

// WithMoid adds the moid to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) WithMoid(moid string) *PatchBiosUnitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch bios units moid params
func (o *PatchBiosUnitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchBiosUnitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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