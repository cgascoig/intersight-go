// Code generated by go-swagger; DO NOT EDIT.

package niaapi_apic_field_notice

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

// NewGetNiaapiApicFieldNoticesMoidParams creates a new GetNiaapiApicFieldNoticesMoidParams object
// with the default values initialized.
func NewGetNiaapiApicFieldNoticesMoidParams() *GetNiaapiApicFieldNoticesMoidParams {
	var ()
	return &GetNiaapiApicFieldNoticesMoidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNiaapiApicFieldNoticesMoidParamsWithTimeout creates a new GetNiaapiApicFieldNoticesMoidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNiaapiApicFieldNoticesMoidParamsWithTimeout(timeout time.Duration) *GetNiaapiApicFieldNoticesMoidParams {
	var ()
	return &GetNiaapiApicFieldNoticesMoidParams{

		timeout: timeout,
	}
}

// NewGetNiaapiApicFieldNoticesMoidParamsWithContext creates a new GetNiaapiApicFieldNoticesMoidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNiaapiApicFieldNoticesMoidParamsWithContext(ctx context.Context) *GetNiaapiApicFieldNoticesMoidParams {
	var ()
	return &GetNiaapiApicFieldNoticesMoidParams{

		Context: ctx,
	}
}

// NewGetNiaapiApicFieldNoticesMoidParamsWithHTTPClient creates a new GetNiaapiApicFieldNoticesMoidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNiaapiApicFieldNoticesMoidParamsWithHTTPClient(client *http.Client) *GetNiaapiApicFieldNoticesMoidParams {
	var ()
	return &GetNiaapiApicFieldNoticesMoidParams{
		HTTPClient: client,
	}
}

/*GetNiaapiApicFieldNoticesMoidParams contains all the parameters to send to the API endpoint
for the get niaapi apic field notices moid operation typically these are written to a http.Request
*/
type GetNiaapiApicFieldNoticesMoidParams struct {

	/*Moid
	  The moid of the niaapiApicFieldNotice instance.

	*/
	Moid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) WithTimeout(timeout time.Duration) *GetNiaapiApicFieldNoticesMoidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) WithContext(ctx context.Context) *GetNiaapiApicFieldNoticesMoidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) WithHTTPClient(client *http.Client) *GetNiaapiApicFieldNoticesMoidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoid adds the moid to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) WithMoid(moid string) *GetNiaapiApicFieldNoticesMoidParams {
	o.SetMoid(moid)
	return o
}

// SetMoid adds the moid to the get niaapi apic field notices moid params
func (o *GetNiaapiApicFieldNoticesMoidParams) SetMoid(moid string) {
	o.Moid = moid
}

// WriteToRequest writes these params to a swagger request
func (o *GetNiaapiApicFieldNoticesMoidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
