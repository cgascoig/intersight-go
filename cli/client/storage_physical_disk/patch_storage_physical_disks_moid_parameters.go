// Code generated by go-swagger; DO NOT EDIT.

package storage_physical_disk

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

// NewPatchStoragePhysicalDisksMoidParams creates a new PatchStoragePhysicalDisksMoidParams object
// with the default values initialized.
func NewPatchStoragePhysicalDisksMoidParams() *PatchStoragePhysicalDisksMoidParams {
	var ()
	return &PatchStoragePhysicalDisksMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStoragePhysicalDisksMoidParamsWithTimeout creates a new PatchStoragePhysicalDisksMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStoragePhysicalDisksMoidParamsWithTimeout(timeout time.Duration) *PatchStoragePhysicalDisksMoidParams {
	var ()
	return &PatchStoragePhysicalDisksMoidParams{

		timeout: timeout,
	}
}

// NewPatchStoragePhysicalDisksMoidParamsWithContext creates a new PatchStoragePhysicalDisksMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStoragePhysicalDisksMoidParamsWithContext(ctx context.Context) *PatchStoragePhysicalDisksMoidParams {
	var ()
	return &PatchStoragePhysicalDisksMoidParams{

		Context: ctx,
	}
}

// NewPatchStoragePhysicalDisksMoidParamsWithHTTPClient creates a new PatchStoragePhysicalDisksMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStoragePhysicalDisksMoidParamsWithHTTPClient(client *http.Client) *PatchStoragePhysicalDisksMoidParams {
	var ()
	return &PatchStoragePhysicalDisksMoidParams{
		HTTPClient: client,
	}
}

/*PatchStoragePhysicalDisksMoidParams contains all the parameters to send to the API endpoint
for the patch storage physical disks moid operation typically these are written to a http.Request
*/
type PatchStoragePhysicalDisksMoidParams struct {

	/*Body
	  storagePhysicalDisk to update

	*/
	Body *models.StoragePhysicalDisk
	/*Moid
	  The moid of the storagePhysicalDisk instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) WithTimeout(timeout time.Duration) *PatchStoragePhysicalDisksMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) WithContext(ctx context.Context) *PatchStoragePhysicalDisksMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) WithHTTPClient(client *http.Client) *PatchStoragePhysicalDisksMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) WithBody(body *models.StoragePhysicalDisk) *PatchStoragePhysicalDisksMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) SetBody(body *models.StoragePhysicalDisk) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) WithMoid(moid string) *PatchStoragePhysicalDisksMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage physical disks moid params
func (o *PatchStoragePhysicalDisksMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStoragePhysicalDisksMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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