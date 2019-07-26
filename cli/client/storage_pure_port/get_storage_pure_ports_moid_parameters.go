// Code generated by go-swagger; DO NOT EDIT.

package storage_pure_port

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

// NewGetStoragePurePortsMoidParams creates a new GetStoragePurePortsMoidParams object
// with the default values initialized.
func NewGetStoragePurePortsMoidParams() *GetStoragePurePortsMoidParams {
	var ()
	return &GetStoragePurePortsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoragePurePortsMoidParamsWithTimeout creates a new GetStoragePurePortsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStoragePurePortsMoidParamsWithTimeout(timeout time.Duration) *GetStoragePurePortsMoidParams {
	var ()
	return &GetStoragePurePortsMoidParams{

		timeout: timeout,
	}
}

// NewGetStoragePurePortsMoidParamsWithContext creates a new GetStoragePurePortsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStoragePurePortsMoidParamsWithContext(ctx context.Context) *GetStoragePurePortsMoidParams {
	var ()
	return &GetStoragePurePortsMoidParams{

		Context: ctx,
	}
}

// NewGetStoragePurePortsMoidParamsWithHTTPClient creates a new GetStoragePurePortsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStoragePurePortsMoidParamsWithHTTPClient(client *http.Client) *GetStoragePurePortsMoidParams {
	var ()
	return &GetStoragePurePortsMoidParams{
		HTTPClient: client,
	}
}

/*GetStoragePurePortsMoidParams contains all the parameters to send to the API endpoint
for the get storage pure ports moid operation typically these are written to a http.Request
*/
type GetStoragePurePortsMoidParams struct {

	/*Moid
	  The moid of the storagePurePort instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) WithTimeout(timeout time.Duration) *GetStoragePurePortsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) WithContext(ctx context.Context) *GetStoragePurePortsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) WithHTTPClient(client *http.Client) *GetStoragePurePortsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) WithMoid(moid string) *GetStoragePurePortsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get storage pure ports moid params
func (o *GetStoragePurePortsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoragePurePortsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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