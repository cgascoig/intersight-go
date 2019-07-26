// Code generated by go-swagger; DO NOT EDIT.

package storage_physical_disk_extension

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

// NewPatchStoragePhysicalDiskExtensionsMoidParams creates a new PatchStoragePhysicalDiskExtensionsMoidParams object
// with the default values initialized.
func NewPatchStoragePhysicalDiskExtensionsMoidParams() *PatchStoragePhysicalDiskExtensionsMoidParams {
	var ()
	return &PatchStoragePhysicalDiskExtensionsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStoragePhysicalDiskExtensionsMoidParamsWithTimeout creates a new PatchStoragePhysicalDiskExtensionsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStoragePhysicalDiskExtensionsMoidParamsWithTimeout(timeout time.Duration) *PatchStoragePhysicalDiskExtensionsMoidParams {
	var ()
	return &PatchStoragePhysicalDiskExtensionsMoidParams{

		timeout: timeout,
	}
}

// NewPatchStoragePhysicalDiskExtensionsMoidParamsWithContext creates a new PatchStoragePhysicalDiskExtensionsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStoragePhysicalDiskExtensionsMoidParamsWithContext(ctx context.Context) *PatchStoragePhysicalDiskExtensionsMoidParams {
	var ()
	return &PatchStoragePhysicalDiskExtensionsMoidParams{

		Context: ctx,
	}
}

// NewPatchStoragePhysicalDiskExtensionsMoidParamsWithHTTPClient creates a new PatchStoragePhysicalDiskExtensionsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStoragePhysicalDiskExtensionsMoidParamsWithHTTPClient(client *http.Client) *PatchStoragePhysicalDiskExtensionsMoidParams {
	var ()
	return &PatchStoragePhysicalDiskExtensionsMoidParams{
		HTTPClient: client,
	}
}

/*PatchStoragePhysicalDiskExtensionsMoidParams contains all the parameters to send to the API endpoint
for the patch storage physical disk extensions moid operation typically these are written to a http.Request
*/
type PatchStoragePhysicalDiskExtensionsMoidParams struct {

	/*Body
	  storagePhysicalDiskExtension to update

	*/
	Body *models.StoragePhysicalDiskExtension
	/*Moid
	  The moid of the storagePhysicalDiskExtension instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) WithTimeout(timeout time.Duration) *PatchStoragePhysicalDiskExtensionsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) WithContext(ctx context.Context) *PatchStoragePhysicalDiskExtensionsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) WithHTTPClient(client *http.Client) *PatchStoragePhysicalDiskExtensionsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) WithBody(body *models.StoragePhysicalDiskExtension) *PatchStoragePhysicalDiskExtensionsMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) SetBody(body *models.StoragePhysicalDiskExtension) {
	o.Body = body
}

// WithMoid adds the moid to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) WithMoid(moid string) *PatchStoragePhysicalDiskExtensionsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the patch storage physical disk extensions moid params
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStoragePhysicalDiskExtensionsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
