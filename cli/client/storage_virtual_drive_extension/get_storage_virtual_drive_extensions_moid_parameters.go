// Code generated by go-swagger; DO NOT EDIT.

package storage_virtual_drive_extension

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

// NewGetStorageVirtualDriveExtensionsMoidParams creates a new GetStorageVirtualDriveExtensionsMoidParams object
// with the default values initialized.
func NewGetStorageVirtualDriveExtensionsMoidParams() *GetStorageVirtualDriveExtensionsMoidParams {
	var ()
	return &GetStorageVirtualDriveExtensionsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageVirtualDriveExtensionsMoidParamsWithTimeout creates a new GetStorageVirtualDriveExtensionsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStorageVirtualDriveExtensionsMoidParamsWithTimeout(timeout time.Duration) *GetStorageVirtualDriveExtensionsMoidParams {
	var ()
	return &GetStorageVirtualDriveExtensionsMoidParams{

		timeout: timeout,
	}
}

// NewGetStorageVirtualDriveExtensionsMoidParamsWithContext creates a new GetStorageVirtualDriveExtensionsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStorageVirtualDriveExtensionsMoidParamsWithContext(ctx context.Context) *GetStorageVirtualDriveExtensionsMoidParams {
	var ()
	return &GetStorageVirtualDriveExtensionsMoidParams{

		Context: ctx,
	}
}

// NewGetStorageVirtualDriveExtensionsMoidParamsWithHTTPClient creates a new GetStorageVirtualDriveExtensionsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStorageVirtualDriveExtensionsMoidParamsWithHTTPClient(client *http.Client) *GetStorageVirtualDriveExtensionsMoidParams {
	var ()
	return &GetStorageVirtualDriveExtensionsMoidParams{
		HTTPClient: client,
	}
}

/*GetStorageVirtualDriveExtensionsMoidParams contains all the parameters to send to the API endpoint
for the get storage virtual drive extensions moid operation typically these are written to a http.Request
*/
type GetStorageVirtualDriveExtensionsMoidParams struct {

	/*Moid
	  The moid of the storageVirtualDriveExtension instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) WithTimeout(timeout time.Duration) *GetStorageVirtualDriveExtensionsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) WithContext(ctx context.Context) *GetStorageVirtualDriveExtensionsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) WithHTTPClient(client *http.Client) *GetStorageVirtualDriveExtensionsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) WithMoid(moid string) *GetStorageVirtualDriveExtensionsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get storage virtual drive extensions moid params
func (o *GetStorageVirtualDriveExtensionsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageVirtualDriveExtensionsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
