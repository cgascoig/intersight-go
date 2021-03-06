// Code generated by go-swagger; DO NOT EDIT.

package hcl_operating_system_vendor

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

// NewGetHclOperatingSystemVendorsMoidParams creates a new GetHclOperatingSystemVendorsMoidParams object
// with the default values initialized.
func NewGetHclOperatingSystemVendorsMoidParams() *GetHclOperatingSystemVendorsMoidParams {
	var ()
	return &GetHclOperatingSystemVendorsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHclOperatingSystemVendorsMoidParamsWithTimeout creates a new GetHclOperatingSystemVendorsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHclOperatingSystemVendorsMoidParamsWithTimeout(timeout time.Duration) *GetHclOperatingSystemVendorsMoidParams {
	var ()
	return &GetHclOperatingSystemVendorsMoidParams{

		timeout: timeout,
	}
}

// NewGetHclOperatingSystemVendorsMoidParamsWithContext creates a new GetHclOperatingSystemVendorsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHclOperatingSystemVendorsMoidParamsWithContext(ctx context.Context) *GetHclOperatingSystemVendorsMoidParams {
	var ()
	return &GetHclOperatingSystemVendorsMoidParams{

		Context: ctx,
	}
}

// NewGetHclOperatingSystemVendorsMoidParamsWithHTTPClient creates a new GetHclOperatingSystemVendorsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHclOperatingSystemVendorsMoidParamsWithHTTPClient(client *http.Client) *GetHclOperatingSystemVendorsMoidParams {
	var ()
	return &GetHclOperatingSystemVendorsMoidParams{
		HTTPClient: client,
	}
}

/*GetHclOperatingSystemVendorsMoidParams contains all the parameters to send to the API endpoint
for the get hcl operating system vendors moid operation typically these are written to a http.Request
*/
type GetHclOperatingSystemVendorsMoidParams struct {

	/*Moid
	  The moid of the hclOperatingSystemVendor instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) WithTimeout(timeout time.Duration) *GetHclOperatingSystemVendorsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) WithContext(ctx context.Context) *GetHclOperatingSystemVendorsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) WithHTTPClient(client *http.Client) *GetHclOperatingSystemVendorsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) WithMoid(moid string) *GetHclOperatingSystemVendorsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get hcl operating system vendors moid params
func (o *GetHclOperatingSystemVendorsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetHclOperatingSystemVendorsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
