// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_flash_controller

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

// NewPatchStorageFlexFlashControllersMoidParams creates a new PatchStorageFlexFlashControllersMoidParams object
// with the default values initialized.
func NewPatchStorageFlexFlashControllersMoidParams() *PatchStorageFlexFlashControllersMoidParams {
	var ()
	return &PatchStorageFlexFlashControllersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStorageFlexFlashControllersMoidParamsWithTimeout creates a new PatchStorageFlexFlashControllersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStorageFlexFlashControllersMoidParamsWithTimeout(timeout time.Duration) *PatchStorageFlexFlashControllersMoidParams {
	var ()
	return &PatchStorageFlexFlashControllersMoidParams{

		timeout: timeout,
	}
}

// NewPatchStorageFlexFlashControllersMoidParamsWithContext creates a new PatchStorageFlexFlashControllersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStorageFlexFlashControllersMoidParamsWithContext(ctx context.Context) *PatchStorageFlexFlashControllersMoidParams {
	var ()
	return &PatchStorageFlexFlashControllersMoidParams{

		Context: ctx,
	}
}

// NewPatchStorageFlexFlashControllersMoidParamsWithHTTPClient creates a new PatchStorageFlexFlashControllersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStorageFlexFlashControllersMoidParamsWithHTTPClient(client *http.Client) *PatchStorageFlexFlashControllersMoidParams {
	var ()
	return &PatchStorageFlexFlashControllersMoidParams{
		HTTPClient: client,
	}
}

/*PatchStorageFlexFlashControllersMoidParams contains all the parameters to send to the API endpoint
for the patch storage flex flash controllers moid operation typically these are written to a http.Request
*/
type PatchStorageFlexFlashControllersMoidParams struct {

	/*Body
	  storageFlexFlashController to update

	*/
	Body *models.StorageFlexFlashController
	/*Moid
	  The moid of the storageFlexFlashController instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) WithTimeout(timeout time.Duration) *PatchStorageFlexFlashControllersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) WithContext(ctx context.Context) *PatchStorageFlexFlashControllersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) WithHTTPClient(client *http.Client) *PatchStorageFlexFlashControllersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) WithBody(body *models.StorageFlexFlashController) *PatchStorageFlexFlashControllersMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) SetBody(body *models.StorageFlexFlashController) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) WithMoid(moid string) *PatchStorageFlexFlashControllersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage flex flash controllers moid params
func (o *PatchStorageFlexFlashControllersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStorageFlexFlashControllersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
