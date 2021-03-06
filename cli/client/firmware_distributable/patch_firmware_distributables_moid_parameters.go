// Code generated by go-swagger; DO NOT EDIT.

package firmware_distributable

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

// NewPatchFirmwareDistributablesMoidParams creates a new PatchFirmwareDistributablesMoidParams object
// with the default values initialized.
func NewPatchFirmwareDistributablesMoidParams() *PatchFirmwareDistributablesMoidParams {
	var ()
	return &PatchFirmwareDistributablesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchFirmwareDistributablesMoidParamsWithTimeout creates a new PatchFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchFirmwareDistributablesMoidParamsWithTimeout(timeout time.Duration) *PatchFirmwareDistributablesMoidParams {
	var ()
	return &PatchFirmwareDistributablesMoidParams{

		timeout: timeout,
	}
}

// NewPatchFirmwareDistributablesMoidParamsWithContext creates a new PatchFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchFirmwareDistributablesMoidParamsWithContext(ctx context.Context) *PatchFirmwareDistributablesMoidParams {
	var ()
	return &PatchFirmwareDistributablesMoidParams{

		Context: ctx,
	}
}

// NewPatchFirmwareDistributablesMoidParamsWithHTTPClient creates a new PatchFirmwareDistributablesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchFirmwareDistributablesMoidParamsWithHTTPClient(client *http.Client) *PatchFirmwareDistributablesMoidParams {
	var ()
	return &PatchFirmwareDistributablesMoidParams{
		HTTPClient: client,
	}
}

/*PatchFirmwareDistributablesMoidParams contains all the parameters to send to the API endpoint
for the patch firmware distributables moid operation typically these are written to a http.Request
*/
type PatchFirmwareDistributablesMoidParams struct {

	/*Body
	  firmwareDistributable to update

	*/
	Body *models.FirmwareDistributable
	/*Moid
	  The moid of the firmwareDistributable instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) WithTimeout(timeout time.Duration) *PatchFirmwareDistributablesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) WithContext(ctx context.Context) *PatchFirmwareDistributablesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) WithHTTPClient(client *http.Client) *PatchFirmwareDistributablesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) WithBody(body *models.FirmwareDistributable) *PatchFirmwareDistributablesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) SetBody(body *models.FirmwareDistributable) {
	o.Body = body
}

// WithMoid adds the moid to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) WithMoid(moid string) *PatchFirmwareDistributablesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch firmware distributables moid params
func (o *PatchFirmwareDistributablesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchFirmwareDistributablesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
