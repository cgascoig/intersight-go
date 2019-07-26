// Code generated by go-swagger; DO NOT EDIT.

package hcl_driver_image

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

// NewGetHclDriverImagesMoidParams creates a new GetHclDriverImagesMoidParams object
// with the default values initialized.
func NewGetHclDriverImagesMoidParams() *GetHclDriverImagesMoidParams {
	var ()
	return &GetHclDriverImagesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHclDriverImagesMoidParamsWithTimeout creates a new GetHclDriverImagesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHclDriverImagesMoidParamsWithTimeout(timeout time.Duration) *GetHclDriverImagesMoidParams {
	var ()
	return &GetHclDriverImagesMoidParams{

		timeout: timeout,
	}
}

// NewGetHclDriverImagesMoidParamsWithContext creates a new GetHclDriverImagesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHclDriverImagesMoidParamsWithContext(ctx context.Context) *GetHclDriverImagesMoidParams {
	var ()
	return &GetHclDriverImagesMoidParams{

		Context: ctx,
	}
}

// NewGetHclDriverImagesMoidParamsWithHTTPClient creates a new GetHclDriverImagesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHclDriverImagesMoidParamsWithHTTPClient(client *http.Client) *GetHclDriverImagesMoidParams {
	var ()
	return &GetHclDriverImagesMoidParams{
		HTTPClient: client,
	}
}

/*GetHclDriverImagesMoidParams contains all the parameters to send to the API endpoint
for the get hcl driver images moid operation typically these are written to a http.Request
*/
type GetHclDriverImagesMoidParams struct {

	/*Moid
	  The moid of the hclDriverImage instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) WithTimeout(timeout time.Duration) *GetHclDriverImagesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) WithContext(ctx context.Context) *GetHclDriverImagesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) WithHTTPClient(client *http.Client) *GetHclDriverImagesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) WithMoid(moid string) *GetHclDriverImagesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hcl driver images moid params
func (o *GetHclDriverImagesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHclDriverImagesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
