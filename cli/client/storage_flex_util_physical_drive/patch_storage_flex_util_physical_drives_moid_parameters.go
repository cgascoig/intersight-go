// Code generated by go-swagger; DO NOT EDIT.

package storage_flex_util_physical_drive

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

// NewPatchStorageFlexUtilPhysicalDrivesMoidParams creates a new PatchStorageFlexUtilPhysicalDrivesMoidParams object
// with the default values initialized.
func NewPatchStorageFlexUtilPhysicalDrivesMoidParams() *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	var ()
	return &PatchStorageFlexUtilPhysicalDrivesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStorageFlexUtilPhysicalDrivesMoidParamsWithTimeout creates a new PatchStorageFlexUtilPhysicalDrivesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStorageFlexUtilPhysicalDrivesMoidParamsWithTimeout(timeout time.Duration) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	var ()
	return &PatchStorageFlexUtilPhysicalDrivesMoidParams{

		timeout: timeout,
	}
}

// NewPatchStorageFlexUtilPhysicalDrivesMoidParamsWithContext creates a new PatchStorageFlexUtilPhysicalDrivesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStorageFlexUtilPhysicalDrivesMoidParamsWithContext(ctx context.Context) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	var ()
	return &PatchStorageFlexUtilPhysicalDrivesMoidParams{

		Context: ctx,
	}
}

// NewPatchStorageFlexUtilPhysicalDrivesMoidParamsWithHTTPClient creates a new PatchStorageFlexUtilPhysicalDrivesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStorageFlexUtilPhysicalDrivesMoidParamsWithHTTPClient(client *http.Client) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	var ()
	return &PatchStorageFlexUtilPhysicalDrivesMoidParams{
		HTTPClient: client,
	}
}

/*PatchStorageFlexUtilPhysicalDrivesMoidParams contains all the parameters to send to the API endpoint
for the patch storage flex util physical drives moid operation typically these are written to a http.Request
*/
type PatchStorageFlexUtilPhysicalDrivesMoidParams struct {

	/*Body
	  storageFlexUtilPhysicalDrive to update

	*/
	Body *models.StorageFlexUtilPhysicalDrive
	/*Moid
	  The moid of the storageFlexUtilPhysicalDrive instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) WithTimeout(timeout time.Duration) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) WithContext(ctx context.Context) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) WithHTTPClient(client *http.Client) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) WithBody(body *models.StorageFlexUtilPhysicalDrive) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) SetBody(body *models.StorageFlexUtilPhysicalDrive) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) WithMoid(moid string) *PatchStorageFlexUtilPhysicalDrivesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage flex util physical drives moid params
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStorageFlexUtilPhysicalDrivesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
