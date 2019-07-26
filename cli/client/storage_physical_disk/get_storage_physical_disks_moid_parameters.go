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
)

// NewGetStoragePhysicalDisksMoidParams creates a new GetStoragePhysicalDisksMoidParams object
// with the default values initialized.
func NewGetStoragePhysicalDisksMoidParams() *GetStoragePhysicalDisksMoidParams {
	var ()
	return &GetStoragePhysicalDisksMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoragePhysicalDisksMoidParamsWithTimeout creates a new GetStoragePhysicalDisksMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStoragePhysicalDisksMoidParamsWithTimeout(timeout time.Duration) *GetStoragePhysicalDisksMoidParams {
	var ()
	return &GetStoragePhysicalDisksMoidParams{

		timeout: timeout,
	}
}

// NewGetStoragePhysicalDisksMoidParamsWithContext creates a new GetStoragePhysicalDisksMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStoragePhysicalDisksMoidParamsWithContext(ctx context.Context) *GetStoragePhysicalDisksMoidParams {
	var ()
	return &GetStoragePhysicalDisksMoidParams{

		Context: ctx,
	}
}

// NewGetStoragePhysicalDisksMoidParamsWithHTTPClient creates a new GetStoragePhysicalDisksMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStoragePhysicalDisksMoidParamsWithHTTPClient(client *http.Client) *GetStoragePhysicalDisksMoidParams {
	var ()
	return &GetStoragePhysicalDisksMoidParams{
		HTTPClient: client,
	}
}

/*GetStoragePhysicalDisksMoidParams contains all the parameters to send to the API endpoint
for the get storage physical disks moid operation typically these are written to a http.Request
*/
type GetStoragePhysicalDisksMoidParams struct {

	/*Moid
	  The moid of the storagePhysicalDisk instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) WithTimeout(timeout time.Duration) *GetStoragePhysicalDisksMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) WithContext(ctx context.Context) *GetStoragePhysicalDisksMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) WithHTTPClient(client *http.Client) *GetStoragePhysicalDisksMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) WithMoid(moid string) *GetStoragePhysicalDisksMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get storage physical disks moid params
func (o *GetStoragePhysicalDisksMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoragePhysicalDisksMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param moid
	if err := r.SetPathParam("moid", o.Moid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
