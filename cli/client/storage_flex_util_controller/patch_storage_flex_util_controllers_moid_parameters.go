// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_util_controller

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

// NewPatchStorageFlexUtilControllersMoidParams creates a new PatchStorageFlexUtilControllersMoidParams object
// with the default values initialized.
func NewPatchStorageFlexUtilControllersMoidParams() *PatchStorageFlexUtilControllersMoidParams {
	var ()
	return &PatchStorageFlexUtilControllersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStorageFlexUtilControllersMoidParamsWithTimeout creates a new PatchStorageFlexUtilControllersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStorageFlexUtilControllersMoidParamsWithTimeout(timeout time.Duration) *PatchStorageFlexUtilControllersMoidParams {
	var ()
	return &PatchStorageFlexUtilControllersMoidParams{

		timeout: timeout,
	}
}

// NewPatchStorageFlexUtilControllersMoidParamsWithContext creates a new PatchStorageFlexUtilControllersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStorageFlexUtilControllersMoidParamsWithContext(ctx context.Context) *PatchStorageFlexUtilControllersMoidParams {
	var ()
	return &PatchStorageFlexUtilControllersMoidParams{

		Context: ctx,
	}
}

// NewPatchStorageFlexUtilControllersMoidParamsWithHTTPClient creates a new PatchStorageFlexUtilControllersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStorageFlexUtilControllersMoidParamsWithHTTPClient(client *http.Client) *PatchStorageFlexUtilControllersMoidParams {
	var ()
	return &PatchStorageFlexUtilControllersMoidParams{
		HTTPClient: client,
	}
}

/*PatchStorageFlexUtilControllersMoidParams contains all the parameters to send to the API endpoint
for the patch storage flex util controllers moid operation typically these are written to a http.Request
*/
type PatchStorageFlexUtilControllersMoidParams struct {

	/*Body
	  storageFlexUtilController to update

	*/
	Body *models.StorageFlexUtilController
	/*Moid
	  The moid of the storageFlexUtilController instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) WithTimeout(timeout time.Duration) *PatchStorageFlexUtilControllersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) WithContext(ctx context.Context) *PatchStorageFlexUtilControllersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) WithHTTPClient(client *http.Client) *PatchStorageFlexUtilControllersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) WithBody(body *models.StorageFlexUtilController) *PatchStorageFlexUtilControllersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) SetBody(body *models.StorageFlexUtilController) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) WithMoid(moid string) *PatchStorageFlexUtilControllersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage flex util controllers moid params
func (o *PatchStorageFlexUtilControllersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStorageFlexUtilControllersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
