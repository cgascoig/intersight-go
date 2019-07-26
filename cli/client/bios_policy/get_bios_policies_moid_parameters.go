// Code generated by go-swagger; DO NOT EDIT.

package bios_policy

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

// NewGetBiosPoliciesMoidParams creates a new GetBiosPoliciesMoidParams object
// with the default values initialized.
func NewGetBiosPoliciesMoidParams() *GetBiosPoliciesMoidParams {
	var ()
	return &GetBiosPoliciesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBiosPoliciesMoidParamsWithTimeout creates a new GetBiosPoliciesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBiosPoliciesMoidParamsWithTimeout(timeout time.Duration) *GetBiosPoliciesMoidParams {
	var ()
	return &GetBiosPoliciesMoidParams{

		timeout: timeout,
	}
}

// NewGetBiosPoliciesMoidParamsWithContext creates a new GetBiosPoliciesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBiosPoliciesMoidParamsWithContext(ctx context.Context) *GetBiosPoliciesMoidParams {
	var ()
	return &GetBiosPoliciesMoidParams{

		Context: ctx,
	}
}

// NewGetBiosPoliciesMoidParamsWithHTTPClient creates a new GetBiosPoliciesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBiosPoliciesMoidParamsWithHTTPClient(client *http.Client) *GetBiosPoliciesMoidParams {
	var ()
	return &GetBiosPoliciesMoidParams{
		HTTPClient: client,
	}
}

/*GetBiosPoliciesMoidParams contains all the parameters to send to the API endpoint
for the get bios policies moid operation typically these are written to a http.Request
*/
type GetBiosPoliciesMoidParams struct {

	/*Moid
	  The moid of the biosPolicy instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) WithTimeout(timeout time.Duration) *GetBiosPoliciesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) WithContext(ctx context.Context) *GetBiosPoliciesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) WithHTTPClient(client *http.Client) *GetBiosPoliciesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) WithMoid(moid string) *GetBiosPoliciesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get bios policies moid params
func (o *GetBiosPoliciesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetBiosPoliciesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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