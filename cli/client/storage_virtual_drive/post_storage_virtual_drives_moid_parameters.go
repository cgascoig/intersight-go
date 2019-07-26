// Code generated by go-swagger; DO NOT EDIT.

package storage_virtual_drive

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

// NewPostStorageVirtualDrivesMoidParams creates a new PostStorageVirtualDrivesMoidParams object
// with the default values initialized.
func NewPostStorageVirtualDrivesMoidParams() *PostStorageVirtualDrivesMoidParams {
	var ()
	return &PostStorageVirtualDrivesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageVirtualDrivesMoidParamsWithTimeout creates a new PostStorageVirtualDrivesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStorageVirtualDrivesMoidParamsWithTimeout(timeout time.Duration) *PostStorageVirtualDrivesMoidParams {
	var ()
	return &PostStorageVirtualDrivesMoidParams{

		timeout: timeout,
	}
}

// NewPostStorageVirtualDrivesMoidParamsWithContext creates a new PostStorageVirtualDrivesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostStorageVirtualDrivesMoidParamsWithContext(ctx context.Context) *PostStorageVirtualDrivesMoidParams {
	var ()
	return &PostStorageVirtualDrivesMoidParams{

		Context: ctx,
	}
}

// NewPostStorageVirtualDrivesMoidParamsWithHTTPClient creates a new PostStorageVirtualDrivesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostStorageVirtualDrivesMoidParamsWithHTTPClient(client *http.Client) *PostStorageVirtualDrivesMoidParams {
	var ()
	return &PostStorageVirtualDrivesMoidParams{
		HTTPClient: client,
	}
}

/*PostStorageVirtualDrivesMoidParams contains all the parameters to send to the API endpoint
for the post storage virtual drives moid operation typically these are written to a http.Request
*/
type PostStorageVirtualDrivesMoidParams struct {

	/*Body
	  storageVirtualDrive to update

	*/
	Body *models.StorageVirtualDrive
	/*Moid
	  The moid of the storageVirtualDrive instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) WithTimeout(timeout time.Duration) *PostStorageVirtualDrivesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) WithContext(ctx context.Context) *PostStorageVirtualDrivesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) WithHTTPClient(client *http.Client) *PostStorageVirtualDrivesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) WithBody(body *models.StorageVirtualDrive) *PostStorageVirtualDrivesMoidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) SetBody(body *models.StorageVirtualDrive) {
	o.Body = body
}

// WithMoid adds the moid to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) WithMoid(moid string) *PostStorageVirtualDrivesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the post storage virtual drives moid params
func (o *PostStorageVirtualDrivesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageVirtualDrivesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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