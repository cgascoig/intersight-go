// Code generated by go-swagger; DO NOT EDIT.

package appliance_backup_policy

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

// NewPatchApplianceBackupPoliciesMoidParams creates a new PatchApplianceBackupPoliciesMoidParams object
// with the default values initialized.
func NewPatchApplianceBackupPoliciesMoidParams() *PatchApplianceBackupPoliciesMoidParams {
	var ()
	return &PatchApplianceBackupPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchApplianceBackupPoliciesMoidParamsWithTimeout creates a new PatchApplianceBackupPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchApplianceBackupPoliciesMoidParamsWithTimeout(timeout time.Duration) *PatchApplianceBackupPoliciesMoidParams {
	var ()
	return &PatchApplianceBackupPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewPatchApplianceBackupPoliciesMoidParamsWithContext creates a new PatchApplianceBackupPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchApplianceBackupPoliciesMoidParamsWithContext(ctx context.Context) *PatchApplianceBackupPoliciesMoidParams {
	var ()
	return &PatchApplianceBackupPoliciesMoidParams{

		Context: ctx,
	}
}

// NewPatchApplianceBackupPoliciesMoidParamsWithHTTPClient creates a new PatchApplianceBackupPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchApplianceBackupPoliciesMoidParamsWithHTTPClient(client *http.Client) *PatchApplianceBackupPoliciesMoidParams {
	var ()
	return &PatchApplianceBackupPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*PatchApplianceBackupPoliciesMoidParams contains all the parameters to send to the API endpoint
for the patch appliance backup policies moid operation typically these are written to a http.Request
*/
type PatchApplianceBackupPoliciesMoidParams struct {

	/*Body
	  applianceBackupPolicy to update

	*/
	Body *models.ApplianceBackupPolicy
	/*Moid
	  The moid of the applianceBackupPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) WithTimeout(timeout time.Duration) *PatchApplianceBackupPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) WithContext(ctx context.Context) *PatchApplianceBackupPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) WithHTTPClient(client *http.Client) *PatchApplianceBackupPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) WithBody(body *models.ApplianceBackupPolicy) *PatchApplianceBackupPoliciesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) SetBody(body *models.ApplianceBackupPolicy) {
	o.Body = body
}

// WithMoid adds the moid to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) WithMoid(moid string) *PatchApplianceBackupPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch appliance backup policies moid params
func (o *PatchApplianceBackupPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchApplianceBackupPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
