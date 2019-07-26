// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_disk

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

// NewGetStoragePureDisksMoidParams creates a new GetStoragePureDisksMoidParams object
// with the default values initialized.
func NewGetStoragePureDisksMoidParams() *GetStoragePureDisksMoidParams {
	var ()
	return &GetStoragePureDisksMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoragePureDisksMoidParamsWithTimeout creates a new GetStoragePureDisksMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStoragePureDisksMoidParamsWithTimeout(timeout time.Duration) *GetStoragePureDisksMoidParams {
	var ()
	return &GetStoragePureDisksMoidParams{

		timeout: timeout,
	}
}

// NewGetStoragePureDisksMoidParamsWithContext creates a new GetStoragePureDisksMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStoragePureDisksMoidParamsWithContext(ctx context.Context) *GetStoragePureDisksMoidParams {
	var ()
	return &GetStoragePureDisksMoidParams{

		Context: ctx,
	}
}

// NewGetStoragePureDisksMoidParamsWithHTTPClient creates a new GetStoragePureDisksMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStoragePureDisksMoidParamsWithHTTPClient(client *http.Client) *GetStoragePureDisksMoidParams {
	var ()
	return &GetStoragePureDisksMoidParams{
		HTTPClient: client,
	}
}

/*GetStoragePureDisksMoidParams contains all the parameters to send to the API endpoint
for the get storage pure disks moid operation typically these are written to a http.Request
*/
type GetStoragePureDisksMoidParams struct {

	/*Moid
	  The moid of the storagePureDisk instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) WithTimeout(timeout time.Duration) *GetStoragePureDisksMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) WithContext(ctx context.Context) *GetStoragePureDisksMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) WithHTTPClient(client *http.Client) *GetStoragePureDisksMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) WithMoid(moid string) *GetStoragePureDisksMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get storage pure disks moid params
func (o *GetStoragePureDisksMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoragePureDisksMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
