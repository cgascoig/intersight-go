// Code generated by go-swagger; DO NOT EDIT.

package appliance_setup_info

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

// NewPatchApplianceSetupInfosMoidParams creates a new PatchApplianceSetupInfosMoidParams object
// with the default values initialized.
func NewPatchApplianceSetupInfosMoidParams() *PatchApplianceSetupInfosMoidParams {
	var ()
	return &PatchApplianceSetupInfosMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchApplianceSetupInfosMoidParamsWithTimeout creates a new PatchApplianceSetupInfosMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchApplianceSetupInfosMoidParamsWithTimeout(timeout time.Duration) *PatchApplianceSetupInfosMoidParams {
	var ()
	return &PatchApplianceSetupInfosMoidParams{

		timeout: timeout,
	}
}

// NewPatchApplianceSetupInfosMoidParamsWithContext creates a new PatchApplianceSetupInfosMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchApplianceSetupInfosMoidParamsWithContext(ctx context.Context) *PatchApplianceSetupInfosMoidParams {
	var ()
	return &PatchApplianceSetupInfosMoidParams{

		Context: ctx,
	}
}

// NewPatchApplianceSetupInfosMoidParamsWithHTTPClient creates a new PatchApplianceSetupInfosMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchApplianceSetupInfosMoidParamsWithHTTPClient(client *http.Client) *PatchApplianceSetupInfosMoidParams {
	var ()
	return &PatchApplianceSetupInfosMoidParams{
		HTTPClient: client,
	}
}

/*PatchApplianceSetupInfosMoidParams contains all the parameters to send to the API endpoint
for the patch appliance setup infos moid operation typically these are written to a http.Request
*/
type PatchApplianceSetupInfosMoidParams struct {

	/*Body
	  applianceSetupInfo to update

	*/
	Body *models.ApplianceSetupInfo
	/*Moid
	  The moid of the applianceSetupInfo instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) WithTimeout(timeout time.Duration) *PatchApplianceSetupInfosMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) WithContext(ctx context.Context) *PatchApplianceSetupInfosMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) WithHTTPClient(client *http.Client) *PatchApplianceSetupInfosMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) WithBody(body *models.ApplianceSetupInfo) *PatchApplianceSetupInfosMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) SetBody(body *models.ApplianceSetupInfo) {
	o.Body = body
}

// WithMoid adds the moid to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) WithMoid(moid string) *PatchApplianceSetupInfosMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch appliance setup infos moid params
func (o *PatchApplianceSetupInfosMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchApplianceSetupInfosMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
