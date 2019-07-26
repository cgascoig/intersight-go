// Code generated by go-swagger; DO NOT EDIT.

package boot_device_boot_mode

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

// NewPatchBootDeviceBootModesMoidParams creates a new PatchBootDeviceBootModesMoidParams object
// with the default values initialized.
func NewPatchBootDeviceBootModesMoidParams() *PatchBootDeviceBootModesMoidParams {
	var ()
	return &PatchBootDeviceBootModesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchBootDeviceBootModesMoidParamsWithTimeout creates a new PatchBootDeviceBootModesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchBootDeviceBootModesMoidParamsWithTimeout(timeout time.Duration) *PatchBootDeviceBootModesMoidParams {
	var ()
	return &PatchBootDeviceBootModesMoidParams{

		timeout: timeout,
	}
}

// NewPatchBootDeviceBootModesMoidParamsWithContext creates a new PatchBootDeviceBootModesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchBootDeviceBootModesMoidParamsWithContext(ctx context.Context) *PatchBootDeviceBootModesMoidParams {
	var ()
	return &PatchBootDeviceBootModesMoidParams{

		Context: ctx,
	}
}

// NewPatchBootDeviceBootModesMoidParamsWithHTTPClient creates a new PatchBootDeviceBootModesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchBootDeviceBootModesMoidParamsWithHTTPClient(client *http.Client) *PatchBootDeviceBootModesMoidParams {
	var ()
	return &PatchBootDeviceBootModesMoidParams{
		HTTPClient: client,
	}
}

/*PatchBootDeviceBootModesMoidParams contains all the parameters to send to the API endpoint
for the patch boot device boot modes moid operation typically these are written to a http.Request
*/
type PatchBootDeviceBootModesMoidParams struct {

	/*Body
	  bootDeviceBootMode to update

	*/
	Body *models.BootDeviceBootMode
	/*Moid
	  The moid of the bootDeviceBootMode instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) WithTimeout(timeout time.Duration) *PatchBootDeviceBootModesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) WithContext(ctx context.Context) *PatchBootDeviceBootModesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) WithHTTPClient(client *http.Client) *PatchBootDeviceBootModesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) WithBody(body *models.BootDeviceBootMode) *PatchBootDeviceBootModesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) SetBody(body *models.BootDeviceBootMode) {
	o.Body = body
}

// WithMoid adds the moid to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) WithMoid(moid string) *PatchBootDeviceBootModesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch boot device boot modes moid params
func (o *PatchBootDeviceBootModesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchBootDeviceBootModesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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