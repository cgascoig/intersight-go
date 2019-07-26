// Code generated by go-swagger; DO NOT EDIT.

package storage_sas_expander

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

// NewGetStorageSasExpandersMoidParams creates a new GetStorageSasExpandersMoidParams object
// with the default values initialized.
func NewGetStorageSasExpandersMoidParams() *GetStorageSasExpandersMoidParams {
	var ()
	return &GetStorageSasExpandersMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageSasExpandersMoidParamsWithTimeout creates a new GetStorageSasExpandersMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStorageSasExpandersMoidParamsWithTimeout(timeout time.Duration) *GetStorageSasExpandersMoidParams {
	var ()
	return &GetStorageSasExpandersMoidParams{

		timeout: timeout,
	}
}

// NewGetStorageSasExpandersMoidParamsWithContext creates a new GetStorageSasExpandersMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStorageSasExpandersMoidParamsWithContext(ctx context.Context) *GetStorageSasExpandersMoidParams {
	var ()
	return &GetStorageSasExpandersMoidParams{

		Context: ctx,
	}
}

// NewGetStorageSasExpandersMoidParamsWithHTTPClient creates a new GetStorageSasExpandersMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStorageSasExpandersMoidParamsWithHTTPClient(client *http.Client) *GetStorageSasExpandersMoidParams {
	var ()
	return &GetStorageSasExpandersMoidParams{
		HTTPClient: client,
	}
}

/*GetStorageSasExpandersMoidParams contains all the parameters to send to the API endpoint
for the get storage sas expanders moid operation typically these are written to a http.Request
*/
type GetStorageSasExpandersMoidParams struct {

	/*Moid
	  The moid of the storageSasExpander instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) WithTimeout(timeout time.Duration) *GetStorageSasExpandersMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) WithContext(ctx context.Context) *GetStorageSasExpandersMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) WithHTTPClient(client *http.Client) *GetStorageSasExpandersMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) WithMoid(moid string) *GetStorageSasExpandersMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get storage sas expanders moid params
func (o *GetStorageSasExpandersMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageSasExpandersMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
