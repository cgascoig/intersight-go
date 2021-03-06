// Code generated by go-swagger; DO NOT EDIT.

package softwarerepository_operating_system_file

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

// NewPatchSoftwarerepositoryOperatingSystemFilesMoidParams creates a new PatchSoftwarerepositoryOperatingSystemFilesMoidParams object
// with the default values initialized.
func NewPatchSoftwarerepositoryOperatingSystemFilesMoidParams() *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	var ()
	return &PatchSoftwarerepositoryOperatingSystemFilesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSoftwarerepositoryOperatingSystemFilesMoidParamsWithTimeout creates a new PatchSoftwarerepositoryOperatingSystemFilesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSoftwarerepositoryOperatingSystemFilesMoidParamsWithTimeout(timeout time.Duration) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	var ()
	return &PatchSoftwarerepositoryOperatingSystemFilesMoidParams{

		timeout: timeout,
	}
}

// NewPatchSoftwarerepositoryOperatingSystemFilesMoidParamsWithContext creates a new PatchSoftwarerepositoryOperatingSystemFilesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSoftwarerepositoryOperatingSystemFilesMoidParamsWithContext(ctx context.Context) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	var ()
	return &PatchSoftwarerepositoryOperatingSystemFilesMoidParams{

		Context: ctx,
	}
}

// NewPatchSoftwarerepositoryOperatingSystemFilesMoidParamsWithHTTPClient creates a new PatchSoftwarerepositoryOperatingSystemFilesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSoftwarerepositoryOperatingSystemFilesMoidParamsWithHTTPClient(client *http.Client) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	var ()
	return &PatchSoftwarerepositoryOperatingSystemFilesMoidParams{
		HTTPClient: client,
	}
}

/*PatchSoftwarerepositoryOperatingSystemFilesMoidParams contains all the parameters to send to the API endpoint
for the patch softwarerepository operating system files moid operation typically these are written to a http.Request
*/
type PatchSoftwarerepositoryOperatingSystemFilesMoidParams struct {

	/*Body
	  softwarerepositoryOperatingSystemFile to update

	*/
	Body *models.SoftwarerepositoryOperatingSystemFile
	/*Moid
	  The moid of the softwarerepositoryOperatingSystemFile instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) WithTimeout(timeout time.Duration) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) WithContext(ctx context.Context) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) WithHTTPClient(client *http.Client) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) WithBody(body *models.SoftwarerepositoryOperatingSystemFile) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) SetBody(body *models.SoftwarerepositoryOperatingSystemFile) {
	o.Body = body
}

// WithMoid adds the moid to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) WithMoid(moid string) *PatchSoftwarerepositoryOperatingSystemFilesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch softwarerepository operating system files moid params
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSoftwarerepositoryOperatingSystemFilesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
