// Code generated by go-swagger; DO NOT EDIT.

package ls_service_profile

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

// NewPatchLsServiceProfilesMoidParams creates a new PatchLsServiceProfilesMoidParams object
// with the default values initialized.
func NewPatchLsServiceProfilesMoidParams() *PatchLsServiceProfilesMoidParams {
	var ()
	return &PatchLsServiceProfilesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLsServiceProfilesMoidParamsWithTimeout creates a new PatchLsServiceProfilesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLsServiceProfilesMoidParamsWithTimeout(timeout time.Duration) *PatchLsServiceProfilesMoidParams {
	var ()
	return &PatchLsServiceProfilesMoidParams{

		timeout: timeout,
	}
}

// NewPatchLsServiceProfilesMoidParamsWithContext creates a new PatchLsServiceProfilesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLsServiceProfilesMoidParamsWithContext(ctx context.Context) *PatchLsServiceProfilesMoidParams {
	var ()
	return &PatchLsServiceProfilesMoidParams{

		Context: ctx,
	}
}

// NewPatchLsServiceProfilesMoidParamsWithHTTPClient creates a new PatchLsServiceProfilesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLsServiceProfilesMoidParamsWithHTTPClient(client *http.Client) *PatchLsServiceProfilesMoidParams {
	var ()
	return &PatchLsServiceProfilesMoidParams{
		HTTPClient: client,
	}
}

/*PatchLsServiceProfilesMoidParams contains all the parameters to send to the API endpoint
for the patch ls service profiles moid operation typically these are written to a http.Request
*/
type PatchLsServiceProfilesMoidParams struct {

	/*Body
	  lsServiceProfile to update

	*/
	Body *models.LsServiceProfile
	/*Moid
	  The moid of the lsServiceProfile instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) WithTimeout(timeout time.Duration) *PatchLsServiceProfilesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) WithContext(ctx context.Context) *PatchLsServiceProfilesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) WithHTTPClient(client *http.Client) *PatchLsServiceProfilesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) WithBody(body *models.LsServiceProfile) *PatchLsServiceProfilesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) SetBody(body *models.LsServiceProfile) {
	o.Body = body
}

// WithMoid adds the moid to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) WithMoid(moid string) *PatchLsServiceProfilesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch ls service profiles moid params
func (o *PatchLsServiceProfilesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLsServiceProfilesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
