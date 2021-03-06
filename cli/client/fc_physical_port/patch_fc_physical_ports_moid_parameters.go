// Code generated by go-swagger; DO NOT EDIT.

package fc_physical_port

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

// NewPatchFcPhysicalPortsMoidParams creates a new PatchFcPhysicalPortsMoidParams object
// with the default values initialized.
func NewPatchFcPhysicalPortsMoidParams() *PatchFcPhysicalPortsMoidParams {
	var ()
	return &PatchFcPhysicalPortsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchFcPhysicalPortsMoidParamsWithTimeout creates a new PatchFcPhysicalPortsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchFcPhysicalPortsMoidParamsWithTimeout(timeout time.Duration) *PatchFcPhysicalPortsMoidParams {
	var ()
	return &PatchFcPhysicalPortsMoidParams{

		timeout: timeout,
	}
}

// NewPatchFcPhysicalPortsMoidParamsWithContext creates a new PatchFcPhysicalPortsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchFcPhysicalPortsMoidParamsWithContext(ctx context.Context) *PatchFcPhysicalPortsMoidParams {
	var ()
	return &PatchFcPhysicalPortsMoidParams{

		Context: ctx,
	}
}

// NewPatchFcPhysicalPortsMoidParamsWithHTTPClient creates a new PatchFcPhysicalPortsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchFcPhysicalPortsMoidParamsWithHTTPClient(client *http.Client) *PatchFcPhysicalPortsMoidParams {
	var ()
	return &PatchFcPhysicalPortsMoidParams{
		HTTPClient: client,
	}
}

/*PatchFcPhysicalPortsMoidParams contains all the parameters to send to the API endpoint
for the patch fc physical ports moid operation typically these are written to a http.Request
*/
type PatchFcPhysicalPortsMoidParams struct {

	/*Body
	  fcPhysicalPort to update

	*/
	Body *models.FcPhysicalPort
	/*Moid
	  The moid of the fcPhysicalPort instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) WithTimeout(timeout time.Duration) *PatchFcPhysicalPortsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) WithContext(ctx context.Context) *PatchFcPhysicalPortsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) WithHTTPClient(client *http.Client) *PatchFcPhysicalPortsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) WithBody(body *models.FcPhysicalPort) *PatchFcPhysicalPortsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) SetBody(body *models.FcPhysicalPort) {
	o.Body = body
}

// WithMoid adds the moid to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) WithMoid(moid string) *PatchFcPhysicalPortsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch fc physical ports moid params
func (o *PatchFcPhysicalPortsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchFcPhysicalPortsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
