// Code generated by go-swagger; DO NOT EDIT.

package security_unit

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

// NewPatchSecurityUnitsMoidParams creates a new PatchSecurityUnitsMoidParams object
// with the default values initialized.
func NewPatchSecurityUnitsMoidParams() *PatchSecurityUnitsMoidParams {
	var ()
	return &PatchSecurityUnitsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSecurityUnitsMoidParamsWithTimeout creates a new PatchSecurityUnitsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSecurityUnitsMoidParamsWithTimeout(timeout time.Duration) *PatchSecurityUnitsMoidParams {
	var ()
	return &PatchSecurityUnitsMoidParams{

		timeout: timeout,
	}
}

// NewPatchSecurityUnitsMoidParamsWithContext creates a new PatchSecurityUnitsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSecurityUnitsMoidParamsWithContext(ctx context.Context) *PatchSecurityUnitsMoidParams {
	var ()
	return &PatchSecurityUnitsMoidParams{

		Context: ctx,
	}
}

// NewPatchSecurityUnitsMoidParamsWithHTTPClient creates a new PatchSecurityUnitsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSecurityUnitsMoidParamsWithHTTPClient(client *http.Client) *PatchSecurityUnitsMoidParams {
	var ()
	return &PatchSecurityUnitsMoidParams{
		HTTPClient: client,
	}
}

/*PatchSecurityUnitsMoidParams contains all the parameters to send to the API endpoint
for the patch security units moid operation typically these are written to a http.Request
*/
type PatchSecurityUnitsMoidParams struct {

	/*Body
	  securityUnit to update

	*/
	Body *models.SecurityUnit
	/*Moid
	  The moid of the securityUnit instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) WithTimeout(timeout time.Duration) *PatchSecurityUnitsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) WithContext(ctx context.Context) *PatchSecurityUnitsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) WithHTTPClient(client *http.Client) *PatchSecurityUnitsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) WithBody(body *models.SecurityUnit) *PatchSecurityUnitsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) SetBody(body *models.SecurityUnit) {
	o.Body = body
}

// WithMoid adds the moid to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) WithMoid(moid string) *PatchSecurityUnitsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch security units moid params
func (o *PatchSecurityUnitsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSecurityUnitsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
