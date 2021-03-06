// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_array

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

// NewPatchStoragePureArraysMoidParams creates a new PatchStoragePureArraysMoidParams object
// with the default values initialized.
func NewPatchStoragePureArraysMoidParams() *PatchStoragePureArraysMoidParams {
	var ()
	return &PatchStoragePureArraysMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStoragePureArraysMoidParamsWithTimeout creates a new PatchStoragePureArraysMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStoragePureArraysMoidParamsWithTimeout(timeout time.Duration) *PatchStoragePureArraysMoidParams {
	var ()
	return &PatchStoragePureArraysMoidParams{

		timeout: timeout,
	}
}

// NewPatchStoragePureArraysMoidParamsWithContext creates a new PatchStoragePureArraysMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStoragePureArraysMoidParamsWithContext(ctx context.Context) *PatchStoragePureArraysMoidParams {
	var ()
	return &PatchStoragePureArraysMoidParams{

		Context: ctx,
	}
}

// NewPatchStoragePureArraysMoidParamsWithHTTPClient creates a new PatchStoragePureArraysMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStoragePureArraysMoidParamsWithHTTPClient(client *http.Client) *PatchStoragePureArraysMoidParams {
	var ()
	return &PatchStoragePureArraysMoidParams{
		HTTPClient: client,
	}
}

/*PatchStoragePureArraysMoidParams contains all the parameters to send to the API endpoint
for the patch storage pure arrays moid operation typically these are written to a http.Request
*/
type PatchStoragePureArraysMoidParams struct {

	/*Body
	  storagePureArray to update

	*/
	Body *models.StoragePureArray
	/*Moid
	  The moid of the storagePureArray instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) WithTimeout(timeout time.Duration) *PatchStoragePureArraysMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) WithContext(ctx context.Context) *PatchStoragePureArraysMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) WithHTTPClient(client *http.Client) *PatchStoragePureArraysMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) WithBody(body *models.StoragePureArray) *PatchStoragePureArraysMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) SetBody(body *models.StoragePureArray) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) WithMoid(moid string) *PatchStoragePureArraysMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage pure arrays moid params
func (o *PatchStoragePureArraysMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStoragePureArraysMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
