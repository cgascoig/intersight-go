// Code generated by go-swagger; DO NOT EDIT.

package tam_security_advisory

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

// NewPatchTamSecurityAdvisoriesMoidParams creates a new PatchTamSecurityAdvisoriesMoidParams object
// with the default values initialized.
func NewPatchTamSecurityAdvisoriesMoidParams() *PatchTamSecurityAdvisoriesMoidParams {
	var ()
	return &PatchTamSecurityAdvisoriesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTamSecurityAdvisoriesMoidParamsWithTimeout creates a new PatchTamSecurityAdvisoriesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchTamSecurityAdvisoriesMoidParamsWithTimeout(timeout time.Duration) *PatchTamSecurityAdvisoriesMoidParams {
	var ()
	return &PatchTamSecurityAdvisoriesMoidParams{

		timeout: timeout,
	}
}

// NewPatchTamSecurityAdvisoriesMoidParamsWithContext creates a new PatchTamSecurityAdvisoriesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchTamSecurityAdvisoriesMoidParamsWithContext(ctx context.Context) *PatchTamSecurityAdvisoriesMoidParams {
	var ()
	return &PatchTamSecurityAdvisoriesMoidParams{

		Context: ctx,
	}
}

// NewPatchTamSecurityAdvisoriesMoidParamsWithHTTPClient creates a new PatchTamSecurityAdvisoriesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchTamSecurityAdvisoriesMoidParamsWithHTTPClient(client *http.Client) *PatchTamSecurityAdvisoriesMoidParams {
	var ()
	return &PatchTamSecurityAdvisoriesMoidParams{
		HTTPClient: client,
	}
}

/*PatchTamSecurityAdvisoriesMoidParams contains all the parameters to send to the API endpoint
for the patch tam security advisories moid operation typically these are written to a http.Request
*/
type PatchTamSecurityAdvisoriesMoidParams struct {

	/*Body
	  tamSecurityAdvisory to update

	*/
	Body *models.TamSecurityAdvisory
	/*Moid
	  The moid of the tamSecurityAdvisory instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) WithTimeout(timeout time.Duration) *PatchTamSecurityAdvisoriesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) WithContext(ctx context.Context) *PatchTamSecurityAdvisoriesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) WithHTTPClient(client *http.Client) *PatchTamSecurityAdvisoriesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) WithBody(body *models.TamSecurityAdvisory) *PatchTamSecurityAdvisoriesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) SetBody(body *models.TamSecurityAdvisory) {
	o.Body = body
}

// WithMoid adds the moid to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) WithMoid(moid string) *PatchTamSecurityAdvisoriesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch tam security advisories moid params
func (o *PatchTamSecurityAdvisoriesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTamSecurityAdvisoriesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
