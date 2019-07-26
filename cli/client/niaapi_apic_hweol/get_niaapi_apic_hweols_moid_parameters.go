// Code generated by go-swagger; DO NOT EDIT.

package niaapi_apic_hweol

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

// NewGetNiaapiApicHweolsMoidParams creates a new GetNiaapiApicHweolsMoidParams object
// with the default values initialized.
func NewGetNiaapiApicHweolsMoidParams() *GetNiaapiApicHweolsMoidParams {
	var ()
	return &GetNiaapiApicHweolsMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNiaapiApicHweolsMoidParamsWithTimeout creates a new GetNiaapiApicHweolsMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNiaapiApicHweolsMoidParamsWithTimeout(timeout time.Duration) *GetNiaapiApicHweolsMoidParams {
	var ()
	return &GetNiaapiApicHweolsMoidParams{

		timeout: timeout,
	}
}

// NewGetNiaapiApicHweolsMoidParamsWithContext creates a new GetNiaapiApicHweolsMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNiaapiApicHweolsMoidParamsWithContext(ctx context.Context) *GetNiaapiApicHweolsMoidParams {
	var ()
	return &GetNiaapiApicHweolsMoidParams{

		Context: ctx,
	}
}

// NewGetNiaapiApicHweolsMoidParamsWithHTTPClient creates a new GetNiaapiApicHweolsMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNiaapiApicHweolsMoidParamsWithHTTPClient(client *http.Client) *GetNiaapiApicHweolsMoidParams {
	var ()
	return &GetNiaapiApicHweolsMoidParams{
		HTTPClient: client,
	}
}

/*GetNiaapiApicHweolsMoidParams contains all the parameters to send to the API endpoint
for the get niaapi apic hweols moid operation typically these are written to a http.Request
*/
type GetNiaapiApicHweolsMoidParams struct {

	/*Moid
	  The moid of the niaapiApicHweol instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) WithTimeout(timeout time.Duration) *GetNiaapiApicHweolsMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) WithContext(ctx context.Context) *GetNiaapiApicHweolsMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) WithHTTPClient(client *http.Client) *GetNiaapiApicHweolsMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) WithMoid(moid string) *GetNiaapiApicHweolsMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get niaapi apic hweols moid params
func (o *GetNiaapiApicHweolsMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetNiaapiApicHweolsMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
